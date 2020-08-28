"use strict";
require("./programModal.sass");
import React from "react";
import { Link } from "react-router-dom";
import { getFullStateName } from "../utils/utils.js";

export class ProgramModal extends React.Component {
    render = () => {
        const semester = this.props.semester || {};
        const program = this.props.program || {};
        const classes = this.props.classes || [];

        const classList = classes.map((c, index) => (
            <ProgramClass key={index} classObj={c} />
        ));

        return (
            <div className="program-modal">
                <h1>{"Classes for " + program.name}</h1>
                <h4>{semester.title}</h4>
                <ul>{classList}</ul>
            </div>
        );
    };
}

class ProgramClass extends React.Component {
    getTimes = (times) => {
        return times
            .split(",")
            .map((time, index) => <span key={index}>{time.trim()}</span>);
    };

    render = () => {
        const classObj = this.props.classObj;
        const fullState = classObj.fullState;
        const name = classObj.classKey + getFullStateName(fullState, true);
        const url = "/class/" + classObj.classId;

        const times = this.getTimes(classObj.times);

        return (
            <li>
                <div className={fullState ? "name full" : "name"}>{name}</div>
                <div className="times">{times}</div>
                <Link to={url}>{"Details >"}</Link>
            </li>
        );
    };
}
