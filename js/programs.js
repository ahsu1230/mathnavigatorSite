'use strict';
require('./../styl/programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
	getAvailablePrograms,
	getClasses
} from './repos/mainRepo.js';
import { ProgramClassModal } from './programClassModal.js';
import { Modal } from './modal.js';

export class ProgramsPage extends React.Component {
	componentDidMount() {
	  window.scrollTo(0, 0);
	}

	render() {
		var programs = getAvailablePrograms();
		var programsAvail = Object.values(programs.available);
		var programsSoon = Object.values(programs.soon);

		return (
      <div id="view-program">
        <div id="view-program-container">
					<ProgramSection title={"Available Programs"} programs={programsAvail}/>
					<ProgramSection title={"Coming Soon"} programs={programsSoon}/>
        </div>
      </div>
		);
	}
}

class ProgramSection extends React.Component {
	render() {
		var title = this.props.title;
		var programs = this.props.programs;

		const cards = programs.map((program, index) =>
      <ProgramCard key={index} program={program}/>
    );

		return (
			<div className="section">
				<h2 className="section-title">{title}</h2>
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
			this.setState({
				showModal: true
			});
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
		const programId = program.programId;
		this.classes = getClasses(programId);

		const grades = "Grades " + program.grade1 + " - " + program.grade2;
		var modalDiv;
		if (this.classes.length > 1) {
			const modalContent = <ProgramClassModal programObj={program}/>;
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

		return (
			<div className="program-card-container">
				{/*
					^We require an outer div `program-card-container` because
					of modal styling conflicts?
				*/}
				<div className="program-card" onClick={this.handleClick}>
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
