"use strict";
require("./userSearcher.sass");
import React from "react";
import { Searcher } from "./searcher.js";
import { UserRowCard } from "./userRowCard.js";

export default class SingleUserSearcher extends React.Component {
    state = {
        account: {},
        user: {},
        show: false,
    };

    onFoundAccount = (account) => {
        this.setState({
            account: account,
            show: true,
        });
        this.props.onFoundAccount && this.props.onFoundAccount(account);
    };

    onFoundUser = (user) => {
        this.setState({
            user: user,
            show: true,
        });
        this.props.onFoundUser && this.props.onFoundUser(user);
    };

    render() {
        return (
            <div className="single-user-searcher">
                <Searcher
                    type="user"
                    onFoundAccount={this.onFoundAccount}
                    onFoundUser={this.onFoundUser}
                />

                {this.state.show && (
                    <Content
                        account={this.state.account}
                        user={this.state.user}
                    />
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
