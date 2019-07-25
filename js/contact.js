'use strict';
require('./../styl/contact.styl');
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
					<ContactInputForm/>
        </div>
      </div>
		);
	}
}
