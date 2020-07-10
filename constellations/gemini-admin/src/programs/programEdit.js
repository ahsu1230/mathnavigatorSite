"use strict";
require("./programEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { TextInput } from "../utils/textInput.js";

export class ProgramEditPage extends React.Component {
    state = {
        isEdit: false,
        showDeleteModal: false,
        showSaveModal: false,
        oldProgramId: "",
        inputProgramId: "",
        inputProgramName: "",
        inputGrade1: 0,
        inputGrade2: 0,
        inputDescription: "",
    };

    componentDidMount = () => {
        const programId = this.props.programId;
        if (programId) {
            API.get("api/programs/program/" + programId).then((res) => {
                const program = res.data;
                this.setState({
                    oldProgramId: program.programId,
                    inputProgramId: program.programId,
                    inputProgramName: program.name,
                    inputGrade1: program.grade1,
                    inputGrade2: program.grade2,
                    inputDescription: program.description,
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
            programId: this.state.inputProgramId,
            name: this.state.inputProgramName,
            grade1: parseInt(this.state.inputGrade1),
            grade2: parseInt(this.state.inputGrade2),
            description: this.state.inputDescription,
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

                <TextInput
                    label="Program Id"
                    value={this.state.inputProgramId}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputProgramId")
                    }
                    required={true}
                    description="Enter the program ID. Examples: ap_calculus, sat1, ap_java"
                    validators={[
                        {
                            validate: (programId) => programId != "",
                            message: "You must input a programId",
                        },
                    ]}
                />

                <TextInput
                    label="Program Name"
                    value={this.state.inputProgramName}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputProgramName")
                    }
                    required={true}
                    description="Enter the program name. This name will be present to users. Example: AP Calculus, SAT2 Subject Math"
                    validators={[
                        {
                            validate: (name) => name != "",
                            message: "You must input a name",
                        },
                    ]}
                />

                <TextInput
                    label="Grade1"
                    value={this.state.inputGrade1}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputGrade1")
                    }
                    required={true}
                    description="Enter the lower grade"
                    validators={[
                        {
                            validate: (grade1) => grade1 != "",
                            message: "You must input a grade",
                        },
                        {
                            validate: (grade1) =>
                                parseInt(grade1) >= 1 && parseInt(grade1) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade1) =>
                                this.state.inputGrade2 >= parseInt(grade1),
                            message:
                                "Grade1 must be less than or equal to Grade2",
                        },
                    ]}
                />

                <TextInput
                    label="Grade2"
                    value={this.state.inputGrade2}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputGrade2")
                    }
                    required={true}
                    description="Enter the higher grade"
                    validators={[
                        {
                            validate: (grade2) => grade2 != "",
                            message: "You must input a grade",
                        },
                        {
                            validate: (grade2) =>
                                parseInt(grade2) >= 1 && parseInt(grade2) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade2) =>
                                this.state.inputGrade1 <= parseInt(grade2),
                            message:
                                "Grade2 must be greater than or equal to Grade1",
                        },
                    ]}
                />

                <TextInput
                    label="Description"
                    isTextBox={true}
                    value={this.state.inputDescription}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputDescription")
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
