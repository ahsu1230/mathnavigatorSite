"use strict";
require("./sessionList.sass");
import React from "react";
import { Link } from "react-router-dom";

export class SessionList extends React.Component {
    render = () => {
        const classId = this.props.classId;
        const sessions = this.props.sessions;
        const rows = sessions.map((session, index) => {
            return (
                <div className="row" key={index}>
                    <span className="column">
                        {session.startsAt.format("l")}
                    </span>
                    <span className="wide-column">
                        {session.startsAt.format("LT")}
                        {" - "}
                        {session.endsAt.format("LT")}
                    </span>
                    <span className="column">
                        {session.canceled ? "true" : ""}
                    </span>
                    <span className="wide-column">{session.notes}</span>
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
                    <span className="wide-column">Time</span>
                    <span className="column">Canceled</span>
                    <span className="wide-column">Notes</span>
                    <span className="edit"></span>
                </div>
                {rows}
            </section>
        );
    };
}
