'use strict';
require('./../styl/contactInterestModal.styl');
import _ from 'lodash';
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { Modal } from './modal.js';
import { getProgramClass, getAvailableClasses } from './repos/mainRepo.js';
const classnames = require('classnames');

export class ContactInterestModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showModal: false,
      interested: []
    };

    this.onSelectProgram = this.onSelectProgram.bind(this);
    this.onToggleShowModal = this.onToggleShowModal.bind(this);
  }

  onToggleShowModal() {
    const show = this.state.showModal;
    var newShow = !show;
    this.setState({
      showModal: newShow
    });
  }

  onSelectProgram(classKey, selected) {
    var interestedList = this.state.interested;
    if (selected) {
      interestedList.push(classKey);
    } else {
      interestedList = _.pull(interestedList, classKey);
    }
    this.setState({
      interested: interestedList
    });
  }

  render() {
    const interestedClasses = this.state.interested;
    const interestedSection = generateInterested(interestedClasses, this.onToggleShowModal);
    const interestModal = <InterestModal
                  onSelectProgram={this.onSelectProgram}
                  interested={interestedClasses}
                  onDismiss={this.onToggleShowModal}/>;
    return (
      <div>
        <h2>Interested Programs</h2>
        {interestedSection}
        <Modal content={interestModal}
          show={this.state.showModal}
          onDismiss={this.onToggleShowModal}>
        </Modal>
      </div>
    );
  }
}

class InterestModal extends React.Component {
  render() {
    const interestedClasses = this.props.interested;
    const numClasses = interestedClasses.length;

    var classesPair = getAvailableClasses();
    var classesAvail = classesPair.available;
    var classesSoon = classesPair.soon;

    var selectedLineClassNames = classnames("selected-line", {
      highlight: numClasses > 0
    });
    var selectedLineText = numClasses;
    selectedLineText += (numClasses == 1 ? " class " : " classes ");
    selectedLineText += "selected";

    const listAvail = classesAvail.map((classObj, index) =>
      <InterestList key={index} classObj={classObj} onSelect={this.props.onSelectProgram}/>
    );
    const listSoon = classesSoon.map((classObj, index) =>
      <InterestList key={index} classObj={classObj} onSelect={this.props.onSelectProgram}/>
    );

    return (
      <div id="interest-modal">
        <h1>Interested Programs</h1>
        <div id="interest-headers">
          <div className="header class-name">Class</div>
          <div className="header times">Times</div>
          <div className="header start-date">Starting Date</div>
        </div>
        <div id="interest-view">
          <ul>
            <h2>Classes Available</h2>
            {listAvail}
            <h2>Classes Coming Soon</h2>
            {listSoon}
          </ul>
        </div>
        <div className={selectedLineClassNames}>{selectedLineText}</div>
        <button className="btn-done" onClick={this.props.onDismiss}>Done</button>
      </div>
    );
  }
}

class InterestList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      highlighted: false
    }
    this.handleSelected = this.handleSelected.bind(this);
  }

  handleSelected(classKey) {
    var selected = this.state.highlighted;
    var newSelected = !selected;
    this.setState({
      highlighted: newSelected
    });
    if (this.props.onSelect) {
      var classKey = this.props.classObj.key;
      this.props.onSelect(classKey, newSelected);
    }
  }

  render() {
    const classObj = this.props.classObj;
    const className = classObj.fullClassName;
    const startingDate = classObj.startDate;

    const liClassNames = classnames("", {
      "highlight": this.state.highlighted
    });

    const times = classObj.times.map((time, index) =>
      <div key={index}>{time}</div>
    );

    return (
      <li className={liClassNames}>
        <input type="checkbox" name="interest" value="program" onClick={this.handleSelected}/>
        <div className="class-name">{className}</div>
        <div className="times">{times}</div>
        <div className="start-date">{startingDate}</div>
      </li>
    );
  }
}

/* Helper Functions */
function generateInterested(interestedList, onClick) {
  var nonEmpty = interestedList.length > 0;
  var interestedButtonText = nonEmpty ? "Select More Programs..." : "Select Programs...";
  if (nonEmpty) {
    var numClasses = interestedList.length;
    var numClassText = numClasses == 1 ? numClasses + " class." : numClasses + " classes."

    const selectedText = interestedList.map(function(classKey, index) {
      var pair = getProgramClass(classKey);
      var programTitle = pair.programObj.title;
      var className = pair.classObj.className;
      var fullName = className ? programTitle + " " + className : programTitle;
      var url = "/class/" + classKey;
      return (
        <Link key={index} to={url}>{fullName}</Link>
      );
    });

    return (
      <div id="contact-section-interested">
        <p>
          You have selected {numClassText}<br/>
          {selectedText}
        </p>
        <button className="inverted" onClick={onClick}>
          {interestedButtonText}
        </button>
      </div>
    );
  } else {
    return (
      <button className="inverted" onClick={onClick}>
        {interestedButtonText}
      </button>
    );
  }
}
