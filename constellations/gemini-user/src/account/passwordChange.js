"use strict";
require("./passwordChange.sass");
import React from "react";
import API from "../utils/api.js";

export class PasswordChange extends React.Component {
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
            API.post("api/accounts/account/" + this.props.accountId, account)
                .then((res) => {
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
                })
                .catch((err) => alert("Could not save password: " + err));
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
