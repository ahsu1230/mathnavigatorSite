"use strict";
require("./achieveEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";
import { emptyValidator } from "../utils/inputText.js";

export class AchieveEditPage extends React.Component {
    state = {
        isEdit: false,
        year: "",
        position: "",
        message: "",
    };

    componentDidMount = () => {
        console.log(this.props.id);
        const id = this.props.id;
        if (id) {
            API.get("api/achievements/achievement/" + id).then((res) => {
                const achieve = res.data;
                this.setState({
                    isEdit: true,
                    year: achieve.year,
                    position: achieve.position,
                    message: achieve.message,
                });
            });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = "achievements";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onClickSave = () => {
        let achieve = {
            year: parseInt(this.state.year),
            position: parseInt(this.state.position),
            message: this.state.message,
        };
        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save achievement: " + err.response.data);

        if (this.state.isEdit) {
            API.post("api/achievements/achievement/" + this.props.id, achieve)
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/achievements/create", achieve)
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        }
    };

    onSaved = () => {
        this.onDismissModal();
        window.location.hash = "achievements";
    };

    onDeleted = () => {
        API.delete("api/achievements/achievement/" + this.props.id)
            .then(() => (window.location.hash = "achievements"))
            .finally(() => this.onDismissModal());
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    renderModal = () => {
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
        return modalDiv;
    };

    render = () => {
        const isEdit = this.state.isEdit;
        const title = isEdit ? "Edit Achievement" : "Add Achievement";

        let deleteButton = isEdit ? (
            <button className="btn-delete" onClick={this.onClickDelete}>
                Delete
            </button>
        ) : (
            <div></div>
        );

        return (
            <div id="view-achieve-edit">
                {this.renderModal()}
                <h2>{title}</h2>

                <InputText
                    label="Year"
                    description="Enter the achievement year"
                    required={true}
                    value={this.state.year}
                    onChangeCallback={(e) => this.handleChange(e, "year")}
                    validators={[
                        {
                            validate: (number) => parseInt(number) > 2000,
                            message: "You must input a year greater than 2000",
                        },
                    ]}
                />

                <InputText
                    label="Position"
                    description="Enter the position (Lower position numbers are shown first in that year)"
                    required={true}
                    value={this.state.position}
                    onChangeCallback={(e) => this.handleChange(e, "position")}
                    validators={[
                        {
                            validate: (number) =>
                                Number.isInteger(parseInt(number)) &&
                                parseInt(number) > 0,
                            message: "You must input a positive integer",
                        },
                    ]}
                />

                <InputText
                    label="Message"
                    description="Enter the achievement message"
                    required={true}
                    isTextBox={true}
                    value={this.state.message}
                    onChangeCallback={(e) => this.handleChange(e, "message")}
                    validators={[emptyValidator("message")]}
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
    };
}
