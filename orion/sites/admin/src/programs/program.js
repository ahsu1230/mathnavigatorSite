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
    API.get("api/programs/v1/all")
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
        <ul>
            <h1>All Programs ({numRows}) </h1>
            <div> You have 5 unpublished items. </div>
            <div> You have selected 0 items to publish. </div>
        </ul>
        <ul id="list-heading">
          <li className="li-med">ProgramKey</li>
          <li className="li-med">Name</li>
          <li className="li-small">Grade1</li>
          <li className="li-small">Grade2</li>
        </ul>
        <ul id="list-rows">
          {rows}
        </ul>
        <ul id="list-buttons">
            <div className="li-med">
                <button>
                <Link className="add-program" to={"/programs/add"}>Add Program</Link>
                </button>
            </div>
            <div className="li-med">
                <button>
                <Link className="publish" to={"/programs/add"}>Publish</Link>
                </button>
            </div>
        </ul>
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
        <input type="checkbox" name="unpublished"/>
        <div className="li-med">{row.programId}</div>
        <div className="li-med">{row.name}</div>
        <div className="li-small">{row.grade1}</div>
        <div className="li-small">{row.grade2}</div>
        <Link to={url}>Edit</Link>
      </li>
    );
  }
}
