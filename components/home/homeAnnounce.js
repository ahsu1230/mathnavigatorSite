'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getAnnounceList } from '../repos/mainRepo.js';
const classnames = require('classnames');
const srcClose = require('../../assets/close_black.svg');

export class HomeAnnounce extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: false
    }
    this.announcements = getAnnounceList();

    this.handleDismiss = this.handleDismiss.bind(this);
  }

  componentDidMount() {
    this.unbindTimeout = setTimeout(function() {
      this.setState({ show: true });
    }.bind(this), 2500);
  }

  componentWillUnmount() {
    clearTimeout(this.unbindTimeout);
  }

  handleDismiss() {
    this.setState({show: false});
  }

	render() {
    const announce = this.announcements[0]; // take the first announcement
    const show = this.state.show;
    const announceClass = classnames({
      show: this.state.show
    });


		return (
			<div id="home-announce" className={announceClass}>
        <h3>New Announcement!</h3>
        <button className="close-x" onClick={this.handleDismiss}>
          <img src={srcClose}/>
        </button>
        <div className="text-container">
          <p>
            {announce.message}
          </p>
        </div>
        <span>...</span>
        <Link to="/announcements">Read more &#62;</Link>
			</div>
		);
	}
}
