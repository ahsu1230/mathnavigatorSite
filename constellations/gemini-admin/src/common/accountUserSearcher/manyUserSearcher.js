"use strict";
require("./userSearcher.sass");
import React from "react";
import { Searcher } from "./searcher.js";
import { UserRowCard } from "../rowCards/userRowCard.js";

export default class ManyUserSearcher extends React.Component {
    state = {
        users: [],
        searched: false, // true means we have attempted a search
        found: false, // true means the search was a success
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
            users: [],
            searched: true,
            found: false,
        });
    };

    render() {
        return (
            <div className="many-user-searcher">
                <Searcher
                    type="user_list"
                    onFoundUsers={this.onFoundUsers}
                    onSearchSuccess={this.onSearchSuccess}
                    onSearchFailed={this.onSearchFailed}
                />

                {this.state.searched && this.state.found && (
                    <Content users={this.state.users} />
                )}
                {this.state.searched && !this.state.found && (
                    <p>No users found from search query.</p>
                )}
            </div>
        );
    }
}

class Content extends React.Component {
    render() {
        const users = (this.props.users || []).map((user, index) => {
            const editUrl =
                "/account/" +
                user.accountId +
                "?view=edit-users&userId=" +
                user.id;
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
                <div className="users-info">
                    <h4>Users found from search</h4>
                    {users}
                </div>
            </div>
        );
    }
}
