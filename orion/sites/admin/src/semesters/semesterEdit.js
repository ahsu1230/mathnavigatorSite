"use strict";
require("./semesterEdit.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";

export class SemesterEditPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isEdit: false,
      inputSemesterId: "",
      inputTitle: "",
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

        <h4>SemesterID</h4>
        <input
          value={this.state.inputDate}
          onChange={(e) => this.handleChange(e, "inputDate")}
        />

        <h4>Title</h4>
        <input
          value={this.state.inputAuthor}
          onChange={(e) => this.handleChange(e, "inputAuthor")}
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
