'use strict';
require('./../styl/class.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { getProgramClass } from './repos/programRepo.js';
import { getLocation } from './repos/locationRepo.js';
import { getSessions } from './repos/sessionRepo.js';
import { getPrereqs } from './repos/prereqRepo.js';

export class ClassPage extends React.Component {
	render() {
    const key = this.props.slug;

    const pair = getProgramClass(key);
    const programObj = pair.programObj;
    const classObj = pair.classObj;
    const locationObj = getLocation(classObj.locationId);
    const sessions = getSessions(key);

    const schedules = sessions.map((session, index) =>
      <ScheduleSession
        key = {session.key + "_" + index}
        date = {session.date}
        canceled = {session.canceled}
        time = {session.time}
        notes = {session.notes}
      />
    );

		return (
      <div id="view-class">
        <div id="view-class-container">
          Class {key}
          <p>{programObj.title}</p>
          <p>{classObj.className}</p>
          <p>{programObj.description}</p>
          <p>{programObj.grade1} - {programObj.grade2}</p>
          <p>{locationObj.name}</p>
          <p>times: {JSON.stringify(classObj.times)}</p>
          <p>start: {classObj.startDate}</p>
          <p>end: {classObj.endDate}</p>
          <p>num sessions: {classObj.times.length} </p>
          <p>price per session: ${classObj.pricePerSession} </p>
          <p>total: ${classObj.times.length * classObj.pricePerSession}</p>

          <div id="view-schedule">
            {schedules}
          </div>
        </div>
      </div>
		);
	}
}

class ScheduleSession extends React.Component {
  render() {
    return (
      <div>
        Schedule:
        <ul>
          <li>{this.props.date}</li>
          <li>{this.props.canceled.toString()}</li>
          <li>{this.props.time}</li>
          <li>{this.props.notes}</li>
        </ul>
      </div>
    );
  }
}
