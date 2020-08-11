"use strict";
require("./classSchedule.sass");
import React from "react";
import moment from "moment";

export class ClassSchedule extends React.Component {
    getSchedules = (sessions) => {
        const schedules = sessions.map((session, index) => {
            const startTime = moment(session.startsAt);
            const endTime = moment(session.endsAt);
            const times =
                startTime.format("dddd h:mma") +
                " - " +
                endTime.format("h:mma");

            const canceled = session.canceled;
            var state = "row";
            if (moment().isAfter(endTime))
                state += canceled ? " pink" : " gray";
            else state += canceled ? " red" : "";

            return (
                <div key={index} className={state}>
                    <span className="index-column">{index + 1}</span>
                    <span className="date-column">{startTime.format("l")}</span>
                    <span className="time-column">{times}</span>
                    <span className="canceled">
                        {canceled ? "Canceled" : ""}
                    </span>
                </div>
            );
        });
        return schedules;
    };

    render = () => {
        const sessions = this.props.sessions;

        return (
            <div id="class-session">
                <h3>Schedule</h3>
                {this.getSchedules(sessions)}
            </div>
        );
    };
}
