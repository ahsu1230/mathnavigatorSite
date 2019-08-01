'use strict';
require('./footer.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { NavLinks } from '../constants.js';
const headerIcon = require('../../assets/navigate_white.png');

export class Footer extends React.Component {
	render() {
		const links1 = NavLinks.map((link, index) =>
			<li>
				<Link key={link.id} to={link.url}>{link.name}</Link>
			</li>
		);

		return (
      <div id="view-footer">
        <div id="view-footer-container">
					<ul>{links1}</ul>
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
