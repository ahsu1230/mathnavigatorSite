'use strict';
require('./../styl/programClassModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
	getClasses
} from './repos/mainRepo.js';
const srcClose = require('../assets/close_black.svg');

export class ProgramClassModal extends React.Component {
  render() {
    const show = this.props.show;
    const modalClass = show ? "show" : "";

    // Replace later ////////
    const programId = "sat1";
    const programTitle = "SAT";
    const classes = getClasses(programId);

    // const programTitle = this.props.programTitle;
    // const classes = this.props.classes;
    const onDismissFn = this.props.onDismiss;
    const onDismiss = function() {
      console.log("Dismiss modal");
      // onDismissFn();
    }
    /////////////////////////

    const classList = classes.map((c, index) =>
      <ProgramClassLine key={index} classObj={c}/>
    );

    return (
      <div id="program-class-modal-view" className={modalClass}>
        <div id="program-class-overlay" onClick={onDismiss}></div>
        <div id="program-class-modal">
          <h1>{programTitle}</h1>
          <button onClick={onDismiss}>
            <img src={srcClose}/>
          </button>
          <ul>
            {classList}
          </ul>
        </div>
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

    return (
      <li>
        <div className="class-name">{classObj.className}</div>
        <div className="class-times">{times}</div>
        <Link to="/">Details &#62;</Link>
      </li>
    );
  }
}
