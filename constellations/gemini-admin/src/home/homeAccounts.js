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
        // flexbox for headers (Name, Email, Account Number?, Unpaid Amount? )

        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">Unpaid Accounts</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>View All Accounts ??</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="list-header">Account Number</div>
                </div>
            </div>
        );
    }
}
