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

    render() {
        let currentSub = this.state.sessions.filter(
            (session) => session.subject == this.state.currentTab
        );
        let showSessions = currentSub.map((row, index) => {
            return (
                <li key={index}>
                    {row.date}
                    {row.timeString} <br />
                    {row.title} {row.notes} <br /> {row.locationId}
                </li>
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
                <div class="tab">
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
                    <div>{showSessions}</div>
                </div>

                <div
                    className={
                        this.state.currentTab == "English" ? "showTab" : "hide"
                    }>
                    <div>{showSessions}</div>
                </div>

                <div
                    className={
                        this.state.currentTab == "Computer Programming"
                            ? "showTab"
                            : "hide"
                    }>
                    <div>{showSessions}</div>
                </div>
            </div>
        );
    }
}
