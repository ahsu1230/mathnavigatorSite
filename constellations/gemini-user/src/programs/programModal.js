"use strict";
require("./programModal.sass");
import React from "react";
import { Link } from "react-router-dom";
import { capitalizeWord } from "../utils/utils.js";

export class ProgramModal extends React.Component {
    render = () => {
        const semester = this.props.semester || {};
        const program = this.props.program || {};
        const classes = this.props.classes || [];
        const fullStates = this.props.fullStates || [];

        const classList = classes.map((c, index) => (
            <ProgramClass key={index} classObj={c} fullStates={fullStates} />
        ));

        return (
            <div className="program-modal">
                <h1>{"Classes for " + program.title}</h1>
                <h4>{semester.title}</h4>
                <ul>{classList}</ul>
            </div>
        );
    };
}

class ProgramClass extends React.Component {
    render = () => {
        const classObj = this.props.classObj;
        const classTitle = capitalizeWord(classObj.classKey);
        const url = "/class/" + classObj.classId;

        const fullState = classObj.fullState;
        let fullStateStr = "";
        if (this.props.fullStates.length > 0 && fullState) {
            fullStateStr = "(" + this.props.fullStates[fullState] + ")";
        }

        const priceStr = classObj.priceLumpSum
            ? classObj.priceLumpSum + " total"
            : classObj.pricePerSession + " per session";

        return (
            <li>
                <div>
                    <div className={"name" + (fullState ? " full" : "")}>
                        {classTitle}
                    </div>
                    <div className="full-state">{fullStateStr}</div>
                </div>
                <div>
                    <div className="times">{classObj.timesStr}</div>
                    <div className="price">${priceStr}</div>
                </div>

                <Link to={url}>{"Details >"}</Link>
            </li>
        );
    };
}
