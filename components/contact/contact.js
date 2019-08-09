'use strict';
require('./contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactForm } from './contactForm.js';
const classnames = require('classnames');


export class ContactPage extends React.Component {
	render() {
		return (
      <div id="view-contact">
        <div id="view-contact-container">
          <h1>Contact Us</h1>
					<h3>
						Please fill the following form<br/>
						to register for a program or to contact us.
					</h3>
					<ContactForm/>
        </div>
      </div>
		);
	}
}
