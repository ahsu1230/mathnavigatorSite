"use strict";
require("./accountSearcher.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import { Searcher } from "./searcher.js";
import { UserRowCard } from "./userRowCard.js";

export default class AccountSearcher extends React.Component {
    state = {
        account: {},
        users: [],
        searched: false, // true means we have attempted a search
        found: false, // true means the search was a success
    };

    onFoundAccount = (account) => {
        this.setState({
            account: account,
        });
        this.props.onFoundAccount && this.props.onFoundAccount(account);
    };

    onFoundUsers = (users) => {
        this.setState({
            users: users,
        });
        this.props.onFoundUsers && this.props.onFoundUsers(users);
    };

    onSearchSuccess = () => {
        this.setState({
            searched: true,
            found: true,
        });
    };

    onSearchFailed = () => {
        this.setState({
            account: {},
            users: [],
            searched: true,
            found: false,
        });
    };

    render() {
        return (
            <div className="account-searcher">
                <Searcher
                    type="account"
                    onFoundAccount={this.onFoundAccount}
                    onFoundUsers={this.onFoundUsers}
                    onSearchSuccess={this.onSearchSuccess}
                    onSearchFailed={this.onSearchFailed}
                />

                {this.state.searched && this.state.found && (
                    <Content
                        account={this.state.account}
                        users={this.state.users}
                    />
                )}
                {this.state.searched && !this.state.found && (
                    <p>The account you're looking for does not exist.</p>
                )}
            </div>
        );
    }
}

class Content extends React.Component {
    render() {
        const account = this.props.account || {};
        const users = (this.props.users || []).map((user, index) => {
            const editUrl = "/account/" + account.id + "?view=edit-users";
            return (
                <div className="" key={index}>
                    <UserRowCard
                        user={user}
                        editTitle={"View/Edit User"}
                        editUrl={editUrl}
                    />
                </div>
            );
        });
        return (
            <div className="content">
                <div className="account-info">
                    <h4>Account Information</h4>
                    <p>AccountId: {account.id}</p>
                    <p>Primary Email Contact: {account.primaryEmail}</p>
                    <p>
                        Account Created on{" "}
                        {moment(account.createdAt).format("l")}
                    </p>
                    <p>Last updated {moment(account.updatedAt).fromNow()}</p>
                    <Link to={"/account/" + account.id}>
                        View Account Details
                    </Link>
                </div>
                <div className="users-info">
                    <h4>Users found in Account</h4>
                    {users}
                </div>
            </div>
        );
    }
}
