'use strict';
require('./programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
	getClassesBySemester,
	getClassesByProgramAndSemester,
	getProgramsBySemester,
	getSemester,
	getSemesterIds
} from '../repos/mainRepo.js';
import { ProgramClassModal } from './programClassModal.js';
import { Modal } from '../modals/modal.js';
import { history } from '../app/history.js';

export class ProgramsPage extends React.Component {
	componentDidMount() {
		if (process.env.NODE_ENV === 'production') {
			mixpanel.track("programs");
		}
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
					<div id="star-legend">
						<div className="star-container">
							<div className="star-img"></div>
						</div>
					  = Featured Programs
					</div>
					{sections}
        </div>
      </div>
		);
	}
}

class ProgramSection extends React.Component {
	render() {
		const programs = this.props.programs;
		const semester = this.props.semester;
		const title = semester.title;

		const cards = programs.map((program, index) =>
      <ProgramCard key={index} semester={semester} program={program}/>
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
		//	history.push("/class/" + slug); // Use with BrowserRouter
			window.location.hash = "/class/" + slug; // Use with HashRouter
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
		if (isFeaturedProgram(program.programId)) {
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

/* Helper functions */
function isFeaturedProgram(programId) {
	return programId === "sat1" ||
			programId === "sat2" ||
			programId === "amc8" ||
			programId === "ap_calc";
}
