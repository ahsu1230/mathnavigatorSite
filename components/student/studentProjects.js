'use strict';
require('./studentProjects.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
const classnames = require('classnames');

export class StudentProjectsPage extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
		if (process.env.NODE_ENV === 'production') {
      mixpanel.track("student-projects");
    }
  }

	render() {
		return (
      <div id="view-students-projects">
        <div id="view-students-projects-container">
          <h1>Student Website Projects</h1>
          <h3>Coming Soon</h3>
        </div>
      </div>
		);
	}
}
