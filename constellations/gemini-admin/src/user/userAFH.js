"use strict";
require("./userAFH.sass");
import React from "react";
import moment from "moment";
import axios from "axios";
import { Link } from "react-router-dom";
import API from "../api.js";
import Axios from "axios";

export class UserAFHPage extends React.Component {
    state = {
        id: 0,
        user: {},
        userAFHs: [],
        afhs: [],
        afhId: 0,
    };

    componentDidMount = () => {
        this.fetchUser();
    };

    fetchUser = () => {
        const id = this.props.id;

        const apiCalls = [
            API.get("api/users/user/" + id),
            API.get("api/userafhs/users/" + id),
        ];

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    let afhIds = [];
                    responses[1].data.forEach((userAFH) => {
                        afhIds.push(userAFH.afhId);
                    });

                    this.setState({
                        id: id,
                        user: responses[0].data,
                    });

                    if (afhIds.length > 0) {
                        this.fetchAFHs(afhIds);
                    }
                })
            )
            .catch((err) => {
                alert("Could not fetch user: " + err.response.data);
            });
    };

    fetchAFHs = (afhIds) => {
        let searchArray = new Set(afhIds);
        API.get("api/askforhelp/all")
            .then((res) => {
                var userAFHs = [];
                var afhs = [];
                res.data.forEach((afh) => {
                    if (searchArray.has(afh.id)) {
                        userAFHs.push(afh);
                    } else {
                        afhs.push(afh);
                    }
                });

                userAFHs = _.sortBy(userAFHs, ["date"]);
                afhs = _.sortBy(afhs, ["date"]);
                this.setState({
                    userAFHs: userAFHs,
                    afhs: afhs,
                });
            })
            .catch((err) => {
                alert("Could not fetch afhs: " + err.response.data);
            });
    };

    onAFHChange = (e) => {
        this.setState({
            afhId: e.target.value,
        });
    };

    onClickSchedule = () => {
        const userAFH = {
            userId: parseInt(this.state.id),
            afhId: parseInt(this.state.afhId),
        };

        API.post("api/userafhs/create", userAFH)
            .then(() => {
                this.fetchUser();
            })
            .catch((err) => {
                alert("Could not schedule AFH: " + err.response.data);
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

        const rows = this.state.userAFHs.map((afh, index) => {
            const status = moment().isBefore(afh.date)
                ? "Will Attend"
                : "Attended";
            return (
                <div className="row" key={index}>
                    <span className="column">
                        {moment(afh.date).format("l")}
                    </span>
                    <span className="large-column">{afh.title}</span>
                    <span className="column">{afh.subject}</span>
                    <span className="column status">{status}</span>
                </div>
            );
        });

        const afhOptions = this.state.afhs.map((afh, index) => {
            return (
                <option key={index} value={afh.id}>
                    {moment(afh.date).format("l") +
                        " " +
                        afh.subject +
                        " " +
                        afh.timeString}
                </option>
            );
        });

        return (
            <div id="view-user-afh">
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

                <h2>User AskForHelp Sessions</h2>
                <div id="user-afh">
                    <div className="header row">
                        <span className="column">AskForHelp Date</span>
                        <span className="large-column">Title</span>
                        <span className="column">Subject</span>
                        <span className="column status">Status</span>
                    </div>
                    {rows}
                </div>

                <h2>Schedule AskForHelp for User</h2>
                <p>Select a AFH session for user:</p>
                <select
                    value={this.state.afhId}
                    onChange={(e) => this.onAFHChange(e)}>
                    <option default hidden>
                        Select an AFH session
                    </option>
                    {afhOptions}
                </select>

                <button onClick={this.onClickSchedule}>Schedule</button>
            </div>
        );
    };
}
