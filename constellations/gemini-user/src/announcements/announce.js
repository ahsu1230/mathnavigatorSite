"use strict";
require("./announce.sass");
import React from "react";
import API from "../utils/api.js";
import { keys, sortBy, groupBy } from "lodash";
import moment from "moment";
import srcPin from "../../assets/pin_green.svg";

export class AnnouncePage extends React.Component {
    state = {
        pinnedAnnouncement: null,
        announcementList: [],
    };
    componentDidMount() {
        console.log("api attempt");
        API.get("api/announcements/all").then((res) => {
            const list = res.data;
            const pinnedAnnouncement = list.find((a) => a.onHomePage);
            this.setState({
                announcementList: list,
                pinnedAnnouncement: pinnedAnnouncement,
            });
        });
    }
    render() {
        const announcements = this.state.announcementList;
        let sorted = groupBy(announcements, (a) =>
            moment(a.postedAt).format("l")
        );
        const dates = sortBy(keys(sorted)).reverse();
        const items = dates.map((date, index) => (
            <AnnouncementGroup
                key={index}
                postedAtDate={date}
                announcements={sorted[date]}
            />
        ));
        return (
            <div id="view-announce">
                <PinnedAnnouncement
                    announcement={this.state.pinnedAnnouncement}
                />
                <h1>All Announcements</h1>
                {items}
            </div>
        );
    }
}

class AnnouncementGroup extends React.Component {
    render() {
        const postedAtDate = moment(this.props.postedAtDate);
        const announcements = this.props.announcements.map(
            (announcement, index) => (
                <li key={announcement.id}>
                    <p className="message">{announcement.message}</p>
                    <p className="posted-by">
                        Posted at{" "}
                        {moment(announcement.postedAt).format("h:mm a")} by{" "}
                        {announcement.author}
                    </p>
                </li>
            )
        );

        return (
            <div className="announcement-card">
                <h2>{postedAtDate.format("dddd MMM. Do, YYYY")}</h2>
                <ul>{announcements}</ul>
                <div className="bar"></div>
            </div>
        );
    }
}

class PinnedAnnouncement extends React.Component {
    render() {
        const announcement = this.props.announcement || null;
        let content = <div></div>;
        if (announcement) {
            const postedAt = moment(announcement.postedAt);
            content = (
                <div className="announcement-card pinned">
                    <div className="pinned-header">
                        <img src={srcPin} />
                        <h2>Reminder</h2>
                    </div>
                    <h2>{postedAt.format("dddd MMM. Do, YYYY")}</h2>
                    <p className="message">{announcement.message}</p>
                    <p className="posted-by">
                        Posted at {postedAt.format("h:mm a")} by{" "}
                        {announcement.author}
                    </p>
                </div>
            );
        }

        return (
            <div>
                {this.props.pinnedAnnouncement}
                {content}
            </div>
        );
    }
}
