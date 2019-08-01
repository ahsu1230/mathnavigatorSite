'use strict';
require('./afh.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
const classnames = require('classnames');
import {
  getLocation
} from '../repos/mainRepo.js';

export class AFHPage extends React.Component {
	componentDidMount() {
	  window.scrollTo(0, 0);
	}

	render() {
		// Switch to real values later
		const textLocation = generateLocation("loc_wchs");
		const textTimes = generateTimes();

		return (
      <div id="view-afh">
        <div id="view-afh-container">

					<h1>Ask For Help</h1>
					<p>
						We provide free sessions for students to ask for additional assistance on any of our program subjects.
						Please fill the below form to let us know you are attending.
						You must be registered with one of our programs to attend.
					</p>

					<h2>Location & Time</h2>
					<div className="location">{textLocation}</div>
					<div className="times">{textTimes}</div>

					<h2>Student Information</h2>

					<Link className="back-programs" to="/programs">&#60; Back to Programs</Link>
        </div>
      </div>
		);
	}
}

/* Helper functions */
function generateLocation(locationId) {
	var locationObj = getLocation(locationId);
  return (
    <div className="class-lines">
      {locationObj.name}<br/>
      {locationObj.address1}<br/>
      {locationObj.address2}<br/>
      {locationObj.address3}
    </div>
  );
}

function generateTimes() {
	return (
    <div className="class-lines">
      Mon. 6:30pm - 9:00pm
    </div>
  );
}
