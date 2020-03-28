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
             oldLocId: "",
             inputLocId: "",
             inputStreet: "",
             inputCity: "",
             inputState: "",
             inputZip: "",
             inputRoom: "",
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
              oldLocId: location.locId,
              inputLocId: location.locId,
              inputStreet: location.street,
              inputCity: location.city,
              inputState: location.state,
              inputZip: location.zipcode,
              inputRoom: location.room,
              isEdit: true
            });
          });
      }
    }

    handleChange(event,value) {
        this.setState({[value]: event.target.value});
    }

    onClickCancel() {
        window.location.hash = "locations";
    }

    onClickDelete() {
        const locId = this.props.locId;
        API.delete("api/locations/v1/location/" + locId)
            .then(res => {
                window.location.hash = "locations";
            })
    }

    onClickSave() {
        let location = {
            locId: this.state.inputLocId,
            street: this.state.inputStreet,
            city: this.state.inputCity,
            state: this.state.inputState,
            zipcode: this.state.inputZip,
            room: this.state.inputRoom
        };

        let successCallback = () => this.onSaved();
        let failCallback = (err) => alert("Could not save location: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/locations/v1/location/" + this.state.oldLocId, location)
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
        alert("Location has been saved!");
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
                <input value={this.state.inputLocId}
                    onChange={(e) => this.handleChange(e, "inputLocId")}/>
                <h4>Street</h4>
                <input value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputStreet")}/>
                <h4>City</h4>
                <input value={this.state.inputCity}
                    onChange={(e) => this.handleChange(e, "inputCity")}/>
                <h4>State</h4>
                <input value={this.state.inputState}
                    onChange={(e) => this.handleChange(e, "inputState")}/>
                <h4>Zipcode</h4>
                <input value={this.state.inputZip}
                    onChange={(e) => this.handleChange(e, "inputZip")}/>
                <h4>Room</h4>
                <input value={this.state.inputRoom}
                    onChange={(e) => this.handleChange(e, "inputRoom")}/>
                <div className="buttons">
                    <button className="btn-save" onClick={this.onClickSave}>Save</button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
                    {deleteButton}
                </div>
            </div>
        );
    }
}
