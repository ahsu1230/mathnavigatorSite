'use strict';
require('./announceEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class AnnounceEditPage extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        isEdit: false,
        inputDate: "",
        inputAuthor: "",
        inputMessage: ""
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
        const title = "New Announcement";
        let deleteButton = <div></div>;
        return (
            <div id ="view-announce-edit">
            <h2>Add Announcement</h2>

            <h4>Date</h4> {/* TODO: Make this "Posted Date" */}
            <input value={this.state.inputDate}
                    onChange={(e) => this.handleChange(e, "inputDate")}/>

            <h4>Author</h4>
            <input value={this.state.inputAuthor}
                    onChange={(e) => this.handleChange(e, "inputAuthor")}/>

            <h4>Message</h4>
            <input value={this.state.inputMessage}
                    onChange={(e) => this.handleChange(e, "inputMessage")}/>

            <div className="buttons">
                <button className="btn-save" onClick={this.onClickSave}>Save</button>
                <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
                {deleteButton}
            </div>
      </div>
    );
  }
}
