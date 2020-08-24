"use strict";
require("./accountEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { InputText } from "../utils/inputText.js";
import { generatePassword } from "../utils/userUtils.js";
import { setCurrentAccountId } from "../localStorage.js";
import { UserInput } from "../user/userInput.js";

export class AccountEditPage extends React.Component {
    state = {
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
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onChangeGuardian = () => {
        this.setState({ isGuardian: !this.state.isGuardian });
    };

    onClickCancel = () => {
        window.location.hash = "accounts";
    };

    onClickSave = () => {
        const accountUser = {
            account: {
                primaryEmail: this.state.email,
                password: generatePassword(),
            },
            user: {
                firstName: this.state.firstName,
                middleName: this.state.middleName,
                lastName: this.state.lastName,
                email: this.state.email,
                phone: this.state.phone,
                isGuardian: this.state.isGuardian,
                notes: this.state.notes,
            },
        };

        console.log(accountUser);

        API.post("api/accounts/create", accountUser)
            .then(() => {
                this.setState({ showSaveModal: true });
                setCurrentAccountId(id);
            })
            .catch((err) =>
                alert("Could not create account or user: " + err.response.data)
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

                <InputText
                    label="Notes"
                    description="Add any additional information about this account"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
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
