"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

const sectionDisplayNames = {
    class: "Unpublished Classes",
    registration: "New Registrations",
    user: "New Users",
    unpaid: "Unpaid Accounts",
};

export class HomePage extends React.Component {
    state = {
        classes: [],
        currentSection: "class",
    };

    changeSection = (sectionName) => {
        this.setState({
            currentSection: sectionName,
        });
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData = () => {
        API.get("api/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                classes: unpublishedList.classes,
            });
        });
    };

    render() {
        let unpublishedClasses = this.state.classes.map((row, index) => {
            return <li key={index}> {row.classId} </li>;
        });

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

                <div className="section">
                    <div className="container-class">
                        <h3 className="section-header">Unpublished Classes</h3>{" "}
                        <button id="publish">
                            <Link to={"/classes"}>
                                View All Classes to Publish
                            </Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="list-header">Class ID</div>
                        <ul>{unpublishedClasses}</ul>
                    </div>
                </div>

                <div className="section">
                    <h3 className="section-header">New Users</h3>
                </div>

                <div className="section">
                    <h3 className="section-header">New Registrations</h3>
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
