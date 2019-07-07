'use strict';
require('./../styl/home.styl');
import React from 'react';
import ReactDOM from 'react-dom';

const mnIcon = require('../assets/navigate_white.png');
const mnLogo = require('../assets/logo_white.png');

export class HomePage extends React.Component {
	render() {
		return (
      <div id="view-home">
        <div id="view-home-container">
          <HomeBanner/>
					<HomeSection/>
					<HomeSection/>
					<HomeSection/>
        </div>
      </div>
		);
	}
}

class HomeBanner extends React.Component {
	render() {
		return (
			<div id="home-banner-container">
				<div id="banner-bg-img"></div>
				<div id="banner-bg-overlay"></div>
				<div id="banner-logo-wrapper">
					<img src={mnIcon}/>
					<img src={mnLogo}/>
				</div>
			</div>
		);
	}
}

class HomeSection extends React.Component {
	render() {
		return (
			<div className="section">
				<h2>Some Text 1</h2>
				<p>
				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
				</p>
			</div>
		);
	}
}
