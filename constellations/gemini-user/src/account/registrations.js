"use strict";
require("./registrations.sass");
import React from "react";
import API from "../utils/api.js";
import { union } from "lodash";
import moment from "moment";

import {
    subjectDisplayNames,
    seasonOrder,
    renderMultiline,
    fetchError,
} from "./accountUtils.js";

export class RegistrationsTab extends React.Component {
    state = {
        userClasses: [],
        upcomingAFHs: [],
        locationsById: {},
        viewAllEnrolledClasses: false,
    };

    componentDidMount = () => {
        Promise.all([
            API.get("api/classes/all"),
            API.get("api/programs/all"),
            API.get("api/semesters/all"),

            API.get("api/askforhelp/all"),

            API.get("api/locations/all"),
        ])
            .then((res) => {
                const allClasses = res[0].data;
                const allPrograms = res[1].data;
                const allSemesters = res[2].data;

                const allAFHs = res[3].data;

                const users = this.props.users;

                const allLocations = res[4].data;
                const locationsById = {};
                allLocations.map((loc, index) => {
                    locationsById[loc.locationId] = loc;
                });
                this.setState({ locationsById: locationsById });

                let userClasses = [];
                let upcomingAFHs = [];
                users.map((user, index) => {
                    API.get("api/user-classes/user/" + user.id).then((res) => {
                        let classes = res.data.map((c, index) => {
                            let matchedClass = allClasses.find(
                                (element) => element.classId == c.classId
                            );
                            let matchedProgram = allPrograms.find(
                                (element) =>
                                    element.programId == matchedClass.programId
                            );
                            let matchedSemester = allSemesters.find(
                                (element) =>
                                    element.semesterId ==
                                    matchedClass.semesterId
                            );
                            return {
                                program: matchedProgram,
                                semester: matchedSemester,
                                enrollDate: c.createdAt,
                            };
                        });

                        userClasses.push({
                            name: user.firstName + " " + user.lastName,
                            classes: classes,
                        });
                        this.setState({ userClasses: userClasses });
                    });
                    API.get("api/userafhs/users/" + user.id).then((res) => {
                        let afhs = res.data.map((afh, index) => {
                            let matchedAFH = allAFHs.find(
                                (element) => element.id == afh.afhId
                            );
                            upcomingAFHs.push(matchedAFH);
                        });
                        upcomingAFHs = union(upcomingAFHs);
                        this.setState({ upcomingAFHs: upcomingAFHs });
                    });
                });
            })
            .catch((err) => alert("Could not fetch data: " + err));
    };

    toggleViewAllClasses = () => {
        this.setState({
            viewAllEnrolledClasses: !this.state.viewAllEnrolledClasses,
        });
    };

    render = () => {
        return this.state.viewAllEnrolledClasses ? (
            <RegistrationsTabAllClasses
                userClasses={this.state.userClasses}
                toggleTabCallback={this.toggleViewAllClasses}
            />
        ) : (
            <RegistrationsTabMain
                userClasses={this.state.userClasses}
                upcomingAFHs={this.state.upcomingAFHs}
                locationsById={this.state.locationsById}
                toggleTabCallback={this.toggleViewAllClasses}
            />
        );
    };
}

class RegistrationsTabMain extends React.Component {
    renderClassList = (classes) => {
        if (!classes.length) {
            return <span>(No classes registered)</span>;
        }
        return classes.map((c, index) => {
            return (
                <div key={index} className="classList-item space-between">
                    <span>
                        {c.program.name + " (" + c.semester.title + ")"}
                    </span>
                    <span>Enrolled on: {moment(c.enrollDate).format("l")}</span>
                </div>
            );
        });
    };

    render = () => {
        const classRegistrationList = this.props.userClasses.map(
            (user, index) => {
                return (
                    <ul key={index}>
                        <li className="li-med">{user.name}</li>
                        <li className="li-large">
                            {this.renderClassList(user.classes)}
                        </li>
                    </ul>
                );
            }
        );

        const afhList = this.props.upcomingAFHs.map((afh, index) => {
            let titleInfo = renderMultiline([
                afh.title,
                subjectDisplayNames[afh.subject],
            ]);
            let dateInfo = renderMultiline([
                moment(afh.date).format("MMMM Do, YYYY"),
                afh.timeString,
            ]);

            let loc = this.props.locationsById[afh.locationId];
            let locInfo = renderMultiline([
                loc.street,
                loc.city + ", " + loc.state,
                loc.room,
            ]);

            return (
                <ul key={index} className="afh-table three-columns">
                    <li className="li-med">{titleInfo}</li>
                    <li className="li-med">{dateInfo}</li>
                    <li className="li-large">{locInfo}</li>
                </ul>
            );
        });

        const afhSection =
            this.props.upcomingAFHs.length > 0 ? (
                <div>
                    <h2>Upcoming Ask For Help Sessions</h2>
                    <ul className="three-columns header hide-mobile">
                        <li className="li-med">Title</li>
                        <li className="li-med">Date</li>
                        <li className="li-large">Location</li>
                    </ul>
                    {afhList}
                </div>
            ) : (
                <div>
                    <h2>Upcoming Ask For Help Sessions</h2>
                    <p>
                        There are no scheduled Ask For Help sessions for this
                        account.
                    </p>
                </div>
            );

        return (
            <div className="tab-content" id="reg-tab-main">
                <div>
                    <h2>Currently Enrolled Classes</h2>
                    {classRegistrationList}
                    <p>
                        <a
                            className="orange"
                            onClick={this.props.toggleTabCallback}>
                            View all enrolled classes
                        </a>
                    </p>
                </div>
                {afhSection}
            </div>
        );
    };
}

class RegistrationsTabAllClasses extends React.Component {
    render = () => {
        const allClasses = [];
        this.props.userClasses.map((user, index) => {
            user.classes.map((c, index) => {
                allClasses.push({
                    user: user.name,
                    classInfo: c,
                });
            });
        });

        allClasses.sort((a, b) => {
            let semA = a.classInfo.semester.semesterId.split("_");
            let semB = b.classInfo.semester.semesterId.split("_");

            if (semA[0] < semB[0]) {
                return 1;
            } else if (semA[0] > semB[0]) {
                return -1;
            } else {
                return seasonOrder.indexOf(semA[1]) <
                    seasonOrder.indexOf(semB[1])
                    ? 1
                    : -1;
            }
        });

        const classRegistrationList = allClasses.map((c, index) => {
            return (
                <ul key={index}>
                    <li className="li-med">{c.user}</li>
                    <li className="li-large">
                        {c.classInfo.program.name +
                            " (" +
                            c.classInfo.semester.title +
                            ")"}
                    </li>
                    <li>
                        Enrolled on:{" "}
                        {moment(c.classInfo.enrollDate).format("l")}
                    </li>
                </ul>
            );
        });

        return (
            <div className="tab-content" id="reg-tab-all">
                <div>
                    <div className="vertical-centered space-between">
                        <h2>All Enrolled Classes</h2>
                        <a
                            className="orange"
                            onClick={this.props.toggleTabCallback}>
                            View current enrollments
                        </a>
                    </div>
                    <div>{classRegistrationList}</div>
                </div>
            </div>
        );
    };
}
