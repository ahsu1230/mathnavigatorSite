'use strict';
require('./contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactForm } from './contactForm.js';
const classnames = require('classnames');

export class ContactPage extends React.Component {
	constructor(props) {
		super(props);

		const { match, location, history } = this.props;
		this.interested = [];
		var parsed = parseQuery((location || {}).search);
		if (parsed && parsed.interest) {
			this.interested.push(parsed.interest);
		}
	}

	componentDidMount() {
		if (process.env.NODE_ENV === 'production') {
			mixpanel.track("contact");
		}
	}

	render() {
		return (
      <div id="view-contact">
        <div id="view-contact-container">
          <h1>Contact Us</h1>
					<h3>
						Please fill the following form<br/>
						to register for a program or to contact us.
					</h3>
					<ContactForm startingInterest={this.interested}/>
        </div>
      </div>
		);
	}
}

function parseQuery(hash) {
	if (!hash) {
		return "";
	}
  var i = hash.indexOf("?");
  var parsed = {};
  if (i >= 0) {
    hash = hash.slice(i + 1);

    // parse Query String
    var params = hash.split("&");
    for (var i = 0; i < params.length; i++) {
      var pair = params[i].split("=");
      var pairKey = decodeURIComponent(pair[0]);
      var pairValue = decodeURIComponent(pair[1]);
      parsed[pairKey] = pairValue;
    }
  }
  return parsed;
}
