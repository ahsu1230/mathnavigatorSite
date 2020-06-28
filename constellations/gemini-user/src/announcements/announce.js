"use strict";
require("./announce.sass");
import React from "react";
import API from "../utils/api.js";
import { keys, sortBy, groupBy } from "lodash";
import moment from "moment";

export class AnnouncementGroup extends React.Component {
    render() {
        let a = this.props.announcements;
        var postedAt = moment(this.props.postedAt).format('l');
        let announcements = a.map((announcement, index) => (
            <li key={announcement.id}>
                <p>{announcement.message}</p>
            </li>
        ));

        return (
            <div className="announcement-card">
                <h2>{postedAt}</h2>
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
        console.log("api attempt ");
        API.get("api/announcements/all").then((res) => {
            const list = res.data;
            console.log("api success!");
            this.setState({ announcementList: list });
        });
    }
    render() {
        const announcements = this.state.announcementList;
        let sorted = groupBy(announcements, (a) => a.postedAt.substring(0, 10));
        let dates = sortBy(keys(sorted)).reverse();
        let items = [];
        dates.forEach((date) => {
            items.push(
                <AnnouncementGroup
                    key={date}
                    postedAt={date}
                    announcements={sorted[date]}
                />
            );
        });
        return (
            <div id="view-announce">
                <h1>Announcements</h1>
                {items}
            </div>
        );
    }
}
