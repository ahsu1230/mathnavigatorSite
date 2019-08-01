'use strict';
require('./../styl/emailModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
const classnames = require('classnames');

const srcMail = require("../assets/mail_green.svg");
const srcError = require("../assets/error_orange.svg");
const srcCheck = require("../assets/checkmark_white.svg");

export const STATE_FAIL = -1;
export const STATE_NONE = 0;
export const STATE_EMPTY = 1; // hack for smoothly animating between states
export const STATE_LOADING = 2;
export const STATE_SUCCESS = 3;


export class EmailModal extends React.Component {
  render() {
    const loadingState = this.props.loadingState || STATE_EMPTY;

    const modalClassName = classnames({
      loading: loadingState == STATE_LOADING,
      success: loadingState == STATE_SUCCESS,
      fail: loadingState == STATE_FAIL
    });
    var showLoading = false;
    var showSuccess = false;
    var showFail = false;
    switch(loadingState) {
      case STATE_LOADING:
        showLoading = true;
        break;
      case STATE_SUCCESS:
        showSuccess = true;
        break;
      case STATE_FAIL:
        showFail = true;
        break;
    }
    return (
      <div id="email-modal" className={modalClassName}>
        <SubmitLoading show={showLoading}/>
        <SubmitSuccess show={showSuccess}/>
        <SubmitFail show={showFail} failText={this.props.failText}/>
      </div>
    );
  }
}

class SubmitLoading extends React.Component {
  render() {
    const containerClassNames = classnames("container loading", {
      show: this.props.show
    });
    return (
      <div className={containerClassNames}>
        <h1>Sending...</h1>
        <div className="content">
          <div className="dots-container">
            <div className="dots">
              <div className="dot dot1"></div>
              <div className="dot dot2"></div>
              <div className="dot dot3"></div>
            </div>
            <img src={srcMail}/>
          </div>
        </div>
      </div>
    );
  }
}

class SubmitSuccess extends React.Component {
  render() {
    const containerClassNames = classnames("container success", {
      show: this.props.show
    });
    return (
      <div className={containerClassNames}>
        <h1>Success</h1>
        <div className="content">
          <div className="img-wrapper">
            <img src={srcCheck}/>
          </div>
          <p>
            Thank you for contacting us.<br/>
            You will hear from us soon!<br/>
            <Link to="/">Back to Home</Link>
          </p>
        </div>
      </div>
    );
  }
}

class SubmitFail extends React.Component {
  render() {
    const containerClassNames = classnames("container fail", {
      show: this.props.show
    });
    const textAreaValue = this.props.failText || "";
    const handleOnChange = function() {};

    return (
      <div className={containerClassNames}>
        <h1>Email request failed...</h1>
        <div className="content">
          <img src={srcError}/>
          <p>
            Please copy the following text,<br/>
            paste it into a new email message,<br/>
            and send it to <b>andymathnavigator@gmail.com</b>.<br/>
            We apologize for any inconveniences.
          </p>
          <textarea value={textAreaValue} onChange={handleOnChange}/>
          <Link to="/">Back to Home</Link>
        </div>
      </div>
    );
  }
}
