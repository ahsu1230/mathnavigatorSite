"use strict";
require("./classEdit.sass");
import axios from "axios";
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API, { reduceApiCalls } from "../api.js";
import { InputText, emptyValidator } from "../common/inputs/inputText.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

export class ClassEditPage extends React.Component {
    state = {
        isEdit: false,

        // class object
        classKey: "",
        timesStr: "",
        programId: "",
        semesterId: "",
        locationId: "",
        fullState: 0,
        googleClassCode: "",
        priceLumpSum: 0,
        pricePerSession: 0,
        paymentNotes: "",

        programs: [],
        semesters: [],
        locations: [],
        sessions: [],
        fullStates: [],
    };

    componentDidMount = () => {
        const classId = this.props.classId;
        const apiCalls = [
            API.get("api/programs/all"),
            API.get("api/semesters/all"),
            API.get("api/locations/all"),
            API.get("api/classes/full-states"),
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
                    const fullStates = responses[3].data;

                    const hasClassId = responses.length > 4;
                    let classObj = hasClassId ? responses[4].data : {};
                    const sessions = (hasClassId ? responses[5].data : []).map(
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
                        timesStr: classObj.timesStr || "",
                        programId: programId,
                        semesterId: semesterId,
                        locationId: locationId,
                        fullState: classObj.fullState || 0,
                        googleClassCode: classObj.googleClassCode || "",
                        priceLumpSum: classObj.priceLumpSum || 0,
                        pricePerSession: classObj.pricePerSession || 0,
                        paymentNotes: classObj.paymentNotes || "",

                        programs: programs,
                        semesters: semesters,
                        locations: locations,
                        sessions: sessions,
                        fullStates: fullStates,
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
        const values = [this.state.priceLumpSum, this.state.pricePerSession];
        return {
            validate: () => values.filter((x) => x).length == 1,
            message:
                "Only one of PriceLump or PricePerSession should be filled.",
        };
    };

    onSave = () => {
        const classId = this.createClassId();
        let classObj = {
            classId: classId,
            programId: this.state.programId,
            semesterId: this.state.semesterId,
            locationId: this.state.locationId,
            classKey: this.state.classKey,
            timesStr: this.state.timesStr,
            fullState: this.state.fullState,
            googleClassCode: this.state.googleClassCode,
            priceLumpSum: parseInt(this.state.priceLumpSum),
            pricePerSession: parseInt(this.state.pricePerSession),
            paymentNotes: this.state.paymentNotes,
        };

        if (this.state.isEdit) {
            return API.post("api/classes/class/" + classId, classObj);
        } else {
            return API.post("api/classes/create", classObj);
        }
    };

    onDelete = () => {
        const classId = this.createClassId();
        // return API.delete("api/classes/class/" + classId);   // hard-delete
        return API.delete("api/classes/archive/" + classId); // soft-delete
    };

    renderClassInformation = () => {
        const classId = this.createClassId();

        const programOptions = this.state.programs.map((program) => ({
            value: program.programId,
            displayName: program.title,
        }));

        const semesterOptions = this.state.semesters.map((semester) => ({
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
                        description="Enter the class key. (Example: class1, sectionA)"
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

    renderEditSection = () => {
        const locationOptions = this.state.locations.map((location) => ({
            value: location.locationId,
            displayName: location.locationId,
        }));
        const fullStateOptions = this.state.fullStates.map((item, index) => ({
            value: index,
            displayName: item,
        }));

        return (
            <div id="edit-section">
                <InputSelect
                    label="LocationId"
                    description="Select a location id"
                    required={true}
                    value={this.state.locationId}
                    onChangeCallback={(e) => this.handleChange(e, "locationId")}
                    options={locationOptions}
                    errorMessageIfEmpty={
                        <span>
                            There are no locations to choose from. Please add
                            one <Link to="/locations/add">here</Link>
                        </span>
                    }
                />

                <InputText
                    label="Display Times"
                    description="A display string to convey to users the class session
                                time every week. Each class session should be separated
                                by a comma. (Example: Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm)"
                    required={true}
                    value={this.state.timesStr}
                    onChangeCallback={(e) => this.handleChange(e, "timesStr")}
                    validators={[emptyValidator("time")]}
                />

                <InputSelect
                    label="Class Availability (Full State)"
                    description="Select a level of availability (i.e. full, almost full, or not full)."
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
                    value={this.state.priceLumpSum}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "priceLumpSum")
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
        );
    };

    render = () => {
        const classId = this.createClassId();
        const title = this.state.isEdit
            ? "Edit Class: " + classId
            : "Add Class";

        const content = (
            <div>
                {this.renderClassInformation()}
                {this.renderEditSection()}
            </div>
        );

        return (
            <div id="view-class-edit">
                <EditPageWrapper
                    isEdit={this.state.isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"classes"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityId={classId}
                    entityName={"class"}
                />
            </div>
        );
    };
}
