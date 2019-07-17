'use strict';
require('./../styl/announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getList } from './repos/announceRepo.js';
import { getProgramClass } from './repos/programRepo.js';
const classnames = require('classnames');

export class AnnouncePage extends React.Component {
	render() {
		const announcements = getList();

		const cards = announcements.map((obj, index) =>
			<AnnounceCard key={index} announcement={obj} />
		);

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
		const announcement = this.props.announcement;
		const alertClassNames = classnames("card-alert", {
			"show" : announcement.important
		});
		const links = announcement.classKeys.map(function(classKey, index) {
			var pair = getProgramClass(classKey);
			var programName = pair.programObj.title;
			var className = pair.classObj.className || "";
			var fullName = (programName + " " + className).trim();
			const url = "/class/" + classKey;

			return (
				<Link key={index} to={url}>View {fullName}</Link>
			);
		});

		return (
			<div className="announce-card">
				<div className={alertClassNames}>Important!</div>
				<h3>{announcement.date}</h3>
				<p>{announcement.message}</p>
				<div className="card-author">~ {announcement.author}</div>
				<div className="card-links">{links}</div>
			</div>
		);
	}
}
