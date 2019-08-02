'use strict';
require('./afhForm.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  EmailModal,
  STATE_NONE,
  STATE_EMPTY,
  STATE_LOADING,
  STATE_SUCCESS,
  STATE_FAIL
} from '../modals/emailModal.js';
import { sendEmail, sendTestEmail } from '../repos/emailRepo.js';
import { FormInput } from '../forms/formInput.js';
import {
  EmailCheck,
  NameCheck,
  PhoneCheck
} from '../forms/formInputChecks.js';
import {
  getKeyValue,
  getProgramsBySemester
} from '../repos/mainRepo.js';
import { Modal } from '../modals/modal.js';

import { keys, remove } from 'lodash';
const classnames = require('classnames');

export class AfhForm extends React.Component {
	constructor(props) {
    super(props);
		this.state = {
      submitState: STATE_NONE,
			studentFirstName: "",
			studentLastName: "",
			studentPhone: "",
			studentEmail: "",
			programs: {},
			additionalText: "",
      generatedEmail: null
		};

    var currentSemester = getKeyValue("current_semester_id");
    this.semesterPrograms = getProgramsBySemester()[currentSemester];

    this.handleSubmit = this.handleSubmit.bind(this);
    this.onSubmitSuccess = this.onSubmitSuccess.bind(this);
    this.onSubmitFail = this.onSubmitFail.bind(this);

		this.getInputInfo = this.getInputInfo.bind(this);
    this.checkAllInputs = this.checkAllInputs.bind(this);
		this.updateCb = this.updateCb.bind(this);
    this.updatePrograms = this.updatePrograms.bind(this);
		this.updateTextArea = this.updateTextArea.bind(this);
  }

	updateCb(propertyName, newValue) {
		var obj = {};
		obj[propertyName] = newValue;
		this.setState(obj);
	}

  updatePrograms(event) {
    var programId = event.target.value;
    var checked = event.target.checked;

    var newPrograms = this.state.programs || {};
    if (checked) {
      newPrograms[programId] = true;
    } else {
      newPrograms[programId] = undefined;
    }
    this.setState({programs: newPrograms});
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

    const programs = this.state.programs;
    const onChange = this.updatePrograms;
    const list = this.semesterPrograms.map(function(p, index) {
      var inputClasses = classnames({
        highlight: programs[p.programId]
      });
      return (
        <li key={index} className={inputClasses}>
          <input type="checkbox" name="program" value={p.programId}
            onChange={onChange}/>
            {p.title}
        </li>
      );
    });

		return (
      <div id="afh-form">
        <Modal content={modalContent}
                show={showModal}
                persistent={true}/>
        <div className="section input">
          <h2>Student Information</h2>
          <div className="afh-input-container">
            <FormInput addClasses="student-fname" title="First Name" propertyName="studentFirstName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
            <FormInput addClasses="student-lname" title="Last Name" propertyName="studentLastName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
          </div>
          <div className="afh-input-container">
            <FormInput addClasses="student-phone" title="Phone" propertyName="studentPhone"
                  onUpdate={this.updateCb} validator={PhoneCheck}/>
            <FormInput addClasses="student-email" title="Email" propertyName="studentEmail"
                  onUpdate={this.updateCb} validator={EmailCheck}/>
          </div>
        </div>

        <div className="section programs">
					<h2>Program</h2>
          <p>What programs do you need assistance with?</p>
          <ul>{list}</ul>
        </div>

				<div className="section additional">
					<h2>Additional Information</h2>
					<textarea onChange={this.updateTextArea}
                    placeholder="Please specify any topics or questions you may have."/>
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
			studentPhone: this.state.studentPhone,
			studentEmail: this.state.studentEmail,
			programs: this.state.programs,
			additionalText: this.state.additionalText
		};
	}

  checkAllInputs() {
    return NameCheck.validate(this.state.studentFirstName)
                    && NameCheck.validate(this.state.studentLastName)
                    && PhoneCheck.validate(this.state.studentPhone)
                    && EmailCheck.validate(this.state.studentEmail)
                    && keys(this.state.programs).length > 0;
  }

	handleSubmit(event) {
    event.preventDefault();

		const template = "mathnavigatorwebsitecontact";
		const receiverEmail = "andymathnavigator@gmail.com";
		const senderEmail = "anonymous@andymathnavigator.com";

		const emailMessage = generateEmailMessage(this.getInputInfo());
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
		"Phone: " + info.studentPhone,
		"Email: " + info.studentEmail,
    "",
		"Programs: " + JSON.stringify(info.programs),
		"Additional Info: " + info.additionalText
	].join("\n");
}
