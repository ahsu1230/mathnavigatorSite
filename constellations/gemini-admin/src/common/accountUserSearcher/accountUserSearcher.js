"use strict";
require("./accountUserSearcher.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import Handler from "./searchHandler.js";
import { getFullName } from "../userUtils.js";

const SEARCH_TYPE_ACCOUNT = "account";
const SEARCH_TYPE_USER = "user";
const SEARCH_TYPE_USER_LIST = "user_list";

const SEARCH_BY_ID = "search_by_id";
const SEARCH_BY_EMAIL = "search_by_email";

export default class AccountUserSearcher extends React.Component {
    state = {
        search: "",
        searchBy: this.props.type == SEARCH_TYPE_USER_LIST ? "" : SEARCH_BY_ID,
        accountInfo: {},
        userInfos: [],
        searched: false,
        found: false,
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
                (accountInfo) => this.setState({ accountInfo: accountInfo }),
                (users) =>
                    this.setState({
                        userInfos: users,
                        searched: true,
                        found: true,
                    }),
                () => this.setState({ searched: true, found: false })
            );
        } else if (searchBy == SEARCH_BY_EMAIL) {
            Handler.searchAccountByEmail(
                searchQuery,
                (accountInfo) => this.setState({ accountInfo: accountInfo }),
                (users) =>
                    this.setState({
                        userInfos: users,
                        searched: true,
                        found: true,
                    }),
                () => this.setState({ searched: true, found: false })
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
                (user) => this.setState({ userInfos: [user] }),
                (accountInfo) => {
                    this.setState({
                        searched: true,
                        found: true,
                        accountInfo: accountInfo,
                    });
                },
                () => this.setState({ searched: true, found: false })
            );
        } else if (searchBy == SEARCH_BY_EMAIL) {
            Handler.searchUserByEmail(
                searchQuery,
                (user) => this.setState({ userInfos: [user] }),
                (accountInfo) => {
                    this.setState({
                        searched: true,
                        found: true,
                        accountInfo: accountInfo,
                    });
                },
                () => this.setState({ searched: true, found: false })
            );
        } else {
            window.alert("Unrecognized search query.");
        }
    };

    searchForAllUsers = (searchQuery) => {
        Handler.searchUsers(
            searchQuery,
            (users) =>
                this.setState({
                    searched: true,
                    found: true,
                    userInfos: users,
                }),
            () => this.setState({ searched: true, found: false })
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
        const searched = this.state.searched;
        const found = this.state.found;

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
                <section className="content">
                    {searchType == SEARCH_TYPE_ACCOUNT &&
                        searched &&
                        !found && (
                            <p>The account you're looking for is not found.</p>
                        )}

                    {(searchType == SEARCH_TYPE_USER ||
                        searchType == SEARCH_TYPE_USER_LIST) &&
                        searched &&
                        !found && (
                            <p>The user you're looking for is not found.</p>
                        )}

                    {(searchType == SEARCH_TYPE_ACCOUNT ||
                        searchType == SEARCH_TYPE_USER) &&
                        searched &&
                        found && (
                            <div>
                                <AccountInfo
                                    accountInfo={this.state.accountInfo}
                                />
                                <UserInfos
                                    userInfos={this.state.userInfos}
                                    type={SEARCH_TYPE_USER}
                                />
                            </div>
                        )}

                    {searchType == SEARCH_TYPE_USER_LIST &&
                        searched &&
                        found && (
                            <UserInfos
                                userInfos={this.state.userInfos}
                                type={SEARCH_TYPE_USER_LIST}
                            />
                        )}
                </section>
            </article>
        );
    }
}

class AccountInfo extends React.Component {
    render() {
        const accountInfo = this.props.accountInfo;
        return (
            <div className="account">
                <h3>Account Information</h3>
                <div>Account Id: {accountInfo.id}</div>
                <div>Primary Email: {accountInfo.primaryEmail}</div>
                <div>Created: {moment(accountInfo.createdAt).fromNow()}</div>
                <Link to={"/account/" + accountInfo.id}>View</Link>
            </div>
        );
    }
}

class UserInfos extends React.Component {
    render() {
        const users = this.props.userInfos.map((user, index) => {
            return (
                <div key={index} className="user">
                    <h4>{getFullName(user)}</h4>
                    <div>UserId: {user.id}</div>
                    <div>Email: {user.email}</div>
                    <div>Phone: {user.phone}</div>
                    {user.isGuardian && <div>IsGuardian: true</div>}
                    {user.school && <div>School: {user.school}</div>}
                    {user.graduationYear && (
                        <div>Graduation Year: {user.graduationYear}</div>
                    )}
                    <div>Created: {moment(user.createdAt).fromNow()}</div>
                    {user.notes && <p>Notes: {user.notes}</p>}
                    <Link to={"/account/" + user.accountId}>View</Link>
                </div>
            );
        });

        const title =
            this.props.type == SEARCH_TYPE_USER_LIST
                ? users.length + " User(s) found"
                : users.length + " User(s) in Account";

        return (
            <div className="users">
                <h3>{title}</h3>
                {users}
            </div>
        );
    }
}
