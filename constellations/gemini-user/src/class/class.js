"use strict";
require("./class.sass");
import React from "react";
import moment from "moment";
import axios from "axios";
import API from "../utils/api.js";
import { formatCurrency, getFullStateName } from "../utils/utils.js";
import { Link } from "react-router-dom";
import { ClassSchedule } from "./classSchedule.js";
import { ErrorPage } from "../error/error.js";

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
                        classObj: classObj,
                        sessions: sessions,
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

    getTitle = (program, semester, classObj) => {
        let name = program.name + " " + semester.title;
        name += classObj.classKey ? " " + classObj.classKey : "";

        return (
            <h1>
                <Link to="/programs">Programs</Link>
                {" > " + name}
            </h1>
        );
    };

    getLocation = (location) => {
        const address =
            location.city + ", " + location.state + " " + location.zipcode;
        return (
            <div id="class-location">
                <p>{location.street}</p>
                <p>{address}</p>
                <p>{location.room}</p>
            </div>
        );
    };

    getTimes = (classObj, sessions) => {
        const times = (classObj.times || "")
            .split(", ")
            .map((time, index) => <p key={index}>{time}</p>);

        let startEndDate = <p>No sessions scheduled</p>;
        if (sessions.length > 0) {
            const startTime = moment(sessions[0].startsAt);
            const endTime = moment(sessions[sessions.length - 1].startsAt);
            startEndDate = (
                <div id="class-times">
                    <p>{"First Session: " + startTime.format("l")}</p>
                    <p>{"Last Session: " + endTime.format("l")}</p>
                </div>
            );
        }

        return (
            <div>
                {times}
                {startEndDate}
            </div>
        );
    };

    getPricing = (classObj) => {
        const isLump = !!classObj.priceLump;
        const title = isLump ? "Total Price: " : "Price per session: ";
        const price = formatCurrency(
            isLump ? classObj.priceLump : classObj.pricePerSession
        );

        return (
            <div id="class-pricing">
                <p>{title + price}</p>
                <p>{classObj.paymentNotes}</p>
            </div>
        );
    };

    render = () => {
        const classObj = this.state.classObj;
        const program = this.state.program;
        const semester = this.state.semester;
        const location = this.state.location;
        const sessions = this.state.sessions;
        const url = "/contact?interest=" + classObj.classId;

        if (_.isEmpty(classObj))
            return <ErrorPage classId={this.props.classId} />;
        else
            return (
                <div id="view-class">
                    {this.getTitle(program, semester, classObj)}

                    <div id="class-description">
                        <h4 className="red">
                            {getFullStateName(classObj.fullState)}
                        </h4>
                        <h4>
                            Grades: {program.grade1} - {program.grade2}
                        </h4>
                        <p>{program.description}</p>
                    </div>

                    <div id="class-info">
                        <Link to={url}>
                            <button>Register</button>
                        </Link>

                        <h3>Location</h3>
                        {this.getLocation(location)}
                        <h3>Times</h3>
                        {this.getTimes(classObj, sessions)}
                        <h3>Pricing</h3>
                        {this.getPricing(classObj)}
                    </div>

                    <ClassSchedule sessions={sessions} />

                    <div id="class-questions">
                        <div>
                            Questions? <Link to={url}>Contact Us</Link>
                        </div>
                        <Link to="/programs">
                            <button className="inverted">
                                {"< More Programs"}
                            </button>
                        </Link>
                    </div>
                </div>
            );
    };
}
