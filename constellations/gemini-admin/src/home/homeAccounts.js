"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

// does not work

const sectionDisplayNames = {
    unpaid: "Unpaid Accounts",
};

export class HomeTabSectionAccounts extends React.Component {
    state = {
        unpaidAcc: [],
    };

    // counter to keep track of the number of unpaid accounts => unpaidAcc.length

    //unpaid accounts
    /*  componentDidMount() {
        API.get("api/transactions").then((res) => {
            const transaction = res.data;
            this.setState({
                unpaidAcc: transaction,
            });
        });
    } */

    render() {
        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">Unpaid Accounts</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>View All Accounts</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="header-flex">
                        <div className={"list-header" + " name"}>Name</div>
                        <div className={"list-header" + " email"}>Email</div>
                        <div className={"list-header" + " other"}>
                            Unpaid Amount
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
