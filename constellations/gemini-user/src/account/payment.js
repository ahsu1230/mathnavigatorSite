"use strict";
require("./account.sass");
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
                            {moment(0).format("l") /*Fake data*/}
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

        return (
            <div className="tab-content" id="payment-tab">
                <div>
                    <h2>Account Balance: {formatCurrency(balance)}</h2>
                    <p>
                        You currently owe {formatCurrency(-balance)}. Please pay
                        through any of the following methods:
                    </p>

                    <span>
                        - <Link to="">Cash</Link>
                        <br />
                    </span>
                    <span>
                        - <Link to="">Check</Link> (written to Math Navigator)
                        <br />
                    </span>
                    <span>
                        - <Link to="">Paypal</Link>
                    </span>

                    <p>
                        For questions about your account, please contact us at{" "}
                        <a
                            href="mailto:andymathnavigator@gmail.com"
                            className="orange">
                            andymathnavigator@gmail.com
                        </a>
                    </p>
                </div>
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
            </div>
        );
    };
}
