"use strict";
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API from "../../api.js";
import { getFullName } from "../../common/userUtils.js";

export default class AccountInfo extends React.Component {
    state = {
        account: {},
        users: [],
    };

    componentDidMount() {
        const accountId = this.props.accountId;
        API.get("api/accounts/account/" + accountId)
            .then((res) => {
                const account = res.data || {};
                this.setState({ account: account });
            })
            .catch((err) => {
                console.log("Error searching account " + err);
            });

        API.get("api/users/account/" + accountId)
            .then((res) => {
                const users = res.data || [];
                this.setState({ users: users });
            })
            .catch((err) => {
                console.log("Error searching account users" + err);
            });
    }

    render() {
        const account = this.state.account;
        const users = this.state.users.map((user, index) => (
            <div key={index}>
                <div>Id: {user.id}</div>
                <div>{getFullName(user)}</div>
                <div>{user.email}</div>
                <div>{user.isGuardian ? "Guardian" : "Student"}</div>
            </div>
        ));
        const addNewUserUrl = "/account/" + account.id + "/user/add";

        return (
            <section>
                <h2>Account No. {account.id}</h2>
                <div>Primary email: {account.primaryEmail}</div>
                <div>Created {moment(account.createdAt).format("l")}</div>

                <h2>{users.length} Users found in Account</h2>
                {users}
                <div>
                    <Link to={addNewUserUrl}>
                        <button>Add New User to Account</button>
                    </Link>
                </div>

                <div>
                    <button>Delete Account</button>
                    <p>
                        Warning: Deleting this account will delete all user
                        information associated with account!
                    </p>
                </div>
            </section>
        );
    }
}
