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
  return prereqsNames.join(", ");
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
  var sessionCounter = 0;
  return sessions.map(function(session, index) {
    var text1 = "";
    var text2 = ""
    if (session.canceled) {
      text1 = "Canceled";
      text2 = "";
    } else {
      text1 = session.time;
      text2 = session.notes;
      sessionCounter++;
    }
    return (
      <SessionLine
        key = {session.key + index}
        sessionIndex = {sessionCounter}
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
    let prereqsLine;
    if (prereqs) {
      prereqsLine = <div>{"Prequirements: " + prereqs}<br/></div>;
    }
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
            <div className="class-info-1">
              Grades: {programObj.grade1} - {programObj.grade2}<br/>
              {prereqsLine}
              {programObj.description}
            </div>

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
    var sessionIndex = this.props.sessionIndex;
    var date = this.props.date;
    var canceled = this.props.canceled;
    var text1 = this.props.text1;
    var text2 = this.props.text2;

    const classText1 = classNames("", {
      alert: canceled
    });
    return (
      <li>
        <div className="line-left">
          <div className="line-index">{canceled ? "" : sessionIndex}</div>
          <div className="line-date">{date}</div>
        </div>
        <div className="line-right">
          <p className={classText1}>{text1}</p>
          <p className="alert">{text2}</p>
        </div>
      </li>
    );
  }
}
