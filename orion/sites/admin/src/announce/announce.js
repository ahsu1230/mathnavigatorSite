"use strict";
require("./announce.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";

export class AnnouncePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    render() {
        const rows = this.state.list.map((row, index) => {
            return <AnnounceRow key={index} row={row} />;
        });
        const numRows = rows.length;
        return (
            <div id="view-announce">
                <h1>All Announcements ({numRows}) </h1>

                <ul className="announce-lists">
                    <li className="li-med">Date</li>
                    <li className="li-med">Author</li>
                    <li className="li-large">Message</li>
                    <li className="li-small"> </li>
                </ul>

                <ul id="announce-list">{rows}</ul>
                <Link to={"/announcements/add"}>
                    {" "}
                    <button className="announcement-button">
                        {" "}
                        Add Announcement
                    </button>{" "}
                </Link>
            </div>
        );
    }
}

class AnnounceRow extends React.Component {
    render() {
        const date = this.props.announceObj.date;
        const author = this.props.announceObj.author;
        const message = this.props.announceObj.message;
        const url = "/announcement/" + "/edit";
        return (
            <ul className="announce-lists">
                <li className="li-med"> {date} </li>
                <li className="li-med"> {author} </li>
                <li className="li-large"> {message} </li>
                <Link to={url}> Edit </Link>
            </ul>
        );
    }
}
