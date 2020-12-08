"use strict";
import React from "react";
import { getFullName } from "../../common/userUtils.js";

export default class UserSelector extends React.Component {
    onChange = (e) => {
        const newUserId = parseInt(e.target.value);
        this.props.onChange(newUserId);
    };

    render() {
        const userOptions = this.props.users.map((user, index) => (
            <option key={index} value={user.id}>
                {user.id + " " + getFullName(user) + " " + user.email}
            </option>
        ));

        return (
            <section className="user-selector">
                <p>{this.props.description}</p>
                <select
                    value={this.props.selectedUserId}
                    onChange={this.onChange}>
                    {userOptions}
                </select>
            </section>
        );
    }
}
