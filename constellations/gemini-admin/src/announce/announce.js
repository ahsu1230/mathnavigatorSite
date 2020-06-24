"use strict";
require("./announce.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { Link } from "react-router-dom";

export class AnnouncePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: []
        };
    }

    componentDidMount() {
        API.get("api/announcements/all").then((res) => {
            const announcements = res.data;
            this.setState({ list: announcements });
        });
    }
	
	onChangeCheckbox = (e) => {
		let successCallback = () => console.log('api success')
        let failCallback = (err) =>
            alert("Could not save announcement: " + err.response.data);
			
		let previous = this.state.list.findIndex(announcement => announcement.onHomePage)
		let current = this.state.list.findIndex(announcement => announcement.id == e.target.id)
		
		this.state.list[current].onHomePage = true
		if (previous >= 0) { this.state.list[previous].onHomePage = false }
		
		[previous, current].forEach(index => {
			if (index >= 0) {
				API.post(
					"api/announcements/announcement/" + this.state.list[index].id,
					this.state.list[index]
				).then((res) => successCallback())
				.catch((err) => failCallback(err));
			}
		})
		this.setState(this.state)
	}

    render() {
        const rows = this.state.list.map((row, index) => {
            return (
                <li key={index}>
                    <AnnounceRow key={index} row={row} selected={this.state.selected} onChangeCheckbox={this.onChangeCheckbox}/>
                </li>
            );
        });
        const numRows = rows.length;
        return (
            <div id="view-announce">
                <h1>All Announcements ({numRows}) </h1>

                <ul className="announce-list-row subheader">
					<li className="li-med">On Home Page</li>
                    <li className="li-med">State</li>
                    <li className="li-med">Date</li>
                    <li className="li-med">Author</li>
                    <li className="li-large">Message</li>
                </ul>

                <ul id="announce-list">{rows}</ul>
                <Link to={"/announcements/add"}>
                    <button className="announcement-button">
                        Add Announcement
                    </button>
                </Link>
            </div>
        );
    }
}

class AnnounceRow extends React.Component {
    render() {
        const announceId = this.props.row.id;
        const postedAt = moment(this.props.row.postedAt);
		
		const checked = this.props.row.onHomePage || false;

        const now = moment();
        const isScheduled = postedAt.isAfter(now);

        const author = this.props.row.author;
        const message = this.props.row.message;

        const url = "/announcements/" + announceId + "/edit";
        return (
            <ul className="announce-list-row">
				<li className="li-med">
					<input type="checkbox" onChange={this.props.onChangeCheckbox} id={announceId} checked={checked}/>
				</li>
                <li
                    className={
                        "li-med " + (isScheduled ? " scheduled" : " published")
                    }>
                    <div>{isScheduled ? "Scheduled" : "Published"}</div>
                    <div>{postedAt.fromNow()}</div>
                </li>
                <li className="li-med">
                    <div>{postedAt.format("M/D/YYYY")}</div>
                    <div>{postedAt.format("hh:mm a")}</div>
                </li>
                <li className="li-med"> {author} </li>
                <li className="li-large"> {message} </li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
