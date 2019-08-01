'use strict';
require('./contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactInputForm } from './contactInputForm.js';
const classnames = require('classnames');


export class ContactPage extends React.Component {
	componentDidMount() {
	  window.scrollTo(0, 0);
	}

	render() {
		return (
      <div id="view-contact">
        <div id="view-contact-container">
          <h1>Contact Us</h1>
					<h3>
						Please fill the following form.<br/>
						All fields are required.
					</h3>
					<ContactInputForm/>
        </div>
      </div>
		);
	}
}
