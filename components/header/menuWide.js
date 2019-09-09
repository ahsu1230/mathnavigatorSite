'use strict';
require('./headerWide.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import MenuLink from './menuLink.js';
import { NavLinks } from '../constants.js';

export default class MenuWide extends React.Component {
	render() {
		const location = this.props.location;
		const onClick = this.props.onClick;
    const items = NavLinks.map((link, i) =>
			<MenuLink key={link.id} title={link.name} url={link.url} location={location}/>
		);
		return (
			<ul className="header-menu-wide">{items}</ul>
		);
	}
}
