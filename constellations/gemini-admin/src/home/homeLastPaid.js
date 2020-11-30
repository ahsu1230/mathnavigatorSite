"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import { keyBy, values } from "lodash";
import API from "../api.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import { formatCurrency } from "../common/displayUtils.js";
import {
    getAccountBalance,
    sortTransactionsLatestFirst,
} from "../common/transactionUtils.js";
import { getFullName } from "../common/userUtils.js";
import SrcArrowDown from "../../assets/triangle_down_black.svg";

export class HomeTabSectionLastPaid extends React.Component {
    state = {
        allClasses: [],
        classMap: {},
        selectedClassId: "",
        userClasses: [],
        userClassMap: {},
        userMap: {},
        users: [],
    };

    componentDidMount() {
        API.get("api/classes/all").then((res) => {
            const classes = res.data;
            this.setState({
                allClasses: classes,
                classMap: keyBy(classes, "classId"),
            });
        });
    }

    onSwitchClass = (classId) => {
        this.setState({
            selectedClassId: classId,
        });
        API.get("api/user-classes/class/" + classId).then((resUserClasses) => {
            const userClasses = resUserClasses.data || [];
            this.setState({
                userClasses: userClasses,
                userClassMap: keyBy(userClasses, "userId"),
            });
            if (userClasses.length > 0) {
                const userIds = userClasses.map(
                    (userClass) => userClass.userId
                );
                API.post("api/users/map", userIds).then((resUsers) => {
                    const userMap = resUsers.data;
                    const userList = values(userMap).map((user) => {
                        // TODO (aaron): fix bug on backend side!
                        user.middleName = user.middleName.String;
                        user.school = user.school.String;
                        user.graduationYear = user.graduationYear.Uint;
                        return user;
                    });
                    this.setState({
                        userMap: userMap,
                        users: userList,
                    });
                });
            } else {
                this.setState({
                    userMap: {},
                    users: [],
                });
            }
        });
    };

    render() {
        const classOptions = this.state.allClasses.map((classObj) => {
            return {
                displayName: classObj.classId,
                value: classObj.classId,
            };
        });
        const users = this.state.users.map((user, index) => (
            <UserClassRow
                key={index}
                user={user}
                userClass={this.state.userClassMap[user.id]}
            />
        ));
        return (
            <div className="section-details">
                <InputSelect
                    label={"Select a class"}
                    description={
                        "Select a class to view students who are in that class."
                    }
                    value={this.state.selectedClassId}
                    onChangeCallback={(e) => this.onSwitchClass(e.target.value)}
                    options={classOptions}
                    hasNoDefault={true}
                    required={true}
                />
                {users}
                {users.length == 0 && (
                    <p>
                        There are currently no users registered for this class.
                    </p>
                )}
            </div>
        );
    }
}

class UserClassRow extends React.Component {
    state = {
        transactions: [],
        expand: false,
    };

    componentDidMount() {
        const user = this.props.user;
        this.fetchAccountTransactions(user);
    }

    componentDidUpdate(prevProps) {
        const prevUser = prevProps.user || {};
        const currUser = this.props.user || {};
        if (prevUser.id !== currUser.id) {
            this.fetchAccountTransactions(currUser);
        }
    }

    fetchAccountTransactions = (user) => {
        API.get("api/transactions/account/" + user.accountId).then((res) => {
            const transactions = res.data;
            this.setState({
                transactions: sortTransactionsLatestFirst(transactions),
            });
        });
    };

    toggleExpand = () => {
        this.setState({
            expand: !this.state.expand,
        });
    };

    render() {
        const user = this.props.user || {};
        const userClass = this.props.userClass || {};
        const transactions = this.state.transactions;
        const viewAccountUrl =
            "/account/" + user.accountId + "?view=transactions";
        const expand = this.state.expand;
        const balance = getAccountBalance(transactions);
        return (
            <div className="last-paid-user-row">
                <h3>{getFullName(user)}</h3>
                <div className="info-container">
                    <div>
                        <div className="line">Email: {user.email}</div>
                        <div className="line">
                            Enrollment Status: {userClass.state}{" "}
                            {moment(userClass.updatedAt).format("l")}
                        </div>
                    </div>
                    <div>
                        <div className="line">{user.school}</div>
                        <div className="line">
                            Graduation Year: {user.graduationYear}
                        </div>
                    </div>
                    <div className="account-info">
                        <h4 className={balance < 0 ? "warning" : ""}>
                            Account Balance: {formatCurrency(balance)}
                        </h4>
                        <Link to={viewAccountUrl}>View Account Details</Link>
                        <button onClick={this.toggleExpand}>
                            <img
                                src={SrcArrowDown}
                                className={expand ? "flip" : ""}
                            />
                            {expand ? (
                                <span>Collapse</span>
                            ) : (
                                <span>Expand</span>
                            )}
                        </button>
                    </div>
                </div>
                {expand && <LastTransactions transactions={transactions} />}
            </div>
        );
    }
}

class LastTransactions extends React.Component {
    render() {
        const lastTransactions = (this.props.transactions || []).slice(0, 3);
        const lastFewTransactions = lastTransactions.map((trans, index) => (
            <div className="transact-container" key={index}>
                <div>{moment(trans.createdAt).format("l")}</div>
                <div>{trans.type}</div>
                <div>{formatCurrency(trans.amount)}</div>
            </div>
        ));
        return (
            <div>
                {lastTransactions.length == 0 && <p>No transactions yet.</p>}
                {lastTransactions.length > 0 && (
                    <div>
                        <h4>Last Few Transactions</h4>
                        {lastFewTransactions}
                        <p className="note">
                            To see more transactions for this account, click on
                            "View Account Details"
                        </p>
                    </div>
                )}
            </div>
        );
    }
}
