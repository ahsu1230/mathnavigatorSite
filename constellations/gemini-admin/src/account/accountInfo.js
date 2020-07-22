"use strict";
require("./accountInfo.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class AccountInfo extends React.Component {
    formatCurrency = (amount) => {
        if (amount < 0) {
            return "-$" + (-amount).toString();
        } else {
            return "$" + amount;
        }
    };

    render = () => {
        const id = this.props.id;
        const email = this.props.email;
        const users = this.props.users;
        const transactions = this.props.transactions;
        const userAddLink = "/users/" + this.props.id + "/add";

        const userRows = users.map((user, index) => {
            var name = user.firstName + " ";
            if (user.middleName) {
                name += user.middleName + " " + user.lastName;
            } else {
                name += user.lastName;
            }
            return (
                <div className="row" key={index}>
                    <span className="column">{name}</span>
                    <span className="column">{user.email}</span>
                    <span className="column status">
                        {user.isGuardian
                            ? email == user.email
                                ? "(guardian, primary contact)"
                                : "(guardian)"
                            : "(student)"}
                    </span>
                </div>
            );
        });

        var balance = 0;
        const transactionRows = transactions.map((transaction, index) => {
            const url = "/accounts/transaction/" + transaction.id + "/edit";
            const amount = transaction.amount;
            balance += parseInt(amount);
            return (
                <div className="row" key={index}>
                    <span className="medium-column">
                        {transaction.date.format("MM-DD-YYYY")}
                    </span>
                    <span className="column">{transaction.type}</span>
                    <span className="column">
                        {this.formatCurrency(amount)}
                    </span>
                    <span className="large-column">{transaction.notes}</span>
                    <span className="edit">
                        <Link to={url}>{"Edit >"}</Link>
                    </span>
                </div>
            );
        });

        return (
            <section id="account-info">
                <span id="account-number">Account No. {id}</span>

                <div id="account-users">
                    <h1>Users in Account</h1>
                    {userRows}

                    <button id="add-user">
                        <Link className="button" to={userAddLink}>
                            Add New User to Account
                        </Link>
                    </button>
                </div>

                <div id="account-transactions">
                    <h1>Transaction History</h1>
                    <div className="header row">
                        <span className="medium-column">Date</span>
                        <span className="column">Type</span>
                        <span className="column">Amount</span>
                        <span className="large-column">Notes</span>
                        <span className="edit"></span>
                    </div>
                    {transactionRows}
                    <div id="transaction-footer">
                        <span>
                            Account Balance:{" "}
                            <b>{this.formatCurrency(balance)}</b>
                        </span>
                        <button id="add-transaction">
                            <Link
                                className="button"
                                to="/accounts/transaction/add">
                                Add Transaction
                            </Link>
                        </button>
                    </div>
                </div>
            </section>
        );
    };
}
