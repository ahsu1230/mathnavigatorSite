'use strict';
require('./afhPopup.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { FTU_AFH_POPUP } from '../storage.js';
const classnames = require('classnames');
const srcQuestion = require('../../assets/question_white.svg');
const srcClose = require('../../assets/close_white.svg');

export class AfhPopup extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			firstTime: false,
			showMessage: false
		};
    this.directToAfh = this.directToAfh.bind(this);
		this.dismissMessage = this.dismissMessage.bind(this);
	}

	componentDidMount() {
    var localFtuAfhPopup = localStorage.getItem(FTU_AFH_POPUP);
    if (!localFtuAfhPopup) {
      this.unbindTimeout = setTimeout(function() {
  			this.setState({ showMessage: true });
  		}.bind(this), 2500);
    }
	}

	componentWillUnmount() {
    if (this.unbindTimeout) {
      clearTimeout(this.unbindTimeout);
    }
	}

  directToAfh() {
    // history.push("/askforhelp"); // use with BrowserRouter
		window.location.hash = "/ask-for-help"; // use with HashRouter
  }

	dismissMessage() {
    localStorage.setItem(FTU_AFH_POPUP, true);
		this.setState({ showMessage: false });
	}

	render() {
		const messageClasses = classnames("afh-message", {
			show: this.state.showMessage
		});
		return (
			<div className="afh-popup-container">
				<button className="afh-question" onClick={this.directToAfh}>
					<img src={srcQuestion}/>
				</button>
				<div className={messageClasses}>
					<button className="afh-dismiss-btn" onClick={this.dismissMessage}>
						<img src={srcClose}/>
					</button>
					<p>
						Need additional help on any topics?<br/>
						Register for free sessions to ask your questions!
					</p>
          <Link to={"/ask-for-help"}>Details &#62;</Link>
				</div>
			</div>
		);
	}
}
