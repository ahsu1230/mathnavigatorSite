'use strict';
require('./programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Promise } from 'bluebird';
import {
	getClassesBySemester,
	getClassesByProgramAndSemester,
	getProgramsBySemesters
} from '../repos/apiRepo.js';
import { keys } from 'lodash';
import { ProgramClassModal } from './programClassModal.js';
import { Modal } from '../modals/modal.js';
import { history } from '../app/history.js';

export class ProgramsPage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			semesterIds: [], 				// list of semesterIds
			semesters: {}, 					// semesterId => semesterObj
			programsBySemester: {} 	// semesterId => list of programObjs
		};
	}


	componentDidMount() {
		getProgramsBySemesters().then(data => {
			this.setState({
				semesterIds: keys(data.semesterMap),
				semesterMap: data.semesterMap,
				programsBySemester: data.programSemesterMap
			})
		});

		if (process.env.NODE_ENV === 'production') {
			mixpanel.track("programs");
		}
	}

	render() {
		const sections = this.state.semesterIds.map((semesterId, index) => {
			var semester = this.state.semesterMap[semesterId];
			var programs = this.state.programsBySemester[semesterId];
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

export class ProgramSection extends React.Component {
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

export class ProgramCard extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			classes: [],
			showModal: false
		}
		this.handleClick = this.handleClick.bind(this);
		this.dismissModal = this.dismissModal.bind(this);
	}

	componentDidMount() {
		var program = this.props.program;
		var semester = this.props.semester;
		var classes = getClassesByProgramAndSemester(program.programId,
			semester.semesterId).then(classes => {
				this.setState({ classes: classes });
			});
	}

	handleClick() {
		var classes = this.state.classes;
		var program = this.props.program;
		var programId = program.programId;

		if (classes.length == 0) {
			return; // do nothing
		} else if (classes.length == 1) {
			var slug = classes[0].key;
			// history.push("/class/" + slug); // Use with BrowserRouter
			window.location.hash = "/class/" + slug; // Use with HashRouter
		} else if (classes.length > 1) {
			this.setState({ showModal: true });
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
		const classes = this.state.classes;

		const grades = "Grades " + program.grade1 + " - " + program.grade2;
		var modalDiv;
		if (classes.length > 1) {
			const modalContent = <ProgramClassModal programObj={program}
															classList={classes}
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
