'use strict';
require('./program.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import API from '../api.js';
import { Modal } from '../modals/modal.js';
import { Link } from 'react-router-dom';

export class ProgramPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      list: []
    };
  }

  componentDidMount() {
    API.get("api/programs/v1/")
      .then(res => {
        const programs = res.data;
        this.setState({ list: programs });
      });
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
        <Link className="add-program" to={"/programs/add"}>Add Program</Link>
      </div>
		);
	}
}

class ProgramRow extends React.Component {
  render() {
    const row = this.props.row;
    const url = "/program/" + row.programId + "/edit";
    return (
      <li className="program-row">
        <div className="li-med">{row.programId}</div>
        <div className="li-med">{row.name}</div>
        <div className="li-small">{row.grade1}</div>
        <div className="li-small">{row.grade2}</div>
        <Link to={url}>Edit</Link>
      </li>
    );
  }
}
