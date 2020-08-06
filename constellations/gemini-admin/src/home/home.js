"use strict";
require("./home.sass");
import React from "react";

const sectionDisplayNames = {
    class: "Unpublished Classes",
    registration: "New Registrations",
    user: "New Users",
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
        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>

                <div className="tabs">
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "class"}
                        section={"class"}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "registration"}
                        section={"registration"}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "user"}
                        section={"user"}
                    />
                    <TabButton
                        onChangeTab={this.changeSection}
                        highlight={this.state.currentSection == "unpaid"}
                        section={"unpaid"}
                    />
                </div>

                <div className="showSection">
                    <HomeTabSection section={this.state.currentSection} />
                </div>
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
