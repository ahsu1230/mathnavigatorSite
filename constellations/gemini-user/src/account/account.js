"use strict";
require("./account.sass");
import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import API from "../utils/api.js";
import { union } from "lodash";
import moment from "moment";

const chargeDisplayNames = {
    charge: "Charge",
    refund: "Refund",
    pay_check: "Paid (Check)",
    pay_cash: "Paid (Cash)",
    pay_paypal: "Paid (Paypal)",
};
const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Computer Programming",
};
const seasonOrder = ["spring", "summer", "fall", "winter"];

function renderMultiline(lines) {
    return lines.map((line, index) => {
        return (
            <span key={index}>
                {line}
                <br />
            </span>
        );
    });
}

function formatCurrency(amount) {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(amount);
}
function fetchError(err) {
    alert("Could not fetch data: " + err);
}

export class AccountPage extends React.Component {
    state = {
        id: 1,
        primaryEmail: "",
        password: "",

        primaryEmail: "",
        users: [],
        transactions: [],
        userClasses: [],

        upcomingAFHs: [],
        locationsById: {},

        selectedTab: "settings",

        viewAllEnrolledClasses: false,
    };

    componentDidMount = () => {
        const id = this.state.id; // Replace with signin later

        if (id) {
            API.get("api/accounts/account/" + id)
                .then((res) => this.fetchData(res))
                .catch((err) => fetchError(err));
        }
    };

    fetchData = (res) => {
        const id = res.data.id;
        this.setState({
            id: id,
            primaryEmail: res.data.primaryEmail,
            password: res.data.password,
        });
        Promise.all([
            API.get("api/classes/all"),
            API.get("api/programs/all"),
            API.get("api/semesters/all"),

            API.get("api/askforhelp/all"),

            API.get("api/users/account/" + id),

            API.get("api/locations/all"),
        ])
            .then((res) => {
                const allClasses = res[0].data;
                const allPrograms = res[1].data;
                const allSemesters = res[2].data;

                const allAFHs = res[3].data;

                const users = res[4].data;
                this.setState({ users: users });

                const allLocations = res[5].data;
                const locationsById = {};
                allLocations.map((loc, index) => {
                    locationsById[loc.locationId] = loc;
                });
                this.setState({ locationsById: locationsById });

                let userClasses = [];
                let upcomingAFHs = [];
                users.map((user, index) => {
                    API.get("api/user-classes/user/" + user.id).then((res) => {
                        let classes = res.data.map((c, index) => {
                            let matchedClass = allClasses.find(
                                (element) => element.classId == c.classId
                            );
                            let matchedProgram = allPrograms.find(
                                (element) =>
                                    element.programId == matchedClass.programId
                            );
                            let matchedSemester = allSemesters.find(
                                (element) =>
                                    element.semesterId ==
                                    matchedClass.semesterId
                            );
                            return {
                                program: matchedProgram,
                                semester: matchedSemester,
                                enrollDate: c.updatedAt,
                            };
                        });

                        userClasses.push({
                            id: user.id,
                            name: user.firstName + " " + user.lastName,
                            classes: classes,
                        });
                        this.setState({ userClasses: userClasses });
                    });
                    API.get("api/userafhs/users/" + user.id).then((res) => {
                        let afhs = res.data.map((afh, index) => {
                            let matchedAFH = allAFHs.find(
                                (element) => element.id == afh.afhId
                            );
                            upcomingAFHs.push(matchedAFH);
                        });
                        upcomingAFHs = union(upcomingAFHs);
                        this.setState({ upcomingAFHs: upcomingAFHs });
                    });
                });
            })
            .catch((err) => alert("Could not fetch data: " + err));
        API.get("api/transactions/account/" + id).then((res) =>
            this.setState({ transactions: res.data })
        );
    };

    toggleViewAllClasses = () => {
        this.setState({
            viewAllEnrolledClasses: !this.state.viewAllEnrolledClasses,
        });
    };

    onTabSelect = (tab) => {
        this.setState({ selectedTab: tab.toLowerCase() });
    };

    onPasswordChange = (password) => {
        this.setState({ password: password });
    };

    render = () => {
        const tabButtons = ["Settings", "Registrations", "Payment"].map(
            (item, index) => {
                return (
                    <div
                        className={
                            this.state.selectedTab == item.toLowerCase()
                                ? "selected"
                                : ""
                        }
                        onClick={(e) => this.onTabSelect(item)}
                        key={index}>
                        {item}
                    </div>
                );
            }
        );

        let tabContent;
        switch (this.state.selectedTab) {
            case "settings":
                tabContent = (
                    <SettingsTab
                        users={this.state.users}
                        primaryEmail={this.state.primaryEmail}
                        accountId={this.state.id}
                        password={this.state.password}
                        passwordChangeCallback={this.onPasswordChange}
                    />
                );
                break;
            case "registrations":
                tabContent = this.state.viewAllEnrolledClasses ? (
                    <RegistrationsTabAllClasses
                        userClasses={this.state.userClasses}
                        toggleTabCallback={this.toggleViewAllClasses}
                    />
                ) : (
                    <RegistrationsTabMain
                        userClasses={this.state.userClasses}
                        upcomingAFHs={this.state.upcomingAFHs}
                        locationsById={this.state.locationsById}
                        toggleTabCallback={this.toggleViewAllClasses}
                    />
                );
                break;
            case "payment":
                tabContent = (
                    <PaymentTab transactions={this.state.transactions} />
                );
                break;
        }

        return (
            <div id="view-account">
                <h1>Your Account</h1>

                <div id="tab-container">{tabButtons}</div>

                {tabContent}
            </div>
        );
    };
}

class SettingsTab extends React.Component {
    render = () => {
        const currentYear = new Date().getFullYear();

        const usersList = this.props.users.map((user, index) => {
            let contactInfo = [user.email];
            if (user.phone) {
                contactInfo.push(user.phone);
            }
            contactInfo = renderMultiline(contactInfo);

            let otherInfo = [
                (user.isGuardian ? "Guardian" : "Student") +
                    (user.email == this.props.primaryEmail
                        ? " (Primary Contact)"
                        : ""),
            ];
            if (user.school) {
                otherInfo.push(user.school);
            }
            if (user.graduationYear) {
                otherInfo.push(
                    12 -
                        (user.graduationYear - currentYear) +
                        "th Grade, " +
                        "Graduation Year: " +
                        user.graduationYear
                );
            }
            otherInfo = renderMultiline(otherInfo);

            return (
                <ul key={index} className="users-table">
                    <li className="li-med">
                        {user.firstName + " " + user.lastName}
                    </li>
                    <li className="li-med">{contactInfo}</li>
                    <li className="li-large">{otherInfo}</li>
                </ul>
            );
        });

        return (
            <div className="tab-content" id="settings-tab">
                <div>
                    <h2>Your Account Information</h2>
                    <p className="vertical-mobile">
                        <span>Primary email: {this.props.primaryEmail}</span>
                        <a className="edit orange"> Change primary contact</a>
                    </p>
                    <PasswordChange
                        accountId={this.props.accountId}
                        primaryEmail={this.props.primaryEmail}
                        oldPassword={this.props.password}
                        passwordChangeCallback={
                            this.props.passwordChangeCallback
                        }
                    />
                </div>

                <div>
                    <h2>User Information</h2>
                    <ul className="header hide-mobile">
                        <li className="li-med">Name</li>
                        <li className="li-med">Contact</li>
                        <li className="li-large">Other Information</li>
                    </ul>
                    {usersList}
                    <p>
                        <a className="orange">Edit users for this account</a>
                    </p>
                </div>

                <div>
                    <p>
                        You may opt to delete your Math Navigator account.
                        <br />
                        However, doing so will delete all your data with Math
                        Navigator, including all user and class information.
                    </p>
                    <a className="red">Request to Delete Account...</a>
                </div>
            </div>
        );
    };
}

class PasswordChange extends React.Component {
    state = {
        tabOpen: false,
        message: "",

        oldPassword: "",
        newPassword: "",
        confirmPassword: "",
    };

    onClickChange = () => {
        this.setState({
            tabOpen: !this.state.tabOpen,
            message: "",
        });
    };

    onClickSave = () => {
        if (
            this.state.oldPassword == this.props.oldPassword &&
            this.state.newPassword == this.state.confirmPassword
        ) {
            let account = {
                primaryEmail: this.props.primaryEmail,
                password: this.state.newPassword,
            };
            API.post(
                "api/accounts/account/" + this.props.accountId,
                account
            ).then((res) => {
                this.props.passwordChangeCallback(this.state.newPassword);
                this.setState({
                    tabOpen: false,
                    message: "New password saved!",
                    oldPassword: "",
                    newPassword: "",
                    confirmPassword: "",
                });
            });
        } else if (this.state.oldPassword == this.props.oldPassword) {
            this.setState({
                message: "New password does not match confirmation",
            });
        } else {
            this.setState({ message: "Old password is incorrect" });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    render = () => {
        const changePasswordDialog = this.state.tabOpen ? (
            <div>
                <ul className="vertical-centered no-border">
                    <li className="li-med">Old password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "oldPassword")
                            }
                            value={this.state.oldPassword}
                        />
                    </li>
                </ul>
                <ul className="vertical-centered no-border">
                    <li className="li-med">New password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "newPassword")
                            }
                            value={this.state.newPassword}
                        />
                    </li>
                </ul>
                <ul className="vertical-centered no-border">
                    <li className="li-med">Confirm new password</li>
                    <li className="li-large">
                        <input
                            type="password"
                            onChange={(e) =>
                                this.handleChange(e, "confirmPassword")
                            }
                            value={this.state.confirmPassword}
                        />
                    </li>
                </ul>
                <div className="password-buttons space-between">
                    <button className="btn-cancel" onClick={this.onClickChange}>
                        Cancel
                    </button>
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                </div>
            </div>
        ) : (
            ""
        );

        const message = <span> {this.state.message}</span>;

        return (
            <div id="password-change">
                <p className="vertical-mobile">
                    <a className="orange" onClick={this.onClickChange}>
                        Change password...
                    </a>
                    {message}
                </p>
                {changePasswordDialog}
            </div>
        );
    };
}

class RegistrationsTabMain extends React.Component {
    renderClassList = (classes) => {
        if (!classes.length) {
            return <span>(No classes registered)</span>;
        }
        return classes.map((c, index) => {
            return (
                <div key={index} className="classList-item space-between">
                    <span>
                        {c.program.name + " (" + c.semester.title + ")"}
                    </span>
                    <span>Enrolled on: {moment(c.enrollDate).format("l")}</span>
                </div>
            );
        });
    };

    render = () => {
        const classRegistrationList = this.props.userClasses.map(
            (user, index) => {
                return (
                    <ul key={index}>
                        <li className="li-med">{user.name}</li>
                        <li className="li-large">
                            {this.renderClassList(user.classes)}
                        </li>
                    </ul>
                );
            }
        );

        const afhList = this.props.upcomingAFHs.map((afh, index) => {
            let titleInfo = renderMultiline([
                afh.title,
                subjectDisplayNames[afh.subject],
            ]);
            let dateInfo = renderMultiline([
                moment(afh.date).format("MMMM Do, YYYY"),
                afh.timeString,
            ]);

            let loc = this.props.locationsById[afh.locationId];
            let locInfo = renderMultiline([
                loc.street,
                loc.city + ", " + loc.state,
                loc.room,
            ]);

            return (
                <ul key={index} className="afh-table three-columns">
                    <li className="li-med">{titleInfo}</li>
                    <li className="li-med">{dateInfo}</li>
                    <li className="li-large">{locInfo}</li>
                </ul>
            );
        });

        return (
            <div className="tab-content" id="reg-tab-main">
                <div>
                    <h2>Currently Enrolled Classes</h2>
                    {classRegistrationList}
                    <p>
                        <a
                            className="orange"
                            onClick={this.props.toggleTabCallback}>
                            View all enrolled classes
                        </a>
                    </p>
                </div>
                <div>
                    <h2>Upcoming Ask For Help Sessions</h2>
                    <ul className="three-columns header hide-mobile">
                        <li className="li-med">Title</li>
                        <li className="li-med">Date</li>
                        <li className="li-large">Location</li>
                    </ul>
                    {afhList}
                </div>
            </div>
        );
    };
}

class RegistrationsTabAllClasses extends React.Component {
    render = () => {
        const allClasses = [];
        this.props.userClasses.map((user, index) => {
            user.classes.map((c, index) => {
                allClasses.push({
                    user: user.name,
                    classInfo: c,
                });
            });
        });

        allClasses.sort((a, b) => {
            let semA = a.classInfo.semester.semesterId.split("_");
            let semB = b.classInfo.semester.semesterId.split("_");

            if (semA[0] < semB[0]) {
                return 1;
            } else if (semA[0] > semB[0]) {
                return -1;
            } else {
                return seasonOrder.indexOf(semA[1]) <
                    seasonOrder.indexOf(semB[1])
                    ? 1
                    : -1;
            }
        });

        const classRegistrationList = allClasses.map((c, index) => {
            return (
                <ul key={index}>
                    <li className="li-med">{c.user}</li>
                    <li className="li-large">
                        {c.classInfo.program.name +
                            " (" +
                            c.classInfo.semester.title +
                            ")"}
                    </li>
                    <li>Enrolled on: {moment(c.enrollDate).format("l")}</li>
                </ul>
            );
        });

        return (
            <div className="tab-content" id="reg-tab-all">
                <div>
                    <div className="vertical-centered space-between">
                        <h2>All Enrolled Classes</h2>
                        <a
                            className="orange"
                            onClick={this.props.toggleTabCallback}>
                            View current enrollments
                        </a>
                    </div>
                    <div>{classRegistrationList}</div>
                </div>
            </div>
        );
    };
}

class PaymentTab extends React.Component {
    render = () => {
        let balance = 0;
        const transactionsList = this.props.transactions.map(
            (transaction, index) => {
                balance += parseInt(transaction.amount);
                return (
                    <ul key={index}>
                        <li className="li-med">
                            {moment(0).format("l") /*Fake data*/}
                        </li>
                        <li className="li-med">
                            {chargeDisplayNames[transaction.paymentType]}
                        </li>
                        <li className="li-med">
                            {formatCurrency(transaction.amount)}
                        </li>
                        <li className="li-large">{formatCurrency(balance)}</li>
                    </ul>
                );
            }
        );
        transactionsList.reverse();

        const formattedBalance = formatCurrency(balance);

        return (
            <div className="tab-content" id="payment-tab">
                <div>
                    <h2>Account Balance: {formattedBalance}</h2>
                    <p>
                        You currently owe {formattedBalance}. Please pay through
                        any of the following methods:
                    </p>

                    <span>
                        - <Link to="">Cash</Link>
                        <br />
                    </span>
                    <span>
                        - <Link to="">Check</Link> (written to Math Navigator)
                        <br />
                    </span>
                    <span>
                        - <Link to="">Paypal</Link>
                    </span>

                    <p>
                        For questions about your account, please contact us at{" "}
                        <a
                            href="mailto:andymathnavigator@gmail.com"
                            className="orange">
                            andymathnavigator@gmail.com
                        </a>
                    </p>
                </div>
                <div>
                    <h2>Your Payment History</h2>
                    <ul className="header">
                        <li className="li-med">Date</li>
                        <li className="li-med">Transaction</li>
                        <li className="li-med">Amount</li>
                        <li className="li-large">Balance</li>
                    </ul>
                    {transactionsList}
                </div>
            </div>
        );
    };
}
