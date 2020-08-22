"use strict";
require("./payment.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../utils/api.js";
import moment from "moment";

import {
    chargeDisplayNames,
    formatCurrency,
    fetchError,
} from "./accountUtils.js";

export class PaymentTab extends React.Component {
    state = {
        transactions: [],
    };

    componentDidMount = () => {
        if (this.props.accountId) {
            API.get(
                "api/transactions/account/" + this.props.accountId
            ).then((res) => this.setState({ transactions: res.data }));
        }
    };

    render = () => {
        let balance = 0;
        const transactionsList = this.state.transactions.map(
            (transaction, index) => {
                balance += parseInt(transaction.amount);
                return (
                    <ul key={index}>
                        <li className="li-med">
                            {moment(transaction.createdAt).format("l")}
                        </li>
                        <li className="li-med">
                            {chargeDisplayNames[transaction.paymentType]}
                        </li>
                        <li className="li-med">
                            {formatCurrency(transaction.amount)}
                        </li>
                        <li className="li-large">{formatCurrency(balance)}</li>
                    </ul>
                );
            }
        );
        transactionsList.reverse();

        const balanceMessage =
            balance >= 0
                ? "You currently have " +
                  formatCurrency(balance) +
                  " in your account. Our payment options are:"
                : "You currently owe " +
                  formatCurrency(-balance) +
                  ". Please pay through any of the following methods:";

        const transactionsSection =
            this.state.transactions.length > 0 ? (
                <div>
                    <h2>Your Payment History</h2>
                    <ul className="header">
                        <li className="li-med">Date</li>
                        <li className="li-med">Transaction</li>
                        <li className="li-med">Amount</li>
                        <li className="li-large">Balance</li>
                    </ul>
                    {transactionsList}
                </div>
            ) : (
                <div>
                    <h2>Your Payment History</h2>
                    <p>There are no transactions for this account.</p>
                </div>
            );

        return (
            <div className="tab-content" id="payment-tab">
                <div>
                    <h2>Account Balance: {formatCurrency(balance)}</h2>
                    <p>{balanceMessage}</p>

                    <div>
                        - <Link to="">Cash</Link>
                    </div>
                    <div>
                        - <Link to="">Check</Link> (written to Math Navigator)
                    </div>
                    <div>
                        - <Link to="">Paypal</Link>
                    </div>

                    <p>
                        For questions about your account, please contact us at{" "}
                        <a
                            href="mailto:andymathnavigator@gmail.com"
                            className="orange">
                            andymathnavigator@gmail.com
                        </a>
                    </p>
                    {transactionsSection}
                </div>
            </div>
        );
    };
}
