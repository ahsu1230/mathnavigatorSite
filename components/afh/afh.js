'use strict';
require('./afh.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { AfhForm } from './afhForm.js';

export class AFHPage extends React.Component {
	componentDidMount() {
	  window.scrollTo(0, 0);
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
