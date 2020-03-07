'use strict';
require('./achieveEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class AchieveEditPage extends React.Component {
    constructor(props) {
         super(props);
         this.state = {
             inputYear: "",
             inputMessage: "",
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
        const title = "Edit Achievement";
        let deleteButton = <div></div>;
        return (
            <div id="view-achieve-edit">
            <h2>{title}</h2>
            <h4>Year</h4>
            <input value={this.state.inputYear}
                onChange={(e) => this.handleChange(e, "inputYear")}/>
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
