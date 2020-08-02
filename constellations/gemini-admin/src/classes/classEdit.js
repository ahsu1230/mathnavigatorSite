"use strict";
require("./classEdit.sass");
import axios from "axios";
import React from "react";
import moment from "moment";
import API, { executeApiCalls } from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";
import { InputSelect } from "../utils/inputSelect.js";
import { Link } from "react-router-dom";

export class ClassEditPage extends React.Component {
    state = {
        isEdit: false,

        // class object
        oldClassId: "",
        inputClassKey: "",
        inputTimeString: "",

        selectProgramId: "",
        selectSemesterId: "",
        selectLocationId: "",

        programs: [],
        semesters: [],
        locations: [],
        sessions: [],

        fullState: 0,
        googleClassCode: "",
        priceLump: 0,
        pricePerSession: 0,
        paymentNotes: "",
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

                    let selectedProgramId = hasClassId
                        ? classObj.programId
                        : programs[0].programId;
                    let selectedSemesterId = hasClassId
                        ? classObj.semesterId
                        : semesters[0].semesterId;
                    let selectedLocationId = hasClassId
                        ? classObj.locationId
                        : locations[0].locationId;

                    this.setState({
                        isEdit: !!classId,
                        oldClassId: classObj.classId,
                        inputClassKey: classObj.classKey || "",
                        inputTimeString: classObj.times || "",

                        selectProgramId: selectedProgramId,
                        selectSemesterId: selectedSemesterId,
                        selectLocationId: selectedLocationId,

                        programs: programs,
                        semesters: semesters,
                        locations: locations,
                        sessions: sessions,

                        fullState: classObj.fullState,
                    });
                })
            )
            .catch((err) => {
                console.log("Error: api call failed. " + err.message);
            });
    };

    createClassId = () => {
        let classId =
            this.state.selectProgramId + "_" + this.state.selectSemesterId;
        classId = this.state.inputClassKey
            ? classId + "_" + this.state.inputClassKey
            : classId;
        return classId;
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    handleIntegerChange = (event, value) => {
        let number = event.target.value;
        this.setState({ [value]: parseInt(number) });
    };

    onChangeFullState = (e) => {
        const value = e.target.value;
        this.setState({
            fullState: parseInt(value),
        });
    };

    onClickSave = () => {
        const oldClassId = this.state.oldClassId;
        const newClassId = this.createClassId();
        let classObj = {
            classId: newClassId,
            programId: this.state.selectProgramId,
            semesterId: this.state.selectSemesterId,
            locationId: this.state.selectLocationId,
            classKey: this.state.inputClassKey,
            times: this.state.inputTimeString,
            fullState: this.state.fullState,
            googleClassCode: this.state.googleClassCode,
            priceLump: this.state.priceLump,
            pricePerSession: this.state.pricePerSession,
            paymentNotes: this.state.paymentNotes,
            startDate: moment().toJSON(), // TODO: need to remove
            endDate: moment().add(30, "d").toJSON(), // TODO: need to remove
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) => alert("Could not save class: " + err);

        let apiCalls = [];
        if (this.state.isEdit) {
            apiCalls.push(
                API.post("api/classes/class/" + oldClassId, classObj)
            );
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

    render = () => {
        const title = this.state.isEdit ? "Edit Class" : "Add Class";

        const programOptions = this.state.programs.map((program, index) => ({
            value: program.programId,
            displayName: program.name,
        }));

        const semesterOptions = this.state.semesters.map((semester, index) => ({
            value: semester.semesterId,
            displayName: semester.title,
        }));

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

        const classId = this.createClassId();

        let classInformation = <h3 className="class-id">ClassId: {classId}</h3>;
        if (!this.state.isEdit) {
            classInformation = (
                <div className="edit-section">
                    <h3>Class Information</h3>

                    <InputSelect
                        label="ProgramId"
                        description="Select a program id"
                        value={this.state.selectProgramId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "selectProgramId")
                        }
                        required={true}
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
                        value={this.state.selectSemesterId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "selectSemesterId")
                        }
                        required={true}
                        options={semesterOptions}
                        errorMessageIfEmpty={
                            <span>
                                There are no semesters to choose from. Please
                                add one <Link to="/semesters/add">here</Link>
                            </span>
                        }
                    />

                    <h4>ClassKey</h4>
                    <input
                        value={this.state.inputClassKey}
                        placeholder="Optional"
                        onChange={(e) => this.handleChange(e, "inputClassKey")}
                    />

                    <h3 className="class-id">ClassId: {classId}</h3>
                </div>
            );
        }

        let deleteButton = <div></div>;
        if (this.state.isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        const modalDiv = renderModal(
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
                {classInformation}
                <div className="edit-section">
                    <h3>Class Schedule</h3>

                    <InputSelect
                        label="LocationId"
                        description="Select a location id"
                        value={this.state.selectLocationId}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "selectLocationId")
                        }
                        required={true}
                        options={locationOptions}
                        errorMessageIfEmpty={
                            <span>
                                There are no locations to choose from. Please
                                add one <Link to="/locations/add">here</Link>
                            </span>
                        }
                    />

                    <h4>Display Time</h4>
                    <p>
                        A display string to convey to users the class session
                        time every week. Each class session should be separated
                        by a comma.
                        <br />
                        Example: Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm
                    </p>
                    <input
                        value={this.state.inputTimeString}
                        placeholder="i.e. Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm"
                        onChange={(e) =>
                            this.handleChange(e, "inputTimeString")
                        }
                    />

                    <InputSelect
                        label="Class Availability"
                        description="Select a level of availability"
                        value={this.state.fullState}
                        onChangeCallback={(e) => this.onChangeFullState(e)}
                        required={true}
                        options={fullStateOptions}
                    />

                    <InputText
                        label="Google Classroom Code"
                        required={false}
                        description="Enter the google classroom code"
                        value={this.state.googleClassCode}
                        onChangeCallback={(e) =>
                            this.handleChange(e, "googleClassCode")
                        }
                    />

                    <InputText
                        label="Price Lump"
                        required={false}
                        description="Enter price for one time payment (Either enter only in this field or only in the price per session field)"
                        onChangeCallback={(e) =>
                            this.handleIntegerChange(e, "priceLump")
                        }
                        validators={[
                            {
                                validate: (input) => parseInt(input) != NaN,
                                message: "Price must be a valid number",
                            },
                            {
                                validate: (input) => input != "",
                                validate: this.state.pricePerSession != 0,
                                message:
                                    "Only one of PriceLump or PricePerSession should be filled. Both cannot be filled",
                            },
                        ]}
                    />

                    <InputText
                        label="Price Per Session"
                        required={false}
                        description="Enter price for one time payment (Either enter only in this field or only in the price lump field)"
                        onChangeCallback={(e) =>
                            this.handleIntegerChange(e, "pricePerSession")
                        }
                        validators={[
                            {
                                validate: (input) => parseInt(input) != "NaN",
                                message: "Price must be a valid number",
                            },
                            {
                                validate: (input) => input != "",
                                validate: this.state.priceLump != 0,
                                message:
                                    "Only one of PriceLump or PricePerSession should be filled. Both cannot be filled",
                            },
                        ]}
                    />

                    <InputText
                        label="Payment Notes"
                        isTextBox={true}
                        required={false}
                        description="Enter payment notes"
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

function renderModal(
    showSaveModal,
    showDeleteModal,
    onModalOkSaved,
    onModalDeleteConfirm,
    onModalDismiss
) {
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
}
