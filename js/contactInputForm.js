'use strict';
require('./../styl/contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactInput } from './contactInput.js';
import { ContactInterestSection } from './contactInterest.js';
import {
  ContactSubmitModal,
  STATE_NONE,
  STATE_EMPTY,
  STATE_LOADING,
  STATE_SUCCESS,
  STATE_FAIL
} from './contactSubmitModal.js';
import { Modal } from './modal.js';
const classnames = require('classnames');

export class ContactInputForm extends React.Component {
	constructor(props) {
    super(props);
		this.state = {
      submitState: STATE_NONE,
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
			additionalText: "",
      generatedEmail: null
		};

    this.handleSubmit = this.handleSubmit.bind(this);
    this.onSubmitSuccess = this.onSubmitSuccess.bind(this);
    this.onSubmitFail = this.onSubmitFail.bind(this);
    this.dismissModal = this.dismissModal.bind(this);

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
    const submitState = this.state.submitState;
    const modalContent = <ContactSubmitModal
                            loadingState={submitState}
                            failText={this.state.generatedEmail}/>;
    const showModal = submitState != STATE_NONE;
		return (
      <div>
        <Modal content={modalContent}
                show={showModal}
                persistent={true}
                onDismiss={this.dismissModal}/>
        <div className="section input">
          <h2>Student Information</h2>
          <div className="contact-input-container">
            <ContactInput addClasses="student-fname" title="First Name" propertyName="studentFirstName" onUpdate={this.updateCb}/>
            <ContactInput addClasses="student-lname" title="Last Name" propertyName="studentLastName" onUpdate={this.updateCb}/>
          </div>
          <div className="contact-input-container">
            <ContactInput addClasses="student-age" title="Age" propertyName="studentAge" onUpdate={this.updateCb}/>
            <ContactInput addClasses="student-grade" title="Grade" propertyName="studentGrade" onUpdate={this.updateCb}/>
            <ContactInput addClasses="student-school" title="School" propertyName="studentSchool" onUpdate={this.updateCb}/>
          </div>
          <div className="contact-input-container">
            <ContactInput addClasses="student-phone" title="Phone" propertyName="studentPhone" onUpdate={this.updateCb}/>
            <ContactInput addClasses="student-email" title="Email" propertyName="studentEmail" onUpdate={this.updateCb}/>
          </div>

          <h2>Guardian Information</h2>
          <div className="contact-input-container">
            <ContactInput addClasses="guard-fname" title="First Name" propertyName="guardFirstName" onUpdate={this.updateCb}/>
            <ContactInput addClasses="guard-lname" title="Last Name" propertyName="guardLastName" onUpdate={this.updateCb}/>
          </div>
          <div className="contact-input-container">
            <ContactInput addClasses="guard-phone" title="Phone" propertyName="guardPhone" onUpdate={this.updateCb}/>
            <ContactInput addClasses="guard-email" title="Email" propertyName="guardEmail" onUpdate={this.updateCb}/>
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
		const emailMessage = generateEmailMessage(inputInfo);

    console.log("Sending email... " + emailMessage);
    this.setState({
      submitState: STATE_LOADING,
      generatedEmail: emailMessage
    });

		sendTestEmail(this.onSubmitSuccess, this.onSubmitFail, true);
    // sendEmail(
    // 	template,
    // 	senderEmail,
    // 	receiverEmail,
    // 	emailMessage
    // );
	}

	onSubmitSuccess() {
    setTimeout(() => {
      this.setState({ submitState: STATE_EMPTY });
      setTimeout(() => {
        console.log("Email success!");
        this.setState({ submitState: STATE_SUCCESS });
      }, 400);
    }, 3600);

	}

	onSubmitFail() {
    setTimeout(() => {
      this.setState({ submitState: STATE_EMPTY });
      setTimeout(() => {
        console.log("Email fail!");
        this.setState({ submitState: STATE_FAIL });
      }, 400);
    }, 3600);
	}

  dismissModal() {
    console.log("Dismiss modal");
    this.setState({
      submitState: STATE_LOADING
    });
  }
}

/* Helper functions */

function generateEmailMessage(info) {
	if (!info) {
		return null;
	}
	return [
    "To Math Navigator,",
    "",
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
    "",
		"Interested Programs: " + info.interestedPrograms,
		"Additional Info: " + info.additionalText
	].join("\n");
}

function sendEmail(templateId, senderEmail, receiverEmail, emailMessage,
	onSuccess, onFail) {
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
		if (onSuccess) {
			onSuccess();
		}
  }).catch(err => {
		console.error('Failed to send email. Error: ', err);
		if (onFail) {
			onFail();
		}
	});
}

function sendTestEmail(onSuccess, onFail, success) {
  if (success && onSuccess) {
    onSuccess();
  } else if (onFail) {
    onFail();
  }
}
