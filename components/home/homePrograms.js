'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getNav } from '../constants.js';

const srcSat = require('../../assets/cb_sat.jpg');
const srcApCalc = require('../../assets/ap_calc.jpg');
const srcAmc = require('../../assets/amc.jpg');

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
			<div className="section programs">
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
    var url = getNav("programs").url;
    // var url = getNav(this.props.programId).url;
    return (
      <div className="home-tile-card program">
        <div className="img-container">
          <img src={program.imgSrc}/>
        </div>
        <h3>{program.title}</h3>
        <p>{program.description}</p>
        <Link to={url}>
          <button className="inverted">Sign Up</button>
        </Link>
      </div>
    );
  }
}
