'use strict';
require('./announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';


export class AnnouncePage extends React.Component {
	render() {
		var numAnnouncements = 5; 	{/* this needs to be a variable later */}

		{/* these fake announcements will eventually be replaced with real announcements */}
		var fakeAnnounceA = {
			date: "02/15/2020",
			author: "Melon King",
			message: "Only evil people waste watermelon"
		};
		var fakeAnnounceB = {
			date: "02/16/2020",
			author: "Fruit Hater",
			message: "Watermelons are fruit, therefore they deserve to dry on the vine"
		};
		{/* will use this to test text wrapping later */}
		var fakeAnnounceC = {
			date: "02/23/2020",
			author: "Talkative",
			message: "This will be a longer message later"
		};
		return (
      <div id="view-announce">
        <h1>All Announcements ({numAnnouncements}) </h1>
		<ul> {/* this is the subheading */}
			<li className="announce-lists">
				<div className="li-med"> Date </div>
				<div className="li-med"> Author </div>
				<div className="li-large"> Message </div>
				<div className="li-small"> </div>
			</li>
		</ul>

		{/* these are the lists under the subheading */}

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

		<div className="announcement-button"> Add Announcement</div>

      </div>
		);
	}
}
