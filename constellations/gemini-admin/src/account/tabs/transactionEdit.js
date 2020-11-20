"use strict";
import React from "react";
import API from "../../api.js";
import EditPageWrapper from "../../common/editPages/editPageWrapper.js";
import { InputText } from "../../common/inputs/inputText.js";
import { InputSelect } from "../../common/inputs/inputSelect.js";

export class TransactionEditPage extends React.Component {
    state = {
        isEdit: false,
        transactionId: 0,
        createdAt: null,
        type: "",
        amount: 0,
        notes: "",
        allTypes: [],
    };

    componentDidMount = () => {
        const transactionId = this.props.transactionId;
        if (transactionId) {
            API.get("api/transactions/transaction/" + transactionId).then(
                (res) => {
                    const transaction = res.data;
                    this.setState({
                        isEdit: true,
                        transactionId: transaction.id,
                        createdAt: transaction.createdAt,
                        type: transaction.type,
                        amount: transaction.amount,
                        notes: transaction.notes,
                    });
                }
            );
        }
        API.get("api/transactions/types").then((res) => {
            this.setState({ allTypes: res.data });
        });
    };

    handleChange = (e, fieldName) => {
        this.setState({
            [fieldName]: e.target.value,
        });
    };

    onSave = () => {
        const transaction = {
            accountId: parseInt(this.props.accountId),
            type: this.state.type,
            amount: parseInt(this.state.amount),
            notes: this.state.notes,
        };
        const transactionId = this.state.transactionId;
        if (this.state.isEdit) {
            return API.post(
                "api/transactions/transaction/" + transactionId,
                transaction
            );
        } else {
            return API.post("api/transactions/create", transaction);
        }
    };

    onDelete = () => {
        const transactionId = this.state.transactionId;
        return API.delete("api/transactions/transaction/" + transactionId);
    };

    renderContent = () => {
        const options = this.state.allTypes.map((option) => {
            return {
                value: option,
                displayName: option,
            };
        });
        return (
            <div>
                <InputSelect
                    label="Transaction Type"
                    value={this.state.type}
                    onChangeCallback={(e) => this.handleChange(e, "type")}
                    required={true}
                    options={options}
                />
                <InputText
                    label="Amount"
                    description="Enter the amount of this transaction. The amount MUST be negative if the transaction type is 'charge'. Otherwise, if it is a payment, the amount should be positive."
                    value={this.state.amount}
                    onChangeCallback={(e) => this.handleChange(e, "amount")}
                    required={true}
                />
                <InputText
                    label="Notes"
                    description="Enter any additional information about this transaction."
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                />
            </div>
        );
    };

    render() {
        const accountId = this.props.accountId;
        const transactionId = this.props.transactionId;
        const isEdit = this.state.isEdit;
        const title = isEdit ? "Edit Transaction" : "Add Transaction";
        return (
            <div id="view-transaction-edit">
                <EditPageWrapper
                    isEdit={isEdit}
                    title={title}
                    content={this.renderContent()}
                    prevPageUrl={"account/" + accountId + "?view=transactions"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityName={"transaction"}
                />
            </div>
        );
    }
}
