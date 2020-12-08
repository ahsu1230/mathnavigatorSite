"use strict";
require("./searcher.sass");
import React from "react";
import Handler from "./searchHandler.js";

const SEARCH_TYPE_ACCOUNT = "account";
const SEARCH_TYPE_USER = "user";
const SEARCH_TYPE_USER_LIST = "user_list";

const SEARCH_BY_ID = "search_by_id";
const SEARCH_BY_EMAIL = "search_by_email";

/**
 * A general searcher that can be used for searching for accounts or users.
 *
 * Note: This component is a helper component and should ONLY be used by searchers in this package
 * (i.e. accountSearcher, singleUserSearcher, manyUserSearcher).
 * You should import those other components instead of importing this one.
 *
 * Parameters:
 *  - type = ["account", "user", "user_list"]. Based on the type, the searcher will look slightly different and offer different functionalities.
 *  - onFoundUser = a function that when a single user is found, invoke this function. Typically used when type is "user".
 *  - onFoundUsers = a function that when many users are found, invoke this function. Typically used when type is "user_list".
 *  - onFoundAccount = a function when when an account is found, invoke this function.
 */
export class Searcher extends React.Component {
    state = {
        search: "",
        searchBy: this.props.type == SEARCH_TYPE_USER_LIST ? "" : SEARCH_BY_ID,
    };

    onFoundAccount = (account) => {
        if (this.props.onFoundAccount) {
            this.props.onFoundAccount(account);
        }
    };

    onFoundUser = (user) => {
        if (this.props.onFoundUser) {
            this.props.onFoundUser(user);
        }
    };

    onFoundUsers = (users) => {
        if (this.props.onFoundUsers) {
            this.props.onFoundUsers(users);
        }
    };

    onSearchSuccess = () => {
        if (this.props.onSearchSuccess) {
            this.props.onSearchSuccess();
        }
    };

    onSearchFailed = () => {
        if (this.props.onSearchFailed) {
            this.props.onSearchFailed();
        }
    };

    onChangeSelection = (e) => {
        this.setState({ searchBy: e.target.value });
    };

    onChangeInput = (e) => {
        this.setState({ search: e.target.value });
    };

    onClickSearch = () => {
        const searchQuery = this.state.search;
        const searchBy = this.state.searchBy;
        const searchType = this.props.type;

        this.setState({ searched: false, found: false });

        if (searchBy == SEARCH_BY_ID && !parseInt(searchQuery)) {
            alert("If searching by an id, the id must be a positive number!");
            return;
        }

        if (searchType == SEARCH_TYPE_ACCOUNT) {
            this.searchForAccount(searchBy, searchQuery);
        } else if (searchType == SEARCH_TYPE_USER) {
            this.searchForOneUser(searchBy, searchQuery);
        } else if (searchType == SEARCH_TYPE_USER_LIST) {
            this.searchForAllUsers(searchQuery);
        } else {
            window.alert(
                "Something went wrong with your search. Unrecognized search type."
            );
        }
    };

    searchForAccount = (searchBy, searchQuery) => {
        if (searchBy == SEARCH_BY_ID) {
            let accountId = parseInt(searchQuery);
            Handler.searchAccountById(
                accountId,
                (account) => this.onFoundAccount(account),
                (users) => this.onFoundUsers(users),
                this.onSearchSuccess,
                this.onSearchFailed
            );
        } else if (searchBy == SEARCH_BY_EMAIL) {
            Handler.searchAccountByEmail(
                searchQuery,
                (account) => this.onFoundAccount(account),
                (users) => this.onFoundUsers(users),
                this.onSearchSuccess,
                this.onSearchFailed
            );
        } else {
            window.alert("Unrecognized search query.");
        }
    };

    searchForOneUser = (searchBy, searchQuery) => {
        if (searchBy == SEARCH_BY_ID) {
            let userId = parseInt(searchQuery);
            Handler.searchUserById(
                userId,
                (user) => this.onFoundUser(user),
                (account) => this.onFoundAccount(account),
                this.onSearchSuccess,
                this.onSearchFailed
            );
        } else if (searchBy == SEARCH_BY_EMAIL) {
            Handler.searchUserByEmail(
                searchQuery,
                (user) => this.onFoundUser(user),
                (account) => this.onFoundAccount(account),
                this.onSearchSuccess,
                this.onSearchFailed
            );
        } else {
            window.alert("Unrecognized search query.");
        }
    };

    searchForAllUsers = (searchQuery) => {
        Handler.searchUsers(
            searchQuery,
            (users) => this.onFoundUsers(users),
            this.onSearchSuccess,
            this.onSearchFailed
        );
    };

    getTitle = () => {
        const searchType = this.props.type;
        return searchType == SEARCH_TYPE_USER_LIST
            ? "Search for any user"
            : searchType == SEARCH_TYPE_ACCOUNT
            ? "Search for an Account"
            : "Search for a user";
    };

    getPlaceholder = () => {
        return this.props.type == SEARCH_TYPE_USER_LIST
            ? "Enter a name or an email"
            : this.state.searchBy == SEARCH_BY_ID
            ? "Enter an id"
            : "Enter an email";
    };

    render() {
        const searchType = this.props.type;
        const title = this.getTitle();
        const placeholder = this.getPlaceholder();

        return (
            <article className="user-account-searcher">
                <section className="searcher">
                    <h2>{title}</h2>
                    {(searchType == SEARCH_TYPE_ACCOUNT ||
                        searchType == SEARCH_TYPE_USER) && (
                        <select
                            value={this.state.searchBy}
                            onChange={this.onChangeSelection}>
                            <option value={SEARCH_BY_ID}>Search by ID</option>
                            <option value={SEARCH_BY_EMAIL}>
                                Search by Email
                            </option>
                        </select>
                    )}
                    <input
                        type="text"
                        placeholder={placeholder}
                        value={this.state.search}
                        onChange={this.onChangeInput}
                    />
                    <button onClick={this.onClickSearch}>Search</button>
                </section>
            </article>
        );
    }
}
