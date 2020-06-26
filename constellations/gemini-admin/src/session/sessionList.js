"use strict";
require("./sessionList.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class SessionList extends React.Component {
    render = () => {
        const classId = this.props.classId;
        const sessions = this.props.sessions;
        const rows = sessions.map((session, index) => {
            return (
                <div className="row" key={index}>
                    <span className="column">
                        {moment(session.startsAt).format("l")}
                    </span>
                    <span className="medium-column">
                        {moment(session.startsAt).format("LT")}
                        {" - "}
                        {moment(session.endsAt).format("LT")}
                    </span>
                    <span className="column">
                        {session.canceled
                            ? "Canceled"
                            : moment().isAfter(session.startsAt)
                            ? "Done"
                            : "Scheduled"}
                    </span>
                    <span className="large-column">{session.notes}</span>
                    <span className="edit">
                        <Link
                            to={
                                "/sessions/" +
                                classId +
                                "/" +
                                session.id +
                                "/edit"
                            }>
                            {"Edit >"}
                        </Link>
                    </span>
                </div>
            );
        });

        return (
            <section id="session-rows">
                <div id="header" className="row">
                    <span className="column">Date</span>
                    <span className="medium-column">Time</span>
                    <span className="column">Status</span>
                    <span className="large-column">Notes</span>
                    <span className="edit"></span>
                </div>
                {rows}
            </section>
        );
    };
}
