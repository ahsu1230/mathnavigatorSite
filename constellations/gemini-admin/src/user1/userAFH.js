"use strict";
require("./userAFH.sass");
import React from "react";
import moment from "moment";
import axios from "axios";
import { Link } from "react-router-dom";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import { InputSelect } from "../common/inputs/inputSelect.js";

export class UserAFHPage extends React.Component {
    state = {
        id: 0,
        user: {},
        userAFHs: [],
        otherAFHs: [],
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
            API.get("api/askforhelp/all"),
        ];

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    let userAFHs = [];
                    let otherAFHs = [];

                    let afhIds = new Set();
                    responses[1].data.forEach((userAFH) => {
                        afhIds.add(userAFH.afhId);
                    });

                    // Separate afhs into selected and unselected
                    responses[2].data.forEach((afh) => {
                        if (afhIds.has(afh.id)) {
                            userAFHs.push(afh);
                        } else {
                            otherAFHs.push(afh);
                        }
                    });
                    userAFHs = _.sortBy(userAFHs, ["date"]);
                    otherAFHs = _.sortBy(otherAFHs, ["date"]);

                    this.setState({
                        id: id,
                        user: responses[0].data,
                        userAFHs: userAFHs,
                        otherAFHs: otherAFHs,
                        afhId: otherAFHs[0] ? otherAFHs[0].id : 0,
                    });
                })
            )
            .catch((err) => alert("Could not fetch user: " + err));
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
            .catch((err) => alert("Could not schedule AFH: " + err));
    };

    renderScheduleSection = () => {
        const afhOptions = this.state.otherAFHs.map((afh, index) => {
            return {
                value: afh.id,
                displayName:
                    moment(afh.date).format("l") +
                    " " +
                    afh.subject +
                    " " +
                    afh.timeString,
            };
        });
        const enrollButton = afhOptions.length ? (
            <button onClick={this.onClickSchedule}>Schedule</button>
        ) : (
            ""
        );

        return (
            <div>
                <InputSelect
                    label="Schedule AskForHelp for User"
                    description="Select a AFH session for user:"
                    value={this.state.afhId}
                    onChangeCallback={(e) => this.onAFHChange(e)}
                    required={true}
                    options={afhOptions}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no AFH sessions to choose from. Please add
                            one <Link to="/afh/add">here</Link>
                        </span>
                    }
                />
                {enrollButton}
            </div>
        );
    };

    render = () => {
        const user = this.state.user;
        const userAFHs = this.state.userAFHs.map((afh, index) => {
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

        return (
            <div id="view-user-afh">
                <h2>
                    <Link className="users-back" to="/users">
                        {"< Back to Users"}
                    </Link>
                </h2>

                <div>
                    <h2>User Information</h2>
                    <p>{getFullName(user)}</p>
                    <p>{user.email}</p>
                    <p>{user.phone}</p>
                </div>

                <div id="user-afh">
                    <h2>User AskForHelp Sessions</h2>
                    <div className="header row">
                        <span className="column">AskForHelp Date</span>
                        <span className="large-column">Title</span>
                        <span className="column">Subject</span>
                        <span className="column status">Status</span>
                    </div>
                    {userAFHs}
                </div>

                <div id="user-schedule">
                    <h2>Schedule AskForHelp for User</h2>
                    {this.renderScheduleSection()}
                </div>
            </div>
        );
    };
}
