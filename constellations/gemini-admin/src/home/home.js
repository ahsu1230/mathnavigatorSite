"use strict";
require("./home.sass");
import React from "react";
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
    };

    changeSection = (sectionName) => {
        this.setState({
            currentSection: sectionName,
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
        } else if (this.state.currentSection == "unpaid") {
            sectionComponent = <HomeTabSectionAccounts />;
        }

        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>

                <div className="tabs">
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "class"}
                        section={"class"}
                        //unpubClasses.length in homeClasses
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "user"}
                        section={"user"}
                        //newUsers.length in homeUsers
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "registration"}
                        section={"registration"}
                        //pendingReg.length + afhReg.length in homeRegistrations
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "unpaid"}
                        section={"unpaid"}
                        //unpaidAcc.length in homeAccounts
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

        return (
            <button
                className={highlight ? "active" : ""}
                onClick={() => this.props.onChangeTab(section)}>
                {displayName}
            </button>
        );
    }
}
