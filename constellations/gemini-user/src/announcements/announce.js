"use strict";
require("./announce.sass");
import React from "react";
import API from "../utils/api.js";
import { keys, sortBy, groupBy } from "lodash";
import moment from "moment";

export class AnnouncementGroup extends React.Component {
    render() {
        const postedAt = moment(this.props.postedAt);
        const announcements = this.props.announcements.map(
            (announcement, index) => (
                <li key={announcement.id}>
                    <h4>
                        Posted at {postedAt.format("hh:mm a")} by{" "}
                        {announcement.author}
                    </h4>
                    <p>{announcement.message}</p>
                </li>
            )
        );

        return (
            <div className="announcement-card">
                <h2>{postedAt.format("l")}</h2>
                <ul>{announcements}</ul>
            </div>
        );
    }
}

export class AnnouncePage extends React.Component {
    state = {
        announcementList: [],
    };
    componentDidMount() {
        console.log("api attempt");
        API.get("api/announcements/all").then((res) => {
            const list = res.data;
            console.log("api success!");
            this.setState({ announcementList: list });
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
                postedAt={date}
                announcements={sorted[date]}
            />
        ));
        return (
            <div id="view-announce">
                <h1>Announcements</h1>
                {items}
            </div>
        );
    }
}
