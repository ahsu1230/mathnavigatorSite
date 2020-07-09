"use strict";
require("./afhEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { AFHEditDateTime } from "./afhEditDateTime.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AskForHelpEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEdit: false,
            id: 0,
            date: null,
            locationId: "",
            notes: "",
            subject: "",
            locations: [],
        };

        this.handleChange = this.handleChange.bind(this);

        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        this.onConfirmDelete = this.onConfirmDelete.bind(this);
        this.onSavedOk = this.onSavedOk.bind(this);
        this.onDismissModal = this.onDismissModal.bind(this);

        this.onMomentChange = this.onMomentChange.bind(this);
    }

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
                    console.log(afh);
                    this.setState({
                        isEdit: true,

                        id: afh.id,
                        date: moment(afh.date),
                        locationId: afh.locationId,
                        notes: afh.notes,
                        subject: afh.subject,
                        timeString: afh.timeString,
                        title: afh.title,
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

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
        console.log(this.state);
    }

    onClickSave() {
        let afh = {
            id: this.state.id,
            date: this.state.date.toJSON(),
            locationId: this.state.locationId,
            notes: this.state.notes,
            subject: this.state.subject,
            timeString: this.state.timeString,
            title: this.state.title,
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
    }

    onClickCancel() {
        window.location.hash = "afh";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onConfirmDelete() {
        const afhId = this.state.id;
        API.delete("api/askforhelp/afh/" + afhId).then((res) => {
            window.location.hash = "afh";
        });
    }

    onSavedOk() {
        this.onDismissModal();
        window.location.hash = "afh";
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    }

    onMomentChange(newMoment) {
        this.setState({ inputPostedAt: newMoment, date: newMoment });
    }

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

                <h3>Time</h3>
                <input
                    value={this.state.timeString}
                    onChange={(e) => this.handleChange(e, "timeString")}
                />

                <h3>Title</h3>
                <input
                    value={this.state.title}
                    onChange={(e) => this.handleChange(e, "title")}
                />

                <h3>Subject</h3>
                <select
                    value={this.state.subject}
                    onChange={(e) => this.handleChange(e, "subject")}>
                    <option>Math</option>
                    <option>English</option>
                    <option>Computer Programming</option>
                </select>

                <h3>Location ID</h3>
                <select
                    value={this.state.locationId}
                    onChange={(e) => this.handleChange(e, "locationId")}>
                    {optLocations}
                </select>

                <h3>Notes</h3>
                <textarea
                    value={this.state.notes}
                    onChange={(e) => this.handleChange(e, "notes")}
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
