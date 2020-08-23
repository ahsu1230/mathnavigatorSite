"use strict";
require("./programEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText, emptyValidator } from "../utils/inputText.js";

export class ProgramEditPage extends React.Component {
    state = {
        isEdit: false,
        oldProgramId: "",
        programId: "",
        name: "",
        grade1: 0,
        grade2: 0,
        description: "",
    };

    componentDidMount = () => {
        const programId = this.props.programId;
        if (programId) {
            API.get("api/programs/program/" + programId).then((res) => {
                const program = res.data;
                this.setState({
                    oldProgramId: program.programId,
                    programId: program.programId,
                    name: program.name,
                    grade1: program.grade1,
                    grade2: program.grade2,
                    description: program.description,
                    isEdit: true,
                });
            });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = "programs";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onClickSave = () => {
        let program = {
            programId: this.state.programId,
            name: this.state.name,
            grade1: parseInt(this.state.grade1),
            grade2: parseInt(this.state.grade2),
            description: this.state.description,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save program: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/programs/program/" + this.state.oldProgramId, program)
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/programs/create", program)
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        }
    };

    onDeleted = () => {
        const programId = this.props.programId;
        API.delete("api/programs/program/" + programId)
            .then(() => {
                window.location.hash = "programs";
            })
            .catch((err) => {
                alert("Could not delete program: " + err.response.data);
            })
            .finally(() => this.onDismissModal());
    };

    onSaved = () => {
        this.onDismissModal();
        window.location.hash = "programs";
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    render = () => {
        const isEdit = this.state.isEdit;
        const program = this.state.program;
        const title = isEdit ? "Edit Program" : "Add Program";

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
                    text={"Program information saved!"}
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
            <div id="view-program-edit">
                {modalDiv}
                <h2>{title}</h2>

                <InputText
                    label="Program ID"
                    value={this.state.programId}
                    onChangeCallback={(e) => this.handleChange(e, "programId")}
                    required={true}
                    description="Enter the program ID. Examples: ap_calculus, sat1, ap_java"
                    validators={[emptyValidator("program ID")]}
                />

                <InputText
                    label="Program Name"
                    value={this.state.name}
                    onChangeCallback={(e) => this.handleChange(e, "name")}
                    required={true}
                    description="Enter the program name. This name will be present to users. Example: AP Calculus, SAT2 Subject Math"
                    validators={[emptyValidator("program name")]}
                />

                <InputText
                    label="Grade1"
                    value={this.state.grade1}
                    onChangeCallback={(e) => this.handleChange(e, "grade1")}
                    required={true}
                    description="Enter the lower grade"
                    validators={[
                        emptyValidator("grade"),
                        {
                            validate: (grade1) =>
                                parseInt(grade1) >= 1 && parseInt(grade1) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade1) =>
                                this.state.grade2 >= parseInt(grade1),
                            message:
                                "Grade1 must be less than or equal to Grade2",
                        },
                    ]}
                />

                <InputText
                    label="Grade2"
                    value={this.state.grade2}
                    onChangeCallback={(e) => this.handleChange(e, "grade2")}
                    required={true}
                    description="Enter the higher grade"
                    validators={[
                        emptyValidator("grade"),
                        {
                            validate: (grade2) =>
                                parseInt(grade2) >= 1 && parseInt(grade2) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade2) =>
                                this.state.grade1 <= parseInt(grade2),
                            message:
                                "Grade2 must be greater than or equal to Grade1",
                        },
                    ]}
                />

                <InputText
                    label="Description"
                    isTextBox={true}
                    value={this.state.description}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "description")
                    }
                    required={true}
                    description="Enter the description"
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a description",
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
    };
}
