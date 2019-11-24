'use strict';
require('./header.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { MenuSlim } from './menuSlim.js';
import MenuWide from './menuWide.js';
const headerIcon = require('../../assets/navigate_white.png');

export class Header extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			location: {},
		};
	}

	componentDidMount() {
		this.unlisten = this.props.history.listen((location, action) => {
			this.setState({location: location});
		});
		this.setState({location: this.props.history.location});
	}

	componentWillUnmount() {
		this.unlisten();
	}

	render() {
		const location = this.state.location;
		return (
      <div id="view-header">
        <div id="view-header-container">
          <HeaderLogo/>
          <MenuWide location={location}/>
					<MenuSlim location={location}/>
        </div>
      </div>
		);
	}
}

class HeaderLogo extends React.Component {
  render() {
    return (
      <a id="header-logo" href="/">
        <img src={headerIcon}/>
        <h1 className="logo"></h1>
      </a>
    );
  }
}
