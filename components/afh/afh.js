'use strict';
require('./afh.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { AfhForm } from './afhForm.js';
import {
  getAFH,
  getLocation
} from '../repos/mainRepo.js';
const classnames = require('classnames');

export class AFHPage extends React.Component {
  constructor(props) {
    super(props);
    this.afh = getAFH();
  }

	componentDidMount() {
	  window.scrollTo(0, 0);
	}

	render() {
    const afh = this.afh;
		const textLocation = generateLocation(afh.locationId);

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
					<div className="times class-lines">{afh.date + " " + afh.time}</div>
          <div className="notes class-lines">{afh.notes || ""}</div>

					<AfhForm/>

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
