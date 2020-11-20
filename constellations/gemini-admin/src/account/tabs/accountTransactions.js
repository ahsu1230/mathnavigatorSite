"use strict";
import React from "react";
require("./accountTransactions.sass");
import { Link } from "react-router-dom";
import moment from "moment";
import { sortBy } from "lodash";
import classnames from "classnames";
import API from "../../api.js";
import RowCardBasic from "../../common/rowCards/rowCardBasic.js";
import RowCardColumns from "../../common/rowCards/rowCardColumns.js";
import { formatCurrency } from "../../common/displayUtils.js";

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
                    sortedTransactions: sortBy(
                        transactions,
                        "createdAt"
                    ).reverse(),
                });
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
        const transactions = this.state.sortedTransactions.map(
            (transaction, index) => {
                const editUrl =
                    "/account/" +
                    this.props.accountId +
                    "/transaction/" +
                    transaction.id +
                    "/edit";
                return (
                    <RowCardColumns
                        key={index}
                        title={transaction.type}
                        editUrl={editUrl}
                        fieldsList={[
                            [
                                {
                                    label: "Date",
                                    value: moment(transaction.createdAt).format(
                                        "l"
                                    ),
                                },
                            ],
                            [
                                {
                                    label: "Amount",
                                    value: formatCurrency(transaction.amount),
                                    highlightFn: () => transaction.amount < 0,
                                },
                            ],
                        ]}
                        texts={[
                            {
                                label: "",
                                value: transaction.notes,
                            },
                        ]}
                    />
                );
            }
        );
        const titleClasses = classnames("", { alert: totalBalance < 0 });
        return (
            <section className="account-tab transactions">
                <h3 className={titleClasses}>
                    Account Balance: {formatCurrency(totalBalance)}
                </h3>
                {transactions}
                <Link to={addTransactionUrl}>
                    <button className="add-transaction">
                        Add New Transaction
                    </button>
                </Link>
            </section>
        );
    }
}
