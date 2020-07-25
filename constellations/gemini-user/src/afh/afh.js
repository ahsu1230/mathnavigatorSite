"use strict";
require("./afh.sass");
import React from "react";
import API from "../utils/api.js";

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

    //planning to add a class name to the "selected" sessions so font color of session can be changed to turquoise

    onSelectSession = (title) => {
        var session = title;
        session.checked = true;

        let originalClassName = "sessions-list";
        originalClassName = originalClassName + "active";
    };

    selectCheckbox = (title, onSelectSession) => {
        return (
            <input
                className="select"
                type="checkbox"
                onChange={() => onSelectSession(title)}
            />
        );
    };

    render() {
        let currentSub = this.state.sessions.filter(
            (session) => session.subject == this.state.currentTab
        );
        let showSessions = currentSub.map((row, index) => {
            return (
                <div className="sessions-list" key={index}>
                    {this.selectCheckbox}
                    {row.date} {row.timeString} <br />
                    {row.title} {row.notes} <br /> {row.locationId}
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
