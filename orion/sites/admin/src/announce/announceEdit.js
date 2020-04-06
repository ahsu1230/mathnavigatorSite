"use strict";
require("./announceEdit.styl");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AnnounceEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            announceId: 0,
            inputPostedAt: new Date(Date.now()),
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
    }

    componentDidMount() {
        const announceId = this.props.announceId;
        if (announceId) {
            API.get("api/announcements/v1/announcement/" + announceId).then(
                (res) => {
                    const announce = res.data;
                    this.setState({
                        announceId: announce.Id,
                        inputPostedAt: new Date(announce.postedAt),
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
                "api/announcements/v1/announcement/" + this.state.announceId,
                announcement
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/announcements/v1/create", announcement)
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
        API.delete("api/announcements/v1/announcement/" + announceId).then(
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
                <h2>Add Announcement</h2>

                <h4>Post Date</h4>
                <p>{this.state.inputPostedAt.toLocaleString()}</p>

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
