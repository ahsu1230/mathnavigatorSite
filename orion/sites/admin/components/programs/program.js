'use strict';
require('./program.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Modal } from '../modals/modal.js';
import { ProgramModal } from './programModal.js';

const FAKE_PROGRAM = {
  programKey: "ap_java",
  name: "AP Java",
  grade1: 11,
  grade2: 12,
  description: "Some description blahb lahblah"
};
const FAKE_LIST = [FAKE_PROGRAM, FAKE_PROGRAM, FAKE_PROGRAM, FAKE_PROGRAM];

export class ProgramPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: FAKE_LIST,
      showModal: false,
      targetProgram: {}
    };
    this.handleClick = this.handleClick.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
		this.dismissModal = this.dismissModal.bind(this);
  }

  handleClick() {
    this.setState({
      showModal: true
    });
	}

  handleEdit(program) {
    this.setState({
			showModal: true,
      targetProgram: program
		});
  }

	dismissModal() {
		this.setState({
			showModal: false,
      targetProgram: undefined
		});
	}

	render() {
    const rows = this.state.list.map((row, index) => {
      return <ProgramRow key={index} row={row} onHandleEdit={this.handleEdit}/>
    });
    const numRows = rows.length;
    const modalContent = <ProgramModal
                            program={this.state.targetProgram}
                            onDismiss={this.dismissModal}/>;
    const modalDiv = (
      <Modal content={modalContent}
              show={this.state.showModal}
              withClose={false}
              onDismiss={this.dismissModal}/>
    );

		return (
      <div id="view-program">
        {modalDiv}
        <h1>All Programs ({numRows})</h1>
        <ul id="list-heading">
          <li className="li-med">ProgramKey</li>
          <li className="li-med">Name</li>
          <li className="li-small">Grade1</li>
          <li className="li-small">Grade2</li>
        </ul>
        <ul id="list-rows">
          {rows}
        </ul>
        <button className="btn-program-add" onClick={this.handleClick}>
          Add Program
        </button>
      </div>
		);
	}
}

class ProgramRow extends React.Component {
  render() {
    const row = this.props.row;
    return (
      <li className="program-row">
        <div className="li-med">{row.programKey}</div>
        <div className="li-med">{row.name}</div>
        <div className="li-small">{row.grade1}</div>
        <div className="li-small">{row.grade2}</div>
        <button onClick={() => this.props.onHandleEdit(row)}>Edit</button>
      </li>
    );
  }
}
