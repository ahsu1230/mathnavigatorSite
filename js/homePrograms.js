'use strict';
require('./../styl/home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getNav } from './constants.js';

const srcSat = require('../assets/cb_sat.jpg');
const srcApCalc = require('../assets/ap_calc.jpg');
const srcAmc = require('../assets/amc.jpg');

const programs = [
  {
    title: "SAT",
    description: "Our most popular program. We provide dedicated programs for both the SAT1 Math Section and SAT Subject Test (Math Level 2).",
    imgSrc: srcSat,
    programId: ""
  },
  {
    title: "AP Calculus",
    description: "Get college credits early by acing the AP Calculus Exams. This program is dedicated to both Calculus AB and BC students.",
    imgSrc: srcApCalc,
    programId: ""
  },
  {
    title: "AMC",
    description: "The American Mathematics Competition administered by the Mathematics Association of America. Represent your school and compete with other mathematics athletes.",
    imgSrc: srcAmc,
    programId: ""
  }
];

export class HomeSectionPrograms extends React.Component {
	render() {
    var programsUrl = getNav('programs').url;
		return (
			<div className="section">
				<h2>Programs</h2>
        <div className="card-container">
          <ProgramCard program={programs[0]}/>
          <ProgramCard program={programs[1]}/>
          <ProgramCard program={programs[2]}/>
        </div>

        <div>
          <Link to={programsUrl}>
            <button className="inverted">View More Programs</button>
          </Link>
        </div>
			</div>
		)
	}
}

class ProgramCard extends React.Component {
  render() {
    var program = this.props.program;
    var title = program.title;
    var imgSrc = program.imgSrc;
    var description = program.description;
    // var url = getNav(this.props.programId).url;
    return (
      <div className="home-program-card">
        <div className="img-container">
          <img src={imgSrc}/>
        </div>
        <h3>{title}</h3>
        <p>{description}</p>
        <button className="inverted">Sign Up</button>
      </div>
    );
  }
}
