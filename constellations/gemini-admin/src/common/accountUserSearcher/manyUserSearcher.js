"use strict";
require("./userSearcher.sass");
import React from "react";
import { Searcher } from "./searcher.js";
import { UserRowCard } from "./userRowCard.js";

export default class ManyUserSearcher extends React.Component {
    state = {
        users: [],
        show: false,
    };

    onFoundUsers = (users) => {
        this.setState({
            users: users,
            show: true,
        });
        this.props.onFoundUsers && this.props.onFoundUsers(users);
    };

    render() {
        return (
            <div className="many-user-searcher">
                <Searcher type="user_list" onFoundUsers={this.onFoundUsers} />

                {this.state.show && <Content users={this.state.users} />}
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
