"use strict";
require("./afhEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { AFHEditDateTime } from "./afhEditDateTime.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";

export class AskForHelpEditPage extends React.Component {
    state = {
        isEdit: false,
        id: 0,
        date: null,
        timeString: "",
        subject: "",
        title: "",
        locationId: "",
        notes: "",
        locations: [],
    };

    componentDidMount() {
        const afhId = this.props.afhId;

        API.get("api/locations/all").then((res) => {
            let locations = res.data;
            this.setState({
                locations: res.data,
            });
            if (afhId) {
                API.get("api/askforhelp/afh/" + afhId).then((res2) => {
                    const afh = res2.data;
                    this.setState({
                        isEdit: true,

                        id: afh.id,
                        date: moment(afh.date),
                        timeString: afh.timeString,
                        subject: afh.subject,
                        title: afh.title,
                        locationId: afh.locationId,
                        notes: afh.notes,
                    });
                });
            } else {
                this.setState({
                    locationId: locations[0].locationId,
                    subject: "Math",
                    date: moment(),
                });
            }
        });
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickSave = () => {
        let afh = {
            id: this.state.id,
            date: this.state.date.toJSON(),
            timeString: this.state.timeString,
            subject: this.state.subject,
            title: this.state.title,
            locationId: this.state.locationId,
            notes: this.state.notes,
        };
        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save Ask For Help session: " + err.response.data);

        if (this.state.isEdit) {
            API.post("api/askforhelp/afh/" + this.state.id, afh)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/askforhelp/create", afh)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    };

    onClickCancel = () => {
        window.location.hash = "afh";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onConfirmDelete = () => {
        const afhId = this.state.id;
        API.delete("api/askforhelp/afh/" + afhId).then((res) => {
            window.location.hash = "afh";
        });
    };

    onSavedOk = () => {
        this.onDismissModal();
        window.location.hash = "afh";
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    onMomentChange = (newMoment) => {
        this.setState({ inputPostedAt: newMoment, date: newMoment });
    };

    render() {
        const title =
            (this.state.isEdit ? "Edit" : "Add") + " Ask For Help Session";

        const optLocations = this.state.locations.map((loc, index) => (
            <option key={index}>{loc.locationId}</option>
        ));

        let deleteButton = <div></div>;
        if (this.state.isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        let modalDiv;
        let modalContent;
        let showModal;
        if (this.state.showDeleteModal) {
            showModal = this.state.showDeleteModal;
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={this.onConfirmDelete}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalContent = (
                <OkayModal
                    text={"Ask for help information saved!"}
                    onOkay={this.onSavedOk}
                />
            );
        }
        if (modalContent) {
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showModal}
                    onDismiss={this.onDismissModal}
                />
            );
        }

        return (
            <div id="view-afh-edit">
                {modalDiv}
                <h2>{title}</h2>

                <AFHEditDateTime
                    postedAt={this.state.date}
                    onMomentChange={this.onMomentChange}
                />

                <InputText
                    label="Time"
                    value={this.state.timeString}
                    onChangeCallback={(e) => this.handleChange(e, "timeString")}
                    required={true}
                    description="Enter a time (e.g. 2:00 - 4:00 PM)"
                    validators={[
                        {
                            validate: (name) => name != "",
                            message: "You must input a time",
                        },
                    ]}
                />

                <InputText
                    label="Title"
                    value={this.state.title}
                    onChangeCallback={(e) => this.handleChange(e, "title")}
                    required={true}
                    description="Enter a title (i.e. AP Calculus Practice Exam, Computer Programming Office Hours, English Essay Review Session)"
                    validators={[
                        {
                            validate: (name) => name != "",
                            message: "You must input a title",
                        },
                    ]}
                />

                <h3>Subject</h3>
                <select
                    value={this.state.subject}
                    onChange={(e) => this.handleChange(e, "subject")}>
                    <option>math</option>
                    <option>english</option>
                    <option>computer</option>
                </select>

                <h3>Location ID</h3>
                <select
                    value={this.state.locationId}
                    onChange={(e) => this.handleChange(e, "locationId")}>
                    {optLocations}
                </select>

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                    description="Add any notes"
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
