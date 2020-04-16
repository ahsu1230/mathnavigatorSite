"use strict";
require("./classEdit.styl");
import axios from "axios";
import React from "react";
import moment from "moment";
import API from "../api.js";
import { ClassSessions } from "./classSessions.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

// React DatePicker
import "react-dates/initialize";
import { DateRangePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";

export class ClassEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            classObj: {},
            isEdit: false,
            selectProgramId: "",
            selectSemesterId: "",
            selectLocationId: "",
            inputClassKey: "",
            inputTimeString: "",
            listPrograms: [],
            listSemesters: [],
            listLocations: [],
            startDate: moment(),
            endDate: moment(),
            focusedInput: undefined
        };

        this.handleChange = this.handleChange.bind(this);

        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        this.onModalDeleteConfirm = this.onModalDeleteConfirm.bind(this);
        this.onModalOkSaved = this.onModalOkSaved.bind(this);
        this.onModalDismiss = this.onModalDismiss.bind(this);

        this.onSaveSessions = this.onSaveSessions.bind(this);
    }

    componentDidMount() {
        const classId = this.props.classId;
        const apiCalls = [
            API.get("api/programs/v1/all"),
            API.get("api/semesters/v1/all"),
            API.get("api/locations/v1/all"),
        ];
        if (classId) {
            apiCalls.push(API.get("api/classes/v1/class/" + classId));
        }

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const programs = responses[0].data;
                    const semesters = responses[1].data;
                    const locations = responses[2].data;
                    const firstProgram = programs.length > 0 ? programs[0] : "";
                    const firstSemester =
                        semesters.length > 0 ? semesters[0] : "";
                    const firstLocations =
                        locations.length > 0 ? locations[0] : "";

                    const hasClassId = responses.length > 3;
                    let classObj = hasClassId ? responses[3].data : {};
                    
                    this.setState({
                        listPrograms: programs,
                        listSemesters: semesters,
                        listLocations: locations,
                        classObj: classObj,
                        selectProgramId: firstProgram.programId,
                        selectSemesterId: firstSemester.semesterId,
                        selectLocationId: firstLocations.locId,
                    });
                })
            )
            .catch((errors) => {
                console.log("Error: api call failed. " + errors.message);
            });
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickSave() {
        const oldClassId = this.state.classObj
            ? this.state.classObj.classId
            : undefined;
        const newClassId = createClassId(
            this.state.selectProgramId,
            this.state.selectSemesterId,
            this.state.inputClassKey
        );
        let classObj = {
            classId: newClassId,
            programId: this.state.selectProgramId,
            semesterId: this.state.selectSemesterId,
            locId: this.state.selectLocationId,
            classKey: this.state.inputClassKey,
            times: this.state.inputTimeString,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save class: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/classes/v1/class/" + oldClassId, classObj)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/classes/v1/create", classObj)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    }

    onClickCancel() {
        window.location.hash = "classes";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onModalDeleteConfirm() {
        const classId = this.props.classId;
        API.delete("api/classes/v1/class/" + classId).then((res) => {
            window.location.hash = "classes";
        });
    }

    onModalOkSaved() {
        this.onModalDismiss();
        window.location.hash = "classes";
    }

    onModalDismiss() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    }

    onSaveSessions() {
        console.log("API - send sessions!");
    }

    render() {
        const title = this.state.isEdit ? "Edit Class" : "Add Class";

        const optPrograms = renderOptionsPrograms(this.state.listPrograms);
        const optSemesters = renderOptionsSemesters(this.state.listSemesters);
        const optLocations = renderOptionsLocations(this.state.listLocations);
        const classId = createClassId(
            this.state.selectProgramId,
            this.state.selectSemesterId,
            this.state.inputClassKey
        );

        const deleteButton = renderDeleteButton(
            this.state.isEdit,
            this.onClickDelete
        );
        const modalDiv = renderModal(
            this.state.showSaveModal,
            this.state.showDeleteModal,
            this.onModalOkSaved,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );
        return (
            <div id="view-class-edit">
                <div className="buttons upper-right">
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    {deleteButton}
                </div>
                {modalDiv}
                <h2>{title}</h2>

                <div className="edit-section">
                    <h3>Class Information</h3>

                    <h4>ProgramId</h4>
                    <select
                        value={this.state.selectProgramId}
                        onChange={(e) =>
                            this.handleChange(e, "selectProgramId")
                        }>
                        {optPrograms}
                    </select>

                    <h4>SemesterId</h4>
                    <select
                        value={this.state.selectSemesterId}
                        onChange={(e) =>
                            this.handleChange(e, "selectSemesterId")
                        }>
                        {optSemesters}
                    </select>

                    <h4>ClassKey</h4>
                    <input
                        value={this.state.inputClassKey}
                        onChange={(e) => this.handleChange(e, "inputClassKey")}
                    />

                    <h3 className="class-id">ClassId: {classId}</h3>
                </div>

                <div className="edit-section">
                    <h3>Class Schedule</h3>

                    <h4>LocationId</h4>
                    <select
                        value={this.state.selectLocationId}
                        onChange={(e) =>
                            this.handleChange(e, "selectLocationId")
                        }>
                        {optLocations}
                    </select>

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
                </div>

                <div className="edit-section">
                    <h4>Dates</h4>
                    <DateRangePicker
                        startDate={this.state.startDate} // momentPropTypes.momentObj or null,
                        startDateId="your_unique_start_date_id" // PropTypes.string.isRequired,
                        endDate={this.state.endDate} // momentPropTypes.momentObj or null,
                        endDateId="your_unique_end_date_id" // PropTypes.string.isRequired,
                        onDatesChange={({ startDate, endDate }) =>
                            this.setState({ startDate, endDate })
                        } // PropTypes.func.isRequired,
                        focusedInput={this.state.focusedInput} // PropTypes.oneOf([START_DATE, END_DATE]) or null,
                        onFocusChange={(focusedInput) =>
                            this.setState({ focusedInput })
                        } // PropTypes.func.isRequired,
                    />
                </div>

                <ClassSessions classId={classId} isSaving={this.state.isSavingToRemote}/>

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
    }
}

function createClassId(programId, semesterId, classKey) {
    let classId = programId + "_" + semesterId;
    classId = classKey ? classId + "_" + classKey : classId;
    return classId;
}

function renderOptionsPrograms(programs) {
    return programs.map((program, index) => (
        <option key={index}>{program.programId}</option>
    ));
}

function renderOptionsSemesters(semesters) {
    return semesters.map((semester, index) => (
        <option key={index}>{semester.semesterId}</option>
    ));
}

function renderOptionsLocations(locations) {
    return locations.map((loc, index) => (
        <option key={index}>{loc.locId}</option>
    ));
}

function renderDeleteButton(isEdit, onClickDelete) {
    let deleteButton = <div></div>;
    if (isEdit) {
        deleteButton = (
            <button className="btn-delete" onClick={onClickDelete}>
                Delete
            </button>
        );
    }
    return deleteButton;
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
                text={"class information saved!"}
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
