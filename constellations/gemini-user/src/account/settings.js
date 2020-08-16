"use strict";
require("./account.sass");
import React from "react";
import API from "../utils/api.js";

import {
    chargeDisplayNames,
    subjectDisplayNames,
    seasonOrder,
    renderMultiline,
    formatCurrency,
    fetchError,
} from "./accountUtils.js";

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
                .catch((err) => fetchError(err));
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

class PasswordChange extends React.Component {
    state = {
        tabOpen: false,
        message: "",
        validated: false,
        successMessage: "",

        oldPassword: "",
        newPassword: "",
        confirmPassword: "",
    };

    onClickChange = () => {
        this.setState({
            tabOpen: !this.state.tabOpen,
            message: "",
            successMessage: "",
        });
    };

    onClickSave = () => {
        if (this.state.validated) {
            let account = {
                primaryEmail: this.props.primaryEmail,
                password: this.state.newPassword,
            };
            API.post(
                "api/accounts/account/" + this.props.accountId,
                account
            ).then((res) => {
                this.props.passwordChangeCallback(this.state.newPassword);
                this.setState({
                    tabOpen: false,
                    message: "",
                    validated: false,
                    successMessage: " New password saved!",
                    oldPassword: "",
                    newPassword: "",
                    confirmPassword: "",
                });
            });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value }, () =>
            this.validateInput()
        );
    };

    validateInput = () => {
        if (this.state.oldPassword != this.props.oldPassword) {
            this.setState({
                message: "Old password is incorrect",
                validated: false,
            });
        } else if (this.state.newPassword.length < 8) {
            this.setState({
                message: "New password must be at least 8 characters long",
                validated: false,
            });
        } else if (this.state.newPassword != this.state.confirmPassword) {
            this.setState({
                message: "New password does not match confirmation",
                validated: false,
            });
        } else {
            this.setState({ message: "", validated: true });
        }
    };

    render = () => {
        const message = <p>{this.state.message}</p>;

        const changePasswordDialog = this.state.tabOpen ? (
            <div>
                <ul className="vertical-centered no-border">
                    <li className="li-med">Old password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "oldPassword")
                            }
                            value={this.state.oldPassword}
                        />
                    </li>
                </ul>
                <ul className="vertical-centered no-border">
                    <li className="li-med">New password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "newPassword")
                            }
                            value={this.state.newPassword}
                        />
                    </li>
                </ul>
                <ul className="vertical-centered no-border">
                    <li className="li-med">Confirm new password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "confirmPassword")
                            }
                            value={this.state.confirmPassword}
                        />
                    </li>
                </ul>
                <span className="red">{message}</span>
                <div className="password-buttons space-between">
                    <button className="btn-cancel" onClick={this.onClickChange}>
                        Cancel
                    </button>
                    <button
                        className={
                            this.state.validated ? "btn-save" : "btn-cancel"
                        }
                        onClick={this.onClickSave}>
                        Save
                    </button>
                </div>
            </div>
        ) : (
            ""
        );

        return (
            <div id="password-change">
                <p className="vertical-mobile">
                    <a className="orange" onClick={this.onClickChange}>
                        Change password...
                    </a>
                    {this.state.successMessage}
                </p>
                {changePasswordDialog}
            </div>
        );
    };
}
