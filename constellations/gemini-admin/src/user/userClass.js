"use strict";
require("./userClass.sass");
import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import { UserClassRow } from "./userClassRow.js";
import { InputSelect } from "../common/inputs/inputSelect.js";

export class UserClassPage extends React.Component {
    state = {
        id: 0,
        user: {},
        userClasses: [],
        otherClassIds: [],
        classId: "",
    };

    componentDidMount = () => {
        this.fetchUser();
    };

    fetchUser = () => {
        const id = this.props.id;

        const apiCalls = [
            API.get("api/users/user/" + id),
            API.get("api/classes/all"),
            API.get("api/user-classes/user/" + id),
        ];

        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    let classMap = {};
                    let userClasses = [];
                    let classIds = new Set();

                    responses[1].data.forEach((c) => {
                        classMap[c.classId] = c;
                        classIds.add(c.classId);
                    });

                    // Separate classes into selected and unselected
                    responses[2].data.forEach((userClass) => {
                        const newUserClass = {
                            id: userClass.id,
                            // TODO: backend needs to add updatedAt back into JSON
                            // updatedAt: userClass.updatedAt
                            userId: userClass.userId,
                            classObject: classMap[userClass.classId],
                            accountId: userClass.accountId,
                            state: userClass.state,
                        };
                        userClasses.push(newUserClass);
                        classIds.delete(userClass.classId);
                    });

                    const classIdArray = Array.from(classIds);
                    this.setState({
                        id: id,
                        user: responses[0].data,
                        userClasses: userClasses,
                        otherClassIds: classIdArray,
                        classId: classIdArray[0],
                    });
                })
            )
            .catch((err) => alert("Could not fetch user: " + err));
    };

    onClassChange = (e) => {
        this.setState({
            classId: e.target.value,
        });
    };

    onClickEnroll = () => {
        const userClass = {
            userId: parseInt(this.state.id),
            classId: this.state.classId,
            accountId: parseInt(this.state.user.accountId),
            state: 0,
        };

        API.post("api/user-classes/create", userClass)
            .then(() => {
                this.fetchUser();
            })
            .catch((err) =>
                alert("Could not enroll in class: " + err.response.data)
            );
    };

    renderEnrollSection = () => {
        const classOptions = this.state.otherClassIds.map((classId, index) => {
            return { value: classId, displayName: classId };
        });
        const enrollButton = classOptions.length ? (
            <button onClick={this.onClickEnroll}>Enroll</button>
        ) : (
            ""
        );

        return (
            <div>
                <InputSelect
                    label="Enroll User for Class"
                    description="Select a Class ID to enroll user into:"
                    value={this.state.classId}
                    onChangeCallback={(e) => this.onClassChange(e)}
                    required={true}
                    options={classOptions}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no classes to choose from. Please add one{" "}
                            <Link to="/classes/add">here</Link>
                        </span>
                    }
                />
                {enrollButton}
            </div>
        );
    };

    render = () => {
        const user = this.state.user;
        const userClasses = this.state.userClasses.map((userClass, index) => {
            return (
                <UserClassRow
                    userClass={userClass}
                    key={index}
                    updateCallback={this.fetchUser}
                />
            );
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
                    <p>{getFullName(user)}</p>
                    <p>{user.email}</p>
                    <p>{user.phone}</p>
                </div>

                <div id="user-class">
                    <h2>User Classes</h2>
                    <div className="header row">
                        <span className="large-column">Class ID</span>
                        <span className="column">Enroll Status</span>
                        <span className="column">Full State</span>
                        <span className="space"></span>
                        <span className="space"></span>
                    </div>
                    {userClasses}
                </div>

                <div id="user-enroll">{this.renderEnrollSection()}</div>
            </div>
        );
    };
}
