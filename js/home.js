'use strict';
require('./../styl/home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { HomeSectionPrograms } from './homePrograms.js';
import { HomeSectionSuccess } from './homeSuccess.js';
import { HomeSectionStories } from './homeStories.js';
import { getNav } from './constants.js';

export class HomePage extends React.Component {
	render() {
		return (
      <div id="view-home">
        <div id="view-home-container">
          <HomeBanner/>
					<HomeSectionPrograms/>
					<HomeSectionSuccess/>
					<HomeSectionStories/>
        </div>
      </div>
		);
	}
}

class HomeBanner extends React.Component {
	render() {
		var link = getNav('programs');
		return (
			<div id="home-banner-container">
				<div id="banner-bg-img"></div>
				<div id="banner-bg-overlay"></div>
				<div id="banner-content">
					<h2>Montgomery County, MD</h2>
					<h1>
						Providing affordable, high quality education<br/>
						to thousands of students for 15 years.
					</h1>
				</div>
				<Link to={link.url}><button>Join a Program today</button></Link>
			</div>
		);
	}
}
