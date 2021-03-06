"use strict";
import React from "react";
require("./accountUsers.sass");
import { Link } from "react-router-dom";
import API from "../../api.js";
import UserSelector from "./userSelector.js";
import { UserRowCard } from "../../common/rowCards/userRowCard.js";

export default class UserInfos extends React.Component {
    onDeleteUser = (e) => {
        const user = this.props.selectedUser;
        const userId = user.id;
        API.delete("api/users/full/user/" + userId)
            .then((res) => {
                window.alert("User deleted!");
                this.props.refreshAccount();
            })
            .catch((err) => window.alert("Could not delete account " + err));
    };

    render() {
        const accountId = this.props.accountId;
        const selectedUser = this.props.selectedUser;
        const userId = selectedUser.id;
        const editUserUrl =
            "/account/" + accountId + "/user/" + userId + "/edit";
        const moveUserUrl =
            "/account/" + accountId + "/user/" + userId + "/move";
        return (
            <section className="account-tab account-users">
                <h3>Select a User in this Account</h3>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={userId}
                    onChange={this.props.onSwitchUser}
                />
                <UserRowCard user={selectedUser} editUrl={editUserUrl} />

                <div>
                    <Link to={moveUserUrl}>
                        Move User to a different account
                    </Link>
                </div>

                <section className="delete">
                    <button onClick={this.onDeleteUser}>
                        Delete User from account
                    </button>
                    <p>
                        Warning: Deleting a user will delete all user
                        information including contacts, classes enrollments,
                        etc.
                    </p>
                </section>
            </section>
        );
    }
}
