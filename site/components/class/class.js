'use strict';
require('./class.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { Promise } from 'bluebird';
import { find, isEmpty, pick } from 'lodash';
import { ErrorPage } from '../errorPage/error.js';
import {
  getAnnouncements,
  getLocation,
  getPrereqs,
  getProgramByIds,
  getProgramAndClass,
  getSessions
} from '../repos/apiRepo.js';
import { createFullClassName, createPageTitle } from '../constants.js';
const classnames = require('classnames');
const srcClose = require('../../assets/close_black.svg');

export class ClassPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      announcements: [],
      classKey: undefined,
      classAnnounce: {},
      classObj: {},
      location: {},
      prereqPrograms: [],
      programObj: {},
      sessions: []
    };
  }

  componentDidMount() {
    const classKey = this.props.slug;

    Promise.join(
      getAnnouncements(),
      getProgramAndClass(classKey), // -> get prereqs(programId) -> getProgramByIds(arr)
      getSessions(classKey),
      (announcements, programClassObj, sessions) => {
        var classObj = programClassObj.classObj;
        var programObj = programClassObj.programObj;
        var classAnnounce = find(announcements, function(o) {
          return find(o.classKeys, cKey => (classKey === cKey));
        });
        this.setState({
          announcements: announcements,
          classAnnounce: classAnnounce,
          classObj: classObj,
          programObj: programObj,
          sessions: sessions
        });

        getLocation(classObj.locationId).then(location => {
          this.setState({ location: location });
        });

        getPrereqs(programObj.programId).then(prereqObj => {
          if (prereqObj) {
            getProgramByIds(prereqObj.requiredProgramIds).then(prereqPrograms => {
              if (prereqPrograms) {
                this.setState({ prereqPrograms: prereqPrograms });
              }
            });
          }
        });
      }
    );

    this.setState({
      classKey: classKey
    });

    if (process.env.NODE_ENV === 'production') {
      mixpanel.track("class", {"key": classKey});
    }
  }

  render () {
    const classKey = this.state.classKey;
    var content;
    if (isEmpty(this.state.classObj)) {
      content = <ErrorPage classDNE={this.state.classKey}/>;
    } else {
      const classInfo = pick(this.state, [
        'classKey',
        'classAnnounce',
        'classObj',
        'programObj',
        'location',
        'prereqPrograms',
        'sessions'
      ]);
      content = <ClassContent info={classInfo}/>
    }

    return (
      <div> {content} </div>
    );
  }
}

export class ClassContent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showAnnounce: true
    }
    this.handleDismissAnnounce = this.handleDismissAnnounce.bind(this);
  }

  componentDidUpdate() {
    const info = this.props.info;
    const fullClassName = createFullClassName(info.programObj, info.classObj);
    document.title = createPageTitle(fullClassName);
  }

  handleDismissAnnounce() {
    this.setState({showAnnounce: false});
  }

	render() {
    // Variables
    const info = this.props.info || {};
    const announcements = info.classAnnounce;
    const classKey = info.classKey;
    const classObj = info.classObj;
    const locationObj = info.location;
    const programObj = info.programObj;
    const sessions = info.sessions || [];
    const prereqPrograms = info.prereqPrograms || [];

    // All Components
    const announce = generateAnnouncement(announcements,
      this.state.showAnnounce, this.handleDismissAnnounce);
    const classFullName = createFullClassName(programObj, classObj);
    const prereqsLine = generatePrereqs(prereqPrograms);
    const textLocation = generateLocation(locationObj);
    const textTimes = generateTimes(classObj);
    const textPricing = generatePricing(classObj.priceLump,
        classObj.pricePerSession, sessions, classObj.paymentNotes);
    const schedules = generateSchedules(sessions, classObj.allYear, classObj.times);

		return (
      <div id="view-class">
        {announce}
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
              <Link to={"/contact?interest=" + classKey}>
                <button>Register</button>
              </Link>
              <b>Location</b>
              {textLocation}
              <b>Times</b>
              {textTimes}
              <b>Pricing</b>
              {textPricing}
            </div>
          </div>

          <div id="view-schedule">
            <b>Schedule</b>
            {schedules}
          </div>

          <div id="view-questions">
            Questions? <Link to={"/contact?interest=" + classKey}>Contact Us</Link>
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

    const classText1 = classnames("text-1", {
      alert: canceled
    });
    return (
      <li>
        <div className="line-left">
          <div className="line-index">{canceled ? "" : sessionIndex}</div>
          <div className="line-date">{date}</div>
        </div>
        <div className="line-right">
          <span className={classText1}>{text1}</span>
          <span className="text-2 alert">{text2}</span>
        </div>
      </li>
    );
  }
}

/* Helper Methods */
function generatePrereqs(prereqPrograms) {
  const validPrereqs = prereqPrograms.filter(p => p && p.title);

  let prereqsLine = <div></div>;
  if (validPrereqs.length > 0) {
    prereqsLine = <div>
      {"Pre-requirements: " + validPrereqs.map(p => p.title).join(", ")}
    </div>;
  } else {
    return <div></div>;
  }
  return prereqsLine;
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

function generatePricing(priceLump, pricePerSession, sessions, paymentNotes) {
  var numSessions = 0;
  sessions.forEach(function(session) {
    if (!session.canceled) {
      numSessions++;
    }
  });

  let line1;
  let line2;
  let line3;
  let line4;
  if (priceLump) {
    line1 = <div>{"Price: $" + priceLump}</div>
  } else if (numSessions > 0) {
    line1 = <div>{"Number of sessions: " + numSessions}</div>;
    line2 = <div>{"Price per session: $" + pricePerSession}</div>;
    line3 = <div>{"Total: $" + numSessions * pricePerSession}</div>;
  } else if (pricePerSession > 0) {
    line1 = <div>{"Price per session: $" + pricePerSession}</div>;
  }

  if (paymentNotes) {
    line4 = <div>{"Payment Due: " + paymentNotes}</div>;
  }

  return (
    <div className="class-lines">
      {line1}
      {line2}
      {line3}
      {line4}
    </div>
  );
}

function generateSchedules(sessions, allYear, classTimes) {
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
    var text2 = "";
    if (session.canceled) {
      text1 = "No Class";
      text2 = "";
    } else {
      text1 = session.time;
      text2 = session.notes;
      sessionIndex++;
    }
    return (
      <SessionLine
        key = {index}
        sessionIndex = {sessionIndex}
        date = {session.date}
        canceled = {session.canceled}
        text1 = {text1}
        text2 = {text2}
      />
    );
  });

  var allYearText = "";
  if (allYear && classTimes.length > 0) {
    allYearText = "Classes are held every week on ";
    const timesText = classTimes.map((c, i) => {
      return (<div key={i}>{c}</div>);
    });
    allYearText = (
      <div>
        {allYearText}
        {timesText}
        <h4>Except the following times:</h4>
      </div>
    );
  }

  return (
    <div>
      {allYearText}
      <ul>
        {sessionLines}
      </ul>
    </div>
  );
}

function generateAnnouncement(announce, showAnnounce, onDismiss) {
  if (isEmpty(announce)) {
    return (<div></div>);
  }

  var classView = classnames({
      dismiss: !showAnnounce
  });
  return (
    <div id="view-class-announce" className={classView}>
      <div className="announce-container">
        <h2>Announcement</h2>
        <p>{announce.message}</p>
        <button className="close-x" onClick={onDismiss}>
          <span>Dismiss</span><img src={srcClose}/>
        </button>
      </div>
    </div>
  );
}
