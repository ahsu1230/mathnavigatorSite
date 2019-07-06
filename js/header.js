'use strict';
require('./../styl/header.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  Link
} from 'react-router-dom';
import { NavLinks } from './constants.js';
const headerIcon = require('../assets/navigate_white.png');


export class Header extends React.Component {
	render() {
		return (
      <div id="view-header">
        <div id="view-header-container">
          <HeaderLogo/>
          <HeaderMenu/>
        </div>
      </div>
		);
	}
}

class HeaderLogo extends React.Component {
  render() {
    return (
      <a id="header-logo" href="/">
        <img className="icon" src={headerIcon}/>
      </a>
    );
  }
}

class HeaderMenu extends React.Component {
  render() {
    return (
      <div>
        <button className="header-menu-btn">
          Menu
          <div className="icon-arrow"></div>
        </button>
        <HeaderMenuList/>
      </div>
    );
  }
}

class HeaderMenuList extends React.Component {
  render() {
    const numLinks = NavLinks.length;
    const items = NavLinks.map((link, i) =>
      <li key={link.id}>
        <MenuLink title={link.name} url={link.url}/>
      </li>
    );
    return (
      <div className="header-menu-list">
        <ul>
          {items}
        </ul>
      </div>
    );
  }
}

class MenuLink extends React.Component {
  render() {
    return (
    	<Link to={this.props.url}>
    		<span>{this.props.title}</span>
    	</Link>
		);
  }
}
