"use strict";
require("./announceEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { AnnounceEditDateTime } from "./announceEditDateTime.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AnnounceEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            announceId: 0,
            datePickerFocused: false,
            inputPostedAt: null,
            inputAuthor: "",
            inputMessage: "",
            isEdit: false,
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
        const announceId = this.props.announceId;
        if (announceId) {
            API.get("api/announcements/announcement/" + announceId).then(
                (res) => {
                    const announce = res.data;
                    this.setState({
                        announceId: announce.id,
                        inputPostedAt: moment(announce.postedAt),
                        inputAuthor: announce.author,
                        inputMessage: announce.message,
                        isEdit: true,
                        showDeleteModal: false,
                        showSaveModal: false,
                    });
                }
            );
        }
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickSave() {
        let announcement = {
            postedAt: this.state.inputPostedAt.toJSON(),
            author: this.state.inputAuthor,
            message: this.state.inputMessage,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save announcement: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/announcements/announcement/" + this.state.announceId,
                announcement
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/announcements/create", announcement)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    }

    onClickCancel() {
        window.location.hash = "announcements";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onConfirmDelete() {
        const announceId = this.props.announceId;
        API.delete("api/announcements/announcement/" + announceId).then(
            (res) => {
                window.location.hash = "announcements";
            }
        );
    }

    onSavedOk() {
        this.onDismissModal();
        window.location.hash = "announcements";
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    }

    onMomentChange(newMoment) {
        this.setState({ inputPostedAt: newMoment });
    }

    render() {
        const title = this.state.isEdit
            ? "Edit Announcement"
            : "Add Announcement";
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
                    text={"Announcement information saved!"}
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
            <div id="view-announce-edit">
                {modalDiv}
                <h2>{title}</h2>

                <AnnounceEditDateTime
                    postedAt={this.state.inputPostedAt}
                    onMomentChange={this.onMomentChange}
                />

                <h4>Author</h4>
                <input
                    value={this.state.inputAuthor}
                    onChange={(e) => this.handleChange(e, "inputAuthor")}
                />

                <h4>Message</h4>
                <textarea
                    value={this.state.inputMessage}
                    onChange={(e) => this.handleChange(e, "inputMessage")}
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
