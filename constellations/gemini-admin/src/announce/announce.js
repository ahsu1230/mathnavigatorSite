"use strict";
require("./announce.styl");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { Link } from "react-router-dom";

export class AnnouncePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/announcements/all").then((res) => {
            const announcements = res.data;
            this.setState({ list: announcements });
        });
    }

    render() {
        const rows = this.state.list.map((row, index) => {
            return (
                <li key={index}>
                    <AnnounceRow key={index} row={row} />
                </li>
            );
        });
        const numRows = rows.length;
        return (
            <div id="view-announce">
                <h1>All Announcements ({numRows}) </h1>

                <ul className="announce-list-row subheader">
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

        const now = moment();
        const isScheduled = postedAt.isAfter(now);

        const author = this.props.row.author;
        const message = this.props.row.message;

        const url = "/announcements/" + announceId + "/edit";
        return (
            <ul className="announce-list-row">
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
