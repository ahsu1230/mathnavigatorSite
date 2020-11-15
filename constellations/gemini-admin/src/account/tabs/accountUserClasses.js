"use strict";
import React from "react";
import API from "../../api.js";
import { getFullName } from "../../common/userUtils.js";
import UserSelector from "./userSelector.js";

export default class UserClasses extends React.Component {
    state = {
        userClasses: [],
    };

    componentDidMount() {
        // get all user classes for account or for all users
        // make a map userId -> [userClasses]
    }

    render() {
        const userClasses = [];
        const selectedUser = this.props.selectedUser;

        return (
            <section>
                <h2>User Class Registrations</h2>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={selectedUser.id}
                    onChange={this.props.onSwitchUser}
                />

                <div>
                    <div>Id: {selectedUser.id}</div>
                    <div>Name: {getFullName(selectedUser)}</div>
                    <div>Email: {selectedUser.email}</div>
                </div>

                {userClasses.length > 0 && userClasses}
                {userClasses.length == 0 && (
                    <p>This user has no class registrations.</p>
                )}
                <button>Register a class for user</button>
            </section>
        );
    }
}
