'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getNav } from '../constants.js';
import { programsText } from './homeText.js';

export class HomeSectionPrograms extends React.Component {
	render() {
    var programsUrl = getNav('programs').url;
		return (
			<div className="section programs">
				<h2>Programs</h2>
        <div className="card-container">
          <ProgramCard program={programsText[0]}/>
          <ProgramCard program={programsText[1]}/>
          <ProgramCard program={programsText[2]}/>
        </div>

        <div>
          <Link to={programsUrl}>
            <button>View More Programs</button>
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
