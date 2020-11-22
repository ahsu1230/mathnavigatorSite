"use strict";
require("./userMove.sass");
import React from "react";
import { assign, isEmpty } from "lodash";
import API from "../../api.js";
import AccountSearcher from "../../common/accountUserSearcher/accountSearcher.js";
import { AccountRowCard } from "../../common/rowCards/accountRowCard.js";
import { UserRowCard } from "../../common/rowCards/userRowCard.js";
import Checkmark from "../../../assets/checkmark_green.svg";

export class UserMovePage extends React.Component {
    state = {
        user: {},
        account: {}, // current account for user
        newAccount: {}, // proposed new account to move user to
    };

    componentDidMount() {
        const userId = this.props.userId;
        if (userId) {
            API.get("api/users/user/" + userId).then((res) => {
                const user = res.data;
                this.setState({
                    user: user,
                });

                const accountId = user.accountId;
                API.get("api/accounts/account/" + accountId).then((res) => {
                    const account = res.data;
                    this.setState({
                        account: account,
                    });
                });
            });
        }
    }

    onChooseNewAccount = (account) => {
        this.setState({
            newAccount: account,
        });
    };

    onConfirmSave = () => {
        const newAccount = this.state.newAccount;
        const user = assign(this.state.user, { accountId: newAccount.id });
        API.post("api/users/user/" + user.id, user)
            .then((res) => {
                window.alert("Successfully moved user to other account!");
                window.location.hash =
                    "/account/" +
                    newAccount.id +
                    "?view=edit-users&userId=" +
                    user.id;
            })
            .catch((err) => window.alert("Error moving user: " + err));
    };

    render() {
        const user = this.state.user;
        const account = this.state.account;
        const newAccount = this.state.newAccount;
        return (
            <main id="view-user-move">
                <article>
                    <h1>Move User to another Account</h1>
                    <p>
                        Use this page to move a user from one account to
                        another.
                    </p>
                    <section>
                        <div className="step">
                            <h3>Step 1) Selected User and Account</h3>
                            {!isEmpty(user) && !isEmpty(account) && (
                                <img src={Checkmark} />
                            )}
                        </div>
                        <UserRowCard user={user} />
                        <AccountRowCard account={account} />
                    </section>

                    <section>
                        <div className="step">
                            <h3>Step 2) Select another Account</h3>
                            {!isEmpty(newAccount) &&
                                newAccount.id != account.id && (
                                    <img src={Checkmark} />
                                )}
                        </div>
                        {!isEmpty(newAccount) && newAccount.id == account.id && (
                            <p className="error">
                                Both selected accounts are the same!
                                <br />
                                Please choose a different account!
                            </p>
                        )}
                        <AccountSearcher
                            hideUsers={true}
                            onFoundAccount={this.onChooseNewAccount}
                        />
                    </section>

                    {!isEmpty(user) &&
                        !isEmpty(account) &&
                        !isEmpty(newAccount) &&
                        newAccount.id != account.id && (
                            <div>
                                <h3>
                                    Are you sure you want to move this user?
                                </h3>
                                <p>
                                    Click on the "Save" button below to confirm.
                                </p>
                                <button
                                    className="confirm"
                                    onClick={this.onConfirmSave}>
                                    Confirm move
                                </button>
                            </div>
                        )}
                </article>
            </main>
        );
    }
}
