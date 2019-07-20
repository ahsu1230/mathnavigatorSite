'use strict';
require('./../styl/contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactInterestSection } from './contactInterest.js';
const classnames = require('classnames');


export class ContactPage extends React.Component {
	constructor(props){
    super(props);
		this.state = {
			formSubmitted: false,
			studentFirstName: "",
			studentLastName: "",
			studentAge: 0,
			studentGrade: 0,
			studentSchool: "",
			studentPhone: "",
			studentEmail: "",
			guardFirstName: "",
			guardLastName: "",
			guardPhone: "",
			guardEmail: "",
			interestedPrograms: [],
			additionalText: ""
		};
    this.handleSubmit = this.handleSubmit.bind(this);
		this.getInputInfo = this.getInputInfo.bind(this);
		this.updateCb = this.updateCb.bind(this);
		this.updateTextArea = this.updateTextArea.bind(this);
  }

	updateCb(propertyName, newValue) {
		var obj = {};
		obj[propertyName] = newValue;
		this.setState(obj);
	}

	updateTextArea(event) {
		this.updateCb("additionalText", event.target.value);
	}

	render() {
		return (
      <div id="view-contact">
        <div id="view-contact-container">
          <h1>Contact Us</h1>

					<div className="section input">
						<h2>Student Information</h2>
						<div className="contact-input-container">
							<ContactInput addClasses="student-fname" title="First Name" propertyName="studentFirstName" updateCallback={this.updateCb}/>
							<ContactInput addClasses="student-lname" title="Last Name" propertyName="studentLastName" updateCallback={this.updateCb}/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="student-age" title="Age" propertyName="studentAge" updateCallback={this.updateCb}/>
							<ContactInput addClasses="student-grade" title="Grade" propertyName="studentGrade" updateCallback={this.updateCb}/>
							<ContactInput addClasses="student-school" title="School" propertyName="studentSchool" updateCallback={this.updateCb}/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="student-phone" title="Phone" propertyName="studentPhone" updateCallback={this.updateCb}/>
							<ContactInput addClasses="student-email" title="Email" propertyName="studentEmail" updateCallback={this.updateCb}/>
						</div>

						<h2>Guardian Information</h2>
						<div className="contact-input-container">
							<ContactInput addClasses="guard-fname" title="First Name" propertyName="guardFirstName" updateCallback={this.updateCb}/>
							<ContactInput addClasses="guard-lname" title="Last Name" propertyName="guardLastName" updateCallback={this.updateCb}/>
						</div>
						<div className="contact-input-container">
							<ContactInput addClasses="guard-phone" title="Phone" propertyName="guardPhone" updateCallback={this.updateCb}/>
							<ContactInput addClasses="guard-email" title="Email" propertyName="guardEmail" updateCallback={this.updateCb}/>
						</div>
					</div>

					<div className="section interested">
						<ContactInterestSection/>
					</div>

					<div className="section additional">
						<h2>Additional Information</h2>
						<div className="textarea-container">
							<textarea onChange={this.updateTextArea}/>
							<p>
								Information will be sent to:<br/>
								<a>andymathnavigator@gmail.com</a>
							</p>
							<button onClick={this.handleSubmit}>Submit</button>
						</div>
					</div>

        </div>
      </div>
		);
	}

	getInputInfo() {
		return {
			studentFirstName: this.state.studentFirstName,
			studentLastName: this.state.studentLastName,
			studentAge: this.state.studentAge,
			studentGrade: this.state.studentGrade,
			studentSchool: this.state.studentSchool,
			studentPhone: this.state.studentPhone,
			studentEmail: this.state.studentEmail,
			guardFirstName: this.state.guardFirstName,
			guardLastName: this.state.guardLastName,
			guardPhone: this.state.guardPhone,
			guardEmail: this.state.guardEmail,
			interestedPrograms: this.state.interestedPrograms,
			additionalText: this.state.additionalText
		};
	}

	handleSubmit(event) {
    event.preventDefault();
		const template = "mathnavigatorwebsitecontact";
		const receiverEmail = "andymathnavigator@gmail.com";
		const senderEmail = "anonymous@andymathnavigator.com";

		var inputInfo = this.getInputInfo();
		var emailMessage = generateEmailMessage(inputInfo);

		console.log("Sending email... " + emailMessage);
		// sendEmail(
		// 	template,
		// 	senderEmail,
		// 	receiverEmail,
		// 	emailMessage
		// );

		this.setState({
			formSubmitted: true
		});
		console.log("Email request sent.");
	}
}


class ContactInput extends React.Component {
	constructor(props){
    super(props);
		this.state = {
			inputValue: "",
			errorMsg: ""
		};
		this.updateInputValue = this.updateInputValue.bind(this);
  }

	updateInputValue(event) {
		const newValue = event.target.value;
		const propertyName = this.props.propertyName;
		this.setState({ inputValue: newValue });
		this.props.updateCallback(propertyName, newValue);
	}

	render() {
		const classNames = classnames("contact-input", this.props.addClasses);
		const title = this.props.title;
		const placeholder = this.props.placeholder;
		const onErrorFn = this.props.onError;
		return (
			<div className={classNames}>
				<label>{title}</label>
				<input placeholder={placeholder}
							value={this.state.inputValue}
							onChange={this.updateInputValue}/>
				<label>{this.state.errorMsg}</label>
			</div>
		);
	}
}


/* Helper functions */

function generateEmailMessage(info) {
	if (!info) {
		return "<blank email message>";
	}
	return [
		"Student: " + info.studentFirstName + " " + info.studentLastName,
		"Age: " + info.studentAge,
		"Grade: " + info.studentGrade,
		"School: " + info.studentSchool,
		"Phone: " + info.studentPhone,
		"Email: " + info.studentEmail,
		"",
		"Guardian: " + info.guardFirstName + " " + info.guardLastName,
		"Phone: " + info.guardPhone,
		"Email: " + info.guardEmail,
		"Interested Programs: " + info.interestedPrograms,
		"Additional Info: " + info.additionalText
	].join("\n");
}

function sendEmail(templateId, senderEmail, receiverEmail, emailMessage,
	successFunc, errorFunc) {
  window.emailjs.send(
    'mailgun',
    templateId,
    {
      senderEmail,
      receiverEmail,
      emailMessage
    }
	).then(res => {
    console.log("Email success!");
		if (successFunc) {
			successFunc();
		}
  }).catch(err => {
		console.error('Failed to send email. Error: ', err);
		if (errorFunc) {
			errorFunc();
		}
	});
}
