"use strict";
import React from "react";
import { Link } from "react-router-dom";
import UserSelector from "./userSelector.js";
import { getFullName } from "../../common/userUtils.js";

export default class UserInfos extends React.Component {
    render() {
        const accountId = this.props.accountId;
        const selectedUser = this.props.selectedUser;
        const userId = selectedUser.id;
        const editUserUrl =
            "/account/" + accountId + "/user/" + userId + "/edit";
        const moveUserUrl =
            "/account/" + accountId + "/user/" + userId + "/move";
        return (
            <section>
                <h2>Users in Account</h2>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={userId}
                    onChange={this.props.onSwitchUser}
                />

                <div>
                    <div>Id: {userId}</div>
                    <div>Name: {getFullName(selectedUser)}</div>
                    <div>Email: {selectedUser.email}</div>
                    <div>
                        IsGuardian: {selectedUser.isGuardian ? "Yes" : "No"}
                    </div>
                    {selectedUser.phone && (
                        <div>Phone: {selectedUser.phone}</div>
                    )}
                    {selectedUser.school && (
                        <div>School: {selectedUser.school}</div>
                    )}
                    {selectedUser.graduationYear && (
                        <div>GraduationYear: {selectedUser.graduationYear}</div>
                    )}
                    <div>Notes: {selectedUser.notes}</div>
                    <Link to={editUserUrl}>Edit User</Link>
                </div>

                <div>
                    <Link to={moveUserUrl}>
                        Move User to a different account
                    </Link>
                </div>
            </section>
        );
    }
}
