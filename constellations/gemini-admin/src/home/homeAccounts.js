"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";
import { EmptyMessage } from "./home.js";
import moment from "moment";

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
                <li className="container-flex" key={index}>
                    <div className="name">
                        <p id="account"> Account Id: {row.account.id} </p>
                        <AccountInfo accountId={row.account.id} />
                    </div>
                    <div className="email">{row.account.primaryEmail} </div>
                    <div className="balance">
                        {"-$"}
                        {Math.abs(row.balance)}
                    </div>
                    <div className="from-now">
                        {moment(row.updatedAt).fromNow()}{" "}
                    </div>
                    <div className="view">
                        <Link to={"/accounts?accountId=" + row.accountId}>
                            {"View >"}
                        </Link>
                    </div>
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
                    <div className="container-flex">
                        <div className={"list-header name"}>Account</div>
                        <div className={"list-header email"}>Email</div>
                        <div className={"list-header balance"}>Balance</div>
                        <div className={"list-header from-now"}>
                            Last Updated
                        </div>
                        <div className={"list-header view"}> </div>
                    </div>
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
        API.get("api/users/account/" + this.props.accountId).then((res) => {
            const userData = res.data;
            this.setState({
                users: userData,
            });
        });
    }

    render() {
        let returnName = this.state.users.map((row, index) => {
            return (
                <div key={index} className="list-names">
                    {getFullName(row)}{" "}
                </div>
            );
        });

        return <div> {returnName} </div>;
    }
}
