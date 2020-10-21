"use strict";
require("./classRegister.sass");
import React from "react";
import { Link } from "react-router-dom";
const FULL_STATE_VALUE = 2;

export class ClassRegister extends React.Component {
    render() {
        const classObj = this.props.classObj;
        const isFull = classObj.fullState == FULL_STATE_VALUE;

        let url = "";
        let message = "";
        if (isFull) {
            message =
                "Unfortunately, this class is full. You will not be able to enroll into this class. " +
                "Please consider enrolling into a different class.";
        } else {
            url = "/register?classId=" + classObj.classId;
            message =
                "If you are interested in this course, please click on Enroll.";
        }

        const fullState = classObj.fullState;
        let fullStateSection = <div></div>;
        if (fullState == 1) {
            // Almost Full
            fullStateSection = (
                <h4 className="full">
                    This class is almost full! Enroll now to reserve your spot.
                </h4>
            );
        } else if (fullState == FULL_STATE_VALUE) {
            // Full
            fullStateSection = (
                <h4 className="full">
                    Unfortunately, this class is full. Please consider
                    registering for another class.
                </h4>
                // TODO: show other available classes in same program/semester (if any)
            );
        }

        return (
            <section id="register">
                {fullStateSection}
                <p className={isFull ? "full" : ""}>{message}</p>
                <Link to={url} className={isFull ? "full" : ""}>
                    <button>Enroll</button>
                </Link>
            </section>
        );
    }
}
