"use strict";
require("./programEdit.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class ProgramEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEdit: false,
            showDeleteModal: false,
            showSaveModal: false,
            oldProgramId: "",
            inputProgramId: "",
            inputProgramName: "",
            inputGrade1: 0,
            inputGrade2: 0,
            inputDescription: ""
        };

        // input onChange
        this.handleChange = this.handleChange.bind(this);

        // click on button
        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        // after action
        this.onDeleted = this.onDeleted.bind(this);
        this.onSaved = this.onSaved.bind(this);
        this.onDismissModal = this.onDismissModal.bind(this);
    }

    componentDidMount() {
        const programId = this.props.programId;
        if (programId) {
            API.get("api/programs/v1/program/" + programId).then(res => {
                const program = res.data;
                this.setState({
                    oldProgramId: program.programId,
                    inputProgramId: program.programId,
                    inputProgramName: program.name,
                    inputGrade1: program.grade1,
                    inputGrade2: program.grade2,
                    inputDescription: program.description,
                    isEdit: true
                });
            });
        }
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickCancel() {
        window.location.hash = "programs";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onClickSave() {
        let program = {
            programId: this.state.inputProgramId,
            name: this.state.inputProgramName,
            grade1: parseInt(this.state.inputGrade1),
            grade2: parseInt(this.state.inputGrade2),
            description: this.state.inputDescription
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = err =>
            alert("Could not save program: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/programs/v1/program/" + this.state.oldProgramId,
                program
            )
                .then(res => successCallback())
                .catch(err => failCallback(err));
        } else {
            API.post("api/programs/v1/create", program)
                .then(res => successCallback())
                .catch(err => failCallback(err));
        }
    }

    onDeleted() {
        const programId = this.props.programId;
        API.delete("api/programs/v1/program/" + programId)
            .then(res => {
                window.location.hash = "programs";
            })
            .finally(() => this.onDismissModal());
    }

    onSaved() {
        this.onDismissModal();
        window.location.hash = "programs";
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false
        });
    }

    render() {
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
                <h4>Program Id</h4>
                <input
                    value={this.state.inputProgramId}
                    onChange={e => this.handleChange(e, "inputProgramId")}
                />
                <h4>Program Name</h4>
                <input
                    value={this.state.inputProgramName}
                    onChange={e => this.handleChange(e, "inputProgramName")}
                />
                <h4>Grade1</h4>
                <input
                    value={this.state.inputGrade1}
                    onChange={e => this.handleChange(e, "inputGrade1")}
                />
                <h4>Grade2</h4>
                <input
                    value={this.state.inputGrade2}
                    onChange={e => this.handleChange(e, "inputGrade2")}
                />
                <h4>Description</h4>
                <textarea
                    value={this.state.inputDescription}
                    onChange={e => this.handleChange(e, "inputDescription")}
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
