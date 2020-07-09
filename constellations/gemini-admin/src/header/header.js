"use strict";
require("./header.styl");
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
                    <HeaderLink title={"Help"} url={"/help"} />
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
        const title = this.props.title;
        const links = LinkMap[this.props.id].map((link, index) => {
            return <HeaderDropdownRow key={index} link={link} />;
        });
        const listStyle = {
            height: this.state.hover ? 32 * links.length + "px" : "0px",
        };

        return (
            <div
                className="header-dropdown header-section"
                onMouseOver={this.onHoverHeader}
                onMouseOut={this.onExitHeader}>
                <div className="title">{title}</div>
                <ul style={listStyle}>{links}</ul>
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
