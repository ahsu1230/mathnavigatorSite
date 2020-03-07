'use strict';
require('./announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class AnnouncePage extends React.Component {
	render() {
		var numAnnouncements = 3; 	/* this needs to be a variable later */

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
				<ul> {/* this is the subheading */}
					<div className="announce-lists">
						<li className="li-med"> Date </li>
						<li className="li-med"> Author </li>
						<li className="li-large"> Message </li>
						<li className="li-small"> </li>
					</div>
				</ul>

				{/* these are the lists under the subheading */}
				<ul> {/*I think there's a problem having li under li, but I'll fix it later*/}
					<li className="announce-lists">
						<li className="li-med"> {fakeAnnounceA.date} </li>
						<li className="li-med"> {fakeAnnounceA.author} </li>
						<p className="li-large"> {fakeAnnounceA.message} </p>
						<li className="li-small"> Edit > </li>
					</li>
					<li className="announce-lists">
						<li className="li-med"> {fakeAnnounceB.date} </li>
						<li className="li-med"> {fakeAnnounceB.author} </li>
						<p className="li-large"> {fakeAnnounceB.message} </p>
						<li className="li-small"> Edit > </li>
					</li>
					<li className="announce-lists">
						<li className="li-med"> {fakeAnnounceC.date} </li>
						<li className="li-med"> {fakeAnnounceC.author} </li>
						<p className="li-large"> {fakeAnnounceC.message} </p>
						<li className="li-small"> Edit > </li>
					</li>

				</ul>

			{/* I think this should be <Link> not <a> but <Link> makes page not load. fix later*/}
			{/*okay now can't find page error. idk why it's not working? */}
				<a href="/announceEdit.js"> <button className="announcement-button"> Add Announcement</button> </a>

      </div>
		);
	}
}
