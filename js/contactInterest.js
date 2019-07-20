'use strict';
require('./../styl/contactInterestModal.styl');
import _ from 'lodash';
import React from 'react';
import ReactDOM from 'react-dom';
import { Modal } from './modal.js';
import { getAllClasses, getAvailableClasses } from './repos/mainRepo.js';
const classnames = require('classnames');

export class ContactInterestModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showModal: false,
      interested: []
    };

    this.onSelectProgram = this.onSelectProgram.bind(this);
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
    var interestedButtonText = "";
    if (interestedClasses.length > 0) {
      interestedButtonText = "Select More Programs...";
    } else {
      interestedButtonText = "Select Programs...";
    }

    const interestModal = <InterestModal
                  onSelectProgram={this.onSelectProgram}
                  interested={interestedClasses}/>;

    return (
      <div>
        <h2>Interested Programs</h2>
        <div className="list-interested">{interestedClasses}</div>
        <button className="inverted">{interestedButtonText}</button>
        <Modal content={interestModal}></Modal>
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
        <button className="btn-done">Done</button>
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
