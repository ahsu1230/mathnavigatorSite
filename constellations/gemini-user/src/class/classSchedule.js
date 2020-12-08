"use strict";
require("./classSchedule.sass");
import React from "react";
import moment from "moment";
import { convertToOrdinal } from "../utils/displayUtils.js";

export class ClassSchedule extends React.Component {
    getSchedules = (sessions) => {
        let sessionIndex = 1;
        const schedules = sessions.map((session, index) => {
            const canceled = session.canceled;
            if (!canceled) {
                session.sessionIndex = sessionIndex;
                sessionIndex++;
            }

            const startTime = moment(session.startsAt);
            const endTime = moment(session.endsAt);
            const times =
                startTime.format("dddd h:mma") +
                " - " +
                endTime.format("h:mma");

            let state = "row";
            let now = moment();
            if (now.isAfter(endTime)) {
                state += canceled ? " canceled-past" : " past";
            } else {
                state += canceled ? " canceled" : "";
            }

            return (
                <div key={index} className={state}>
                    <span className="index-column">
                        {session.sessionIndex &&
                            convertToOrdinal(session.sessionIndex)}
                    </span>
                    <span className="date-column">{startTime.format("l")}</span>
                    <span className="time-column">{!canceled && times}</span>
                    <span className="notes">
                        {canceled ? "No Class. " : ""}
                        {session.notes}
                    </span>
                </div>
            );
        });
        return schedules;
    };

    render = () => {
        const sessions = this.props.sessions || [];

        let content;
        if (sessions.length) {
            content = <div>{this.getSchedules(sessions)}</div>;
        } else {
            content = <p>To be determined. Please check again soon!</p>;
        }

        return (
            <div id="class-session">
                <h3>Schedule</h3>
                {content}
            </div>
        );
    };
}
