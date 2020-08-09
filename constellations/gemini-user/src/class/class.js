"use strict";
require("./class.sass");
import React from "react";
import moment from "moment";
import axios from "axios";
import API from "../utils/api.js";
import { Link } from "react-router-dom";

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
        const apiClassCalls = [
            API.get("api/classes/class/" + classId),
            API.get("api/sessions/class/" + classId),
        ];

        let classObj;
        let sessions;
        axios
            .all(apiClassCalls)
            .then(
                axios.spread((...responses) => {
                    classObj = responses[0].data;
                    sessions = _.sortBy(responses[1].data, ["startsAt"]);

                    const apiCalls = [
                        API.get("api/programs/program/" + classObj.programId),
                        API.get(
                            "api/semesters/semester/" + classObj.semesterId
                        ),
                        API.get(
                            "api/locations/location/" + classObj.locationId
                        ),
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
                            console.log(
                                "Error: could not fetch program, semester, and location: " +
                                    err.message
                            )
                        );
                })
            )
            .catch(
                (err) =>
                    console.log("Error: could not fetch class: " + err.message)
                // TODO: error page?
            );
    };

    formatCurrency = (amount) => {
        return new Intl.NumberFormat("en-US", {
            style: "currency",
            currency: "USD",
        }).format(amount);
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
        const timeString = classObj.times || "";
        const times = timeString
            .split(", ")
            .map((time, index) => <p key={index}>{time}</p>);

        let startEndDate = <p>No sessions scheduled</p>;
        if (sessions.length > 0) {
            startEndDate = (
                <div>
                    <p>{moment(sessions[0].startsAt).format("l")}</p>
                    <p>
                        {moment(sessions[sessions.length - 1].startsAt).format(
                            "l"
                        )}
                    </p>
                </div>
            );
        }

        return (
            <div id="class-times">
                {times}
                {startEndDate}
            </div>
        );
    };

    getPricing = (classObj) => {
        const isLump = !!classObj.priceLump;
        const title = isLump ? "Lump Price: " : "Price per session: ";
        const price = isLump ? classObj.priceLump : classObj.pricePerSession;

        return (
            <div id="class-pricing">
                <p>{title + this.formatCurrency(price)}</p>
                <p>{classObj.paymentNotes}</p>
            </div>
        );
    };

    getSchedules = (sessions) => {
        const schedules = sessions.map((session, index) => {
            const startTime = moment(session.startsAt);
            const endTime = moment(session.endsAt);
            const times =
                startTime.format("dddd") +
                " " +
                startTime.format("h:mma") +
                " - " +
                endTime.format("h:mma");
            return (
                <div key={index} className="row">
                    <span className="index-column">{index + 1}</span>
                    <span className="date-column">{startTime.format("l")}</span>
                    <span className="time-column">{times}</span>
                    <span className="canceled">
                        {session.canceled ? "Canceled" : ""}
                    </span>
                </div>
            );
        });
        return schedules;
    };

    render = () => {
        const classObj = this.state.classObj;
        const program = this.state.program;
        const semester = this.state.semester;
        const location = this.state.location;
        const sessions = this.state.sessions;

        return (
            <div id="view-class">
                {this.getTitle(program, semester, classObj)}

                <div id="class-description">
                    <b>
                        Grades: {program.grade1} - {program.grade2}
                    </b>
                    <p>{program.description}</p>
                </div>

                <div id="class-info">
                    <Link to={"/contact?interest=" + classObj.classId}>
                        <button>Register</button>
                    </Link>

                    <h4>Location</h4>
                    {this.getLocation(location)}
                    <h4>Times</h4>
                    {this.getTimes(classObj, sessions)}
                    <h4>Pricing</h4>
                    {this.getPricing(classObj)}
                </div>

                <div id="class-session">
                    <h4>Schedule</h4>
                    {this.getSchedules(sessions)}
                </div>

                <div id="class-questions">
                    <div>
                        Questions?{" "}
                        <Link to={"/contact?interest=" + classObj.classId}>
                            Contact Us
                        </Link>
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
