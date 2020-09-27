"use strict";
require("./afh.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../utils/api.js";
import moment from "moment";

const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Programming",
};

export class AFHPage extends React.Component {
    state = {
        currentSubject: "math",
        sessions: [],
    };

    componentDidMount() {
        API.get("api/askforhelp/all").then((res) => {
            const afh = res.data;
            this.setState({
                sessions: afh,
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
            return <AfhSessionRow key={index} row={row} />;
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

                <p>
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

        return (
            <button
                className={highlight ? "active" : ""}
                onClick={() => this.props.onChangeTab(subject)}>
                {displayName}
            </button>
        );
    }
}

class LocationAddress extends React.Component {
    state = {
        location: {},
    };
    componentDidMount() {
        API.get("api/locations/location/" + this.props.locationId).then(
            (res) => {
                const currentLocation = res.data;
                this.setState({
                    location: currentLocation,
                });
            }
        );
    }

    render() {
        const address1 = this.state.location.street;
        const address2 =
            this.state.location.city +
            ", " +
            this.state.location.state +
            " " +
            this.state.location.zipcode;
        const room = this.state.location.room;

        return (
            <div className="location">
                {address1} <br /> {address2} <br /> {room}
            </div>
        );
    }
}

class AfhSessionRow extends React.Component {
    render() {
        const row = this.props.row;
        const startsAt = moment(row.startsAt);
        const endsAt = moment(row.endsAt);

        return (
            <div className="afh-row">
                <div className="row-header">
                    <h3>
                        {startsAt.format("dddd M/D/YYYY") +
                            "  (" +
                            startsAt.format("hh:mm a") +
                            " - " +
                            endsAt.format("hh:mm a") +
                            ")"}
                    </h3>
                    <Link to="register">Register to attend</Link>
                </div>
                <div>{row.title}</div>
                <LocationAddress locationId={row.locationId} />

                {row.notes ? (
                    <div className="notes">
                        <b>Note from teacher:</b>
                        <p>{row.notes}</p>
                    </div>
                ) : (
                    <div></div>
                )}
            </div>
        );
    }
}
