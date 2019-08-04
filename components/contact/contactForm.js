'use strict';
require('./contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { ContactInterestSection } from './contactInterest.js';
import {
  EmailModal,
  STATE_NONE,
  STATE_EMPTY,
  STATE_LOADING,
  STATE_SUCCESS,
  STATE_FAIL
} from '../modals/emailModal.js';
import { sendEmail, sendTestEmail } from '../repos/emailRepo.js';
import {
  EmailCheck,
  SchoolCheck,
  GradeCheck,
  NameCheck,
  PhoneCheck
} from '../forms/formInputChecks.js';
import { FormInput } from '../forms/formInput.js';
import { pull } from 'lodash';
import { Modal } from '../modals/modal.js';
const classnames = require('classnames');

export class ContactForm extends React.Component {
	constructor(props) {
    super(props);

    var interested = [];
    var parsed = parseQuery();
    if (parsed && parsed.interest) {
      interested.push(parsed.interest);
    }

		this.state = {
      submitState: STATE_NONE,
			studentFirstName: "",
			studentLastName: "",
			studentGrade: 0,
			studentSchool: "",
			studentPhone: "",
			studentEmail: "",
			guardFirstName: "",
			guardLastName: "",
			guardPhone: "",
			guardEmail: "",
			interestedPrograms: interested,
			additionalText: "",
      generatedEmail: null
		};

    this.handleSubmit = this.handleSubmit.bind(this);
    this.onSubmitSuccess = this.onSubmitSuccess.bind(this);
    this.onSubmitFail = this.onSubmitFail.bind(this);
    this.dismissModal = this.dismissModal.bind(this);

		this.getInputInfo = this.getInputInfo.bind(this);
		this.updateCb = this.updateCb.bind(this);
    this.updateInterested = this.updateInterested.bind(this);
		this.updateTextArea = this.updateTextArea.bind(this);

    this.checkAllInputs = this.checkAllInputs.bind(this);
  }

	updateCb(propertyName, newValue) {
		var obj = {};
		obj[propertyName] = newValue;
		this.setState(obj);
	}

  updateInterested(classKey, selected) {
    var interestedList = this.state.interestedPrograms;
    if (selected) {
      interestedList.push(classKey);
    } else {
      interestedList = pull(interestedList, classKey);
    }
    this.setState({
      interestedPrograms: interestedList
    });
  }

	updateTextArea(event) {
		this.updateCb("additionalText", event.target.value);
	}

	render() {
    const submitState = this.state.submitState;
    const modalContent = <EmailModal
                            loadingState={submitState}
                            failText={this.state.generatedEmail}/>;
    const showModal = submitState != STATE_NONE;

    const formCompleted = this.checkAllInputs();
    const submitBtnClass = classnames({active: formCompleted});
    const onHandleSubmit = formCompleted ? this.handleSubmit : undefined;

		return (
      <div>
        <Modal content={modalContent}
                show={showModal}
                persistent={true}
                onDismiss={this.dismissModal}/>
        <div className="section input">
          <h2>Student Information</h2>
          <div className="contact-input-container">
            <FormInput addClasses="student-fname" title="First Name" propertyName="studentFirstName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
            <FormInput addClasses="student-lname" title="Last Name" propertyName="studentLastName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
          </div>
          <div className="contact-input-container">
            <FormInput addClasses="student-grade" title="Grade" propertyName="studentGrade"
                  onUpdate={this.updateCb} validator={GradeCheck}/>
            <FormInput addClasses="student-school" title="School" propertyName="studentSchool"
                  onUpdate={this.updateCb} validator={SchoolCheck}/>
          </div>
          <div className="contact-input-container">
            <FormInput addClasses="student-phone" title="Phone" propertyName="studentPhone"
                  onUpdate={this.updateCb} validator={PhoneCheck}/>
            <FormInput addClasses="student-email" title="Email" propertyName="studentEmail"
                  onUpdate={this.updateCb} validator={EmailCheck}/>
          </div>

          <h2>Guardian Information</h2>
          <div className="contact-input-container">
            <FormInput addClasses="guard-fname" title="First Name" propertyName="guardFirstName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
            <FormInput addClasses="guard-lname" title="Last Name" propertyName="guardLastName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
          </div>
          <div className="contact-input-container">
            <FormInput addClasses="guard-phone" title="Phone" propertyName="guardPhone"
                  onUpdate={this.updateCb} validator={PhoneCheck}/>
            <FormInput addClasses="guard-email" title="Email" propertyName="guardEmail"
                  onUpdate={this.updateCb} validator={EmailCheck}/>
          </div>
        </div>
				<div className="section interested">
					<ContactInterestSection interested={this.state.interestedPrograms}
                                  onUpdate={this.updateInterested}/>
				</div>

				<div className="section additional">
					<h2>Additional Information</h2>
					<textarea onChange={this.updateTextArea} placeholder="(Optional)"/>
        </div>

        <div className="section submit">
          <div className="submit-container">
            <p>
              Information will be sent to:<br/>
              <a>andymathnavigator@gmail.com</a>
            </p>
            <button className={submitBtnClass} onClick={onHandleSubmit}>
              Submit
            </button>
          </div>
        </div>
      </div>
		);
	}

	getInputInfo() {
		return {
			studentFirstName: this.state.studentFirstName,
			studentLastName: this.state.studentLastName,
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

  checkAllInputs() {
    return NameCheck.validate(this.state.studentFirstName)
                    // && NameCheck.validate(this.state.studentLastName)
                    // && GradeCheck.validate(this.state.studentGrade)
                    // && SchoolCheck.validate(this.state.studentSchool)
                    // && PhoneCheck.validate(this.state.studentPhone)
                    // && EmailCheck.validate(this.state.studentEmail)
                    // && NameCheck.validate(this.state.guardFirstName)
                    // && NameCheck.validate(this.state.guardLastName)
                    // && PhoneCheck.validate(this.state.guardPhone)
                    // && EmailCheck.validate(this.state.guardEmail)
                    && this.state.interestedPrograms.length > 0;
  }

	handleSubmit(event) {
    event.preventDefault();

		var inputInfo = this.getInputInfo();
		const emailMessage = generateEmailMessage(inputInfo);

    console.log("Sending email... " + emailMessage);
    this.setState({
      submitState: STATE_LOADING,
      generatedEmail: emailMessage
    });

    sendEmail(emailMessage, this.onSubmitSuccess, this.onSubmitFail);
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

function parseQuery() {
  var hash = window.location.hash;
  var i = hash.indexOf("?");
  var parsed = {};
  if (i > 0) {
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

function generateEmailMessage(info) {
	if (!info) {
		return null;
	}
	return [
    "<html>",
    "<body>",
    "<h1>To Math Navigator,</h1>",
    "<h2>Contact Us Page</h2>",
    "<h3>Student: " + info.studentFirstName + "	&nbsp; " + info.studentLastName + "</h3>",
    "<h3>Grade:" + info.studentGrade + "</h3>",
    "<h3>School: " + info.studentSchool + "</h3>",
    "<h3>Phone: " + info.studentPhone + "</h3>",
    "<h3>Email: " + info.studentEmail + "</h3>",
    "<br/>",
    "<h3>Guardian: " + info.guardFirstName + "	&nbsp; " + info.guardLastName + "</h3>",
    "<h3>Phone: " + info.guardPhone + "</h3>",
    "<h3>Email: " + info.guardEmail + "</h3>",
    "<br/>",
    "<p>Interested Programs: " + JSON.stringify(info.interestedPrograms) + "</p>",
    "<p>Additional Info: " + info.additionalText + "</p>",
    "</body>",
    "</html>"
	].join("\n");
}
