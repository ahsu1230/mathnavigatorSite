"use strict";
require("./accountSearchPage.sass");
import React from "react";
import { Link } from "react-router-dom";
import AccountSearcher from "../common/accountUserSearcher/accountSearcher.js";

export class AccountSearchPage extends React.Component {
    render() {
        return (
            <main id="view-account-search">
                <p>
                    Click to create a new Account or search for a single account
                    by
                    <br />
                    inputing an AccountId or an account's primary email.
                </p>
                <Link to="/account/create">
                    <button className="create">Create New Account</button>
                </Link>
                <section>
                    <AccountSearcher />
                </section>
            </main>
        );
    }
}
