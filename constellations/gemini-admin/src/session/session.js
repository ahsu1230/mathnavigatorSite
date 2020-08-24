"use strict";
require("./session.sass");
import React from "react";
import API from "../api.js";
import { getCurrentClassId, setCurrentClassId } from "./../localStorage.js";
import { SessionAdd } from "./sessionAdd.js";
import { SessionList } from "./sessionList.js";

export class SessionPage extends React.Component {
    state = {
        classes: [],
        classId: "",
        sessions: [],
    };

    componentDidMount = () => {
        API.get("api/classes/all").then((res) => {
            const classes = res.data;
            const classId =
                getCurrentClassId() ||
                (classes.length > 0 ? classes[0].classId : "");
            if (classId != "") {
                setCurrentClassId(classId);
            }

            this.setState({
                classes: classes,
                classId: classId,
            });
            this.fetchSessionData(classId);
        });
    };

    fetchSessionData = (classId) => {
        API.get("api/sessions/class/" + classId).then((res) =>
            this.setState({
                sessions: res.data,
            })
        );
    };

    onChangeSelect = (e) => {
        const classId = e.target.value;
        this.setState({
            classId: classId,
        });
        setCurrentClassId(classId);
        this.fetchSessionData(classId);
        console.log("Selected classId: " + classId);
    };

    addSessions = (sessions) => {
        var sessionsJSON = [];
        sessions.forEach((session) => {
            sessionsJSON.push({
                classId: this.state.classId,
                startsAt: session.startsAt.toJSON(),
                endsAt: session.endsAt.toJSON(),
                canceled: false,
                notes: "",
            });
        });

        API.post("api/sessions/create", sessionsJSON)
            .then(() => {
                console.log("Successfully created sessions!");
            })
            .catch((err) => {
                window.alert("Could not create sessions: " + err);
            })
            .finally(() => {
                this.fetchSessionData(this.state.classId);
            });
    };

    render = () => {
        const classOptions = this.state.classes.map((row, index) => {
            return (
                <option value={row.classId} key={index}>
                    {row.classId}
                </option>
            );
        });

        let sessionsList = <div></div>;
        if (this.state.sessions.length != 0) {
            sessionsList = (
                <SessionList
                    classId={this.state.classId}
                    sessions={this.state.sessions}
                />
            );
        }

        return (
            <div id="view-session">
                <section id="select-class">
                    <h1>Select Class</h1>
                    <select
                        id="dropdown"
                        value={this.state.classId}
                        onChange={(e) => this.onChangeSelect(e)}>
                        {classOptions}
                    </select>
                </section>
                <SessionAdd
                    classId={this.state.classId}
                    addSessions={this.addSessions}
                />
                {sessionsList}
            </div>
        );
    };
}
