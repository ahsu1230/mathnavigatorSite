'use strict';
require('./../styl/programs.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { locationMap } from './variables/initPrograms.js';

export class ProgramsPage extends React.Component {
	render() {
		var progAvail = {
			title: "Available Programs",
			programs: [

			]};
			var progSoon = {
				title: "Coming Soon",
				programs: []
			};
		return (
      <div id="view-program">
        <div id="view-program-container">
					<ProgramSection title={progAvail.title} programs={progAvail.programs}/>
					<ProgramSection title={progSoon.title} programs={progSoon.programs}/>
        </div>
      </div>
		);
	}
}


class ProgramSection extends React.Component {
	render() {
		var title = this.props.title;
		var programs = this.props.programs;

		var msg;
		if (programs || programs.length == 0) {
			msg = "No Programs available";
		} else {
			msg = programs.length + " Programs Available!";
		}
		return (
			<div className="section">
				<h2 className="section-title">{title}</h2>
				<h3>{msg}</h3>
			</div>
		);
	}
}
