'use strict';
require('./achieveEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import API from '../api.js';
import { Modal } from '../modals/modal.js';
import { OkayModal } from '../modals/okayModal.js';
import { YesNoModal } from '../modals/yesnoModal.js';

export class AchieveEditPage extends React.Component {

  render() {

    return (
      <div id="view-achieve-edit">
        {modalDiv}
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
