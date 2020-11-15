"use strict";
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API from "../../api.js";
import { formatCurrency } from "../../common/userUtils.js";

export default class AccountTransactions extends React.Component {
    state = {
        transactions: [],
    };

    componentDidMount() {
        const accountId = this.props.accountId;
        API.get("api/transactions/account/" + accountId)
            .then((res) => {
                const transactions = res.data || [];
                this.setState({ transactions: transactions });
            })
            .catch((err) => {
                console.log("Error searching account " + err);
            });
    }

    render() {
        const addTransactionUrl =
            "/account/" + this.props.accountId + "/transaction/add";
        const totalBalance = this.state.transactions.reduce(
            (accum, curr) => accum + curr.amount,
            0
        );
        const transactions = this.state.transactions.map((trans, index) => {
            const editTransactionUrl =
                "/account/" +
                this.props.accountId +
                "/transaction/" +
                trans.id +
                "/edit";
            return (
                <div key={index}>
                    <div>{moment(trans.createdAt).format("l")}</div>
                    <div>{trans.type}</div>
                    <div>{formatCurrency(trans.amount)}</div>
                    <div>{trans.notes}</div>
                    <Link to={editTransactionUrl}>Edit</Link>
                </div>
            );
        });

        return (
            <section>
                <h2>Account Balance {formatCurrency(totalBalance)}</h2>
                {transactions}
                <Link to={addTransactionUrl}>
                    <button>Add New Transaction</button>
                </Link>
            </section>
        );
    }
}
