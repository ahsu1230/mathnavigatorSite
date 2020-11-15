"use strict";
require("./accountSearch.sass");
import React from "react";
import { Link } from "react-router-dom";
import AccountUserSearcher from "../common/accountUserSearcher/accountUserSearcher.js";

export class AccountSearchPage extends React.Component {
    render() {
        return (
            <main>
                <section>
                    <AccountUserSearcher type="account" />
                </section>

                <Link to="/account/create">
                    <button>Created New Account</button>
                </Link>

                <section>
                    <AccountUserSearcher type="user" />
                </section>

                <section>
                    <AccountUserSearcher type="user_list" />
                </section>
            </main>
        );
    }
}
