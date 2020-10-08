"use strict";
require("./class.sass");
import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import moment from "moment";
import { isEmpty } from "lodash";
import API from "../utils/api.js";
import {
    getFullTitle,
    displayPrice,
    displayTimeStringOneLine,
    displayFeaturedString,
} from "../utils/classUtils.js";
import { capitalizeWord } from "../utils/displayUtils.js";
import { createLocation } from "../utils/locationUtils.js";
import { ClassSchedule } from "./classSchedule.js";
import { ClassErrorPage } from "./classError.js";

import srcStar from "../../assets/star_green.svg";

const FULL_STATE_VALUE = 2;

export class ClassPage extends React.Component {
    state = {
        classObj: {},
        sessions: [],
        program: {},
        semester: {},
        location: {},
        otherClasses: [],
        allSemesters: [],
        fetchedData: false,
    };

    componentDidMount = () => {
        this.fetchData(this.props.classId);
    };

    fetchData = (classId) => {
        const apiCallsPerClass = [
            API.get("api/classes/class/" + classId),
            API.get("api/sessions/class/" + classId),
        ];

        axios
            .all(apiCallsPerClass)
            .then(
                axios.spread((...responses) => this.fetchOtherData(responses))
            )
            .catch((err) => console.log("Error: could not fetch class: " + err))
            .finally(() => {
                this.setState({ fetchedData: true });
            });
    };

    fetchOtherData = (responses) => {
        const classObj = responses[0].data;
        const sessions = responses[1].data;
        const apiCalls = [
            API.get("api/programs/program/" + classObj.programId),
            API.get("api/semesters/all"),
            API.get("api/locations/location/" + classObj.locationId),
            API.get("api/classes/program/" + classObj.programId),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const program = responses[0].data;
                    const allSemesters = responses[1].data;
                    const semester = allSemesters.find(
                        (semester) => semester.semesterId == classObj.semesterId
                    );
                    this.setState({
                        classObj: classObj || {},
                        sessions: sessions || [],
                        program: program,
                        allSemesters: allSemesters,
                        semester: semester,
                        location: responses[2].data,
                        otherClasses: responses[3].data,
                        fetchedData: true,
                    });
                    document.title = getFullTitle(program, classObj, semester);
                })
            )
            .catch((err) =>
                console.log("Error: could not fetch other data: " + err)
            );
    };

    renderBreadcrumbs = () => {
        const program = this.state.program;
        const fullTitle = getFullTitle(program, this.state.classObj);
        return (
            <section id="breadcrumbs">
                <div>
                    <div>
                        <Link to="/programs">Program Catalog</Link>
                        <span>&middot;</span>
                        <span>{this.state.semester.title}</span>
                    </div>
                    <h1>{fullTitle}</h1>
                </div>

                {program.featured != "none" && (
                    <div className="featured">
                        <div className="header">
                            <img src={srcStar} />
                            <span>{capitalizeWord(program.featured)}</span>
                        </div>
                        <p>{displayFeaturedString(program)}</p>
                    </div>
                )}
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
            <section id="program-info">
                {fullStateSection}
                <p className="grades">
                    Grades: {program.grade1} - {program.grade2}
                </p>
                <p className="description">{program.description}</p>
            </section>
        );
    };

    renderOtherClasses = () => {
        // Unused for now. Looks awkward to be one of the first things you see on page.
        const thisClass = this.state.classObj;
        const otherClasses = this.state.otherClasses.filter((classObj) => {
            return (
                classObj.classId != thisClass.classId &&
                classObj.fullState != FULL_STATE_VALUE
            );
        });

        let semesterMap = {};
        this.state.allSemesters.forEach((semester) => {
            semesterMap[semester.semesterId] = semester;
        });
        const rows = otherClasses.map((classObj, index) => {
            return (
                <div key={index}>
                    <Link
                        to={"/class/" + classObj.classId}
                        onClick={() => this.fetchData(classObj.classId)}>
                        {semesterMap[classObj.semesterId].title}{" "}
                        {capitalizeWord(classObj.classKey)}
                    </Link>
                    {displayTimeStringOneLine(classObj)}
                </div>
            );
        });

        let message = "You may also be interested in these similar classes.";
        if (thisClass.fullState == FULL_STATE_VALUE) {
            message =
                "Unfortunately this class is full. " +
                "Please consider enrolling in one of these available similar classes.";
        }

        const hide = otherClasses.length == 0 ? "hide" : "";

        return (
            <section id="other-classes" className={hide}>
                <p>{message}</p>
                {rows}
            </section>
        );
    };

    renderRegisterBlock = () => {
        const classObj = this.state.classObj;
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
                "If you are interested in this course, please click on Enroll. " +
                "You will be asked to fill out some contact information. " +
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
            firstSession = <p>First session: To be determined</p>;
            lastSession = <p>Last session: To be determined</p>;
        }

        // Pricing information
        const priceString = displayPrice(classObj);

        return (
            <section id="class-info">
                <div className="block">
                    <h3 className="location">Location</h3>
                    <div id="class-location">{createLocation(location)}</div>
                </div>
                <div className="block">
                    <div id="class-times">
                        <h3 className="times">Times</h3>
                        {displayTimeStringOneLine(classObj)}
                        {firstSession}
                        {lastSession}
                    </div>

                    <div id="class-price">
                        <h3 className="pricing">Pricing</h3>
                        <div id="class-pricing">
                            <p>{priceString}</p>
                            <p>{classObj.paymentNotes}</p>
                        </div>
                    </div>
                </div>
            </section>
        );
    };

    render = () => {
        const classObj = this.state.classObj;
        const sessions = this.state.sessions;

        if (this.state.fetchedData && isEmpty(classObj))
            return <ClassErrorPage classId={this.props.classId} />;
        else
            return (
                <div id="view-class">
                    {this.renderBreadcrumbs()}
                    {this.renderProgramInfo()}
                    {/* {this.renderOtherClasses()} */}
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
