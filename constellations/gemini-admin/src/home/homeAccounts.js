"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../api.js";
import moment from "moment";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";

export class HomeTabSectionAccounts extends React.Component {
    render() {
        const accounts = (this.props.unpaidAccounts || []).map((row, index) => {
            return (
                <li key={index}>
                    <AccountInfo account={row.account} balance={row.balance} />
                </li>
            );
        });

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">Unpaid Accounts</h3>
                    <Link to={"/accounts"}>View All Accounts</Link>
                </div>

                {accounts.length > 0 ? (
                    <div className="class-section">
                        <ul>{accounts}</ul>
                    </div>
                ) : (
                    <p className="empty">All accounts have paid. Awesome!</p>
                )}
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
                editTitle="View"
                editUrl={"/account/" + account.id + "?view=transactions"}
                fieldsList={[this.firstColumn(), this.secondColumn()]}
            />
        );
    }
}
