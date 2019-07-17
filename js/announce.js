'use strict';
require('./../styl/announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { getList } from './repos/announceRepo.js';

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
		return (
			<div className="announce-card">
				<h3>{announcement.date}</h3>
				<p>{announcement.message}</p>
				<div className="card-author">~ {announcement.author}</div>
				<div></div>
			</div>
		);
	}
}
