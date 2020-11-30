"use strict";
import React from "react";
require("./accountTransactions.sass");
import { Link } from "react-router-dom";
import moment from "moment";
import classnames from "classnames";
import API from "../../api.js";
import RowCardSlim from "../../common/rowCards/rowCardSlim.js";
import { formatCurrency } from "../../common/displayUtils.js";
import {
    getAccountBalance,
    sortTransactionsLatestFirst,
} from "../../common/transactionUtils.js";

export default class AccountTransactions extends React.Component {
    state = {
        transactions: [],
        sortedTransactions: [],
    };

    componentDidMount() {
        const accountId = this.props.accountId;
        API.get("api/transactions/account/" + accountId)
            .then((res) => {
                const transactions = res.data || [];
                this.setState({
                    transactions: transactions,
                    sortedTransactions: sortTransactionsLatestFirst(
                        transactions
                    ),
                });
            })
            .catch((err) => {
                console.log("Error searching account " + err);
            });
    }

    render() {
        const addTransactionUrl =
            "/account/" + this.props.accountId + "/transaction/add";
        const totalBalance = getAccountBalance(this.state.transactions);
        const transactions = this.state.sortedTransactions.map(
            (transaction, index) => {
                const editUrl =
                    "/account/" +
                    this.props.accountId +
                    "/transaction/" +
                    transaction.id +
                    "/edit";
                const fields = [
                    { value: transaction.type },
                    {
                        value: formatCurrency(transaction.amount),
                        highlightFn: () => transaction.amount < 0,
                    },
                ];
                return (
                    <RowCardSlim
                        key={index}
                        inlineTitle={moment(transaction.createdAt).format("l")}
                        fields={fields}
                        text={transaction.notes}
                        editUrl={editUrl}
                    />
                );
            }
        );
        const titleClasses = classnames("", { alert: totalBalance < 0 });
        return (
            <section className="account-tab transactions">
                <div className="top-container">
                    <h3 className={titleClasses}>
                        Account Balance: {formatCurrency(totalBalance)}
                    </h3>
                    <div>
                        <Link to={addTransactionUrl}>
                            <button className="add-transaction">
                                Add New Transaction
                            </button>
                        </Link>
                    </div>
                </div>

                <div className="transaction-list">{transactions}</div>
            </section>
        );
    }
}
