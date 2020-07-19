"use strict";
require("./userClass.sass");
import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import API from "../api.js";

export class UserClassPage extends React.Component {
    state = {
        id: 0,
        user: {},
        userClasses: [],
        classes: [],
        enrollStatuses: [],
        classId: 0,
    };

    componentDidMount = () => {
        this.fetchUser();
    };

    fetchUser = () => {
        const id = this.props.id;

        const apiCalls = [
            API.get("api/users/user/" + id),
            API.get("api/userclasses/users/" + id),
        ];

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    let classIds = [];
                    let enrollStatuses = [];
                    responses[1].data.forEach((userClass) => {
                        classIds.push(userClass.classId);
                        enrollStatuses.push(userClass.enrollStatus);
                    });

                    this.setState({
                        id: id,
                        user: responses[0].data,
                        enrollStatuses: enrollStatuses,
                    });

                    if (classIds.length > 0) {
                        this.fetchClasses(classIds);
                    }
                })
            )
            .catch((err) => {
                alert("Could not fetch user: " + err.response.data);
            });
    };

    fetchClasses = (classIds) => {
        let searchArray = new Set(classIds);
        API.get("api/classes/all")
            .then((res) => {
                var userClasses = [];
                var classes = [];
                res.data.forEach((c) => {
                    if (searchArray.has(c.classId)) {
                        userClasses.push(c);
                    } else {
                        classes.push(c);
                    }
                });

                this.setState({
                    userClasses: userClasses,
                    classes: classes,
                });
            })
            .catch((err) => {
                alert("Could not fetch classes: " + err.response.data);
            });
    };

    onClassChange = (e) => {
        this.setState({
            classId: e.target.value,
        });
    };

    onClickEnroll = () => {
        const userClass = {
            userId: parseInt(this.state.id),
            classId: parseInt(this.state.classId),
        };

        API.post("api/userclasses/create", userClass)
            .then(() => {
                this.fetchUser();
            })
            .catch((err) => {
                alert("Could not enroll in class: " + err.response.data);
            });
    };

    render = () => {
        const user = this.state.user;

        var fullName = user.firstName;
        if (user.middleName) {
            fullName += " " + user.middleName + " " + user.lastName;
        } else {
            fullName += " " + user.lastName;
        }

        const rows = this.state.userClasses.map((c, index) => {
            const enrollStatus = this.state.enrollStatuses[index];
            const fullState = c.fullState;
            let state = "";
            let firstSpace = <button></button>;
            let secondSpace = <button></button>;
            switch (fullState) {
                case 0:
                    state = "Empty";
                    firstSpace = <button>Reject</button>;
                    secondSpace = <button>Accept</button>;
                    break;
                case 1:
                    status = "Full";
                    firstSpace = <button>Reject</button>;
                    break;
                case 2:
                    state = "Almost Full";
                    firstSpace = <button>Reject</button>;
                    secondSpace = <button>Accept</button>;
                    break;
            }
            if (enrollStatus == "Enrolled") {
                firstSpace = <button></button>;
                secondSpace = <button>Unenroll</button>;
            }

            return (
                <div className="row" key={index}>
                    <span className="large-column">{c.classId}</span>
                    <span className="column">{enrollStatus}</span>
                    <span className="column">{state}</span>
                    <span className="button">{firstSpace}</span>
                    <span className="button">{secondSpace}</span>
                </div>
            );
        });

        const classOptions = this.state.classes.map((c, index) => {
            return <option key={index}>{c.classId}</option>;
        });

        return (
            <div id="view-user-class">
                <h2>
                    <Link className="users-back" to="/users">
                        {"< Back to Users"}
                    </Link>
                </h2>

                <div>
                    <h2>User Information</h2>
                    <p>{fullName}</p>
                    <p>{user.email}</p>
                    <p>{user.phone}</p>
                </div>

                <h2>User Classes</h2>
                <div id="user-class">
                    <div className="header row">
                        <span className="large-column">Class ID</span>
                        <span className="column">Enroll Status</span>
                        <span className="column">Full State</span>
                        <span className="button"></span>
                        <span className="button"></span>
                    </div>
                    {rows}
                </div>

                <h2>Enroll User for Class</h2>
                <p>Select a Class ID to enroll user into:</p>
                <select
                    value={this.state.classId}
                    onChange={(e) => this.onClassChange(e)}>
                    <option default hidden>
                        Select a classId
                    </option>
                    {classOptions}
                </select>

                <button onClick={this.onClickEnroll}>Enroll</button>
            </div>
        );
    };
}
