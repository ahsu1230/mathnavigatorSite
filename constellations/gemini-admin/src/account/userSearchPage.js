"use strict";
require("./userSearchPage.sass");
import React from "react";
import ManyUserSearcher from "../common/accountUserSearcher/manyUserSearcher.js";

export class UserSearchPage extends React.Component {
    render() {
        return (
            <main id="view-user-search">
                <p>
                    Input a user's name or email to search for their
                    information.
                </p>
                <ManyUserSearcher />
            </main>
        );
    }
}
