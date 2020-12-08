"use strict";
require("./userSearcher.sass");
import React from "react";
import { Searcher } from "./searcher.js";
import { UserRowCard } from "../rowCards/userRowCard.js";

export default class SingleUserSearcher extends React.Component {
    state = {
        account: {},
        user: {},
        searched: false, // true means we have attempted a search
        found: false, // true means the search was a success
    };

    onFoundAccount = (account) => {
        this.setState({
            account: account,
        });
        this.props.onFoundAccount && this.props.onFoundAccount(account);
    };

    onFoundUser = (user) => {
        this.setState({
            user: user,
        });
        this.props.onFoundUser && this.props.onFoundUser(user);
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
            user: {},
            searched: true,
            found: false,
        });
    };

    render() {
        return (
            <div className="single-user-searcher">
                <Searcher
                    type="user"
                    onFoundAccount={this.onFoundAccount}
                    onFoundUser={this.onFoundUser}
                    onSearchSuccess={this.onSearchSuccess}
                    onSearchFailed={this.onSearchFailed}
                />

                {this.state.searched && this.state.found && (
                    <Content
                        account={this.state.account}
                        user={this.state.user}
                    />
                )}
                {this.state.searched && !this.state.found && (
                    <p>The user you're looking for does not exist.</p>
                )}
            </div>
        );
    }
}

class Content extends React.Component {
    render() {
        const account = this.props.account;
        const user = this.props.user;
        const editUrl =
            "/account/" + user.accountId + "?view=edit-users&userId=" + user.id;
        return (
            <div className="content">
                <div className="user-info">
                    <h4>Found User Information</h4>
                    <UserRowCard
                        user={user}
                        account={account}
                        editTitle={"View/Edit User"}
                        editUrl={editUrl}
                    />
                </div>
            </div>
        );
    }
}
