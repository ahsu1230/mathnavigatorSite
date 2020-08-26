"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";
import { HomeTabSectionClasses } from "./homeClasses.js";
import { HomeTabSectionUsers } from "./homeUsers.js";
import { HomeTabSectionRegistrations } from "./homeRegistrations.js";
import { HomeTabSectionAccounts } from "./homeAccounts.js";

const sectionDisplayNames = {
    classes: "Unpublished Classes",
    users: "New Users",
    registrations: "New Registrations",
    unpaid: "Unpaid Accounts",
};
const TAB_CLASSES = "classes";
const TAB_USERS = "users";
const TAB_REGISTRATIONS = "registrations";
const TAB_UNPAID = "unpaid";

export class HomePage extends React.Component {
    state = {
        selectedTab: TAB_CLASSES,

        unpublishedClasses: [],
        newUsers: [],
        newUserClasses: [],
        newUserAfh: [],
        unpaidAccounts: [],
    };

    changeSection = (sectionName) => {
        this.setState({
            selectedTab: sectionName,
        });
    };

    componentDidMount = () => {
        API.get("api/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                unpublishedClasses: unpublishedList.classes,
            });
        });

        API.get("api/users/new").then((res) => {
            const users = res.data;
            this.setState({
                newUsers: users,
            });
        });

        API.get("api/user-classes/new").then((res) => {
            const userClass = res.data;
            this.setState({
                newUserClasses: userClass,
            });
        });

        API.get("api/userafhs/new").then((res) => {
            const userAfh = res.data;
            this.setState({
                newUserAfh: userAfh,
            });
        });

        API.get("api/accounts/unpaid").then((res) => {
            const accounts = res.data;
            this.setState({
                unpaidAccounts: accounts,
            });
        });
    };

    render() {
        var sectionComponent = <div></div>;

        switch (this.state.selectedTab) {
            case TAB_CLASSES:
                sectionComponent = <HomeTabSectionClasses />;
                break;
            case TAB_USERS:
                sectionComponent = <HomeTabSectionUsers />;
                break;
            case TAB_REGISTRATIONS:
                sectionComponent = <HomeTabSectionRegistrations />;
                break;
            default:
                sectionComponent = <HomeTabSectionAccounts />;
        }

        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>

                <div className="tabs">
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.selectedTab == TAB_CLASSES}
                        section={TAB_CLASSES}
                        buttonNum={this.state.unpublishedClasses.length}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.selectedTab == TAB_USERS}
                        section={TAB_USERS}
                        buttonNum={this.state.newUsers.length}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.selectedTab == TAB_REGISTRATIONS}
                        section={TAB_REGISTRATIONS}
                        buttonNum={
                            this.state.newUserClasses.length +
                            this.state.newUserAfh.length
                        }
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.selectedTab == TAB_UNPAID}
                        section={TAB_UNPAID}
                        buttonNum={this.state.unpaidAccounts.length}
                    />
                </div>

                <div className="showSection">{sectionComponent}</div>
            </div>
        );
    }
}

class TabButton extends React.Component {
    render() {
        const highlight = this.props.highlight;
        const section = this.props.section;
        const displayName = sectionDisplayNames[section];
        const numNotif = this.props.buttonNum;

        const isZero = numNotif == 0 ? "notif zero" : "notif";

        let displayNotif = <div className={isZero}>{numNotif}</div>;

        return (
            <button
                className={highlight ? "activeTab" : ""}
                onClick={() => this.props.onChangeTab(section)}>
                {displayName} {displayNotif}
            </button>
        );
    }
}

export class EmptyMessage extends React.Component {
    render() {
        const emptySection = this.props.section;
        const length = this.props.length;
        var publishMessage = <div></div>;

        if (length == 0) {
            switch (emptySection) {
                case TAB_CLASSES:
                    publishMessage = (
                        <p>All classes have been successfully published!</p>
                    );
                    break;
                case TAB_USERS:
                    publishMessage = <p>There are no new users!</p>;
                    break;
                case TAB_REGISTRATIONS:
                    publishMessage = <p>There are no registrations!</p>;
                    break;
                default:
                    publishMessage = <p>All accounts have paid!</p>;
            }
        }

        return publishMessage;
    }
}
