"use strict";
require("./account.sass");
import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import API from "../utils/api.js";
import { union } from "lodash";

export class AccountPage extends React.Component {
    state = {
        id: 1,

        primaryEmail: "",
        users: [],
        transactions: [],
        userClasses: [],

        upcomingAFHs: [],
        locationsById: {},

        selectedTab: "registrations",
    };

    componentDidMount = () => {
        const id = this.state.id;

        if (id) {
            API.get("api/accounts/account/" + id)
                .then((res) => this.fetchData(res))
                .catch((err) => this.fetchDataError(err));
        }
    };

    fetchData = (res) => {
        const id = res.data.id;
        this.setState({
            id: id,
            primaryEmail: res.data.primaryEmail,
            users: [],
            transactions: [],
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
                                programName: matchedProgram.name,
                                semester: matchedSemester.title,
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
            .catch((err) => console.log(err));

        API.get("api/transactions/account/" + id).then((res) =>
            this.setState({ transactions: res.data })
        );
    };

    formatCurrency = (amount) => {
        return new Intl.NumberFormat("en-US", {
            style: "currency",
            currency: "USD",
        }).format(amount);
    };

    renderMultiline = (lines) => {
        return lines.map((line, index) => {
            return (
                <span key={index}>
                    {line}
                    <br />
                </span>
            );
        });
    };

    renderSettings = () => {
        const currentYear = new Date().getFullYear();

        const usersList = this.state.users.map((user, index) => {
            let contactInfo = [user.email];
            if (user.phone) {
                contactInfo.push(user.phone);
            }
            contactInfo = this.renderMultiline(contactInfo);

            let otherInfo = [user.isGuardian ? "Guardian" : "Student"];
            if (user.email == this.state.primaryEmail) {
                otherInfo[0] = otherInfo[0] + " (Primary Contact)";
            }
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
            otherInfo = this.renderMultiline(otherInfo);

            return (
                <ul key={index}>
                    <li className="li-med">
                        {user.firstName + " " + user.lastName}
                    </li>
                    <li className="li-med">{contactInfo}</li>
                    <li className="li-large">{otherInfo}</li>
                </ul>
            );
        });
        return (
            <div className="tab-content">
                <h2>Your Account Information</h2>

                <div>
                    <p>
                        <span>Primary email: {this.state.primaryEmail}</span>
                        <Link to="" className="edit orange">
                            Change primary contact
                        </Link>
                    </p>
                    <p>
                        <Link to="" className="orange">
                            Change password...
                        </Link>
                    </p>
                </div>

                <div>
                    <h2>User Information</h2>
                    <ul>
                        <li className="li-med">Name</li>
                        <li className="li-med">Contact</li>
                        <li className="li-large">Other Information</li>
                    </ul>
                    {usersList}
                    <p>
                        <Link to="" className="orange">
                            Edit users for this account
                        </Link>
                    </p>
                </div>

                <div>
                    <p id="delete-message">
                        You may opt to delete your Math Navigator account.
                        <br />
                        However, doing so will delete all your data with Math
                        Navigator, including all user and class information.
                    </p>
                    <Link to="" className="red">
                        Request to Delete Account...
                    </Link>
                </div>
            </div>
        );
    };

    renderClassList = (classes) => {
        if (!classes.length) {
            return <p>(No classes registered)</p>;
        }
        return classes.map((c, index) => {
            return (
                <p key={index} className="classList-item">
                    {c.programName + " (" + c.semester + ")"}
                </p>
            );
        });
    };

    renderRegistrations = () => {
        const classRegistrationList = this.state.userClasses.map(
            (user, index) => {
                return (
                    <ul key={index} className="no-borders">
                        <li className="li-med">
                            <p>{user.name}</p>
                        </li>
                        <li className="li-large classes-list">
                            {this.renderClassList(user.classes)}
                        </li>
                    </ul>
                );
            }
        );

        const afhList = this.state.upcomingAFHs.map((afh, index) => {
            let titleInfo = this.renderMultiline([afh.title, afh.subject]);
            let dateInfo = this.renderMultiline([afh.date, afh.timeString]);

            let loc = this.state.locationsById[afh.locationId];
            let locInfo = this.renderMultiline([
                loc.street,
                loc.city + ", " + loc.state,
                loc.room,
            ]);

            return (
                <ul key={index} className="no-borders three-columns">
                    <li className="li-med">{titleInfo}</li>
                    <li className="li-med">{dateInfo}</li>
                    <li className="li-large">{locInfo}</li>
                </ul>
            );
        });

        return (
            <div className="tab-content">
                <div>
                    <h2>Currently Enrolled Classes</h2>
                    {classRegistrationList}
                    <Link to="" className="orange">
                        View all enrolled classes
                    </Link>
                </div>
                <div>
                    <h2>Upcoming Ask For Help Sessions</h2>
                    <ul className="no-borders three-columns header">
                        <li className="li-med">Title</li>
                        <li className="li-med">Date</li>
                        <li className="li-large">Location</li>
                    </ul>
                    {afhList}
                </div>
            </div>
        );
    };

    renderPayment = () => {
        let balance = 0;
        const transactionsList = this.state.transactions.map(
            (transaction, index) => {
                balance += parseInt(transaction.amount);
                return (
                    <ul key={index} className="no-borders">
                        <li className="li-med">{transaction.paymentType}</li>
                        <li className="li-med">
                            {this.formatCurrency(transaction.amount)}
                        </li>
                        <li className="li-large">
                            {this.formatCurrency(balance)}
                        </li>
                    </ul>
                );
            }
        );

        const formattedBalance = this.formatCurrency(balance);

        return (
            <div className="tab-content">
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
                    <ul className="no-borders header">
                        <li className="li-med">Transaction</li>
                        <li className="li-med">Amount</li>
                        <li className="li-large">Balance</li>
                    </ul>
                    {transactionsList}
                </div>
            </div>
        );
    };

    onTabSelect = (tab) => {
        this.setState({ selectedTab: tab.toLowerCase() });
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
                tabContent = this.renderSettings();
                break;
            case "registrations":
                tabContent = this.renderRegistrations();
                break;
            case "payment":
                tabContent = this.renderPayment();
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
