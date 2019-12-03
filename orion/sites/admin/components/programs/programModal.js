'use strict';
require('./programModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class ProgramModal extends React.Component {
  constructor(props) {
    super(props);
    this.onSave = this.onSave.bind(this);
		this.onCancel = this.onCancel.bind(this);
  }

  onSave() {
    // save!
    if (this.props.onDismiss) {
      this.props.onDismiss();
    }
  }

  onCancel() {
    if (this.props.onDismiss) {
      this.props.onDismiss();
    }
  }

  render() {
    const program = this.props.program;
    const programKey = program ? program.programKey : "";
    const programName = program ? program.name : "";
    const grade1 = program ? program.grade1 : "";
    const grade2 = program ? program.grade2 : "";
    const description = program ? program.description : "";
    return (
      <div id="program-modal">
        <h2>Add Program</h2>
        <h4>Program Key</h4>
        <input value={programKey}/>
        <h4>Program Name</h4>
        <input value={programName}/>
        <h4>Grade1</h4>
        <input value={grade1}/>
        <h4>Grade2</h4>
        <input value={grade2}/>
        <h4>Description</h4>
        <textarea value={description}/>

        <div className="buttons">
          <button className="btn-save" onClick={this.onSave}>Save</button>
          <button className="btn-cancel" onClick={this.onCancel}>Cancel</button>
        </div>
      </div>
    );
  }
}
