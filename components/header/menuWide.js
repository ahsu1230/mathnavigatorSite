'use strict';
require('./headerWide.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { NavLinks } from '../constants.js';

export default class MenuWide extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			links: NavLinks
		}
	}
	render() {
		const location = this.props.location;
		const onClick = this.props.onClick;
    const items = this.state.links.map((link, i) => {
			if (link.subLinks && link.subLinks.length > 0) {
				return (<SubMenu key={link.id} currentLink={link} subLinks={link.subLinks}/>);
			} else {
				return (<MenuLink key={link.id} link={link}/>);
			}
		});
		return (
			<ul className="header-menu-wide">{items}</ul>
		);
	}
}

class MenuLink extends React.Component {
	render() {
		const link = this.props.link;
		return (
			<div className="menu-link-container">
				<Link to={link.url}>{link.name}</Link>
			</div>
		);
	}
}

class SubMenu extends React.Component {
	constructor(props) {
		super(props);
	}

	render() {
		const currentLink = this.props.currentLink;
		const subLinks = this.props.subLinks.map((subLink, i) => (
			<li key={i}>
				<Link to={subLink.url}>{subLink.name}</Link>
			</li>
		));
		return (
			<div className="sub-menu">
				<MenuLink link={currentLink}/>
				<ul>
					{subLinks}
				</ul>
			</div>
		);
	}
}
