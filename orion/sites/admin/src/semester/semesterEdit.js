"use strict";
require("./semesterEdit.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class SemesterEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isEdit: false,
            inputSemesterId: "",
            inputTitle: ""
        };
    }

    onClickSave() {
        console.log("save button clicked");
    }

    onClickCancel() {
        console.log("cancel button clicked");
    }

    handleChange() {
        console.log("handle change");
    }

    render() {
        const title = "New Semester";
        let deleteButton = <div></div>;
        return (
            <div id="view-semester-edit">
                <h2>Add Semester</h2>

                <h4>Semester ID</h4>
                <input
                    value={this.state.inputDate}
                    onChange={e => this.handleChange(e, "inputSemesterId")}
                />

                <h4>Title</h4>
                <input
                    value={this.state.inputAuthor}
                    onChange={e => this.handleChange(e, "inputTitle")}
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
