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
        window.location.hash = "locations";
    }

    onClickDelete() {
        console.log("delete button clicked");
    }

    handleChange() {
        console.log("handle change");
    }

    render() {
        const title = "Edit Location";
        let deleteButton = <div></div>;
        if (this.state.isEdit) {
          deleteButton = (
            <button className="btn-delete" onClick={this.onClickDelete}>Delete</button>
          );
        }
        return (
            <div id="view-location-edit">
                <h2>{title}</h2>
                <h4>Location ID</h4>
                <input value={this.state.inputLocationId}
                    onChange={(e) => this.handleChange(e, "inputLocationId")}/>
                <h4>Street</h4>
                <input value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputStreet")}/>
                <h4>City</h4>
                <input value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputCity")}/>
                <h4>State</h4>
                <input value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputState")}/>
                <h4>Zipcode</h4>
                <input value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputZip")}/>
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
