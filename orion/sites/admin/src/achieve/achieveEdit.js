'use strict';
require('./achieveEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import API from '../api.js';

export class AchieveEditPage extends React.Component {
    constructor(props) {
         super(props);
         this.state = {
             isEdit: false,
             inputYear: 0,
             inputMessage: ""
          };

    this.handleChange = this.handleChange.bind(this);

    this.onClickCancel = this.onClickCancel.bind(this);
    this.onClickDelete = this.onClickDelete.bind(this);
    this.onClickSave = this.onClickSave.bind(this);

    this.onDeleted = this.onDeleted.bind(this);
    this.onSaved = this.onSaved.bind(this);

    }

    componentDidMount() {
      const achieveId = this.props.achieveId;
      if (achieveId) {
        API.get("api/achievements/v1/achievement/" + achieveId)
          .then(res => {
            debugger;
            const achieve = res.data;
            this.setState({
              inputYear: achieve.year,
              inputMessage: achieve.message,
              isEdit: true
            });
          });
      }
    }

    handleChange(event, value) {
      this.setState({[value]: event.target.value});
    }

    onClickCancel() {
        window.location.hash = "achievements";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onClickSave() {
        let achieve = {
            year: parseInt(this.state.inputYear),
            message: this.state.inputMessage
        };
        alert(this.state.isEdit);
        let successCallback = () => alert("Successfully saved!");
        let failCallback = (err) => alert("Could not save achievement: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/achievements/v1/achievement/" + this.props.achieveId, achieve)
            .then (res => successCallback())
            .catch(err => failCallback(err));
        } else {
            API.post("api/achievements/v1/create", achieve)
            .then(res => successCallback())
            .catch(err => failCallback(err));
        }
    }

    onSaved() {
        window.location.hash = "achievements";
    }

    onDeleted() {
        const achieveId = this.props.achieveId;
        API.delete("api/achievements/v1/achievement/" + achieveId)
        .then(res => {
            window.location.hash = "achievements";
        })
    }

    render() {
        const isEdit = this.state.isEdit;
        const achieve = this.state.achieve;
        const title = isEdit ? "Edit Achievement" : "Add Achievement";

        let deleteButton = <div></div>;
        if (isEdit) {
          deleteButton = (
            <button className="btn-delete" onClick={this.onClickDelete}>
              Delete
            </button>
          );
        }

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
