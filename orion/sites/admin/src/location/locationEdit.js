'use strict';
require('./locationEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import API from '../api.js';

export class LocationEditPage extends React.Component {
    constructor(props) {
         super(props);
         this.state = {
             inputlocId: "",
             inputStreet: "",
             inputCity: "",
             inputState: "",
             inputZip: "",
             inputRoomNum: "",
             isEdit: false,
          };
      this.handleChange = this.handleChange.bind(this);

      this.onClickCancel = this.onClickCancel.bind(this);
      this.onClickDelete = this.onClickDelete.bind(this);
      this.onClickSave = this.onClickSave.bind(this);

      this.onDeleted = this.onDeleted.bind(this);
      this.onSaved = this.onSaved.bind(this);
    }



    componentDidMount() {
      const locId = this.props.locId;
      if (locId) {
        API.get("api/locations/v1/location/" + locId)
          .then(res => {
            const location = res.data;
            this.setState({
              oldlocId: location.locId,
              inputlocId: location.locId,
              inputStreet: location.street,
              inputState: location.state1,
              inputZip: location.zip,
              inputRoomNum: location.roomNum,
              isEdit: true
            });
          });
      }
    }

    handleChange(event,inputField) {
        this.setState({[inputField]: event.target.inputField});
    }

    onClickCancel() {
        window.location.hash = "locations";
    }

    onClickDelete() {
        console.log("delete button clicked");
    }

    onClickSave() {
        let location = {
            locId: this.state.inputlocId,
            street: this.state.inputStreet,
            state1: this.state.state1,
            zip: this.state.zip,
            roomNum: this.state.roomNum
        };

        let successCallback = () => this.onSaved();
        let failCallback = (err) => alert("Could not save location: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/locations/v1/location/" + this.state.oldlocId, location)
                .then(res => successCallback())
                .catch(err => failCallback(err));
        } else {
            API.post("api/locations/v1/create", location)
            .then(res => successCallback())
            .catch(err => failCallback(err));
        }
    }

    onDeleted() {
        const locId = this.props.locId;
        API.delete("api/locations/v1/location/" + locId)
            .then(res => {
                window.location.hash = "locations";
            })
    }

    onSaved() {
        window.location.hash = "locations";
        alert(this.state.isEdit);
    }

    render() {
        const isEdit = this.state.isEdit;
        const location = this.state.location;
        const title = this.state.isEdit ? "Edit Location" : "Add Location";
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
                <input inputField={this.state.inputlocId}
                    onChange={(e) => this.handleChange(e, "inputlocId")}/>
                <h4>Street</h4>
                <input inputField={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputStreet")}/>
                <h4>City</h4>
                <input inputField={this.state.inputCity}
                    onChange={(e) => this.handleChange(e, "inputCity")}/>
                <h4>State</h4>
                <input inputField={this.state.inputState}
                    onChange={(e) => this.handleChange(e, "inputState")}/>
                <h4>Zipcode</h4>
                <input inputField={this.state.inputZip}
                    onChange={(e) => this.handleChange(e, "inputZip")}/>
                <h4>Room Number</h4>
                <input inputField={this.state.inputRoomNum}
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
