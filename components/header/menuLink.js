'use strict';
require('./header.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { isPathAt } from '../constants.js';
const classnames = require('classnames');

export default class MenuLink extends React.Component {
  render() {
		const location = this.props.location;
		const url = this.props.url;
		const linkClasses = classnames({
			"active": isPathAt(location.pathname || "", url)
		});

    return (
    	<Link className={linkClasses} to={url} onClick={this.props.onClick}>
        <li>{this.props.title}</li>
    	</Link>
		);
  }
}
