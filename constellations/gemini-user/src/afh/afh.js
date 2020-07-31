"use strict";
require("./afh.sass");
import React from "react";
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

    openSubject = (subjectName) => {
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

        return (
            <div id="view-afh">
                <h1>Ask for Help</h1>
                <div className="description">
                    We provide free sessions for students to ask for additional
                    assistance on any of our program subjects. Please fill the
                    form to let us know you are coming. You must be registered
                    with one of our programs to attend.
                </div>

                <h1>Ask for Help Sessions by Subject</h1>
                <div className="tabs">
                    <TabButton
                        onChangeTab={this.openSubject}
                        highlight={this.state.currentSubject == "math"}
                        subject={"math"}
                    />
                    <TabButton
                        onChangeTab={this.openSubject}
                        highlight={this.state.currentSubject == "english"}
                        subject={"english"}
                    />
                    <TabButton
                        onChangeTab={this.openSubject}
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

class AfhSessionRow extends React.Component {
    state = {
        isActive: false,
    };

    onSelectSession = () => {
        this.setState({
            isActive: !this.state.isActive,
        });
    };

    render() {
        const row = this.props.row;
        let sessionDate = moment(row.date).format("M/D/YYYY dddd");

        return (
            <div
                className={
                    this.state.isActive
                        ? "sessions-list-active"
                        : "sessions-list"
                }>
                <div className="sessions-checkbox">
                    <input
                        className="select"
                        type="checkbox"
                        onChange={this.onSelectSession}
                    />
                </div>

                <div className="session-details">
                    {sessionDate} {row.timeString} <br />
                    {row.title} {row.notes} <br /> {row.locationId}
                </div>
            </div>
        );
    }
}
