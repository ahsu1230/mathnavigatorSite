'use strict';
require('./../styl/programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { getAvailablePrograms, getClasses } from './repos/programRepo.js';

export class ProgramsPage extends React.Component {
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
	determineOnClick(programId) {
		var classes = getClasses(programId);
		var slug = classes[0].key; // For now, just grab first class
		window.location.hash = "/class/" + slug;
	}

	render() {
		var title = this.props.title;
		var programs = this.props.programs;

		programs.forEach(function(program) {
			program.gradesText = "Grades " + program.grade1 + " - " + program.grade2;
		});

		const cards = programs.map((program) =>
      <ProgramCard key={program.programId} // not exposed by reactJs
				programId={program.programId}
				title={program.title}
				grades={program.gradesText}
				onClick={this.determineOnClick}
			/>
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
	render() {
		return (
			<div className="program-card" onClick={() =>
				this.props.onClick(this.props.programId)
			}>
				<h2>{this.props.title}</h2>
				<h3>{this.props.grades}</h3>
					<button>View</button>
			</div>
		);
	}
}
