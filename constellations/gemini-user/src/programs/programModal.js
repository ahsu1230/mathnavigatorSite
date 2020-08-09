"use strict";
require("./programModal.sass");
import React from "react";
import { Link } from "react-router-dom";

export class ProgramModal extends React.Component {
    render = () => {
        const semester = this.props.semester;
        const classes = this.props.classes;

        const classList = classes.map((c, index) => (
            <ProgramClass key={index} classObj={c} />
        ));

        return (
            <div className="program-modal">
                <h1>{semester.title + " Classes"}</h1>
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
        const name =
            classObj.classKey +
            ["", " (ALMOST FULL)", " (FULL)"][classObj.fullState];
        const url = "/class/" + classObj.classId;

        const times = this.getTimes(classObj.times);

        return (
            <li>
                <div className="name">{name}</div>
                <div className="times">{times}</div>
                <Link to={url}>{"Details >"}</Link>
            </li>
        );
    };
}
