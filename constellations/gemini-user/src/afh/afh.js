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
            let sessionDate = moment(row.date).format("M/D/YYYY dddd");
            return (
                <div className="sessions-list" key={index}>
                    <div className="sessions-checkbox">
                        <CheckboxInput
                            key={index}
                            row={row}
                            onChangeCheckbox={this.onChangeCheckbox}
                        />
                    </div>

                    <div className="session-details">
                        {sessionDate} {row.timeString} <br />
                        {row.title} {row.notes} <br /> {row.locationId}
                    </div>
                </div>
            );
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
                <div className="tab">
                    <button
                        className={
                            this.state.currentTab == "Math" ? "active" : ""
                        }
                        onClick={() => this.openSubject("Math")}>
                        Math
                    </button>
                    <button
                        className={
                            this.state.currentTab == "English" ? "active" : ""
                        }
                        onClick={() => this.openSubject("English")}>
                        English
                    </button>
                    <button
                        className={
                            this.state.currentTab == "Computer Programming"
                                ? "active"
                                : ""
                        }
                        onClick={() =>
                            this.openSubject("Computer Programming")
                        }>
                        Computer Programming
                    </button>
                </div>

                <div
                    className={
                        this.state.currentTab == "Math" ? "showTab" : "hide"
                    }>
                    {showSessions}
                </div>

                <div
                    className={
                        this.state.currentTab == "English" ? "showTab" : "hide"
                    }>
                    {showSessions}
                </div>

                <div
                    className={
                        this.state.currentTab == "Computer Programming"
                            ? "showTab"
                            : "hide"
                    }>
                    {showSessions}
                </div>
            </div>
        );
    }
}

class CheckboxInput extends React.Component {
    selectCheckbox = (title) => {
        return (
            <input
                className="select"
                type="checkbox"
                onChange={this.onSelectSession(title)}
            />
        );
    };

    //planning to add a class name to the "selected" sessions so font color of session can be changed to turquoise
    onSelectSession = (title) => {
        let originalClassName = "sessions-list";
        this.className = originalClassName + "active";
    };

    render() {
        const row = this.props.row;
        const checkbox = this.selectCheckbox(row.title);

        return <div className="checkbox">{checkbox}</div>;
    }
}
