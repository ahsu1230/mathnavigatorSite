'use strict';
require('./../styl/programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { programMap } from './repos/initPrograms.js';
import { getAvailablePrograms } from './repos/programRepo.js';

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
	render() {
		var title = this.props.title;
		var programs = this.props.programs;

		programs.forEach(function(program) {
			program.gradesText = "Grades " + program.grade1 + " - " + program.grade2;
			program.onClickFunc = undefined;
		});

		const cards = programs.map((program) =>
      <ProgramCard key={program.programId}
				title={program.title}
				grades={program.gradesText}
				onClick={program.onClickFunc}
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
			<div className="program-card">
				<h2>{this.props.title}</h2>
				<h3>{this.props.grades}</h3>
					<button>View</button>
			</div>
		);
	}
}
