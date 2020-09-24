"use strict";
require("./class.sass");
import React from "react";
import moment from "moment";
import axios from "axios";
import { isEmpty } from "lodash";
import API from "../utils/api.js";
import {
    formatCurrency,
    getFullStateName,
    capitalizeWord,
} from "../utils/utils.js";
import { Link } from "react-router-dom";
import { ClassSchedule } from "./classSchedule.js";
import { ClassErrorPage } from "./classError.js";

export class ClassPage extends React.Component {
    state = {
        classObj: {},
        sessions: [],
        program: {},
        semester: {},
        location: {},
    };

    componentDidMount = () => {
        const classId = this.props.classId;
        const apiCallsPerClass = [
            API.get("api/classes/class/" + classId),
            API.get("api/sessions/class/" + classId),
        ];

        axios
            .all(apiCallsPerClass)
            .then(
                axios.spread((...responses) => this.fetchOtherData(responses))
            )
            .catch((err) =>
                console.log("Error: could not fetch class: " + err)
            );
    };

    fetchOtherData = (responses) => {
        const classObj = responses[0].data;
        const sessions = responses[1].data;
        const apiCalls = [
            API.get("api/programs/program/" + classObj.programId),
            API.get("api/semesters/semester/" + classObj.semesterId),
            API.get("api/locations/location/" + classObj.locationId),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) =>
                    this.setState({
                        classObj: classObj || {},
                        sessions: sessions || [],
                        program: responses[0].data,
                        semester: responses[1].data,
                        location: responses[2].data,
                    })
                )
            )
            .catch((err) =>
                console.log("Error: could not fetch other data: " + err)
            );
    };

    renderBreadcrumbs = () => {
        const fullTitle =
            this.state.program.title +
            " " +
            capitalizeWord(this.state.classObj.classKey);
        return (
            <section id="breadcrumbs">
                <div>
                    <Link to="/programs">Program Catalog</Link>
                    <span>&middot;</span>
                    <span>{this.state.semester.title}</span>
                </div>
                <h1>{fullTitle}</h1>
            </section>
        );
    };

    renderProgramInfo = () => {
        const program = this.state.program;
        const fullState = this.state.classObj.fullState;

        let fullStateSection = <div></div>;
        if (fullState == 1) {
            // Almost Full
            fullStateSection = (
                <h4 className="full">
                    This class is almost full! Enroll now to reserve your spot.
                </h4>
            );
        } else if (fullState == 2) {
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
            <section id="program-info">
                {fullStateSection}
                <p>
                    Grades: {program.grade1} - {program.grade2}
                </p>
                <p>{program.description}</p>
            </section>
        );
    };

    renderRegisterBlock = () => {
        const classObj = this.state.classObj;
        const isFull = classObj.fullState == 2;

        let url = "";
        let message = "";
        if (isFull) {
            message =
                "Unfortunately, this class is full. You will not be able to enroll into this class. " +
                "Please consider enrolling into a different class.";
        } else {
            url = "/contact?interest=" + classObj.classId;
            message =
                'If you are interested in this course, please click on "Enroll". ' +
                "For students to keep their enrollment, payment is due by the first class session.";
        }

        return (
            <section id="register">
                <p className={isFull ? "full" : ""}>{message}</p>
                <Link to={url} className={isFull ? "full" : ""}>
                    <button>Enroll</button>
                </Link>
            </section>
        );
    };

    renderClassInfo = () => {
        const classObj = this.state.classObj;

        // Location information
        const location = this.state.location;
        const address =
            location.city + ", " + location.state + " " + location.zipcode;

        // Timing information
        const sessions = this.state.sessions;
        let firstSession;
        let lastSession;
        if (sessions.length) {
            let firstSessionDate = moment(sessions[0].startsAt).format("l");
            let lastSessionDate = moment(
                sessions[sessions.length - 1].startsAt
            ).format("l");
            firstSession = <p>First session: {firstSessionDate}</p>;
            lastSession = <p>Last session: {lastSessionDate}</p>;
        } else {
            firstSession = <p>To be determined</p>;
            lastSession = <p>To be determined</p>;
        }

        // Pricing information
        const isLump = !!classObj.priceLumpSum;
        const priceLabel = isLump ? "Total Price: " : "Price per session: ";
        const price = formatCurrency(
            isLump ? classObj.priceLumpSum : classObj.pricePerSession
        );

        return (
            <section id="class-info">
                <div className="block">
                    <h3 className="location">Location</h3>
                    <div id="class-location">
                        <p>{location.title}</p>
                        <p>{location.street}</p>
                        <p>{address}</p>
                    </div>
                </div>
                <div className="block">
                    <h3 className="times">Times</h3>
                    <p>{classObj.timesStr}</p>
                    {firstSession}
                    {lastSession}

                    <h3 className="pricing">Pricing</h3>
                    <div id="class-pricing">
                        <p>{priceLabel + price}</p>
                        <p>{classObj.paymentNotes}</p>
                    </div>
                </div>
            </section>
        );
    };

    render = () => {
        const classObj = this.state.classObj;
        const sessions = this.state.sessions;

        if (isEmpty(classObj))
            return <ClassErrorPage classId={this.props.classId} />;
        else
            return (
                <div id="view-class">
                    {this.renderBreadcrumbs()}
                    {this.renderProgramInfo()}
                    {this.renderRegisterBlock()}
                    {this.renderClassInfo()}
                    <ClassSchedule sessions={sessions} />
                    <div id="class-footer">
                        <Link to="/programs">{"< More Programs"}</Link>
                        <p>
                            Still have questions about our programs and classes?
                            <br />
                            Email us at <u>andymathnavigator@gmail.com</u>.
                        </p>
                    </div>
                </div>
            );
    };
}
