"use strict";
require("./accountInfo.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";
import API from "../../api.js";
import { getFullName } from "../../common/userUtils.js";
import { InputRadio } from "../../common/inputs/inputRadio.js";

export default class AccountInfo extends React.Component {
    state = {
        account: {},
        users: [],
        selectedPrimaryEmail: "",
    };

    componentDidMount() {
        const accountId = this.props.accountId;
        API.get("api/accounts/account/" + accountId)
            .then((res) => {
                const account = res.data || {};
                this.setState({
                    account: account,
                    selectedPrimaryEmail: account.primaryEmail,
                });
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

    onChangePrimaryContact = (e) => {
        const newPrimaryEmail = e.target.value;
        this.setState({
            selectedPrimaryEmail: newPrimaryEmail,
        });
        const account = this.state.account;
        account.primaryEmail = newPrimaryEmail;
        API.post("api/accounts/account/" + account.id, account)
            .then((res) => console.log("Primary contact updated!"))
            .catch((err) => console.log("Error updating primary contact"));
    };

    onDeleteAccount = (e) => {
        const accountId = this.state.account.id;
        API.delete("api/accounts/full/account/" + accountId)
            .then((res) => {
                window.alert("Account deleted!");
                window.location.hash = "accounts";
            })
            .catch((err) => window.alert("Could not delete account " + err));
    };

    render() {
        const account = this.state.account;
        const options = this.state.users.map((user) => {
            return {
                value: user.email,
                displayName: getFullName(user) + " (" + user.email + ")",
            };
        });
        const addNewUserUrl = "/account/" + account.id + "/user/add";

        return (
            <section className="account-tab account-info">
                <h3>Primary email: {account.primaryEmail}</h3>
                <h4>
                    Account Created: {moment(account.createdAt).format("l")}
                </h4>

                <InputRadio
                    value={this.state.selectedPrimaryEmail}
                    onChangeCallback={this.onChangePrimaryContact}
                    description={
                        options.length +
                        " Users found in Account. Select one user as the Primary Contact of this account."
                    }
                    options={options}
                />
                <div className="add-user">
                    <Link to={addNewUserUrl}>
                        <button className="add">Add New User to Account</button>
                    </Link>
                </div>

                <div className="delete-account">
                    <button onClick={this.onDeleteAccount}>
                        Delete Account
                    </button>
                    <p>
                        Warning: Deleting this account will delete all user
                        information associated with account! This includes all
                        emails, phone numbers, class registrations, ask-for-help
                        registrations and transactions.
                    </p>
                </div>
            </section>
        );
    }
}
