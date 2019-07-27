'use strict';
require('./../styl/programClassModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
	getClasses
} from './repos/mainRepo.js';

export class ProgramClassModal extends React.Component {
  render() {
		const programObj = this.props.programObj;
		const programId = programObj.programId;
		const classes = this.props.classList;
		const avail = this.props.avail;

		var programTitle;
		if (avail) {
			programTitle = "Available classes for " + programObj.title;
		} else {
			programTitle = "Future classes for " + programObj.title;
		}

    const classList = classes
												.filter((c) => c.isAvailable == avail)
												.map((c, index) =>
										      <ProgramClassLine key={index} classObj={c}/>
										    );

    return (
			<div className="program-class-modal">
				<h1>{programTitle}</h1>
        <ul>
          {classList}
        </ul>
			</div>
    );
  }
}

class ProgramClassLine extends React.Component {
  render() {
    const classObj = this.props.classObj;
    const times = classObj.times.map((time, index) =>
      <div key={index}>{time}</div>
    );
		const url = "/class/" + classObj.key;

    return (
      <li>
        <div className="class-name">{classObj.className}</div>
        <div className="class-times">
					<div className="class-starts">{"Starts on: " + classObj.startDate}</div>
					{times}
				</div>
        <Link to={url}>Details &#62;</Link>
      </li>
    );
  }
}
