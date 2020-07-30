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
import { UserInput } from "./userInput.js";

export class UserEditPage extends React.Component {
    state = {
        isEdit: false,

        id: 0,
        firstName: "",
        middleName: "",
        lastName: "",
        email: "",
        phone: "",
        isGuardian: false,
        accountId: 0,
        school: "",
        graduationYear: "",
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
            API.get("api/users/user/" + userId)
                .then((res) => {
                    const user = res.data;
                    this.setState({
                        isEdit: true,
                        id: user.id,
                        firstName: user.firstName,
                        middleName: user.middleName || "",
                        lastName: user.lastName,
                        email: user.email,
                        phone: user.phone,
                        isGuardian: user.isGuardian,
                        accountId: user.accountId,
                        school: user.school || "",
                        graduationYear: user.graduationYear || "",
                        notes: user.notes || "",

                        originalEmail: user.email,
                    });
                    this.fetchAccountData(user.accountId);
                })
                .catch((err) => {
                    window.alert("Could not fetch user: " + err.response.data);
                });
        }

        const accountId = this.props.accountId;
        if (accountId) {
            this.setState({ accountId: accountId });
            this.fetchAccountData(accountId);
        }
    };

    fetchAccountData = (id) => {
        API.get("api/users/account/" + id)
            .then((res) => {
                this.setState({
                    allUsers: res.data,
                });
            })
            .catch((err) => {
                window.alert("Could not fetch users: " + err.response.data);
            });
        API.get("api/accounts/account/" + id)
            .then((res) => {
                this.setState({
                    primaryEmail: res.data.primaryEmail,
                });
            })
            .catch((err) => {
                window.alert("Could not fetch account: " + err.response.data);
            });
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onChangeGuardian = () => {
        this.setState({ isGuardian: !this.state.isGuardian });
    };

    returnToPage = () => {
        window.location.hash = this.state.isEdit ? "users" : "accounts";
    };

    onClickSave = () => {
        let user = {
            id: this.state.id,
            firstName: this.state.firstName,
            middleName: this.state.middleName,
            lastName: this.state.lastName,
            email: this.state.email,
            phone: this.state.phone,
            isGuardian: this.state.isGuardian,
            accountId: parseInt(this.state.accountId),
            school: this.state.school,
            graduationYear: parseInt(this.state.graduationYear),
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
        API.delete("api/users/user/" + id).then(() => this.returnToPage());
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    onClickAccountDetails = () => {
        setCurrentAccountId(this.state.accountId);
    };

    renderAssociatedAccount = () => {
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

            return (
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
    };

    renderButtons = () => {
        let deleteButton;
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
        return (
            <div className="buttons">
                <div id="buttons-left">
                    <button onClick={this.returnToPage} className="btn-cancel">
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
        );
    };

    renderModal = () => {
        let modalContent;
        if (this.state.showDeleteModal) {
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={this.onConfirmDelete}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            modalContent = (
                <OkayModal
                    text={"User information saved!"}
                    onOkay={this.returnToPage}
                />
            );
        }
        if (modalContent) {
            return (
                <Modal
                    content={modalContent}
                    show={true}
                    onDismiss={this.onDismissModal}
                />
            );
        }
    };

    render = () => {
        let modal = this.renderModal();
        let associatedAccount = this.renderAssociatedAccount();
        let buttonSection = this.renderButtons();

        return (
            <div id="view-user-edit">
                {modal}
                <h1>{this.state.isEdit ? "Edit" : "Add"} User</h1>
                <div id="column-container">
                    <div id="left-column">
                        <UserInput
                            handleChange={this.handleChange}
                            onChangeGuardian={this.onChangeGuardian}
                            firstName={this.state.firstName}
                            middleName={this.state.middleName}
                            lastName={this.state.lastName}
                            email={this.state.email}
                            phone={this.state.phone}
                            isGuardian={this.state.isGuardian}
                            school={this.state.school}
                            graduationYear={this.state.graduationYear}
                        />
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

                {buttonSection}
            </div>
        );
    };
}
