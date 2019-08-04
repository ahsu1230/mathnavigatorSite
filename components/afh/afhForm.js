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
  NameCheck,
} from '../forms/formInputChecks.js';
import {
  getAFH,
  getLocation,
  getProgramsBySemester
} from '../repos/mainRepo.js';
import { Modal } from '../modals/modal.js';

import { find, remove } from 'lodash';
const classnames = require('classnames');

export class AfhForm extends React.Component {
	constructor(props) {
    super(props);
		this.state = {
      submitState: STATE_NONE,
			studentFirstName: "",
			studentLastName: "",
      targetSessions: [],
			additionalText: "",
      generatedEmail: null
		};

    this.afh = getAFH();
    this.afh.map((afh) => afh.location = getLocation(afh.locationId));

    this.handleSubmit = this.handleSubmit.bind(this);
    this.onSubmitSuccess = this.onSubmitSuccess.bind(this);
    this.onSubmitFail = this.onSubmitFail.bind(this);

		this.getInputInfo = this.getInputInfo.bind(this);
    this.checkAllInputs = this.checkAllInputs.bind(this);
		this.updateCb = this.updateCb.bind(this);
    this.updateSessions = this.updateSessions.bind(this);
		this.updateTextArea = this.updateTextArea.bind(this);
  }

	updateCb(propertyName, newValue) {
		var obj = {};
		obj[propertyName] = newValue;
		this.setState(obj);
	}

  updateSessions(event) {
    var afhId = parseInt(event.target.value, 10) || -1;
    var checked = event.target.checked;

    var sessions = this.state.targetSessions || [];
    if (checked) {
      sessions.push(afhId);
    } else {
      remove(sessions, function(a) { return a === afhId;});
    }
    this.setState({targetSessions: sessions});
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

    const onSessionCheck = this.updateSessions;
    const sessions = this.afh.map((afh, index) =>
      <AfhSession key={index} afh={afh} onSelect={onSessionCheck}/>
    );

    const formCompleted = this.checkAllInputs();
    const submitBtnClass = classnames({active: formCompleted});
    const onHandleSubmit = formCompleted ? this.handleSubmit : undefined;

		return (
      <div id="afh-form">
        <Modal content={modalContent}
                show={showModal}
                persistent={true}/>

        <div className="section sessions">
			    <h2>Which session(s) would you like to attend?</h2>
          {sessions}
        </div>

        <div className="section input">
          <h2>Student Information</h2>
          <div className="afh-input-container">
            <FormInput addClasses="student-fname" title="First Name" propertyName="studentFirstName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
            <FormInput addClasses="student-lname" title="Last Name" propertyName="studentLastName"
                  onUpdate={this.updateCb} validator={NameCheck}/>
          </div>
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

        <div className="section errors">
          <AfhErrorReminder formState={this.state}/>
        </div>
      </div>
		);
	}

	getInputInfo() {
		return {
			studentFirstName: this.state.studentFirstName,
			studentLastName: this.state.studentLastName,
      targetSessions: this.state.targetSessions,
			additionalText: this.state.additionalText
		};
	}

  checkAllInputs() {
    return validateName(this.state.studentFirstName, this.state.studentLastName) &&
            validateSessions(this.state.targetSessions);
  }

	handleSubmit(event) {
    event.preventDefault();

		const emailMessage = generateEmailMessage(this.getInputInfo());
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
}


class AfhSession extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      checked: this.props.checked || false
    }
    this.handleSelected = this.handleSelected.bind(this);
  }

  handleSelected(event) {
    this.setState({ checked: !this.state.checked});
    if (this.props.onSelect) {
      this.props.onSelect(event);
    }
  }

  render() {
    const afh = this.props.afh;
    const location = afh.location;
    const textClasses = classnames("afh-session-text", {
      highlight: this.state.checked
    });
    const title = afh.notes || "All math topics";

    return (
      <div className="afh-session">
        <input type="checkbox" value={afh.afhId}
            onChange={this.handleSelected}/>
        <div className={textClasses}>
          <h4>{title}</h4>
          <div>{afh.date} {afh.time}</div>
          <div>{location.name}</div>
          <div>{location.address3}</div>
        </div>
      </div>
    );
  }
}

class AfhErrorReminder extends React.Component {
  render() {
    const formState = this.props.formState;
    var errorNotif;
    var errorName;
    var errorSessions;

    if (!validateName(formState.studentFirstName, formState.studentLastName)) {
      errorName = <li>Please fill your name.</li>;
    }
    if (!validateSessions(formState.targetSessions)) {
      errorSessions = <li>Pick at least one session to attend.</li>;
    }
    if (errorName || errorSessions) {
      errorNotif = <li>Please correctly fill the form in order to submit!</li>;
    }

    return (
      <ul>
        {errorNotif}
        {errorSessions}
        {errorName}
      </ul>
    );
  }
}

/* Helper functions */
function validateName(firstName, lastName) {
  return NameCheck.validate(firstName) && NameCheck.validate(lastName);
}

function validateSessions(sessions) {
  return sessions.length > 0;
}

function generateEmailMessage(info) {
	if (!info) {
		return null;
	}
  var afhList = getAFH();
  const sessions = info.targetSessions.map(function(afhId) {
    var afhObj = find(afhList, {afhId: afhId});
    return { session: afhObj.date + " " + afhObj.time };
  });

  return [
    "<html>",
    "<body>",
    "<h1>To Math Navigator,</h1>",
    "<h2>Ask For Help!</h2>",
    "<h3>Student: " + info.studentFirstName + "	&nbsp; " + info.studentLastName + "</h3>",
    "<p>Coming to: " + JSON.stringify(sessions) + "</p>",
    "<p>Additional Info: " + info.additionalText + "</p>",
    "</body>",
    "</html>"
	].join("\n");
}
