'use strict';
require('./announceEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
//import API from '../api.js';
//import { Modal } from '../modals/modal.js';
import { OkayModal } from '../modals/okayModal.js';
import { YesNoModal } from '../modals/yesnoModal.js';

export class AnnounceEditPage extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        isEdit: false,
        inputDate: "",
        inputAuthor: "",
        inputMessage: ""
      };

      // input onChange
      this.handleChange = this.handleChange.bind(this);

      // click on button
      //this.onClickCancel = this.onClickCancel.bind(this);
      //this.onClickDelete = this.onClickDelete.bind(this);
      //this.onClickSave = this.onClickSave.bind(this);

      // after action
      //this.onDeleted = this.onDeleted.bind(this);
      //this.onSaved = this.onSaved.bind(this);
      //this.onDismissModal = this.onDismissModal.bind(this);
    }

    handleChange(event, value) {
      this.setState({[value]: event.target.value});
    }

      render() {
        return (
            <div id ="view-announce-edit">
            <h2>Add Announcement</h2>

            <h4>Date</h4> {/* TODO: will be post date later*/}
            <input value={this.state.inputDate}
                    onChange={(e) => this.handleChange(e, "inputDate")}/>

            <h4>Author</h4>
            <input value={this.state.inputAuthor}
                    onChange={(e) => this.handleChange(e, "inputAuthor")}/>

            <h4>Message</h4>
            <input value={this.state.inputMessage}
                    onChange={(e) => this.handleChange(e, "inputMessage")}/>

            {/*<div className="buttons">
            <button className="btn-save" onClick={this.onClickSave}>Save</button>
            <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
            {deleteButton}
            </div>*/}
      </div>
    );
  }
}
