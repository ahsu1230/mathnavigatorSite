'use strict';
require('./classEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class ClassEditPage extends React.Component {
    constructor(props) {
         super(props);
         this.state = {
             inputLocationId: "",
             inputStreet: "",
             inputCity: "",
             inputState: "",
             inputZip: "",
             inputRoomNum: "",
             isEdit: false,
          };
    }

    onClickSave() {
        console.log("save button clicked");
    }

    onClickCancel() {
        window.location.hash = "classes";
    }

    onClickDelete() {
        console.log("delete button clicked");
    }

    handleChange() {
        console.log("handle change");
    }

    render() {
        const title = this.state.isEdit ? "Edit Class" : "Add Class";
        let deleteButton = <div></div>;
        if (this.state.isEdit) {
          deleteButton = (
            <button className="btn-delete" onClick={this.onClickDelete}>Delete</button>
          );
        }
        return (
            <div id="view-class-edit">
                <h2>{title}</h2>
                <h4>ClassId</h4>
                <input value={this.state.inputLocationId}
                    onChange={(e) => this.handleChange(e, "inputClassId")}/>
                <h4>Program Key</h4>
                <select id = "dropdown">
                    <option value = "calculus">Calculus</option>
                    <option value = "diffeq">DiffEq</option>
                </select>
                <h4>Class Key</h4>
                    <input value= "program id"/>

                <h4>SemesterId</h4>
                <select id>
                    <option value = "calculus">Calculus</option>
                    <option value = "diffeq">DiffEq</option>
                </select>
                <h4>LocationId</h4>
                <select>
                    <option value = "calculus">Calculus</option>
                    <option value = "diffeq">DiffEq</option>
                </select>
                <div className="buttons">
                    <button className="btn-save" onClick={this.onClickSave}>Save</button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
                    {deleteButton}
                </div>
            </div>
        );
    }
}
