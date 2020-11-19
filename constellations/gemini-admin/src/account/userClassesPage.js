"use strict";
// require("./userAFH.sass");
import React from "react";
import moment from "moment";
import { keyBy } from "lodash";
import { Link } from "react-router-dom";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import { InputSelect } from "../common/inputs/inputSelect.js";

export class UserClassesPage extends React.Component {
    state = {
        allClasses: [],
        classMap: {},
        selectedClassId: "",
        usersForClass: [],
    };

    componentDidMount = () => {
        API.get("/api/classes/all")
            .then((res) => {
                const classes = res.data;
                this.setState({
                    allClasses: classes,
                    classMap: keyBy(classes, "classId"),
                });
            })
            .catch((err) => console.log("Could not fetch classes"));
    };

    onClassChange = (e) => {
        const nextClassId = e.target.value;
        this.setState({
            selectedClassId: nextClassId,
        });

        API.get("api/user-classes/class/" + nextClassId)
            .then((res) => {
                this.setState({ usersForClass: res.data });
            })
            .catch((err) => console.log("Could not fetch users"));
    };

    render() {
        const options = this.state.allClasses.map((classObj) => {
            return {
                value: classObj.classId,
                displayName: classObj.classId,
            };
        });
        const users = this.state.usersForClass.map((userClass, index) => (
            <UserRow key={index} userClass={userClass} />
        ));

        return (
            <div id="view-user-classes">
                <InputSelect
                    label="Select a Class"
                    value={this.state.selectedClassId}
                    onChangeCallback={this.onClassChange}
                    options={options}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no Classes to choose from. Please add one{" "}
                            <Link to="/classes/add">here</Link>
                        </span>
                    }
                />

                {users.length > 0 && (
                    <div id="users">
                        <h3>Users in Class</h3>
                        {users}
                    </div>
                )}
                {users.length == 0 && this.state.selectedClassId && (
                    <p>No Users currently registered for this class.</p>
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
        const userClass = this.props.userClass || {};
        const userId = userClass.userId;
        API.get("api/users/user/" + userId)
            .then((res) => {
                this.setState({ user: res.data });
            })
            .catch((err) => console.log("Could not find user " + userId));
    };

    render() {
        const userClass = this.props.userClass || {};
        const user = this.state.user;
        const viewUserUrl = "/account/" + user.accountId + "?view=user-classes";
        const viewAccountUrl = "/account/" + user.accountId;
        return (
            <div className="user-row">
                <div>{getFullName(user)}</div>
                <div>{user.email}</div>
                <div>{moment(userClass.updatedAt).format("l")}</div>
                <div>{userClass.state}</div>
                <Link to={viewUserUrl}>View User Details</Link>
                <Link to={viewAccountUrl}>View Account</Link>
            </div>
        );
    }
}
