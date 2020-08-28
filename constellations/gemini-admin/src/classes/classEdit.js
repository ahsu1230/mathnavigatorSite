"use strict";
require("./classEdit.sass");
import axios from "axios";
import React from "react";
import moment from "moment";
import API, { executeApiCalls } from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText, emptyValidator } from "../utils/inputText.js";
import { InputSelect } from "../utils/inputSelect.js";
import { Link } from "react-router-dom";

export class ClassEditPage extends React.Component {
    state = {
        isEdit: false,

        // class object
        classKey: "",
        times: "",
        programId: "",
        semesterId: "",
        locationId: "",
        fullState: 0,
        googleClassCode: "",
        priceLump: 0,
        pricePerSession: 0,
        paymentNotes: "",

        programs: [],
        semesters: [],
        locations: [],
        sessions: [],
    };

    componentDidMount = () => {
        const classId = this.props.classId;
        const apiCalls = [
            API.get("api/programs/all"),
            API.get("api/semesters/all"),
            API.get("api/locations/all"),
        ];
        if (classId) {
            apiCalls.push(API.get("api/classes/class/" + classId));
            apiCalls.push(API.get("api/sessions/class/" + classId));
        }

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const programs = responses[0].data;
                    const semesters = responses[1].data;
                    const locations = responses[2].data;

                    const hasClassId = responses.length > 3;
                    let classObj = hasClassId ? responses[3].data : {};
                    const sessions = (hasClassId ? responses[4].data : []).map(
                        (session) => {
                            session.startsAt = moment(session.startsAt);
                            session.endsAt = moment(session.endsAt);
                            return session;
                        }
                    );

                    let programId = hasClassId
                        ? classObj.programId
                        : programs.length
                        ? programs[0].programId
                        : undefined;
                    let semesterId = hasClassId
                        ? classObj.semesterId
                        : semesters.length
                        ? semesters[0].semesterId
                        : undefined;
                    let locationId = hasClassId
                        ? classObj.locationId
                        : programs.length
                        ? locations[0].locationId
                        : undefined;

                    this.setState({
                        isEdit: !!classId,
                        classKey: classObj.classKey || "",
                        times: classObj.times || "",
                        programId: programId,
                        semesterId: semesterId,
                        locationId: locationId,
                        fullState: classObj.fullState,
                        googleClassCode: classObj.googleClassCode || "",
                        priceLump: classObj.priceLump || 0,
                        pricePerSession: classObj.pricePerSession || 0,
                        paymentNotes: classObj.paymentNotes || "",

                        programs: programs,
                        semesters: semesters,
                        locations: locations,
                        sessions: sessions,
                    });
                })
            )
            .catch((err) => {
                console.log("Error: api call failed. " + err.message);
            });
    };

    createClassId = () => {
        let classId = this.state.programId + "_" + this.state.semesterId;
        classId = this.state.classKey
            ? classId + "_" + this.state.classKey
            : classId;
        return classId;
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onChangeFullState = (e) => {
        const value = e.target.value;
        this.setState({
            fullState: parseInt(value),
        });
    };

    getPaymentValidator = () => {
        const values = [this.state.priceLump, this.state.pricePerSession];
        return {
            validate: () => values.filter((x) => x).length == 1,
            message:
                "Only one of PriceLump or PricePerSession should be filled.",
        };
    };

    onClickSave = () => {
        const classId = this.createClassId();
        let classObj = {
            classId: classId,
            programId: this.state.programId,
            semesterId: this.state.semesterId,
            locationId: this.state.locationId,
            classKey: this.state.classKey,
            times: this.state.times,
            fullState: this.state.fullState,
            googleClassCode: this.state.googleClassCode,
            priceLump: parseInt(this.state.priceLump),
            pricePerSession: parseInt(this.state.pricePerSession),
            paymentNotes: this.state.paymentNotes,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) => alert("Could not save class: " + err);

        let apiCalls = [];
        if (this.state.isEdit) {
            apiCalls.push(API.post("api/classes/class/" + classId, classObj));
        } else {
            apiCalls.push(API.post("api/classes/create", classObj));
        }

        executeApiCalls(apiCalls, successCallback, failCallback);
    };

    onClickCancel = () => {
        window.location.hash = "classes";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onModalDeleteConfirm = () => {
        const classId = this.props.classId;

        let apiCalls = [];
        let successCallback = () => (window.location.hash = "classes");
        let failCallback = (err) =>
            alert("Could not delete class or sessions: " + err);

        // Must delete sessions before deleting class
        var sessionIds = [];
        this.state.sessions.forEach((session) => {
            sessionIds.push(session.id);
        });

        apiCalls.push(API.delete("api/sessions/delete", { data: sessionIds }));
        apiCalls.push(API.delete("api/classes/class/" + classId));

        executeApiCalls(apiCalls, successCallback, failCallback);
    };

    onModalOkSaved = () => {
        this.onModalDismiss();
        window.location.hash = "classes";
    };

    onModalDismiss = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    renderModal = (
        showSaveModal,
        showDeleteModal,
        onModalOkSaved,
        onModalDeleteConfirm,
        onModalDismiss
    ) => {
        let modalDiv;
        let modalContent;
        let showModal;
        if (showDeleteModal) {
            showModal = showDeleteModal;
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={onModalDeleteConfirm}
                    onReject={onModalDismiss}
                />
            );
        }
        if (showSaveModal) {
            showModal = showSaveModal;
            modalContent = (
                <OkayModal
                    text={"Class information saved!"}
                    onOkay={onModalOkSaved}
                />
            );
        }
        if (modalContent) {
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showModal}
                    onDismiss={onModalDismiss}
                />
            );
        }
        return modalDiv;
    };

    renderClassInformation = () => {
        const classId = this.createClassId();

        const programOptions = this.state.programs.map((program, index) => ({
            value: program.programId,
            displayName: program.name,
        }));

        const semesterOptions = this.state.semesters.map((semester, index) => ({
            value: semester.semesterId,
            displayName: semester.title,
        }));

        let classInformation = <div></div>;
        if (!this.state.isEdit) {
            classInformation = (
                <div id="class-information">
                    <InputSelect
                        label="ProgramId"
                        description="Select a program id"
                        required={true}
                        value={this.state.programId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "programId")
                        }
                        options={programOptions}
                        errorMessageIfEmpty={
                            <span>
                                There are no programs to choose from. Please add
                                one <Link to="/programs/add">here</Link>
                            </span>
                        }
                    />

                    <InputSelect
                        label="SemesterId"
                        description="Select a semester id"
                        required={true}
                        value={this.state.semesterId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "semesterId")
                        }
                        options={semesterOptions}
                        errorMessageIfEmpty={
                            <span>
                                There are no semesters to choose from. Please
                                add one <Link to="/semesters/add">here</Link>
                            </span>
                        }
                    />

                    <InputText
                        label="ClassKey"
                        description="Enter the class key. (Example: class1)"
                        value={this.state.classKey}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "classKey")
                        }
                    />

                    <h3 className="class-id">ClassId: {classId}</h3>
                </div>
            );
        }
        return classInformation;
    };

    render = () => {
        const title = this.state.isEdit
            ? "Edit Class: " + this.createClassId()
            : "Add Class";

        const locationOptions = this.state.locations.map((location, index) => ({
            value: location.locationId,
            displayName: location.locationId,
        }));

        const fullStateOptions = ["Normal", "Almost Full", "Full"].map(
            (item, index) => ({
                value: index,
                displayName: item,
            })
        );

        let deleteButton = <div></div>;
        if (this.state.isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        const modalDiv = this.renderModal(
            this.state.showSaveModal,
            this.state.showDeleteModal,
            this.onModalOkSaved,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        return (
            <div id="view-class-edit">
                {modalDiv}
                <h2>{title}</h2>
                {this.renderClassInformation()}
                <div id="edit-section">
                    <InputSelect
                        label="LocationId"
                        description="Select a location id"
                        required={true}
                        value={this.state.locationId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "locationId")
                        }
                        options={locationOptions}
                        errorMessageIfEmpty={
                            <span>
                                There are no locations to choose from. Please
                                add one <Link to="/locations/add">here</Link>
                            </span>
                        }
                    />

                    <InputText
                        label="Display Time"
                        description="A display string to convey to users the class session
                                    time every week. Each class session should be separated
                                    by a comma. (Example: Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm)"
                        required={true}
                        value={this.state.times}
                        onChangeCallback={(e) => this.handleChange(e, "times")}
                        validators={[emptyValidator("time")]}
                    />

                    <InputSelect
                        label="Class Availability"
                        description="Select a level of availability"
                        required={true}
                        value={this.state.fullState}
                        onChangeCallback={(e) => this.onChangeFullState(e)}
                        options={fullStateOptions}
                    />

                    <InputText
                        label="Google Classroom Code"
                        description="Enter the google classroom code"
                        value={this.state.googleClassCode}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "googleClassCode")
                        }
                    />

                    <InputText
                        label="Price Lump"
                        description="Enter price for one time payment (Either enter only in this field or only in the price per session field)"
                        required={true}
                        value={this.state.priceLump}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "priceLump")
                        }
                        validators={[this.getPaymentValidator()]}
                    />

                    <InputText
                        label="Price Per Session"
                        description="Enter price for one time payment (Either enter only in this field or only in the price lump field)"
                        required={true}
                        value={this.state.pricePerSession}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "pricePerSession")
                        }
                        validators={[this.getPaymentValidator()]}
                    />

                    <InputText
                        label="Payment Notes"
                        description="Enter payment notes"
                        isTextBox={true}
                        value={this.state.paymentNotes}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "paymentNotes")
                        }
                    />
                </div>

                <div className="buttons">
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    {deleteButton}
                </div>
            </div>
        );
    };
}
