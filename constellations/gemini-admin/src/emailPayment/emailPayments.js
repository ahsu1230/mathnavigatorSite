"use strict";
require("./emailPayments.sass");
import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import API, { executeApiCalls } from "../api.js";
import { getCurrentAccountId, setCurrentAccountId } from "../localStorage.js";
import { AccountInfo } from "./accountInfo.js";

export class EmailPaymentsPage extends React.Component {
    state = {
        id: getCurrentAccountId() || 0,
        email: "",
        users: [],
        transactions: [],
        selectedUserEmails: [],

        searchId: "",
        searchEmail: "",
        errorType: "empty",
        errorValue: "",
    };

    componentDidMount = () => {
        const id = this.state.id;

        if (id) {
            API.get("api/accounts/account/" + id)
                .then((res) => this.fetchData(res))
                .catch((err) => this.fetchDataError(err));
        }
    };

    fetchData = (res) => {
        const id = res.data.id;
        this.setState({
            id: id,
            email: res.data.primaryEmail,
            users: [],
            transactions: [],
            errorType: "",
        });
        setCurrentAccountId(id);

        const apiCalls = [
            API.get("api/users/account/" + id),
            API.get("api/transactions/account/" + id),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) =>
                    this.setState({
                        users: responses[0].data,
                        transactions: responses[1].data,
                    })
                )
            )
            .catch((err) =>
                alert("Could not fetch users or transactions: " + err)
            );
    };

    fetchDataError = (err) => {
        this.setState({
            users: [],
            transactions: [],
        });
        console.log("Could not fetch account: " + err.response.data);
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickSearchById = (id) => {
        this.setState({
            errorType: "id",
            errorValue: id,
        });
        API.get("api/accounts/account/" + id)
            .then((res) => this.fetchData(res))
            .catch((err) => this.fetchDataError(err));
    };

    onClickSearchByEmail = (email) => {
        this.setState({
            errorType: "email",
            errorValue: email,
        });
        API.post("api/accounts/search", { primaryEmail: email })
            .then((res) => this.fetchData(res))
            .catch((err) => this.fetchDataError(err));
    };

    generateErrorMessage = (errorType, errorValue) => {
        var errorMessage = <div></div>;
        switch (errorType) {
            case "empty":
                errorMessage = (
                    <h4>
                        Please search for an account using an ID number or a
                        primary email
                    </h4>
                );
                break;
            case "id":
                errorMessage = <h4>No account with id {errorValue} found</h4>;
                break;
            case "email":
                errorMessage = (
                    <h4>
                        No account with primary email {errorValue} found. You
                        may search for a user's email{" "}
                        <Link to="/users">here</Link>
                    </h4>
                );
                break;
        }

        if (["id", "email"].includes(errorType) && errorValue == "") {
            errorMessage = <h4>No {errorType} entered</h4>;
        }

        return errorMessage;
    };

    render = () => {
        const errorMessage = this.generateErrorMessage(
            this.state.errorType,
            this.state.errorValue
        );

        var names = "";

        this.state.users.map((user, index) => {
            if (index === 0) {
                const name = user.firstName;
                names += name;
            }
        });

        return (
            <div id="view-account">
                <section id="search-accounts">
                    <h1>Generate Payment Reminder Email</h1>
                    <div className="container">
                        <AccountSearch
                            placeholder="Search by Account ID"
                            changeCallback={(e) =>
                                this.handleChange(e, "searchId")
                            }
                            searchCallback={() =>
                                this.onClickSearchById(this.state.searchId)
                            }
                        />
                        <AccountSearch
                            placeholder="Search by Primary Email"
                            changeCallback={(e) =>
                                this.handleChange(e, "searchEmail")
                            }
                            searchCallback={() =>
                                this.onClickSearchByEmail(
                                    this.state.searchEmail
                                )
                            }
                        />
                    </div>

                    {errorMessage}
                </section>

                <div className={this.state.errorType == "" ? "" : "hide"}>
                    <AccountInfo
                        id={this.state.id}
                        email={this.state.email}
                        users={this.state.users}
                        transactions={this.state.transactions}
                        name={names}
                    />
                </div>
            </div>
        );
    };
}

class AccountSearch extends React.Component {
    render = () => {
        return (
            <div className="search-input">
                <input
                    type="text"
                    placeholder={this.props.placeholder}
                    onChange={this.props.changeCallback}
                />
                <button
                    className="btn-search"
                    onClick={this.props.searchCallback}>
                    Search
                </button>
            </div>
        );
    };
}
