"use strict";
require("./homeSection.sass");
import React from "react";
import { UserRowCard } from "../common/rowCards/userRowCard.js";

export class HomeTabSectionUsers extends React.Component {
    render() {
        const users = this.props.users || [];
        const list = users.map((user, index) => {
            const editUrl = "/account/" + user.accountId + "?view=edit-users";
            return (
                <li key={index}>
                    <UserRowCard
                        user={user}
                        editTitle={"View/Edit User"}
                        editUrl={editUrl}
                    />
                </li>
            );
        });

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">New Users</h3>
                </div>
                {list.length > 0 ? (
                    <div className="class-section">
                        <ul>{list}</ul>
                    </div>
                ) : (
                    <p className="empty">No new users recently.</p>
                )}
            </div>
        );
    }
}
