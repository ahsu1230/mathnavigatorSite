'use strict';
require('./announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class AnnouncePage extends React.Component {
	render() {
		var numAnnouncements = 3;/* this needs to be a variable later */

		/* these fake announcements will eventually be replaced with real announcements */
		let fakeAnnounceA = {
			date: "02/15/2020",
			author: "Melon King",
			message: "Only evil people waste watermelon"
		};
		let fakeAnnounceB = {
			date: "02/16/2020",
			author: "Fruit Hater",
			message: "Watermelons are fruit, therefore they deserve to dry on the vine"
		};
		/* will use this to test text wrapping later */
		let fakeAnnounceC = {
			date: "02/23/2020",
			author: "Talkative",
			message: "This will be a longer message later"
		};
		return (
      <div id="view-announce">
        <h1>All Announcements ({numAnnouncements}) </h1>
				<ul>
					<div className="announce-lists">
						<li className="li-med"> Date </li>
						<li className="li-med"> Author </li>
						<li className="li-large"> Message </li>
						<li className="li-small"> </li>
					</div>
				</ul>

				<ul>
					<li className="announce-lists">
						<div className="li-med"> {fakeAnnounceA.date} </div>
						<div className="li-med"> {fakeAnnounceA.author} </div>
						<p className="li-large"> {fakeAnnounceA.message} </p>
						<div className="li-small"> Edit > </div>
					</li>
					<li className="announce-lists">
						<div className="li-med"> {fakeAnnounceB.date} </div>
						<div className="li-med"> {fakeAnnounceB.author} </div>
						<p className="li-large"> {fakeAnnounceB.message} </p>
						<div className="li-small"> Edit > </div>
					</li>
					<li className="announce-lists">
						<div className="li-med"> {fakeAnnounceC.date} </div>
						<div className="li-med"> {fakeAnnounceC.author} </div>
						<p className="li-large"> {fakeAnnounceC.message} </p>
						<div className="li-small"> Edit > </div>
					</li>

				</ul>

				<Link to={"/announce/add"}> <button className="announcement-button"> Add Announcement</button> </Link>
      </div>
		);
	}
}
