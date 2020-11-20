"use strict";
require("./accountSearch.sass");
import React from "react";
import AccountUserSearcher from "../common/accountUserSearcher/accountUserSearcher.js";

export class UserSearchPage extends React.Component {
    render() {
        return (
            <main>
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
