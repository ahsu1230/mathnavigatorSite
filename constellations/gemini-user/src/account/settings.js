"use strict";
require("./settings.sass");
import React from "react";
import API from "../utils/api.js";

import { renderMultiline } from "../utils/utils.js";
import { PasswordChange } from "./passwordChange.js";

export class SettingsTab extends React.Component {
    state = {
        id: 0,
        primaryEmail: "",
        password: "",
    };

    componentDidMount = () => {
        const id = this.props.accountId;

        if (id) {
            this.setState({ id: id });
            API.get("api/accounts/account/" + id)
                .then((res) => {
                    this.setState({
                        primaryEmail: res.data.primaryEmail,
                        password: res.data.password,
                    });
                })
                .catch((err) => alert("Could not fetch data: " + err));
        }
    };

    onPasswordChange = (password) => {
        this.setState({ password: password });
    };

    render = () => {
        const currentYear = new Date().getFullYear();

        const usersList = this.props.users.map((user, index) => {
            let contactInfo = [user.email];
            if (user.phone) {
                contactInfo.push(user.phone);
            }
            contactInfo = renderMultiline(contactInfo);

            let otherInfo = [
                (user.isGuardian ? "Guardian" : "Student") +
                    (user.email == this.state.primaryEmail
                        ? " (Primary Contact)"
                        : ""),
            ];
            if (user.school) {
                otherInfo.push(user.school);
            }
            if (user.graduationYear) {
                otherInfo.push(
                    12 -
                        (user.graduationYear - currentYear) +
                        "th Grade, " +
                        "Graduation Year: " +
                        user.graduationYear
                );
            }
            otherInfo = renderMultiline(otherInfo);

            return (
                <ul key={index} className="users-table">
                    <li className="li-med">
                        {user.firstName + " " + user.lastName}
                    </li>
                    <li className="li-med">{contactInfo}</li>
                    <li className="li-large">{otherInfo}</li>
                </ul>
            );
        });

        return (
            <div className="tab-content" id="settings-tab">
                <div>
                    <h2>Your Account Information</h2>
                    <p className="vertical-mobile">
                        <span>Primary email: {this.state.primaryEmail}</span>
                        <a className="edit orange"> Change primary contact</a>
                    </p>
                    <PasswordChange
                        accountId={this.state.id}
                        primaryEmail={this.state.primaryEmail}
                        oldPassword={this.state.password}
                        passwordChangeCallback={this.onPasswordChange}
                    />
                </div>

                <div>
                    <h2>User Information</h2>
                    <ul className="header hide-mobile">
                        <li className="li-med">Name</li>
                        <li className="li-med">Contact</li>
                        <li className="li-large">Other Information</li>
                    </ul>
                    {usersList}
                    <p>
                        <a className="orange">Edit users for this account</a>
                    </p>
                </div>

                <div>
                    <p>
                        You may opt to delete your Math Navigator account.
                        <br />
                        However, doing so will delete all your data with Math
                        Navigator, including all user and class information.
                    </p>
                    <a className="red">Request to Delete Account...</a>
                </div>
            </div>
        );
    };
}
