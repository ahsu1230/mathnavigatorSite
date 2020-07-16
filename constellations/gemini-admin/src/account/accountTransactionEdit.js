"use strict";
require("./accountTransactionEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { getCurrentAccountId } from "../localStorage.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";

// React DatePicker
import "react-dates/initialize";
import "react-dates/lib/css/_datepicker.css";
import { SingleDatePicker } from "react-dates";

export class TransactionEditPage extends React.Component {
    state = {
        isEdit: false,
        date: moment(),
        type: "",
        amount: "",
        notes: "",
    };

    componentDidMount = () => {
        const id = this.props.id;
        if (id) {
            API.get("api/transactions/transaction/" + id).then((res) => {
                const transaction = res.data;
                this.setState({
                    isEdit: true,
                    date: moment(transaction.programId),
                    type: transaction.programId,
                    amount: transaction.name,
                    notes: transaction.notes || "",
                });
            });
        }
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
            date: this.state.date,
            type: this.state.type,
            amount: parseInt(this.state.amount),
            notes: this.state.notes,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save transaction: " + err.response.data);
        if (this.state.isEdit) {
            API.post("api/transactions/transaction/" + this.state.id, account)
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

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    render = () => {
        const isEdit = this.state.isEdit;
        var title = isEdit ? "Edit Transaction" : "Add Transaction";
        title += " (Account " + getCurrentAccountId() + ")";

        const modalDiv = renderModal(
            this.state.showSaveModal,
            this.state.showDeleteModal,
            this.onModalOkSaved,
            this.onModalDeleteConfirm,
            this.onModalDismiss
        );

        let deleteButton = <div></div>;
        if (isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        return (
            <div id="view-transaction-edit">
                {modalDiv}
                <h2>{title}</h2>

                <div id="date-container">
                    <h2>Date</h2>
                    <SingleDatePicker
                        id="date-picker"
                        date={this.state.date}
                        onDateChange={(date) => this.onDateChange(date)}
                        focused={this.state.dateFocused}
                        onFocusChange={({ focused }) =>
                            this.setState({
                                dateFocused: focused,
                            })
                        }
                        showDefaultInputIcon
                    />
                </div>

                <InputText
                    label="Type"
                    value={this.state.type}
                    onChangeCallback={(e) => this.handleChange(e, "type")}
                    required={true}
                    description="Enter the type of transaction"
                    validators={[
                        {
                            validate: (type) => type != "",
                            message: "You must input a type",
                        },
                    ]}
                />

                <InputText
                    label="Amount"
                    value={this.state.amount}
                    onChangeCallback={(e) => this.handleChange(e, "amount")}
                    required={true}
                    description="Enter the amount of money"
                    validators={[
                        {
                            validate: (amount) =>
                                !isNaN(amount) && amount != "",
                            message: "You must input an amount",
                        },
                    ]}
                />

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                    description="Enter notes"
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
