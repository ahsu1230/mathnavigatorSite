"use strict";
require("./userEdit.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";
import { setCurrentAccountId } from "../localStorage.js";

export class UserEditPage extends React.Component {
    state = {
        isEdit: false,

        id: 0,
        firstName: "",
        lastName: "",
        middleName: "",
        email: "",
        phone: "",
        isGuardian: false,
        accountId: 0,
        notes: "",

        allUsers: [],
        primaryEmail: "",
        originalEmail: "",
    };

    componentDidMount = () => {
        this.fetchData();
    };

    componentDidUpdate = () => {
        if (this.state.isEdit && this.props.id != this.state.id) {
            this.fetchData();
        }
    };

    fetchData = () => {
        const userId = this.props.id;

        if (userId) {
            API.get("api/users/user/" + userId).then((res) => {
                const user = res.data;
                this.setState({
                    isEdit: true,
                    id: user.id,
                    firstName: user.firstName,
                    lastName: user.lastName,
                    middleName: user.middleName || "",
                    email: user.email,
                    phone: user.phone,
                    isGuardian: user.isGuardian,
                    accountId: user.accountId,
                    notes: user.notes || "",

                    originalEmail: user.email,
                });
                API.get("api/users/account/" + user.accountId)
                    .then((res) => {
                        this.setState({
                            allUsers: res.data,
                        });
                    })
                    .catch((err) => {
                        window.alert(
                            "Could not fetch users: " + err.response.data
                        );
                    });
                API.get("api/accounts/account/" + user.accountId)
                    .then((res) => {
                        this.setState({
                            primaryEmail: res.data.primaryEmail,
                        });
                    })
                    .catch((err) => {
                        window.alert(
                            "Could not fetch account: " + err.response.data
                        );
                    });
            });
        }

        const accountId = this.props.accountId;
        if (accountId) {
            this.setState({ accountId: accountId });
            API.get("api/users/account/" + accountId)
                .then((res) => {
                    this.setState({
                        allUsers: res.data,
                    });
                })
                .catch((err) => {
                    window.alert("Could not fetch users: " + err.response.data);
                });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = this.state.isEdit ? "users" : "accounts";
    };

    onClickSave = () => {
        let user = {
            id: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            middleName: this.state.middleName,
            email: this.state.email,
            phone: this.state.phone,
            isGuardian: this.state.isGuardian,
            accountId: parseInt(this.state.accountId),
            notes: this.state.notes,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save user: " + err.response.data);
        API.post(
            this.state.isEdit
                ? "api/users/user/" + this.state.id
                : "api/users/create",
            user
        )
            .then(() => successCallback())
            .catch((err) => failCallback(err));
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onConfirmDelete = () => {
        const id = this.state.id;
        API.delete("api/users/user/" + id).then(
            () =>
                (window.location.hash = this.state.isEdit
                    ? "users"
                    : "accounts")
        );
    };

    onSavedOk = () => {
        this.onDismissModal();
        window.location.hash = this.state.isEdit ? "users" : "accounts";
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    onChangeGuardian = () => {
        this.setState({ isGuardian: !this.state.isGuardian });
    };

    onClickAccountDetails = () => {
        setCurrentAccountId(this.state.accountId);
    };

    render = () => {
        let deleteButton = <div></div>;
        if (
            this.state.isEdit &&
            this.state.originalEmail != this.state.primaryEmail
        ) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete User From Account
                </button>
            );
        }

        let associatedAccount = <div></div>;
        if (this.state.accountId != 0) {
            const otherUsers = this.state.allUsers.filter(
                (item) => item.id != this.state.id
            );

            let otherUsersRows = otherUsers.map((user, index) => {
                const url = "/users/" + user.id + "/edit";
                return (
                    <p key={index}>
                        <Link to={url}>
                            {user.firstName + " " + user.lastName}
                        </Link>
                    </p>
                );
            });

            let otherUsersHeader = null;
            if (otherUsers.length > 0) {
                otherUsersHeader = <h3>Other Users in Account</h3>;
            }

            associatedAccount = (
                <div id="associated-account">
                    <h2>Associated Account</h2>
                    <div className="account-details-wrapper">
                        <p>Account Id: {this.state.accountId}</p>
                        <Link
                            onClick={this.onClickAccountDetails}
                            to="/accounts">
                            View Details
                        </Link>
                    </div>
                    {otherUsersHeader}
                    {otherUsersRows}
                </div>
            );
        }

        let modalDiv;
        let modalContent;
        let showModal;
        if (this.state.showDeleteModal) {
            showModal = this.state.showDeleteModal;
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={this.onConfirmDelete}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalContent = (
                <OkayModal
                    text={"User information saved!"}
                    onOkay={this.onSavedOk}
                />
            );
        }
        if (modalContent) {
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showModal}
                    onDismiss={this.onDismissModal}
                />
            );
        }

        return (
            <div id="view-user-edit">
                {modalDiv}
                <h1>{this.state.isEdit ? "Edit" : "Add"} User</h1>
                <div id="column-container">
                    <div id="left-column">
                        <InputText
                            label="First Name"
                            value={this.state.firstName}
                            onChangeCallback={(e) =>
                                this.handleChange(e, "firstName")
                            }
                            required={true}
                            description="Enter your first name"
                            validators={[
                                {
                                    validate: (name) => name != "",
                                    message: "You must input a name",
                                },
                            ]}
                        />

                        <InputText
                            label="Middle Name"
                            value={this.state.middleName}
                            onChangeCallback={(e) =>
                                this.handleChange(e, "middleName")
                            }
                            description="Enter your middle name"
                        />

                        <InputText
                            label="Last Name"
                            value={this.state.lastName}
                            onChangeCallback={(e) =>
                                this.handleChange(e, "lastName")
                            }
                            required={true}
                            description="Enter your last name"
                            validators={[
                                {
                                    validate: (name) => name != "",
                                    message: "You must input a name",
                                },
                            ]}
                        />

                        <InputText
                            label="Email"
                            value={this.state.email}
                            onChangeCallback={(e) =>
                                this.handleChange(e, "email")
                            }
                            required={true}
                            description="Enter your email address"
                            validators={[
                                {
                                    validate: (email) => email != "",
                                    message: "You must input an email",
                                },
                                {
                                    validate: (email) =>
                                        /^[a-zA-Z0-9]+@[a-zA-Z]+\.[a-zA-Z]+/.test(
                                            email
                                        ),
                                    message: "Invalid email",
                                },
                            ]}
                        />

                        <InputText
                            label="Phone"
                            value={this.state.phone}
                            onChangeCallback={(e) =>
                                this.handleChange(e, "phone")
                            }
                            required={true}
                            description="Enter your phone number"
                            validators={[
                                {
                                    validate: (num) => num != "",
                                    message: "You must input a phone number",
                                },
                                {
                                    validate: (num) => /^\d{10,}$/.test(num),
                                    message: "Invalid phone number",
                                },
                            ]}
                        />

                        <h2 id="guardian">Is this user a guardian?</h2>
                        <input
                            type="checkbox"
                            checked={this.state.isGuardian}
                            onChange={this.onChangeGuardian}
                        />
                        <span>Yes</span>

                        {
                            //TODO: Add the following InputTexts when orion users have school/graduation year data.
                            /*<InputText
                            label="School"
                            value={this.state.school}
                            onChangeCallback={(e) => this.handleChange(e, "school")}
                            required={false}
                            description="Enter your school"
                        />
                        <InputText
                            label="Graduation Year"
                            value={this.state.gradYear}
                            onChangeCallback={(e) => this.handleChange(e, "gradYear")}
                            required={true}
                            description="Enter your graduation year"
                        />*/
                        }
                    </div>
                    {associatedAccount}
                </div>

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                    description="Add any notes"
                />

                <div className="buttons">
                    <div id="buttons-left">
                        <button
                            onClick={this.onClickCancel}
                            className="btn-cancel">
                            Cancel
                        </button>
                        {deleteButton}
                    </div>
                    <div id="buttons-right">
                        <button onClick={this.onClickSave} className="btn-save">
                            Save
                        </button>
                    </div>
                </div>
            </div>
        );
    };
}
