'use strict';
require('./../styl/programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
	getClassesBySemester,
	getClassesByProgramAndSemester,
	getProgramsBySemester,
	getSemester,
	getSemesterIds
} from './repos/mainRepo.js';
import { ProgramClassModal } from './programClassModal.js';
import { Modal } from './modal.js';

export class ProgramsPage extends React.Component {
	componentDidMount() {
	  window.scrollTo(0, 0);
	}

	render() {
		const semesterIds = getSemesterIds();
		const programsBySemesterId = getProgramsBySemester();
		const sections = semesterIds.map(function(semesterId, index) {
			var semester = getSemester(semesterId);
			var programs = programsBySemesterId[semesterId];
			return (
				<ProgramSection key={index} semester={semester} programs={programs}/>
			);
		});

		return (
      <div id="view-program">
        <div id="view-program-container">
					{sections}
        </div>
      </div>
		);
	}
}

class ProgramSection extends React.Component {
	render() {
		const semester = this.props.semester;
		const title = semester.title;
		const programs = this.props.programs;

		const cards = programs.map((program, index) =>
      <ProgramCard key={index} semester={this.props.semester} program={program}/>
    );

		return (
			<div className="section">
				<h1 className="section-title">{title}</h1>
				{cards}
			</div>
		);
	}
}

class ProgramCard extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			showModal: false
		}
		this.handleClick = this.handleClick.bind(this);
		this.dismissModal = this.dismissModal.bind(this);
	}

	handleClick() {
		var classes = this.classes;
		var program = this.props.program;
		var programId = program.programId;

		if (classes.length > 1) {
			this.setState({ showModal: true });
		} else {
			var slug = classes[0].key;
			window.location.hash = "/class/" + slug;
		}
	}

	dismissModal() {
		this.setState({
			showModal: false
		});
	}

	render() {
		const program = this.props.program;
		const semester = this.props.semester;
		this.classes = getClassesByProgramAndSemester(program.programId,
			semester.semesterId);

		const grades = "Grades " + program.grade1 + " - " + program.grade2;
		var modalDiv;
		if (this.classes.length > 1) {
			const modalContent = <ProgramClassModal programObj={program}
															classList={this.classes}
															semester={semester}/>;
			modalDiv = (
				<Modal content={modalContent}
								show={this.state.showModal}
								withClose={true}
								onDismiss={this.dismissModal}/>
			);
		} else {
			modalDiv = (
				<div></div>
			);
		}

		var star = <div></div>;
		if (isImportantProgram(program.programId)) {
			star = (
				<div className="star-container">
					<div className="star-img"></div>
				</div>
			);
		}

		return (
			<div className="program-card-container">
				{/*
					^We require an outer div `program-card-container` because
					of modal styling conflicts?
				*/}
				<div className="program-card" onClick={this.handleClick}>
					{star}
					<div className="program-card-content">
						<h2>{program.title}</h2>
						<h3>{grades}</h3>
						<button>View</button>
					</div>
				</div>
				{modalDiv}
			</div>
		);
	}
}

function isImportantProgram(programId) {
	return programId === "sat1" ||
			programId === "sat2" ||
			programId === "amc8" ||
			programId === "ap_calc";
}
