"use strict";
require("../account/accountInfo.sass");
import React from "react";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";

export class AccountInfo extends React.Component {
    state = {
        id: 0,
        email: "",
        users: [],
        transactions: [],
        selectedUserEmails: [],
        active: false,
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
            users.map((user, index) => {
                if (user.id == userId) {
                    var email = this.state.selectedUserEmails;
                    email.push(user.email);
                    this.setState({
                        selectedUserEmails: email,
                    });
                }
            });
        } else {
            users.map((user, index) => {
                if (user.id == userId) {
                    var email = this.state.selectedUserEmails;
                    email.splice(email.indexOf(user.email), 1);
                    this.setState({
                        selectedUserEmails: email,
                    });
                }
            });
        }
    };

    render = () => {
        const id = this.props.id;
        const email = this.props.email;
        const users = this.props.users;
        const transactions = this.props.transactions;
        const name = this.props.name;

        var emails = [];
        users.map((user, index) => {
            emails.push(user.email);
        });

        const userRows = users.map((user, index) => {
            var status = user.isGuardian ? "(guardian" : "(student";
            status += user.email == email ? ", primary contact)" : ")";

            if (this.state.selectedUserEmails.length > 0) {
                this.state.selectedUserEmails.map((existingEmail, index) => {
                    if (emails.indexOf(existingEmail) == -1) {
                        this.setState({
                            selectedUserEmails: [],
                        });
                    }
                });
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

        var balance = 0;
        transactions.map((transaction, index) => {
            const amount = transaction.amount;
            balance += parseInt(amount);
        });

        return (
            <section id="account-info">
                <span id="account-number">Account Information</span>
                <h2>AccountId: {id}</h2>
                <h2>Primary Email: {email}</h2>
                <div id="account-users">
                    <h2>Users in Account</h2>
                    {userRows}
                </div>
                <div id="transaction-footer">
                    <span>
                        Account Balance: <b>{this.formatCurrency(balance)}</b>
                    </span>
                </div>
                <div className="email-template">
                    <span id="template-title">Generated Email Template</span>
                    <h3>To: {this.state.selectedUserEmails.toString()}</h3>
                    <h3>
                        Subject: Math Navigator: Account Balance Payment
                        Reminder
                    </h3>
                    <h3>Message: </h3>
                    <div className="generated-email">
                        <p>Hello {name},</p>
                        <p>
                            Lorem ipsum dolor sit amet, consectetur adipiscing
                            elit, sed do eiusmod tempor incididunt ut labore et
                            dolore magna aliqua. Ut enim ad minim veniam, quis
                            nostrud exercitation ullamco laboris nisi ut aliquip
                            ex ea commodo consequat. Duis aute irure dolor in
                            reprehenderit in voluptate velit esse cillum dolore
                            eu fugiat nulla pariatur. Excepteur sint occaecat
                            cupidatat non proident, sunt in culpa qui officia
                            deserunt mollit anim id est laborum.
                        </p>
                        <p>Best wishes from the Math Navigator Family</p>
                    </div>
                </div>
            </section>
        );
    };
}
