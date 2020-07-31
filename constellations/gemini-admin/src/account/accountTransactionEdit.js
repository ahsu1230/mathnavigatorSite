"use strict";
require("./accountTransactionEdit.sass");
import React from "react";
import API from "../api.js";
import { getCurrentAccountId } from "../localStorage.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText, emptyValidator } from "../utils/inputText.js";

export class TransactionEditPage extends React.Component {
    state = {
        isEdit: false,
        types: [],
        type: "",
        amount: "",
        notes: "",
        accountId: getCurrentAccountId(),
    };

    componentDidMount = () => {
        const id = this.props.id;
        if (id) {
            API.get("api/transactions/transaction/" + id)
                .then((res) => {
                    const transaction = res.data;
                    this.setState({
                        isEdit: true,
                        type: transaction.paymentType,
                        amount: parseInt(transaction.amount),
                        notes: transaction.paymentNotes || "",
                    });
                })
                .catch((err) =>
                    alert("Could not fetch transaction: " + err.response.data)
                );
        }

        API.get("api/transactions/types")
            .then((res) =>
                this.setState({
                    type: res.data[0],
                    types: res.data,
                })
            )
            .catch((err) =>
                alert("Could not fetch types: " + err.response.data)
            );
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = "accounts";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onModalDeleteConfirm = () => {
        const id = this.props.id;
        API.delete("api/transactions/transaction/" + id)
            .then(() => {
                window.location.hash = "accounts";
            })
            .catch((err) => {
                alert("Could not delete transaction: " + err.response.data);
            })
            .finally(() => this.onDismissModal());
    };

    onClickSave = () => {
        let transaction = {
            paymentType: this.state.type,
            amount: parseInt(this.state.amount),
            paymentNotes: this.state.notes,
            accountId: this.state.accountId,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save transaction: " + err.response.data);

        if (this.state.isEdit) {
            API.post(
                "api/transactions/transaction/" + this.props.id,
                transaction
            )
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/transactions/create", transaction)
                .then(() => successCallback())
                .catch((err) => failCallback(err));
        }
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
        const isEdit = this.state.isEdit;
        var title = isEdit ? "Edit Transaction" : "Add Transaction";
        title += " for Account No. " + this.state.accountId;

        const modalDiv = renderModal(
            this.state.showSaveModal,
            this.state.showDeleteModal,
            this.onModalOkSaved,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        const typeOptions = this.state.types.map((type, index) => (
            <option value={type} key={index}>
                {type}
            </option>
        ));

        let deleteButton = <div></div>;
        if (isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        const chargeValidators = [
            emptyValidator("amount"),
            {
                validate: (amount) =>
                    !(this.state.type != "charge" && amount < 0),
                message: "Payment must be positive",
            },
            {
                validate: (amount) =>
                    !(this.state.type == "charge" && amount > 0),
                message: "Charge must be negative",
            },
        ];

        return (
            <div id="view-transaction-edit">
                {modalDiv}
                <h1>{title}</h1>

                <div id="transaction-type">
                    <h2>Type</h2>
                    <h4>Enter the type of transaction</h4>
                    <select
                        className="dropdown"
                        value={this.state.type}
                        onChange={(e) => this.handleChange(e, "type")}>
                        {typeOptions}
                    </select>
                </div>

                <InputText
                    label="Amount"
                    description="Enter the amount of money"
                    required={true}
                    value={this.state.amount}
                    onChangeCallback={(e) => this.handleChange(e, "amount")}
                    validators={chargeValidators}
                />

                <InputText
                    label="Notes"
                    description="Enter notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                />

                <div className="buttons">
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    {deleteButton}
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                </div>
            </div>
        );
    };
}

function renderModal(
    showSaveModal,
    showDeleteModal,
    onModalOkSaved,
    onModalDeleteConfirm,
    onModalDismiss
) {
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
    if (showSaveModal) {
        showModal = showSaveModal;
        modalContent = (
            <OkayModal
                text={"Transaction information saved!"}
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
