"use strict";
require("./account.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API, { executeApiCalls } from "../api.js";
import { getCurrentAccountId, setCurrentAccountId } from "../localStorage.js";

export class AccountPage extends React.Component {
    state = {
        id: "",
        account: {},
        users: [],
        transactions: [],

        searchId: "",
        searchEmail: "",
    };

    componentDidMount = () => {
        const id = getCurrentAccountId() || "";

        if (id) {
            API.get("api/accounts/" + id)
                .then((res) => {
                    this.setState({
                        id: id,
                        account: res.data,
                    });
                    this.fetchUserData(id);
                    this.fetchTransactionData(id);
                })
                .catch((err) => {
                    window.alert(
                        "Could not fetch account: " + err.response.data
                    );
                });
        }
    };

    fetchUserData = (id) => {
        API.get("api/users/account/" + id).then((res) => {
            this.setState({
                users: res.data,
            });
        });
    };

    // Using fake data until backend is ready
    fetchTransactionData = (id) => {
        const transactions = [
            {
                date: moment("2020-04-22"),
                type: "cash",
                amount: 380,
                notes: "Payment for Aaron's SAT Math 1 Spring 2020",
            },
            {
                date: moment("2020-04-24"),
                type: "charge",
                amount: -800,
                notes: "Charge for Austin's Kindergarten Reading",
            },
        ];

        this.setState({
            transactions: transactions,
        });
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickSearch = (value, email = false) => {
        var url = "api/accounts/account/";
        if (email) {
            url = "api/accounts/search";
        }

        API.get(url + value)
            .then((res) => {
                this.setState({
                    id: value,
                    account: res.data,
                });
                this.fetchUserData(value);
                this.fetchTransactionData(value);
            })
            .catch((err) => {
                window.alert("Could not fetch account: " + err.response.data);
            });
    };

    onClickDeleteAccount = () => {
        const id = this.props.id;

        let apiCalls = [];
        let successCallback = () =>
            console.log("Successfully deleted account and all users!");
        let failCallback = (err) =>
            alert("Could not delete account or users: " + err.response.data);

        // Must delete users before deleting account
        this.state.users.forEach((user) => {
            apiCalls.push(API.delete("api/users/user/" + user.id));
        });
        apiCalls.push(API.delete("api/accounts/account/" + id));

        executeApiCalls(apiCalls, successCallback, failCallback);
    };

    render = () => {
        const users = this.state.users.map((user, index) => {
            return (
                <div className="row" key={index}>
                    <span className="column">
                        {user.firstName + " " + user.lastName}
                    </span>
                    <span className="column">{user.email}</span>
                    <span className="column">{name}</span>
                    <span className="column">
                        {isGuardian ? "(guardian)" : "(student)"}
                    </span>
                </div>
            );
        });

        const transactions = this.state.transactions.map(
            (transaction, index) => {
                return (
                    <div className="row" key={index}>
                        <span className="column">{transaction.date}</span>
                        <span className="column">{transaction.type}</span>
                        <span className="column">{transaction.amount}</span>
                        <span className="column">{transaction.notes}</span>
                        <Link to={"/accounts/transaction/" + transaction.id}>
                            {"Edit >"}
                        </Link>
                    </div>
                );
            }
        );

        return (
            <div id="view-account">
                <section id="search-accounts">
                    <h2>Search Accounts</h2>
                    <div className="search-input">
                        <input
                            type="text"
                            placeholder="Search by Account ID"
                            onChange={(e) => this.handleChange(e, "searchId")}
                        />
                        <button
                            className="btn-search"
                            onClick={() =>
                                this.onClickSearch(this.state.searchId)
                            }>
                            Search
                        </button>
                    </div>
                    <div className="search-input">
                        <input
                            type="text"
                            placeholder="Search by Primary Email"
                            onChange={(e) =>
                                this.handleChange(e, "searchEmail")
                            }
                        />
                        <button
                            className="btn-search"
                            onClick={() =>
                                this.onClickSearch(this.state.searchEmail, true)
                            }>
                            Search
                        </button>
                    </div>

                    <button>
                        <Link id="create-account" to={"/accounts/add"}>
                            Add New User to Account
                        </Link>
                    </button>
                </section>

                <section id="account-users">
                    <span>Account: {this.state.id}</span>
                    <h2>Users in Account</h2>
                    <div id="user-rows">{users}</div>

                    <button>
                        <Link id="add-user" to={"/users/add"}>
                            Add New User to Account
                        </Link>
                    </button>
                </section>

                <section id="account-transactions">
                    <h2>Transaction History</h2>
                    <div className="header row">
                        <span className="column">Date</span>
                        <span className="column">Type</span>
                        <span className="column">Amount</span>
                        <span className="column">Notes</span>
                        <span className="edit"></span>
                    </div>
                    {transactions}
                </section>

                <button
                    id="btn-delete-account"
                    onClick={this.onClickDeleteAccount}>
                    Delete Account and All Users
                </button>
            </div>
        );
    };
}
