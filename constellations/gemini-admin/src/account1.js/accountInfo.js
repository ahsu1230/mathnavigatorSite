"use strict";
require("./accountInfo.sass");
import React from "react";
import { Link } from "react-router-dom";
import { getFullName, formatCurrency } from "../common/userUtils.js";

export class AccountInfo extends React.Component {
    render = () => {
        const id = this.props.id;
        const email = this.props.email;
        const users = this.props.users;
        const transactions = this.props.transactions;
        const userAddLink = "/users/" + this.props.id + "/add";

        const userRows = users.map((user, index) => {
            const url = "/users/" + user.id + "/edit";
            var status = user.isGuardian ? "(guardian" : "(student";
            status += user.email == email ? ", primary contact)" : ")";

            return (
                <div className="row" key={index}>
                    <span className="column">
                        <Link to={url}>{getFullName(user)}</Link>
                    </span>
                    <span className="column">{user.email}</span>
                    <span className="column status">{status}</span>
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
                    <span className="column">{transaction.type}</span>
                    <span className="medium-column">
                        {formatCurrency(amount)}
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
                    <h2>Users in Account</h2>
                    {userRows}

                    <button id="add-user">
                        <Link className="button" to={userAddLink}>
                            Add New User to Account
                        </Link>
                    </button>
                </div>

                <div id="account-transactions">
                    <h2>Transaction History</h2>
                    <div className="header row">
                        <span className="column">Type</span>
                        <span className="medium-column">Amount</span>
                        <span className="large-column">Notes</span>
                        <span className="edit"></span>
                    </div>
                    {transactionRows}
                    <div id="transaction-footer">
                        <span>
                            Account Balance: <b>{formatCurrency(balance)}</b>
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