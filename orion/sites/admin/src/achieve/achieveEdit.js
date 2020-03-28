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

    this.handleChange = this.handleChange.bind(this);

    this.onClickCancel = this.onClickCancel.bind(this);
    this.onClickDelete = this.onClickDelete.bind(this);
    this.onClickSave = this.onClickSave.bind(this);

    this.onDeleted = this.onDeleted.bind(this);
    this.onSaved = this.onSaved.bind(this);

    }

    componentDidMount() {
      const Id = this.props.Id;
      if (Id) {
        API.get("api/achievements/v1/achievement/" + Id)
          .then(res => {
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
        const Id = this.props.Id;
        API.delete("api/achievements/v1/achievement/" + Id)
        .then(res => {
            window.location.hash = "achievements";
        })
    }

    onClickSave() {
        let achieve = {
            year: parseInt(this.state.inputYear),
            message: this.state.inputMessage,
        };
        let successCallback = () => this.onSaved();
        let failCallback = (err) => alert("Could not save achievement: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/achievements/v1/achievement/" + this.props.Id, achieve)
            .then (res => successCallback())
            .catch(err => failCallback(err));
        } else {
            API.post("api/achievements/v1/create", achieve)
            .then(res => successCallback())
            .catch(err => failCallback(err));
        }
    }

    onClickCancel() {
        console.log("cancel button clicked");
    }

    onDeleted() {
        const Id = this.props.Id;
        API.delete("api/achievements/v1/achievement/" + Id)
        .then(res => {
            window.location.hash = "achievements";
        })
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
