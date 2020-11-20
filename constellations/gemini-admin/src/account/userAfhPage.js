"use strict";
// require("./userAFH.sass");
import React from "react";
import moment from "moment";
import { keyBy } from "lodash";
import { Link } from "react-router-dom";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import { getAfhTitle } from "../common/displayUtils.js";

export class UserAFHPage extends React.Component {
    state = {
        allAfhs: [],
        afhMap: {},
        selectedAfhId: "",
        usersForAfh: [],
    };

    componentDidMount = () => {
        API.get("/api/askforhelp/all")
            .then((res) => {
                const afhs = res.data;
                this.setState({
                    allAfhs: afhs,
                    afhMap: keyBy(afhs, "id"),
                });
            })
            .catch((err) => console.log("Could not fetch afh sessions"));
    };

    onAfhChange = (e) => {
        const nextAfhId = e.target.value;
        this.setState({
            selectedAfhId: nextAfhId,
        });

        API.get("api/user-afhs/afh/" + nextAfhId)
            .then((res) => {
                this.setState({ usersForAfh: res.data });
            })
            .catch((err) => console.log("Could not fetch users"));
    };

    render() {
        const options = this.state.allAfhs.map((afh) => {
            const time =
                moment(afh.startsAt).format("MM/DD/yy hh:mm") +
                "-" +
                moment(afh.endsAt).format("hh:mm a");
            return {
                value: afh.id,
                displayName: getAfhTitle(afh),
            };
        });
        const users = this.state.usersForAfh.map((userAfh, index) => (
            <UserRow key={index} userAfh={userAfh} />
        ));

        return (
            <div id="view-user-afhs">
                <InputSelect
                    label="Select an AskForHelp session"
                    value={this.state.selectedAfhId}
                    onChangeCallback={this.onAfhChange}
                    options={options}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no AskForHelp sessions to choose from.
                            Please add one <Link to="/afh/add">here</Link>
                        </span>
                    }
                />

                {users.length > 0 && (
                    <div id="users">
                        <h3>Users in AFH Session</h3>
                        {users}
                    </div>
                )}
                {users.length == 0 && this.state.selectedAfhId && (
                    <p>No Users currently registered for this AFH session.</p>
                )}
            </div>
        );
    }
}

class UserRow extends React.Component {
    state = {
        user: {},
    };

    componentDidMount = () => {
        const userAfh = this.props.userAfh || {};
        const userId = userAfh.userId;
        API.get("api/users/user/" + userId)
            .then((res) => {
                this.setState({ user: res.data });
            })
            .catch((err) => console.log("Could not find user " + userId));
    };

    render() {
        const userAfh = this.props.userAfh || {};
        const user = this.state.user;
        const viewUserUrl = "/account/" + user.accountId + "?view=user-afhs";
        const viewAccountUrl = "/account/" + user.accountId;
        return (
            <div className="user-row">
                <div>{getFullName(user)}</div>
                <div>{user.email}</div>
                <div>{moment(userAfh.updatedAt).format("l")}</div>
                <div>{userAfh.state}</div>
                <Link to={viewUserUrl}>View User Details</Link>
                <Link to={viewAccountUrl}>View Account</Link>
            </div>
        );
    }
}
