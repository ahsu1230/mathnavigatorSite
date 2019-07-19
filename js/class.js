'use strict';
require('./../styl/class.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { ErrorPage } from './error.js';
const classNames = require('classnames');
import {
  getLocation,
  getPrereqs,
  getProgramClass,
  getProgramByIds,
  getSessions
} from './repos/mainRepo.js';

function generatePrereqs(prereqIds) {
  var prereqsPrograms = getProgramByIds(prereqIds);
  var prereqsNames = prereqsPrograms.map(function(program) {
    return program ? program.title : null;
  });
  return prereqsNames.join(", ");
}

function generateLocation(locationObj) {
  return (
    <div className="class-lines">
      {locationObj.name}<br/>
      {locationObj.address1}<br/>
      {locationObj.address2}<br/>
      {locationObj.address3}
    </div>
  );
}

function generateTimes(classObj) {
  const times = classObj.times.map((time, index) =>
    <div key={index}>{time}</div>
  );
  return (
    <div className="class-lines">
      {times}
      <div>
        First session: {classObj.startDate}<br/>
        Last session: {classObj.endDate}
      </div>
    </div>
  );
}

function generatePricing(pricePerSession, numSessions) {
  let line1;
  let line2;
  let line3;
  if (numSessions > 0) {
    line1 = <div>{"Number of sessions: " + numSessions}</div>;
    line3 = <div>{"Total: $" + numSessions * pricePerSession}</div>;
  } else {

  }
  line2 = <div>{"Price per session: $" + pricePerSession}</div>;
  return (
    <div className="class-lines">
      {line1}
      {line2}
      {line3}
    </div>
  );
}

function generateSchedules(sessions) {
  if (!sessions || sessions.length == 0) {
    return (
      <div className="class-lines not-avail">
        No schedule available at the moment. <br/>
        Please check again later.
      </div>
    );
  }

  var sessionIndex = 0;
  const sessionLines = sessions.map(function(session, index) {
    var text1 = "";
    var text2 = ""
    if (session.canceled) {
      text1 = "Canceled";
      text2 = "";
    } else {
      text1 = session.time;
      text2 = session.notes;
      sessionIndex++;
    }
    return (
      <SessionLine
        key = {session.key + index}
        sessionIndex = {sessionIndex}
        date = {session.date}
        canceled = {session.canceled}
        text1 = {text1}
        text2 = {text2}
      />
    );
  });
  return (
    <ul>
      {sessionLines}
    </ul>
  );
}

export class ClassPage extends React.Component {
  render () {
    var valid = true;
    const key = this.props.slug;
    const pair = getProgramClass(key);
    valid = valid && Boolean(pair);
    valid = valid && Boolean(pair.programObj);
    valid = valid && Boolean(pair.classObj);

    var content;
    if (valid) {
      content = <ClassContent classKey={key} pair={pair}/>;
    } else {
      content = <ErrorPage classDNE={key}/>;
    }

    return (
      <div> {content} </div>
    );
  }
}

class ClassContent extends React.Component {
	render() {
    const classKey = this.props.classKey;
    const pair = this.props.pair;

    // Variables
    const programObj = pair.programObj;
    const classObj = pair.classObj;
    const locationObj = getLocation(classObj.locationId);
    let sessions = getSessions(classKey);
    sessions = sessions ? sessions : [];
    var sessionCounter = 0;
    sessions.forEach(function(session) {
      if (!session.canceled) {
        sessionCounter++;
      }
    });
    var programPrereqs = getPrereqs(programObj.programId);
    const prereqIds = programPrereqs ? programPrereqs.requiredProgramIds : [];

    // Components
    const classFullName = programObj.title + " " + classObj.className;
    const prereqs = generatePrereqs(prereqIds);
    let prereqsLine;
    if (prereqs) {
      prereqsLine = <div>{"Prequirements: " + prereqs}<br/></div>;
    }
    const textLocation = generateLocation(locationObj);
    const textTimes = generateTimes(classObj);
    const textPricing = generatePricing(classObj.pricePerSession, sessionCounter);
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
              <Link to={"/contact"}>
                <button className="inverted">Register</button>
              </Link>
            </div>
          </div>

          <div id="view-schedule">
            <b>Schedule</b>
            {schedules}
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
          <div className={classText1}>{text1}</div>
          <div className="class-line alert">{text2}</div>
        </div>
      </li>
    );
  }
}
