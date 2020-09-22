"use strict";
require("../account/accountInfo.sass");
import React from "react";
import { Link } from "react-router-dom";
import { getFullName } from "../common/userUtils.js";

export class AccountInfo extends React.Component {
    state = {
        selectedUsers: [],
    };

    formatCurrency = (amount) => {
        return new Intl.NumberFormat("en-US", {
            style: "currency",
            currency: "USD",
        }).format(amount);
    };

    onCheckUser = (e, userId) => {
        const users = this.props.users;
        if (e.target.checked) {
            var currentSelectedUsers = this.state.selectedUsers;
            const checkedUser = users.find((user) => user.id == userId);
            currentSelectedUsers.push(checkedUser);
            this.setState({
                selectedUsers: currentSelectedUsers,
            });
        } else {
            var currentSelectedUsers = this.state.selectedUsers;
            const uncheckedUser = users.find((user) => user.id == userId);
            currentSelectedUsers.splice(
                currentSelectedUsers.indexOf(uncheckedUser),
                1
            );
            this.setState({
                selectedUsers: currentSelectedUsers,
            });
        }
    };

    createUserRows = (users, selectedUsers) => {
        const userRows = users.map((user, index) => {
            var status = user.isGuardian ? "(guardian" : "(student";
            status +=
                user.email == this.props.email ? ", primary contact)" : ")";

            if (selectedUsers.length > 0) {
                for (var i = 0; i < selectedUsers.length; i++) {
                    var checkedUser = users.find(
                        (user) => user == selectedUsers[i]
                    );
                    if (!checkedUser) {
                        this.setState({
                            selectedUsers: [],
                        });
                        break;
                    }
                }
            }

            return (
                <div className="row" key={index}>
                    <input
                        type="checkbox"
                        onChange={(e) => this.onCheckUser(e, user.id)}
                    />
                    <span className="column">{getFullName(user)}</span>
                    <span className="column">{user.email}</span>
                    <span className="column status">{status}</span>
                </div>
            );
        });
        return userRows;
    };

    render = () => {
        const id = this.props.id;
        const accountEmail = this.props.email;
        const users = this.props.users;
        const transactions = this.props.transactions;
        const name = this.props.name;

        var emails = [];
        emails = users.map((user) => user.email);

        const userRows = this.createUserRows(users, this.state.selectedUsers);

        var balance = 0;
        transactions.map((transaction, index) => {
            const amount = transaction.amount;
            balance += parseInt(amount);
        });

        return (
            <section id="account-info">
                <section id="account-information">
                    <span id="account-number">Account Information</span>
                    <h2>AccountId: {id}</h2>
                    <h2>Primary Email: {accountEmail}</h2>
                    <div id="account-users">
                        <h2>Users in Account</h2>
                        {userRows}
                    </div>
                    <div id="transaction-footer">
                        <span>
                            Account Balance:{" "}
                            <b>{this.formatCurrency(balance)}</b>
                        </span>
                    </div>
                </section>

                <section id="generated-email">
                    <div className="email-template">
                        <span id="template-title">
                            Generated Email Template
                        </span>
                        <h3>
                            To:{" "}
                            {this.state.selectedUsers
                                .map((user) => user.email)
                                .toString()}
                        </h3>
                        <h3>
                            Subject: Math Navigator: Account Balance Payment
                            Reminder
                        </h3>
                        <h3>Message: </h3>
                        <div className="generated-email">
                            <p>Hello {name},</p>
                            <p>
                                Your current balance is ${balance}. Lorem ipsum
                                dolor sit amet, consectetur adipiscing elit, sed
                                do eiusmod tempor incididunt ut labore et dolore
                                magna aliqua. Ut enim ad minim veniam, quis
                                nostrud exercitation ullamco laboris nisi ut
                                aliquip ex ea commodo consequat. Duis aute irure
                                dolor in reprehenderit in voluptate velit esse
                                cillum dolore eu fugiat nulla pariatur.
                                Excepteur sint occaecat cupidatat non proident,
                                sunt in culpa qui officia deserunt mollit anim
                                id est laborum.
                            </p>
                            <p>Best wishes from the Math Navigator Family</p>
                        </div>
                    </div>
                </section>
            </section>
        );
    };
}
