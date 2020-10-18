"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";
import { EmptyMessage } from "./home.js";

const TAB_UNPAID = "unpaid";

export class HomeTabSectionAccounts extends React.Component {
    state = {
        unpaidAccounts: [],
    };

    componentDidMount() {
        API.get("api/accounts/unpaid").then((res) => {
            const accounts = res.data;
            this.setState({
                unpaidAccounts: accounts,
            });
        });
    }

    render() {
        let unpaidAcc = this.state.unpaidAccounts.map((row, index) => {
            return (
                <li key={index}>
                    <AccountInfo account={row.account} balance={row.balance} />
                </li>
            );
        });

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">Unpaid Accounts</h3>{" "}
                    <button className="view-details">
                        <Link to={"/accounts"}>View All Accounts</Link>
                    </button>
                </div>

                <div className="class-section">
                    <EmptyMessage
                        section={TAB_UNPAID}
                        length={this.state.unpaidAccounts.length}
                    />
                    <ul>{unpaidAcc}</ul>
                </div>
            </div>
        );
    }
}

class AccountInfo extends React.Component {
    state = {
        users: [],
    };

    componentDidMount() {
        API.get("api/users/account/" + this.props.account.id).then((res) => {
            const userData = res.data;
            this.setState({
                users: userData,
            });
        });
    }

    firstColumn = () => {
        const account = this.props.account;
        return [
            {
                label: "Created",
                value: moment(account.createdAt).fromNow(),
            },
            {
                label: "Last Updated",
                value: moment(account.updatedAt).fromNow(),
            },
            {
                label: "Balance",
                value: "-$" + Math.abs(this.props.balance),
                highlightFn: () => true,
            },
        ];
    };

    secondColumn = () => {
        return this.state.users.map((user) => {
            return {
                label: getFullName(user),
                value: user.email,
            };
        });
    };

    render() {
        const account = this.props.account;
        return (
            <RowCardColumns
                title={"Account No. " + account.id}
                subtitle={account.primaryEmail}
                editUrl={"/accounts?accountId=" + account.id}
                fieldsList={[this.firstColumn(), this.secondColumn()]}
            />
        );
    }
}
