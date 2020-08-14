"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

export class HomeTabSectionAccounts extends React.Component {
    state = {
        unpaidAccounts: [],
    };

    componentDidMount() {
        API.get("api/accounts/unpaid").then((res) => {
            const transaction = res.data;
            this.setState({
                unpaidAccounts: transaction,
            });
        });
    }

    render() {

        let unpaidAcc = this.state.unpaidAccounts.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">{getFullName(row)} </div>
                    <div className="email">{row.email} </div>
                    <div className="other">{row.accountId} </div>
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
                        <div className={"list-header" + " other"}>Unpaid Amount</div>
                    </div>

                    <ul>{unpaidAcc}</ul>

                </div>
            </div>
        );
    }
}
