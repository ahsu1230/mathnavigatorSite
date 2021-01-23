"use strict";
require("./programModal.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import API from "../utils/api.js";
import { capitalizeWord } from "../utils/displayUtils.js";
import { displayTimeString } from "../utils/classUtils.js";

export class ProgramModal extends React.Component {
    render = () => {
        const semester = this.props.semester || {};
        const program = this.props.program || {};
        const classes = this.props.classes || [];
        const fullStates = this.props.fullStates || [];
        const classList = classes.map((c, index) => (
            <ProgramClass key={index} classObj={c} fullStates={fullStates} />
        ));
        const classesString =
            classes.length == 1 ? "1 class" : classes.length + " classes";

        return (
            <div className="program-modal">
                <h1>{"Classes for " + program.title}</h1>
                <h4>
                    {semester.title} ({classesString} available)
                </h4>
                <ul>{classList}</ul>
            </div>
        );
    };
}

class ProgramClass extends React.Component {
    state = {
        firstSession: {},
    };

    componentDidMount() {
        const classId = this.props.classObj.classId;
        API.get("api/sessions/class/" + classId).then((res) => {
            const sessions = res.data || [];
            this.setState({
                firstSession: sessions[0] || {},
            });
        });
    }

    render = () => {
        const classObj = this.props.classObj;
        const classTitle = capitalizeWord(classObj.classKey);
        const url = "/class/" + classObj.classId;

        const fullState = classObj.fullState;
        let fullStateStr = "";
        if (this.props.fullStates.length > 0 && fullState) {
            fullStateStr = "(" + this.props.fullStates[fullState] + ")";
        }

        const startsOnDate = this.state.firstSession.startsAt
            ? moment(this.state.firstSession.startsAt).format("l")
            : undefined;

        return (
            <li>
                <div>
                    <div className={"name" + (fullState ? " full" : "")}>
                        {classTitle}
                    </div>
                    <div className="full-state">{fullStateStr}</div>
                </div>
                <div>
                    {startsOnDate && <div>Starts on: {startsOnDate}</div>}
                    {displayTimeString(classObj)}
                </div>
                <Link to={url}>{"Details >"}</Link>
            </li>
        );
    };
}
