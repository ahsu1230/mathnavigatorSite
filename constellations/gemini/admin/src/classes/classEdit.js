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
            isEdit: false,

            // class object
            oldClassId: "",
            inputClassKey: "",
            inputTimeString: "",

            selectProgramId: "",
            selectSemesterId: "",
            selectLocationId: "",
            listPrograms: [],
            listSemesters: [],
            listLocations: [],
            listSessionsLocal: [],
            listSessionsRemote: [],

            // other
            focusedInput: undefined,
        };

        this.handleChange = this.handleChange.bind(this);

        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        this.onModalDeleteConfirm = this.onModalDeleteConfirm.bind(this);
        this.onModalOkSaved = this.onModalOkSaved.bind(this);
        this.onModalDismiss = this.onModalDismiss.bind(this);

        this.onAddSessions = this.onAddSessions.bind(this);
        this.onDeleteSession = this.onDeleteSession.bind(this);
    }

    componentDidMount() {
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
                    let sessions = hasClassId ? responses[4].data : [];
                    sessions = sessions.map((s) => {
                        s.startsAt = moment(s.startsAt);
                        s.endsAt = moment(s.endsAt);
                        return s;
                    });

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
                        startDate: moment(classObj.startDate),
                        endDate: moment(classObj.endDate),

                        listPrograms: programs,
                        listSemesters: semesters,
                        listLocations: locations,
                        listSessionsRemote: sessions,
                        listSessionsLocal: sessions,
                        selectProgramId: selectedProgramId,
                        selectSemesterId: selectedSemesterId,
                        selectLocationId: selectedLocationId,
                    });
                })
            )
            .catch((errors) => {
                console.log("Error: api call failed. " + errors.message);
            });
    }

    onAddSessions(newSessions) {
        let newList = _.concat(this.state.listSessionsLocal, newSessions);
        newList = _.sortBy(newList, ["startsAt"]);
        this.setState({ listSessionsLocal: newList });
    }

    onDeleteSession(sessionId) {
        let sessions = _.filter(this.state.listSessionsLocal, (session) => {
            return session.id != sessionId;
        });
        this.setState({
            listSessionsLocal: sessions,
        });
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickSave() {
        const oldClassId = this.state.oldClassId;
        const newClassId = createClassId(
            this.state.selectProgramId,
            this.state.selectSemesterId,
            this.state.inputClassKey
        );
        let classObj = {
            classId: newClassId,
            programId: this.state.selectProgramId,
            semesterId: this.state.selectSemesterId,
            locationId: this.state.selectLocationId,
            classKey: this.state.inputClassKey,
            times: this.state.inputTimeString,
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

        // Find the sessions to persist and add to apiCalls
        let sessionsToAdd = _.difference(
            this.state.listSessionsLocal,
            this.state.listSessionsRemote
        );
        let sessionsToRemove = _.difference(
            this.state.listSessionsRemote,
            this.state.listSessionsLocal
        );

        _.forEach(sessionsToAdd, (session) => {
            let apiSession = _.cloneDeep(session);
            apiSession.startsAt = apiSession.startsAt.toJSON(); // are moment objects
            apiSession.endsAt = apiSession.endsAt.toJSON(); // are moment objects
            apiSession.id = undefined;
            apiCalls.push(API.post("/api/sessions/create", apiSession));
        });

        _.forEach(sessionsToRemove, (session) => {
            apiCalls.push(API.delete("/api/sessions/session/" + session.id));
        });

        executeApiCalls(apiCalls, successCallback, failCallback);
    }

    onClickCancel() {
        window.location.hash = "classes";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onModalDeleteConfirm() {
        const classId = this.props.classId;
        API.delete("api/classes/class/" + classId).then((res) => {
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
        const listSessionsLocal = this.state.listSessionsLocal;
        const startDateString =
            listSessionsLocal.length > 0
                ? listSessionsLocal[0].startsAt.format(
                      "dddd, MMMM Do YYYY, h:mm a"
                  )
                : "Not scheduled yet. Please add new sessions.";
        const endDateString =
            listSessionsLocal.length > 0
                ? listSessionsLocal[listSessionsLocal.length - 1].endsAt.format(
                      "dddd, MMMM Do YYYY, h:mm a"
                  )
                : "Not scheduled yet. Please add new sessions.";

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
                        placeholder="Optional"
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
                    <p>
                        <b>Start:</b> {startDateString}
                    </p>
                    <p>
                        <b>End:</b> {endDateString}
                    </p>
                </div>

                <ClassSessions
                    classId={classId}
                    sessions={this.state.listSessionsLocal}
                    onAddSessions={this.onAddSessions}
                    onDeleteSession={this.onDeleteSession}
                />

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
        <option key={index}>{loc.locationId}</option>
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

function executeApiCalls(apiCalls, successCallback, failCallback) {
    console.log("Reducing " + apiCalls.length);

    let fnResolveTask = function (nextApi) {
        return new Promise((resolve, reject) => {
            nextApi
                .then((resp) => {
                    console.log("Success: " + moment().format("hh:mm:ss"));
                    resolve(resp.data);
                })
                .catch((res) => {
                    console.log("Failure: " + moment().format("hh:mm:ss"));
                    reject(res.response.data);
                });
        });
    };

    let sequence = apiCalls.reduce((accumulatorPromise, nextApi) => {
        console.log(`Loop! ${moment().format("hh:mm:ss")}`);
        return accumulatorPromise.then(() => {
            return fnResolveTask(nextApi);
        });
    }, Promise.resolve());
    sequence
        .then((results) => {
            console.log("All success!");
            successCallback(results);
        })
        .catch((results) => {
            console.log("One error?");
            failCallback(results);
        });
}
