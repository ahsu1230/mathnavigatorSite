"use strict";
require("./header.styl");
import React from "react";
import ReactDOM from "react-dom";
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
            url: "/semester",
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
                    <HeaderLink title={"Accounts"} url={"/accounts"} />
                    <HeaderLink title={"Emails"} url={"/emails"} />
                    <span>Sign Out</span>
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
    render() {
        const title = this.props.title;
        const links = LinkMap[this.props.id].map((link, index) => {
            return <HeaderDropdownRow key={index} link={link} />;
        });

        return (
            <div className="header-dropdown header-section">
                <div className="title">{title}</div>
                <ul>{links}</ul>
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
