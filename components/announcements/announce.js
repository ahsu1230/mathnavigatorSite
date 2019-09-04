'use strict';
require('./announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { Promise } from 'bluebird';
import { groupBy, map, sortBy } from 'lodash';
import { getAnnouncements, getProgramByClass } from '../repos/apiRepo.js';
const classnames = require('classnames');
const srcAttention = require('../../assets/attention_orange.svg');

export class AnnouncePage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			announcements: [],
			groupByDate: {},
			sorted: []
		}
  }

	componentDidMount() {
		getAnnouncements().then(list => {
			var sorted = sortBy(list, ["date", "timestamp"]).reverse();
			var groupByDate = groupBy(sorted, a => a.dateStr);
			this.setState({
				announcements: list,
				groupByDate: groupByDate,
				sorted: sorted
			});
		});

		if (process.env.NODE_ENV === 'production') {
			mixpanel.track("announce");
		}
	}

	render() {
		const announcements = this.state.sorted;
		const groupByDate = this.state.groupByDate;

		var i = 0;
		const cards = map(groupByDate, function(list, date) {
			return <AnnounceCard key={i++} date={date} announcements={list}/>;
		});

		return (
      <div id="view-announce">
        <div id="view-announce-container">
          <h1>Announcements</h1>
					{cards}
        </div>
      </div>
		);
	}
}

class AnnounceCard extends React.Component {
	render() {
		const date = this.props.date;
		const announcements = this.props.announcements;

		const items = announcements.map((a, i) =>
			<AnnounceItem key={i} announcement={a}/>
		);

		return (
			<div className="announce-card">
				<h3>{date}</h3>
				{items}
			</div>
		);
	}
}

class AnnounceItem extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			links: []
		}
	}

	componentDidMount() {
		const announcement = this.props.announcement;
		var promiseList = announcement.classKeys.map(function(classKey) {
			return getProgramByClass(classKey).then(function(pair) {
				var programObj = pair.programObj || {};
				var classObj = pair.classObj || {};
				var fullName = (programObj.title + " " + classObj.className).trim();
				const url = "/class/" + classKey;
				return {
					name: fullName,
					url: url
				};
			});
		});
		Promise.all(promiseList, obj => {return obj; })
			.then(result => {this.setState({links: result})});
	}

	render() {
		const announcement = this.props.announcement;
		const links = this.state.links.map((link, i) =>
			<Link key={i} to={link.url}>View {link.name}</Link>
		);

		var attention;
		if (announcement.important) {
			attention = <img className="img-alert" src={srcAttention}/>;
		}

		return (
			<div className="announce-item">
				<div>
					<p>{announcement.message}</p>
					{attention}
				</div>
				<div className="card-links">{links}</div>
			</div>
		);
	}
}
