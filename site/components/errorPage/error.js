'use strict';
require('../app/app.styl');
require('./error.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { createPageTitle } from '../constants.js';
const srcBroken = require('../../assets/compass_broken.png');

export class ErrorPage extends React.Component {
	componentDidMount() {
		document.title = createPageTitle("Error");
	}

	render() {
    const classDNE = this.props.classDNE;
    var errorMsg = "";
    if (classDNE) {
      errorMsg = "ClassKey '" + classDNE + "' does not exist.";
    }

    return (
      <div id="view-error">
        <h1>Page Not Found</h1>
        <img src={srcBroken}/>
        <p>
          <Link to="/programs">View our Programs</Link>
          To find what you're looking for
        </p>
        <h6>{errorMsg}</h6>
      </div>
    );
  }
}
