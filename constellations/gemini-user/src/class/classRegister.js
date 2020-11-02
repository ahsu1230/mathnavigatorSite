"use strict";
require("./classRegister.sass");
import React from "react";
import { Link } from "react-router-dom";

const DEFAULT_FULL_STATE_VALUE = 0;
const ALMOST_FULL_STATE_VALUE = 1;
const FULL_STATE_VALUE = 2;

export class ClassRegister extends React.Component {
    render() {
        const classObj = this.props.classObj;
        const fullState = classObj.fullState;
        const isFull = classObj.fullState == FULL_STATE_VALUE;

        let url = "/";
        let message = "";

        if (fullState == ALMOST_FULL_STATE_VALUE) {
            message =
                "This class is almost full! Enroll now to reserve your spot.";
        } else if (fullState == FULL_STATE_VALUE) {
            message =
                "Unfortunately, this class is full. Please try registering for a different class.";
            // TODO: show other available classes in same program/semester (if any)
        } else {
            url = "/register?classId=" + classObj.classId;
            message =
                "If you are interested in this course, please click on Enroll.";
        }

        return (
            <section id="register" className={isFull ? "full" : ""}>
                <h4>{message}</h4>
                <Link to={url}>
                    <button>Enroll</button>
                </Link>
            </section>
        );
    }
}
