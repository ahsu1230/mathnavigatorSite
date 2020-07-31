"use strict";
require("./account.sass");
import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import API, { executeApiCalls } from "../api.js";
import { getCurrentAccountId, setCurrentAccountId } from "../localStorage.js";
import { Modal } from "../modals/modal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { AccountInfo } from "./accountInfo.js";

export class AccountPage extends React.Component {
    state = {
        id: getCurrentAccountId() || 0,
        email: "",
        users: [],
        transactions: [],

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

    onClickDeleteAccount = () => {
        this.setState({ showDeleteModal: true });
    };

    onModalDeleteConfirm = () => {
        const id = this.state.id;

        let apiCalls = [];
        let successCallback = () => {
            console.log("Successfully deleted account!");

            this.setState({
                id: 0,
                users: [],
                transactions: [],
                errorType: "empty",
            });

            setCurrentAccountId(0);
            this.onModalDismiss();
        };

        let failCallback = (err) =>
            alert("Could not delete account: " + err.response.data);

        // TODO: the backend will eventually handle deleting all users, transactions, userAFHs, and userClasses
        this.state.users.forEach((user) =>
            apiCalls.push(API.delete("api/users/user/" + user.id))
        );
        this.state.transactions.forEach((transaction) =>
            apiCalls.push(
                API.delete("api/transactions/transaction/" + transaction.id)
            )
        );
        apiCalls.push(API.delete("api/accounts/account/" + id));

        executeApiCalls(apiCalls, successCallback, failCallback);
    };

    onModalDismiss = () => {
        this.setState({
            showDeleteModal: false,
        });
    };

    renderModal = (showDeleteModal, onModalDeleteConfirm, onModalDismiss) => {
        let modalDiv;
        let modalContent;
        if (showDeleteModal) {
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={onModalDeleteConfirm}
                    onReject={onModalDismiss}
                />
            );
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showDeleteModal}
                    onDismiss={onModalDismiss}
                />
            );
        }
        return modalDiv;
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
        const modalDiv = this.renderModal(
            this.state.showDeleteModal,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        const errorMessage = this.generateErrorMessage(
            this.state.errorType,
            this.state.errorValue
        );

        return (
            <div id="view-account">
                {modalDiv}
                <section id="search-accounts">
                    <h1>Search Accounts</h1>
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

                        <button>
                            <Link to="/accounts/add">Create New Account</Link>
                        </button>
                    </div>

                    {errorMessage}
                </section>

                <div className={this.state.errorType == "" ? "" : "hide"}>
                    <AccountInfo
                        id={this.state.id}
                        email={this.state.email}
                        users={this.state.users}
                        transactions={this.state.transactions}
                    />
                    <button
                        id="btn-delete-account"
                        onClick={this.onClickDeleteAccount}>
                        Delete Account and All Users
                    </button>
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
