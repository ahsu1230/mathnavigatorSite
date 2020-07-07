'use strict';
require('./footer.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { NavLinks } from '../constants.js';
const headerIcon = require('../../assets/navigate_white.png');

export default class Footer extends React.Component {
	componentDidMount() {
		if (process.env.NODE_ENV === 'production') {
			mixpanel.track("init");
		} else {
			console.log("Website loaded");
		}
	}

	render() {
		const links = NavLinks.map((link, index) =>
			<li key={link.id}>
				<Link to={link.url}>{link.name}</Link>
			</li>
		);

		return (
      <div id="view-footer">
        <div id="view-footer-container">
					<ul>{links}</ul>
					<div className="logo-container">
						<Link id="footer-logo" to="/">
			        <img src={headerIcon}/>
			        <h1 className="logo"></h1>
			      </Link>
						<h2>Montgomery County, MD</h2>
						<p>
							Program sessions are held off-school hours at local public schools.<br/>
							Math Navigator is not affiliated with those schools.
						</p>
					</div>
        </div>
      </div>
		);
	}
}
