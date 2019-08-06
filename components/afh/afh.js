'use strict';
require('./afh.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { AfhForm } from './afhForm.js';
import { createPageTitle } from '../constants.js';

export class AFHPage extends React.Component {
	componentDidMount() {
    document.title = createPageTitle("Ask For Help");
  }

	render() {
		return (
      <div id="view-afh">
        <div id="view-afh-container">

					<h1>Ask For Help</h1>
					<p>
						We provide free sessions for students to ask for additional assistance on any of our program subjects.
						Please fill the below form to let us know you are attending.
						You must be registered with one of our programs to attend.
					</p>

					<AfhForm/>

					<Link className="back-programs" to="/programs">&#60; Back to Programs</Link>
        </div>
      </div>
		);
	}
}
