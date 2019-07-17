'use strict';
require('./../styl/contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
const classnames = require('classnames');

export class ContactPage extends React.Component {
	render() {
		var interestedPrograms = [];
		const interestedButtonText = interestedPrograms.length > 0 ? "Select More Programs..." : "Select Programs...";

		return (
      <div id="view-contact">
        <div id="view-contact-container">
          <h1>Contact Us</h1>

					<div className="section input">
						<h2>Student Information</h2>
						<div className="contact-input-container">
							<ContactInput addClasses="student-fname" title="First Name" onError="Not valid"/>
							<ContactInput addClasses="student-lname" title="Last Name"/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="student-age" title="Age"/>
							<ContactInput addClasses="student-grade" title="Grade"/>
							<ContactInput addClasses="student-school" title="School"/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="student-phone" title="Phone"/>
							<ContactInput addClasses="student-email" title="Email"/>
						</div>

						<h2>Guardian Information</h2>
						<div className="contact-input-container">
							<ContactInput addClasses="guard-fname" title="First Name"/>
							<ContactInput addClasses="guard-lname" title="Last Name"/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="guard-phone" title="Phone"/>
							<ContactInput addClasses="guard-email" title="Email"/>
						</div>
					</div>


					<div className="section interested">
						<h2>Interested Programs</h2>
						<div className="list-interested">{interestedPrograms}</div>
						<button className="inverted">{interestedButtonText}</button>
					</div>

					<div className="section additional">
						<h2>Additional Information</h2>
						<div className="textarea-container">
							<textarea></textarea>
							<p>
								Information will be sent to:<br/>
								<a>andymathnavigator@gmail.com</a>
							</p>
							<button>Submit</button>
						</div>
					</div>
        </div>
      </div>
		);
	}
}

class ContactInput extends React.Component {
	render() {
		const classNames = classnames("contact-input", this.props.addClasses);
		const title = this.props.title;
		const placeholder = this.props.placeholder;
		const onError = this.props.onError;
		return (
			<div className={classNames}>
				<label>{title}</label>
				<input placeholder={placeholder}/>
				<label>{this.props.onError}</label>
			</div>
		);
	}
}
