'use strict';
require('./header.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import MenuLink from './menuLink.js';
import { NavLinks } from '../constants.js';
const classnames = require('classnames');

export default class HeaderNavVert extends React.Component {
  constructor() {
      super();
      this.state = {
        show: false
      };
      this.toggleMenu = this.toggleMenu.bind(this);
  }

  toggleMenu() {
    this.setState({
      show: !this.state.show
    });
  }

  render() {
		const location = this.props.location;
    const show = this.state.show;
    var buttonClasses = classnames("header-menu-btn", {
      "show": show
    });
    var iconClasses = classnames("icon-arrow", {
      "show": show
    });
    return (
      <div className="header-menu-slim">
        <button className={buttonClasses} onClick={this.toggleMenu}>
          Menu
          <div className={iconClasses}></div>
        </button>
        <HeaderMenuList showMenu={show} toggleMenu={this.toggleMenu} location={location}/>
      </div>
    );
  }
}

class HeaderMenuList extends React.Component {
  render() {
		const location = this.props.location;
		const toggleMenu = this.props.toggleMenu;
    const showMenu = this.props.showMenu;
    const numLinks = NavLinks.length;
    const items = NavLinks.map((link, i) =>
      <MenuLink key={link.id} title={link.name} url={link.url} onClick={toggleMenu} location={location}/>
    );
    const menuClasses = classnames("header-menu-list", {
      "show": showMenu
    });
    return (
      <div className={menuClasses}>
        <ul>
          {items}
        </ul>
      </div>
    );
  }
}
