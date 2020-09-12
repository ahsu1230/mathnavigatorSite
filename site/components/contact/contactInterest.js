'use strict';
require('./contactInterest.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { Modal } from '../modals/modal.js';
import {
  getAllClassesBySemesters,
  getAllProgramsAndClasses
} from '../repos/apiRepo.js';
import { createFullClassName } from '../constants.js';
import { keys } from 'lodash';
const classnames = require('classnames');

export class ContactInterestSection extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showModal: false
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
    if (this.props.onUpdate) {
      this.props.onUpdate(classKey, selected);
    }
  }

  render() {
    const interestedClasses = this.props.interested;
    const modalContent = <InterestModal
                  onSelectProgram={this.onSelectProgram}
                  interested={interestedClasses}
                  onDismiss={this.onToggleShowModal}/>;
    return (
      <div>
        <h2>Interested Programs</h2>
        <InterestSection interestedList={interestedClasses} onClick={this.onToggleShowModal}/>
        <Modal content={modalContent}
          show={this.state.showModal}
          withClose={true}
          onDismiss={this.onToggleShowModal}/>
      </div>
    );
  }
}

class InterestSection extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classProgramMap: {}
    }
  }

  componentDidMount() {
    getAllProgramsAndClasses().then(data => {
      this.setState({classProgramMap: data});
    });
  }

  generateEmptyInterested(buttonText, onButtonClick) {
    return (
      <div id="contact-section-interested">
        <p>
          Please select at least one Program.
        </p>
        <button className="inverted" onClick={onButtonClick}>
          {buttonText}
        </button>
      </div>
    );
  }

  generateNonEmptyInterested(classMap, interestedList, buttonText, onButtonClick) {
    var numClasses = interestedList.length;
    var numClassText = numClasses == 1 ? numClasses + " class." : numClasses + " classes."

    const selectedText = interestedList.map(function(classKey, index) {
      var data = classMap[classKey];
      if (data) {
        var programObj = data.programObj;
        var classObj = data.classObj;
        var fullName = createFullClassName(programObj, classObj);
        var url = "/class/" + classKey;
        return (
          <li key={index}><Link to={url}>{fullName}</Link></li>
        );
      }
    });

    return (
      <div id="contact-section-interested">
        <p>
          You have selected {numClassText}<br/>
          {selectedText}
        </p>
        <button className="inverted" onClick={onButtonClick}>
          {buttonText}
        </button>
      </div>
    );
  }

  render() {
    const interestedList = this.props.interestedList;
    const onClick = this.props.onClick;

    var isEmpty = interestedList.length == 0;
    var interestedButtonText = isEmpty ? "Select Programs..." : "Select More Programs...";
    var comp;
    if (isEmpty) {
      comp = this.generateEmptyInterested(interestedButtonText, onClick);
    } else {
      comp = this.generateNonEmptyInterested(this.state.classProgramMap,
        interestedList, interestedButtonText, onClick);
    }
    return comp;
  }
}

class InterestModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classesBySemester: {},
      semesters: {},
      semesterIds: []
    };
  }

  componentDidMount() {
    getAllClassesBySemesters().then(data => {
      this.setState({
        classesBySemester: data.classSemesterMap,
        semesters: data.semesterMap,
        semesterIds: keys(data.semesterMap)
      });
    });
  }

  render() {
    const interestedClasses = this.props.interested || [];
    var interestClassMap = {};
    interestedClasses.forEach((classKey) => {
      interestClassMap[classKey] = true;
    });
    const numClasses = interestedClasses.length;
    const onSelectProgram = this.props.onSelectProgram;

		const sections = this.state.semesterIds.map((semesterId, index) => {
      var semester = this.state.semesters[semesterId];
			var classes = this.state.classesBySemester[semesterId];
      const list = createInterestItems(classes, interestClassMap, onSelectProgram);
			return (
        <div key={index}>
          <h2>{semester.title}</h2>
          <ul>
            {list}
          </ul>
        </div>
			);
		});

    var selectedLineClassNames = classnames("selected-line", {
      highlight: numClasses > 0
    });
    var selectedLineText = numClasses;
    selectedLineText += (numClasses == 1 ? " class " : " classes ");
    selectedLineText += "selected";

    return (
      <div id="interest-modal">
        <h1>Interested Programs</h1>
        <div id="interest-headers">
          <div className="header class-name">Class</div>
          <div className="header times">Times</div>
          <div className="header start-date">Starting Date</div>
        </div>
        <div id="interest-view" className="use-scrollbar">
          {sections}
        </div>
        <div className={selectedLineClassNames}>{selectedLineText}</div>
        <button className="btn-done" onClick={this.props.onDismiss}>Done</button>
      </div>
    );
  }
}

class InterestItem extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      highlighted: this.props.selected || false
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
    const className = createFullClassName(classObj.programObj, classObj);
    const startingDate = classObj.startDate;
    const isFull = !!classObj.fullState;
    const handleCheck = isFull ? function() {} : this.handleSelected;

    const liClassNames = classnames("", {
      "highlight": this.state.highlighted
    });

    const isFullMessage = isFull ? (
      <div className="warning">
        This class is full. You cannot select this class.
      </div>
    ) : (<div></div>);

    const times = classObj.times.map((time, index) =>
      <div key={index}>{time}</div>
    );

    return (
      <li className={liClassNames}>
        <input type="checkbox" name="interest" value="program"
              onChange={handleCheck}
              checked={this.props.selected}/>
        <div className="class-info">
          <div className="class-name">{className}</div>
          <div className="times">{times}</div>
          <span>Starts on:</span><div className="start-date">{startingDate}</div>
          {isFullMessage}
        </div>
      </li>
    );
  }
}

/* Helper Functions */
function createInterestItems(listClasses, interestedMap, onSelect) {
  return listClasses.map(function(classObj, index) {
    var isSelected = interestedMap[classObj.key] || false;
    return (
      <InterestItem key={index}
                classObj={classObj}
                selected={isSelected}
                onSelect={onSelect}/>
    );
  });
}
