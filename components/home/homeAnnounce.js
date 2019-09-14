'use strict';
require('./homeAnnounce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { isEmpty, filter } from 'lodash';
import { getAnnouncements } from '../repos/apiRepo.js';
import { ANNOUNCE_LAST_DISMISS } from '../storage.js';
const classnames = require('classnames');
const srcClose = require('../../assets/close_black.svg');

export class HomeAnnounce extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: false,
      targetAnnounce: {},
      lastDismissed: 0
    };
    this.handleDismiss = this.handleDismiss.bind(this);
  }

  componentDidMount() {
    getAnnouncements().then(announcements => {
      var valid = filter(announcements, a => a.onHomePage);
      var targetAnnounce = valid.length > 0 ? valid[0] : undefined;
      targetAnnounce.shortMessage = shortenMessage(targetAnnounce.message);

      var localLastDismiss = localStorage.getItem(ANNOUNCE_LAST_DISMISS);
      var lastDismissed = parseInt(localLastDismiss, 10) || 0;

      this.unbindTimeout = setTimeout(function() {
        this.setState({ show: true });
      }.bind(this), 2500);

      this.setState({
        targetAnnounce: targetAnnounce,
        lastDismissed: lastDismissed
      });
    });
  }

  componentWillUnmount() {
    clearTimeout(this.unbindTimeout);
  }

  handleDismiss() {
    this.setState({show: false});
    localStorage.setItem(ANNOUNCE_LAST_DISMISS, this.state.targetAnnounce.timestamp);
  }

	render() {
    const announce = this.state.targetAnnounce;
    var showAnnounce = !isEmpty(announce) && (announce.timestamp > this.state.lastDismissed);
    var component;
    if (showAnnounce) {
      component = <Popup announce={announce}
                          show={this.state.show}
                          announceHeight={this.state.announceHeight}
                          handleDismiss={this.handleDismiss}/>
    } else {
      component = <div></div>;
    }
    return component;
	}
}

class Popup extends React.Component {
  render() {
    const announce = this.props.announce;
    const show = this.props.show;
    const handleDismiss = this.props.handleDismiss;
    const announceClass = classnames({show: show});
    return (
      <div key="real" id="home-announce" className={announceClass}>
        <h3>New Announcement!</h3>
        <button className="close-x" onClick={handleDismiss}>
          <img src={srcClose}/>
        </button>
        <div className="text-container">
          <p>{announce.shortMessage}</p>
        </div>
        <Link to="/announcements">Read more &#62;</Link>
      </div>
    );
  }
}

function shortenMessage(message) {
  var array = message.split(" ");
  var shortMessage = "";
  var needEllipsis = false;

  var i = 0;
  while (i < array.length) {
    var append = shortMessage + array[i];
    if (append.length > 120) {
      needEllipsis = true;
      break;
    }
    shortMessage += array[i] + " ";
    i++;
  }

  if (needEllipsis) {
    shortMessage = shortMessage.slice(0, -1); // remove last "space"
    shortMessage += "...";
  }
  return shortMessage;
}
