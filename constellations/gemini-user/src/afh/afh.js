"use strict";
require("./afh.sass");
import React from "react";
import API from "../utils/api.js";
import moment from "moment";

export class AFHPage extends React.Component {
    state = {
        currentTab: "Math",
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
            currentTab: subjectName,
        });
    };

    render() {
        let currentSub = this.state.sessions.filter(
            (session) => session.subject == this.state.currentTab
        );

        let showSessions = currentSub.map((row, index) => {
            return <AfhSessionRow key={index} row={row} />;
        });

        const subjectDisplayNames = {
            math: "Math",
            english: "English",
            computer: "Computer Programming",
        };

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
                <div className="tab">
                    <TabButton
                        currentTab={this.state.currentTab}
                        subjectDisplayNames={this.subjectDisplayNames}
                    />
                </div>

                <div
                    className={
                        this.state.currentTab == currentSub ? "showTab" : "hide"
                    }>
                    {showSessions}
                </div>
            </div>
        );
    }
}

class TabButton extends React.Component {
    render() {
        let currentTab = this.props.currentTab;
        const subjectDisplayNames = this.props.subjectDisplayNames;

        return (
            <button
                className={
                    currentTab == subjectDisplayNames[currentTab]
                        ? "active"
                        : ""
                }
                onClick={() =>
                    this.openSubject(subjectDisplayNames[currentTab])
                }>
                {subjectDisplayNames[currentTab]}
            </button>
        );
    }
}

class AfhSessionRow extends React.Component {
    state = {
        isActive: false,
    };

    onSelectSession = () => {
        if (isActive == false) {
            this.setState({
                isActive: true,
            });
        } else {
            this.setState({
                isActive: false,
            });
        }
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
