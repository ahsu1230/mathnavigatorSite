"use strict";
require("./session.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { getCurrentClassId, setCurrentClassId } from "./../localStorage.js";
import { SessionAdd } from "./sessionAdd.js";
import RowCardBasic from "../common/rowCards/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    A session represents a single weekly (or bi-weekly) class session that students attend. 
    A class is made up of many sessions so make sure to create a class before creating sessions for it.
    Every session has a date and time and a status.
    You may add a short note for every session (i.e. room change or a topic-of-the-day).
    Sessions are displayed in the "Class" page of the user website.`;
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
            sessionsList = this.state.sessions.map((session, index) => {
                const startsAt = moment(session.startsAt);
                const fields = generateFields(session);
                const texts = generateTexts(session);
                return (
                    <RowCardBasic
                        key={index}
                        title={startsAt.format("l")}
                        editUrl={
                            "/sessions/" +
                            session.classId +
                            "/" +
                            session.id +
                            "/edit"
                        }
                        fields={fields}
                        texts={texts}
                    />
                );
            });
        } else {
            sessionsList = <div>No sessions for this class</div>;
        }

        return (
            <div id="view-session">
                <section className="header-section">
                    <h1>Sessions per class</h1>
                    <p>{PAGE_DESCRIPTION}</p>
                </section>
                <section id="select-class">
                    <h2>Select Class</h2>
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
                <div id="sessions-list">
                    <h3>Sessions for '{this.state.classId}'</h3>
                    {sessionsList}
                </div>
            </div>
        );
    };
}

function generateFields(session) {
    const startTime = moment(session.startsAt).format("h:mm a");
    const endTime = moment(session.endsAt).format("h:mm a");
    let status;
    if (session.canceled) {
        status = "Canceled";
    } else if (moment().isBefore(session.startsAt)) {
        status = "Scheduled";
    } else if (moment().isBetween(session.startsAt, session.endsAt)) {
        status = "In progress";
    } else {
        status = "Done";
    }

    return [
        {
            label: "Times",
            value: startTime + " - " + endTime,
        },
        {
            label: "Status",
            value: status,
            highlightFn: () => session.canceled,
        },
    ];
}

function generateTexts(session) {
    return [
        {
            label: "Notes",
            value: session.notes,
        },
    ];
}
