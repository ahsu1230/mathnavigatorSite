"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";

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
                        {" "}
                        <AccountInfo accountId={row.account.id} />{" "}
                    </div>
                    <div className="email">{row.account.primaryEmail} </div>
                    <div className="other">{row.account.id} </div>
                </li>
            );
        });

        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">Unpaid Accounts</h3>{" "}
                    <button className="view-details">
                        <Link to={"/accounts"}>View All Accounts</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="container-flex">
                        <div className={"list-header" + " name"}>Name</div>
                        <div className={"list-header" + " email"}>Email</div>
                        <div className={"list-header" + " other"}>
                            Account Id
                        </div>
                    </div>

                    <ul>{unpaidAcc}</ul>
                </div>
            </div>
        );
    }
}

class AccountInfo extends React.Component {
    state = {
        account: {},
    };
    componentDidMount() {
        API.get("api/accounts/account/" + this.props.accountId).then((res) => {
            const accountData = res.data;
            this.setState({
                account: accountData,
            });
        });
    }

    render() {
        console.log("account info " + this.state.account);

        let returnName = this.state.account;

        return <div>{returnName}</div>;
    }
}
