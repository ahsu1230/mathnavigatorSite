"use strict";
require("./account.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API, { executeApiCalls } from "../api.js";
import { getCurrentAccountId, setCurrentAccountId } from "../localStorage.js";
import { Modal } from "../modals/modal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AccountPage extends React.Component {
    state = {
        id: 0,
        account: {},
        users: [],
        transactions: [],

        searchId: "",
        searchEmail: "",
        invalid: false,
    };

    componentDidMount = () => {
        const id = getCurrentAccountId() || 0;

        if (id) {
            API.get("api/accounts/account/" + id)
                .then((res) => {
                    this.setState({
                        id: id,
                        account: res.data,
                        invalid: false,
                    });
                    this.fetchUserData(id);
                    this.fetchTransactionData(id);
                })
                .catch((err) => this.getAccountError(err));
        }
    };

    getAccountError = (err) => {
        this.setState({
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
        API.get("api/accounts/account/" + id)
            .then((res) => {
                this.setState({
                    id: id,
                    account: res.data,
                });
                setCurrentAccountId(id);
                this.fetchUserData(id);
                this.fetchTransactionData(id);
            })
            .catch((err) => this.getAccountError(err));
    };

    onClickSearchByEmail = (email) => {
        API.post("api/accounts/search", email)
            .then((res) => {
                const id = res.data.id;
                this.setState({
                    id: id,
                    account: res.data,
                });
                setCurrentAccountId(id);
                this.fetchUserData(id);
                this.fetchTransactionData(id);
            })
            .catch((err) => this.getAccountError(err));
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
        const users = this.state.users.map((user, index) => {
            return (
                <div className="row" key={index}>
                    <span className="column">
                        {user.firstName + " " + user.lastName}
                    </span>
                    <span className="column">{user.email}</span>
                    <span className="column">{name}</span>
                    <span className="column">
                        {user.isGuardian ? "(guardian)" : "(student)"}
                    </span>
                </div>
            );
        });

        const transactions = this.state.transactions.map(
            (transaction, index) => {
                return (
                    <div className="row" key={index}>
                        <span className="column">
                            {transaction.date.format("MM-DD-YYYY")}
                        </span>
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

        const modalDiv = renderModal(
            this.state.showDeleteModal,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        var errorMessage = <div></div>;
        if (this.state.invalid) {
            errorMessage = <h4 className="red">No Account Found</h4>;
        }

        return (
            <div id="view-account" className={this.state.id == 0 ? "hide" : ""}>
                {modalDiv}
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
                                this.onClickSearchById(this.state.searchId)
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
                    {errorMessage}
                    <button>
                        <Link id="create-account" to={"/accounts/add"}>
                            Create New Account
                        </Link>
                    </button>
                </section>

                <section id="account-info">
                    <div id="account-users">
                        <span>Account: {this.state.id}</span>
                        <h2>Users in Account</h2>
                        <div id="user-rows">{users}</div>

                        <button>
                            <Link id="add-user" to={"/users/add"}>
                                Add New User to Account
                            </Link>
                        </button>
                    </div>

                    <div id="account-transactions">
                        <h2>Transaction History</h2>
                        <div className="header row">
                            <span className="column">Date</span>
                            <span className="column">Type</span>
                            <span className="column">Amount</span>
                            <span className="column">Notes</span>
                            <span className="edit"></span>
                        </div>
                        {transactions}
                    </div>

                    <button
                        id="btn-delete-account"
                        onClick={this.onClickDeleteAccount}>
                        Delete Account and All Users
                    </button>
                </section>
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
    }
    if (modalContent) {
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
