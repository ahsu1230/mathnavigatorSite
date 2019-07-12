'use strict';
require('./../styl/class.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
const classNames = require('classnames');
import { getProgramClass, getProgramByIds } from './repos/programRepo.js';
import { getLocation } from './repos/locationRepo.js';
import { getSessions } from './repos/sessionRepo.js';
import { getPrereqs } from './repos/prereqRepo.js';

function generatePrereqs(prereqIds) {
  var prereqsPrograms = getProgramByIds(prereqIds);
  var prereqsNames = prereqsPrograms.map(function(program) {
    return program ? program.title : null;
  });
  var prereqsText = prereqsNames.join(", ");
  return prereqsText ? "Pre-requirements: " + prereqsText : "";
}

function generateLocation(locationObj) {
  return (
    <p>
      {locationObj.name}<br/>
      {locationObj.address1}<br/>
      {locationObj.address2}<br/>
      {locationObj.address3}
    </p>
  );
}

function generateTimes(classObj) {
  const times = classObj.times.map((time, index) =>
    <p className="line-time" key={index}>{time}</p>
  );
  return (
    <div>
      {times}
      <p>
        First session: {classObj.startDate}<br/>
        Last session: {classObj.endDate}
      </p>
    </div>
  );
}

function generatePricing(pricePerSession, numSessions) {
  return (
    <p>
      Number of sessions: {numSessions}<br/>
      Price per session: ${pricePerSession}<br/>
      Total: ${numSessions * pricePerSession}<br/>
    </p>
  );
}

function generateSchedules(sessions) {
  return sessions.map(function(session, index) {
    var text1 = "";
    var text2 = ""
    if (session.canceled) {
      text1 = "Canceled";
      text2 = "";
    } else {
      text1 = session.time;
      text2 = session.notes;
    }
    return (
      <SessionLine
        key = {session.key + index}
        sessionIndex = {index}
        date = {session.date}
        canceled = {session.canceled}
        text1 = {text1}
        text2 = {text2}
      />
    );
  });
}

export class ClassPage extends React.Component {
	render() {
    const key = this.props.slug;

    // Variables
    const pair = getProgramClass(key);
    const programObj = pair.programObj;
    const classObj = pair.classObj;
    const locationObj = getLocation(classObj.locationId);
    const sessions = getSessions(key);
    const prereqIds = getPrereqs(programObj.programId).requiredProgramIds;

    // Components
    const classFullName = programObj.title + " " + classObj.className;
    const prereqs = generatePrereqs(prereqIds);
    const textLocation = generateLocation(locationObj);
    const textTimes = generateTimes(classObj);
    const textPricing = generatePricing(classObj.pricePerSession, sessions.length);
    const schedules = generateSchedules(sessions);

		return (
      <div id="view-class">
        <div id="view-class-container">
          <h1>
            <Link to="/programs">Programs</Link> > {classFullName}
          </h1>

          <div id="class-info-container">
            <p className="class-info-1">
              Grades: {programObj.grade1} - {programObj.grade2}<br/>
              {prereqs}
              {programObj.description}
            </p>

            <div className="class-info-2">
              <b>Location</b>
              {textLocation}
              <b>Times</b>
              {textTimes}
              <b>Pricing</b>
              {textPricing}
              <button className="inverted">Register</button>
            </div>
          </div>

          <div id="view-schedule">
            <b>Schedule</b>
            <ul>
              {schedules}
            </ul>
          </div>

          <div id="view-questions">
            Questions? <Link to="/contact">Contact Us</Link>
          </div>
          <Link to="/programs">
            <button className="inverted">&#60; More Programs</button>
          </Link>
        </div>
      </div>
		);
	}
}

class SessionLine extends React.Component {
  render() {
    const classText1 = classNames("", {
      alert: this.props.canceled
    });
    return (
      <li>
        <div className="line-left">
          <div className="line-index">{this.props.sessionIndex}</div>
          <div className="line-date">{this.props.date}</div>
        </div>
        <div className="line-right">
          <p className={classText1}>{this.props.text1}</p>
          <p className="alert">{this.props.text2}</p>
        </div>
      </li>
    );
  }
}
