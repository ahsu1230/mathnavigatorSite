"use strict";
require("./header.sass");
import React from "react";
import { Link } from "react-router-dom";

const LinkMap = {
    site: [
        {
            title: "Programs",
            url: "/programs",
        },
        {
            title: "Classes",
            url: "/classes",
        },
        {
            title: "Sessions",
            url: "/sessions",
        },
        {
            title: "Announcements",
            url: "/announcements",
        },
        {
            title: "Achievements",
            url: "/achievements",
        },
        {
            title: "Locations",
            url: "/locations",
        },
        {
            title: "Semesters",
            url: "/semesters",
        },
        {
            title: "Ask For Help",
            url: "/afh",
        },
    ],
    accounts: [
        {
            title: "Search Accounts",
            url: "/accounts",
        },
        {
            title: "Search Users",
            url: "/users",
        },
        {
            title: "Users By Class",
            url: "/users/classes",
        },
        {
            title: "Users By Afh",
            url: "/users/afhs",
        },
        {
            title: "Create New Account",
            url: "/account/create",
        },
    ],
    emails: [
        {
            title: "Payment Reminders",
            url: "/emailPayments",
        },
        {
            title: "Program Announcements",
            url: "/emailPrograms",
        },
    ],
};

export class HeaderSection extends React.Component {
    render() {
        return (
            <div id="view-header">
                <h2>
                    <Link to="/">Math Navigator Admin</Link>
                </h2>
                <div id="header-left">
                    <HeaderDropdown id={"site"} title={"Site"} />
                    <HeaderDropdown id={"accounts"} title={"Accounts"} />
                    <HeaderDropdown id={"emails"} title={"Emails"} />
                    <HeaderLink title={"Help"} url={"/help"} />
                </div>
            </div>
        );
    }
}

class HeaderLink extends React.Component {
    render() {
        const title = this.props.title;
        const url = this.props.url;
        return (
            <div className="header-link header-section">
                <Link to={url}>{title}</Link>
            </div>
        );
    }
}

class HeaderDropdown extends React.Component {
    state = {
        hover: false,
    };

    onHoverHeader = () => {
        this.setState({ hover: true });
    };

    onExitHeader = () => {
        this.setState({ hover: false });
    };

    render() {
        const links = LinkMap[this.props.id].map((link, index) => {
            return <HeaderDropdownRow key={index} link={link} />;
        });

        return (
            <div
                className="header-dropdown header-section"
                onMouseOver={this.onHoverHeader}
                onMouseOut={this.onExitHeader}>
                <div className={this.state.hover ? "title orange" : "title"}>
                    {this.props.title}
                </div>
                <ul
                    className={
                        this.state.hover ? "dropdown expand" : "dropdown"
                    }>
                    {links}
                </ul>
            </div>
        );
    }
}

class HeaderDropdownRow extends React.Component {
    render() {
        const link = this.props.link;
        return (
            <li>
                <Link to={link.url}>{link.title}</Link>
            </li>
        );
    }
}
