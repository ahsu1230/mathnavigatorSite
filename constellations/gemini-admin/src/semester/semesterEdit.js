"use strict";
require("./semesterEdit.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";

export class SemesterEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            oldSemesterId: "",
            inputSemesterId: "",
            inputTitle: "",
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
        const semesterId = this.props.semesterId;
        if (semesterId) {
            API.get("api/semesters/semester/" + semesterId).then((res) => {
                const semester = res.data;
                this.setState({
                    oldSemesterId: semester.semesterId,
                    inputSemesterId: semester.semesterId,
                    inputTitle: semester.title,
                    isEdit: true,
                    showDeleteModal: false,
                    showSaveModal: false,
                });
            });
        }
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickSave() {
        let semester = {
            semesterId: this.state.inputSemesterId,
            title: this.state.inputTitle,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save semester: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/semesters/semester/" + this.state.oldSemesterId,
                semester
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/semesters/create", semester)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    }

    onClickCancel() {
        window.location.hash = "semesters";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onConfirmDelete() {
        const semesterId = this.props.semesterId;
        API.delete("api/semesters/semester/" + semesterId).then((res) => {
            window.location.hash = "semesters";
        });
    }

    onSavedOk() {
        this.onDismissModal();
        window.location.hash = "semesters";
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    }

    render() {
        const title = this.state.isEdit ? "Edit Semester" : "Add Semester";
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
                    text={"Semester information saved!"}
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
            <div id="view-semester-edit">
                {modalDiv}
                <h2>{title}</h2>

                <InputText
                    label="Semester ID"
                    required={true}
                    value={this.state.inputSemesterId}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputSemesterId")
                    }
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a semester ID",
                        },
                    ]}
                />

                <InputText
                    label="Title"
                    required={true}
                    value={this.state.inputTitle}
                    onChangeCallback={(e) => this.handleChange(e, "inputTitle")}
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a title",
                        },
                    ]}
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
