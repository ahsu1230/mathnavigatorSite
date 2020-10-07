"use strict";
require("./afh.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../utils/api.js";
import moment from "moment";
import srcPoint from "../../assets/point_right_green.svg";
import srcNotes from "../../assets/lightbulb_white.svg";
import srcMath from "../../assets/icon_math.svg";
import srcWriting from "../../assets/icon_writing.svg";
import srcCoding from "../../assets/icon_coding.svg";
import { keyBy } from "lodash";
import { createLocation } from "../utils/locationUtils.js";

const subjectIconSrc = {
    math: srcMath,
    english: srcWriting,
    programming: srcCoding,
};

const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Programming",
};

export class AFHPage extends React.Component {
    state = {
        currentSubject: "math",
        sessions: [],
        locations: [],
        locationMap: {},
    };

    componentDidMount() {
        API.get("api/askforhelp/all").then((res) => {
            const afh = res.data;
            this.setState({
                sessions: afh,
            });
        });
        API.get("api/locations/all").then((res) => {
            const locations = res.data;
            this.setState({
                locations: locations,
                locationMap: keyBy(locations, "locationId"),
            });
        });
    }

    changeSubject = (subjectName) => {
        this.setState({
            currentSubject: subjectName,
        });
    };

    render() {
        let currentTab = this.state.sessions.filter(
            (session) => session.subject == this.state.currentSubject
        );

        let showSessions = currentTab.map((row, index) => {
            const location = this.state.locationMap[row.locationId];
            return <AfhSessionRow key={index} row={row} location={location} />;
        });
        if (showSessions.length == 0) {
            showSessions = (
                <p>
                    No Ask-for-Help sessions for this subject at the moment.
                    <br />
                    Please check again another time.
                </p>
            );
        }

        return (
            <div id="view-afh">
                <h1>Ask for Help</h1>
                <p className="description">
                    We provide free sessions for students to ask for additional
                    assistance on any of our program subjects. Please fill the
                    form to let us know you are coming. You must be registered
                    with one of our programs to attend.
                </p>

                <p className="directions">
                    Select a subject below to view available ask-for-help
                    sessions.
                </p>
                <div className="tabs">
                    <TabButton
                        onChangeTab={this.changeSubject}
                        highlight={this.state.currentSubject == "math"}
                        subject={"math"}
                    />
                    <TabButton
                        onChangeTab={this.changeSubject}
                        highlight={this.state.currentSubject == "english"}
                        subject={"english"}
                    />
                    <TabButton
                        onChangeTab={this.changeSubject}
                        highlight={this.state.currentSubject == "programming"}
                        subject={"programming"}
                    />
                </div>

                <div className="showTab">{showSessions}</div>
            </div>
        );
    }
}

class TabButton extends React.Component {
    render() {
        let highlight = this.props.highlight;
        let subject = this.props.subject;
        let displayName = subjectDisplayNames[subject];
        let iconSrc = subjectIconSrc[subject];

        return (
            <button
                className={subject + " " + (highlight ? "active" : "")}
                onClick={() => this.props.onChangeTab(subject)}>
                <div className="icon-wrapper">
                    <img className={subject} src={iconSrc} />
                </div>
                {displayName}
            </button>
        );
    }
}

class AfhSessionRow extends React.Component {
    render() {
        const row = this.props.row;
        return (
            <div className="afh-row">
                <div className="row-header">
                    <AfhInfo
                        title={row.title}
                        startsAt={row.startsAt}
                        endsAt={row.endsAt}
                    />
                    {createLocation(this.props.location || {})}
                    <Link
                        to={"/register?afhId=" + row.id}
                        className="link-wrapper">
                        Register to attend
                        <img src={srcPoint} />
                    </Link>
                </div>
                <AfhNotes notes={row.notes} />
            </div>
        );
    }
}

class AfhInfo extends React.Component {
    render() {
        const title = this.props.title;
        const startsAt = moment(this.props.startsAt);
        const endsAt = moment(this.props.endsAt);
        return (
            <div className="info">
                <h4 className="title">{title}</h4>
                <h4 className="time">
                    {startsAt.format("dddd M/D/YYYY")}
                    <br />
                    {startsAt.format("hh:mm a") +
                        " - " +
                        endsAt.format("hh:mm a")}
                </h4>
            </div>
        );
    }
}

class AfhNotes extends React.Component {
    render() {
        const notes = this.props.notes;
        if (notes) {
            return (
                <div className="notes">
                    <div className="icon-wrapper">
                        <img src={srcNotes} />
                    </div>
                    <b>Note from teacher:</b>
                    <p>{notes}</p>
                </div>
            );
        } else {
            return <div></div>;
        }
    }
}
