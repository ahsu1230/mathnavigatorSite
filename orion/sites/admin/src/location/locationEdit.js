'use strict';
require('./locationEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class LocationEditPage extends React.Component {
    constructor(props) {
         super(props);
         this.state = {
             inputLocationId: "",
             inputAddress: "",
             inputRoomNum: 0,
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
        const title = "Edit Location";
        let deleteButton = <div></div>;
        return (
            <div id="view-location-edit">
            <h2>{title}</h2>
            <h4>Location ID</h4>
            <input value={this.state.inputLocationId}
                onChange={(e) => this.handleChange(e, "inputLocationId")}/>
            <h4>Address</h4>
            <input value={this.state.inputAddress}
                onChange={(e) => this.handleChange(e, "inputAddress")}/>
            <h4>Room Number</h4>
            <input value={this.state.inputRoomNum}
                onChange={(e) => this.handleChange(e, "inputRoomNum")}/>
            <div className="buttons">
                <button className="btn-save" onClick={this.onClickSave}>Save</button>
                <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
                {deleteButton}
            </div>
            </div>
        );
    }
}
