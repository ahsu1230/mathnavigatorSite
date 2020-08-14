"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";
import { HomeTabSectionClasses } from "./homeClasses.js";
import { HomeTabSectionUsers } from "./homeUsers.js";
import { HomeTabSectionRegistrations } from "./homeRegistrations.js";
import { HomeTabSectionAccounts } from "./homeAccounts.js";

const sectionDisplayNames = {
    class: "Unpublished Classes",
    user: "New Users",
    registration: "New Registrations",
    unpaid: "Unpaid Accounts",
};

export class HomePage extends React.Component {
    state = {
        currentSection: "class",

        unpublishedClasses: [],
        newUsers: [],
        newUserClasses: [],
        newUserAfh: [],
        unpaidAccounts: [],
    };

    changeSection = (sectionName) => {
        this.setState({
            currentSection: sectionName,
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

        //pending registration for classes
        API.get("api/user-classes/new").then((res) => {
            const userClass = res.data;
            this.setState({
                newUserClasses: userClass,
            });
        });

        //afh registration
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

        if (this.state.currentSection == "class") {
            sectionComponent = <HomeTabSectionClasses />;
        } else if (this.state.currentSection == "user") {
            sectionComponent = <HomeTabSectionUsers />;
        } else if (this.state.currentSection == "registration") {
            sectionComponent = <HomeTabSectionRegistrations />;
        } else {
            sectionComponent = <HomeTabSectionAccounts />;
        } //unpaid account

        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>

                <div className="tabs">
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "class"}
                        section={"class"}
                        buttonNum={this.state.unpublishedClasses.length}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "user"}
                        section={"user"}
                        buttonNum={this.state.newUsers.length}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "registration"}
                        section={"registration"}
                        buttonNum={
                            this.state.newUserClasses.length +
                            this.state.newUserAfh.length
                        }
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "unpaid"}
                        section={"unpaid"}
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
        let highlight = this.props.highlight;
        let section = this.props.section;
        let displayName = sectionDisplayNames[section];
        let numNotif = this.props.buttonNum;

        var isZero = "";

        if (numNotif == 0) {
            isZero = "notif zero";
        } else {
            isZero = "notif";
        }
        console.log("isZero state is " + isZero);

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
