"use strict";
require("./accountUserSearcher.sass");
import React from "react";
import moment from "moment";
import Handler from "./searchHandler.js";
import { getFullName } from "../userUtils.js";

const SEARCH_BY_ID = "search_by_id";
const SEARCH_BY_EMAIL = "search_by_email";

export default class AccountUserSearcher extends React.Component {
    state = {
        search: "",
        searchBy: SEARCH_BY_ID,
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
        const isSearchAccount = this.props.isSearchAccount || false;

        this.setState({ searched: false, found: false });

        if (searchBy == SEARCH_BY_ID && !parseInt(searchQuery)) {
            alert("If searching by an id, the id must be a positive number!");
            return;
        }

        if (isSearchAccount && searchBy == SEARCH_BY_ID) {
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
        } else if (isSearchAccount && searchBy == SEARCH_BY_EMAIL) {
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
        } else if (!isSearchAccount && searchBy == SEARCH_BY_ID) {
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
        } else if (!isSearchAccount && searchBy == SEARCH_BY_EMAIL) {
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
            console.log("Unrecognized search pattern.");
        }
    };

    render() {
        const isSearchAccount = this.props.isSearchAccount || false;
        const title = isSearchAccount
            ? "Search for an Account"
            : "Search for a user";
        const placeholder =
            this.state.searchBy == SEARCH_BY_ID
                ? "Enter an id"
                : "Enter an email";

        const accountInfo = this.state.accountInfo;

        return (
            <div className="user-account-searcher">
                {/* Left Side (with title, searcher, and accountInfo) */}
                <div className="content">
                    <h2>{title}</h2>
                    <section className="searcher">
                        <select
                            value={this.state.searchBy}
                            onChange={this.onChangeSelection}>
                            <option value={SEARCH_BY_ID}>Search by ID</option>
                            <option value={SEARCH_BY_EMAIL}>
                                Search by Email
                            </option>
                        </select>
                        <input
                            type="text"
                            placeholder={placeholder}
                            value={this.state.search}
                            onChange={this.onChangeInput}
                        />
                        <button onClick={this.onClickSearch}>Search</button>
                    </section>

                    {this.state.searched && this.state.found && (
                        <AccountInfo accountInfo={this.state.accountInfo} />
                    )}

                    {this.state.searched &&
                        !this.state.found &&
                        (isSearchAccount ? (
                            <p>The account you're looking for is not found.</p>
                        ) : (
                            <p>The user you're looking for is not found.</p>
                        ))}
                </div>

                {/* Scrollable right side (list of users) */}
                {this.state.searched && this.state.found && (
                    <UserInfos userInfos={this.state.userInfos} />
                )}
            </div>
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
                </div>
            );
        });
        return (
            <div className="users">
                <h3>{users.length} Users in Account</h3>
                {users}
            </div>
        );
    }
}
