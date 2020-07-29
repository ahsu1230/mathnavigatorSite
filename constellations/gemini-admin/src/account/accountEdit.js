"use strict";
require("./accountEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { InputText } from "../utils/inputText.js";
import { generatePassword } from "../utils/utils.js";
import { setCurrentAccountId } from "../localStorage.js";

export class AccountEditPage extends React.Component {
    state = {
        firstName: "",
        middleName: "",
        lastName: "",
        email: "",
        phone: "",
        notes: "",
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = "accounts";
    };

    onClickSave = () => {
        let account = {
            primaryEmail: this.state.email,
            password: generatePassword(),
        };

        // TODO: replace these calls with one endpoint
        API.post("api/accounts/create", account)
            .then(() => {
                API.post("api/accounts/search", {
                    primaryEmail: this.state.email,
                })
                    .then((res) => {
                        const id = res.data.id;

                        let user = {
                            firstName: this.state.firstName,
                            middleName: this.state.middleName,
                            lastName: this.state.lastName,
                            email: this.state.email,
                            phone: this.state.phone,
                            isGuardian: true,
                            accountId: id,
                            notes: this.state.notes,
                        };

                        API.post("api/users/create", user)
                            .then(() => {
                                this.setState({ showSaveModal: true });
                                setCurrentAccountId(id);
                            })
                            .catch((err) =>
                                alert(
                                    "Could not save user: " + err.response.data
                                )
                            );
                    })
                    .catch((err) =>
                        alert("Could not fetch account: " + err.response.data)
                    );
            })
            .catch((err) =>
                alert("Could not create account: " + err.response.data)
            );
    };

    onModalOkSaved = () => {
        this.onModalDismiss();
        window.location.hash = "accounts";
    };

    onModalDismiss = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    render = () => {
        const modalDiv = renderModal(
            this.state.showSaveModal,
            this.onModalOkSaved,
            this.onModalDismiss
        );

        return (
            <div id="view-account-edit">
                {modalDiv}
                <h1>Create Primary Contact for New Account</h1>

                <InputText
                    label="First Name"
                    value={this.state.firstName}
                    onChangeCallback={(e) => this.handleChange(e, "firstName")}
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
                    onChangeCallback={(e) => this.handleChange(e, "middleName")}
                    description="Enter your middle name if applicable"
                />

                <InputText
                    label="Last Name"
                    value={this.state.lastName}
                    onChangeCallback={(e) => this.handleChange(e, "lastName")}
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
                    label="Primary Email"
                    value={this.state.email}
                    onChangeCallback={(e) => this.handleChange(e, "email")}
                    required={true}
                    description="Enter the primary email for this account"
                    validators={[
                        {
                            validate: (email) => email != "",
                            message: "You must input an email",
                        },
                    ]}
                />

                <InputText
                    label="Phone"
                    value={this.state.phone}
                    onChangeCallback={(e) => this.handleChange(e, "phone")}
                    required={true}
                    description="Enter your phone number"
                    validators={[
                        {
                            validate: (phone) => phone != "",
                            message: "You must input a phone number",
                        },
                    ]}
                />

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                    description="Add any notes"
                />

                <div className="buttons">
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                </div>
            </div>
        );
    };
}

function renderModal(showSaveModal, onModalOkSaved, onModalDismiss) {
    let modalDiv;
    let modalContent;
    let showModal;
    if (showSaveModal) {
        showModal = showSaveModal;
        modalContent = (
            <OkayModal
                text={"Account and User information saved!"}
                onOkay={onModalOkSaved}
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
