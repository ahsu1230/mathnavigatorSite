'use strict';
require('./announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class AnnouncePage extends React.Component {
	constructor(props) {
	    super(props);
	    this.state = {
	      list: []
	    };
 	}

	render() {
		var numAnnouncements = 3;
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
		/* TODO: will use this to test text wrapping later */
		let fakeAnnounceC = {
			date: "02/23/2020",
			author: "Talkative",
			message: "This will be a longer message later"
		};

		return (
			<div id="view-announce">
				<h1>All Announcements ({numAnnouncements})</h1>

				<ul className="announce-lists">
					<li className="li-med"> Date </li>
					<li className="li-med"> Author </li>
					<li className="li-large"> Message </li>
					<li className="li-small"> </li>
				</ul>

				<ul>
					<AnnounceRow announceObj = {fakeAnnounceA}/>
					<AnnounceRow announceObj = {fakeAnnounceB}/>
					<AnnounceRow announceObj = {fakeAnnounceC}/>
				</ul>
				<Link to={"/announcements/add"}> <button className="announcement-button"> Add Announcement</button> </Link>
			</div>
		);
	}
}

class AnnounceRow extends React.Component {
  	render() {
  		const date = this.props.announceObj.date;
		const author = this.props.announceObj.author;
		const message = this.props.announceObj.message;
		const url = "/announcement/"  + "/edit";
	    return (
	    	<ul className="announce-lists">
	        	<li className="li-med"> {date} </li>
	        	<li className="li-med"> {author} </li>
				<li className="li-large"> {message} </li>
	        	<Link to={url}> Edit </Link>
	      	</ul>
	    )
  	}
}
