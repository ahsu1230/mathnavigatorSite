"use strict";
require("./classSessionList.styl");
import React from "react";
import moment from "moment";

export class ClassSessionList extends React.Component {
    render() {
        const list = this.props.sessions.map((session, index) => (
            <ClassSessionRow
                key={index}
                session={session}
                onDeleteSession={this.props.onDeleteSession}
            />
        ));
        let content;
        if (list.length > 0) {
            content = (
                <div>
                    <div className="list-row sub-header">
                        <span className="med">Date</span>
                        <span className="med">Start Time</span>
                        <span className="med">End Time</span>
                        <span className="small">Status</span>
                        <span className="big">Notes</span>
                    </div>
                    <ul id="session-list">{list}</ul>
                </div>
            );
        } else {
            content = <h4>No sessions scheduled yet</h4>;
        }

        return (
            <div id="session-list-container">
                <h2 className="list-header">Sessions</h2>
                {content}
            </div>
        );
    }
}

class ClassSessionRow extends React.Component {
    render() {
        const session = this.props.session;
        const startsAt = moment(session.startsAt);
        const endsAt = moment(session.endsAt);
        const canceled = session.canceled ? "Canceled" : "Normal";
        const notes = session.notes || "";
        return (
            <li className="list-row">
                <span className="med">
                    {startsAt.format("ddd, MM/DD/YYYY")}
                </span>
                <span className="med">{startsAt.format("hh:mm a")}</span>
                <span className="med">{endsAt.format("hh:mm a")}</span>
                <span className="small">{canceled}</span>
                <span className="big">{notes}</span>
                <button onClick={() => this.props.onDeleteSession(session.id)}>
                    Delete
                </button>
            </li>
        );
    }
}
