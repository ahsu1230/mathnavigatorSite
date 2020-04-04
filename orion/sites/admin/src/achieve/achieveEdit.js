"use strict";
require("./achieveEdit.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AchieveEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEdit: false,
            inputYear: 0,
            inputMessage: ""
        };

        this.handleChange = this.handleChange.bind(this);

        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        this.onDeleted = this.onDeleted.bind(this);
        this.onSaved = this.onSaved.bind(this);
        this.onDismissModal = this.onDismissModal.bind(this);
    }

    componentDidMount() {
        const Id = this.props.Id;
        if (Id) {
            API.get("api/achievements/v1/achievement/" + Id).then(res => {
                const achieve = res.data;
                this.setState({
                    inputYear: achieve.year,
                    inputMessage: achieve.message,
                    isEdit: true
                });
            });
        }
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickCancel() {
        window.location.hash = "achievements";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onClickSave() {
        let achieve = {
            year: parseInt(this.state.inputYear),
            message: this.state.inputMessage
        };
        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = err =>
            alert("Could not save achievement: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/achievements/v1/achievement/" + this.props.Id,
                achieve
            )
                .then(res => successCallback())
                .catch(err => failCallback(err));
        } else {
            API.post("api/achievements/v1/create", achieve)
                .then(res => successCallback())
                .catch(err => failCallback(err));
        }
    }

    onSaved() {
        this.onDismissModal();
        window.location.hash = "achievements";
    }

    onDeleted() {
        const Id = this.props.Id;
        API.delete("api/achievements/v1/achievement/" + Id)
            .then(res => {
                window.location.hash = "achievements";
            })
            .finally(() => this.onDismissModal());
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false
        });
    }

    render() {
        const isEdit = this.state.isEdit;
        const achieve = this.state.achieve;
        const title = isEdit ? "Edit Achievement" : "Add Achievement";

        let deleteButton = <div></div>;
        if (isEdit) {
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
                    onAccept={this.onDeleted}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalContent = (
                <OkayModal
                    text={"Achievement information saved!"}
                    onOkay={this.onSaved}
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
            <div id="view-achieve-edit">
                {modalDiv}
                <h2>{title}</h2>
                <h4>Year</h4>
                <input
                    value={this.state.inputYear}
                    onChange={e => this.handleChange(e, "inputYear")}
                />
                <h4>Message</h4>
                <input
                    value={this.state.inputMessage}
                    onChange={e => this.handleChange(e, "inputMessage")}
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
