"use strict";
require("./account.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API, { executeApiCalls } from "../api.js";
import { getCurrentAccountId, setCurrentAccountId } from "../localStorage.js";
import { Modal } from "../modals/modal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { AccountInfo } from "./accountInfo.js";

export class AccountPage extends React.Component {
    state = {
        id: 0,
        email: "",
        users: [],
        transactions: [],

        searchId: "",
        searchEmail: "",
        search: "",
        searchValue: "",
        invalid: false,
    };

    componentDidMount = () => {
        const id = getCurrentAccountId() || 0;

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
            invalid: false,
        });
        setCurrentAccountId(id);
        this.fetchUserData(id);
        this.fetchTransactionData(id);
    };

    fetchDataError = (err) => {
        this.setState({
            users: [],
            transactions: [],
            invalid: true,
        });
        console.log("Could not fetch account: " + err.response.data);
    };

    fetchUserData = (id) => {
        API.get("api/users/account/" + id)
            .then((res) => {
                this.setState({
                    users: res.data,
                });
            })
            .catch((err) => {
                window.alert("Could not fetch users: " + err.response.data);
            });
    };

    // Using fake data until backend is ready
    fetchTransactionData = (id) => {
        // API.get("api/transactions/account/" + id)
        //     .then((res) => {
        //         this.setState({
        //             transactions: res.data,
        //         });
        //     })
        //     .catch((err) => {
        //         window.alert("Could not fetch transactions: " + err.response.data);
        //     });

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

    onClickSearchById = (id) => {
        this.setState({
            search: "id",
            searchValue: id,
        });
        API.get("api/accounts/account/" + id)
            .then((res) => this.fetchData(res))
            .catch((err) => this.fetchDataError(err));
    };

    onClickSearchByEmail = (email) => {
        this.setState({
            search: "email",
            searchValue: email,
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
            console.log("Successfully deleted account and all users!");
            this.onModalDismiss();
            this.setState({
                id: 0,
                users: [],
                transactions: [],
                invalid: true,
            });
        };

        let failCallback = (err) =>
            alert("Could not delete account or users: " + err.response.data);

        // Must delete users before deleting account
        this.state.users.forEach((user) => {
            apiCalls.push(API.delete("api/users/user/" + user.id));
        });
        apiCalls.push(API.delete("api/accounts/account/" + id));

        executeApiCalls(apiCalls, successCallback, failCallback);
    };

    onModalDismiss = () => {
        this.setState({
            showDeleteModal: false,
        });
    };

    render = () => {
        const modalDiv = renderModal(
            this.state.showDeleteModal,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        var errorMessage = <div></div>;
        if (this.state.invalid) {
            const value = this.state.search + " " + this.state.searchValue;
            errorMessage = (
                <h4 className="red">No account with {value} found</h4>
            );
        }

        return (
            <div id="view-account" className={this.state.id == 0 ? "hide" : ""}>
                {modalDiv}
                <section id="search-accounts">
                    <h1>Search Accounts</h1>
                    <div className="container">
                        <div>
                            <div className="search-input">
                                <input
                                    type="text"
                                    placeholder="Search by Account ID"
                                    onChange={(e) =>
                                        this.handleChange(e, "searchId")
                                    }
                                />
                                <button
                                    className="btn-search"
                                    onClick={() =>
                                        this.onClickSearchById(
                                            this.state.searchId
                                        )
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
                                        this.onClickSearchByEmail(
                                            this.state.searchEmail
                                        )
                                    }>
                                    Search
                                </button>
                            </div>
                        </div>

                        <button id="create-account">
                            <Link className="button" to="/accounts/add">
                                Create New Account
                            </Link>
                        </button>
                    </div>

                    {errorMessage}
                </section>

                <div className={this.state.invalid ? "hide" : ""}>
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

function renderModal(showDeleteModal, onModalDeleteConfirm, onModalDismiss) {
    let modalDiv;
    let modalContent;
    let showModal;
    if (showDeleteModal) {
        showModal = showDeleteModal;
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
                show={showModal}
                onDismiss={onModalDismiss}
            />
        );
    }
    return modalDiv;
}
