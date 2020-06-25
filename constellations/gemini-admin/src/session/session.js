"use strict";
require("./session.sass");
import React from "react";
import API from "../api.js";
import { SessionAdd } from "./sessionAdd.js";
import { SessionList } from "./sessionList.js";

export class SessionPage extends React.Component {
    state = {
        classes: [],
        classId: "",
        sessions: [],
    };

    componentDidMount = () => {
        this.fetchData();
    };

    fetchData = () => {
        API.get("api/classes/all").then((res) => {
            const classes = res.data;
            this.setState({
                classes: classes,
                classId: classes.length > 0 ? classes[0].classId : "",
            });
        });
    };

    onChangeSelect = (e) => {
        this.setState({
            classId: e.target.value,
        });
        console.log("Selected classId: " + e.target.value);
    };

    addSessions = (sessions) => {
        var newSessions = _.concat(this.state.sessions, sessions);
        newSessions = _.sortBy(newSessions, ["startsAt"]);

        this.setState({
            sessions: newSessions,
        });
        console.log("Added sessions");
    };

    render = () => {
        const classOptions = this.state.classes.map((row, index) => {
            return (
                <option value={row.classId} key={index}>
                    {row.classId}
                </option>
            );
        });

        return (
            <div id="view-session">
                <section id="select-class">
                    <h1>Select Class</h1>
                    <select
                        id="dropdown"
                        onChange={(e) => this.onChangeSelect(e)}>
                        {classOptions}
                    </select>
                </section>
                <SessionAdd addSessions={this.addSessions} />
                <SessionList
                    classId={this.state.classId}
                    sessions={this.state.sessions}
                />
            </div>
        );
    };
}
