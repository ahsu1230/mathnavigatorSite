'use strict';
require('./program.styl');
import React from 'react';
import ReactDOM from 'react-dom';

const FAKE_PROGRAM = {
  programKey: "ap_java",
  name: "AP Java",
  grade1: 11,
  grade2: 12
};
const FAKE_LIST = [FAKE_PROGRAM, FAKE_PROGRAM, FAKE_PROGRAM, FAKE_PROGRAM];

export class ProgramPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: FAKE_LIST
    };
  }

	render() {
    const rows = this.state.list.map((row, index) => {
      return <ProgramRow key={index} row={row}/>
    });
    const numRows = rows.length;

		return (
      <div id="view-program">
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
        <button className="btn-program-add">Add Program</button>
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
        <button>Edit</button>
      </li>
    );
  }
}
