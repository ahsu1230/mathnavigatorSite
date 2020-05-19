"use strict";
require("./menuSlim.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";
import { NavLinks, isPathAt } from "../constants.js";
const classnames = require("classnames");
const srcArrowDown = require("../../assets/arrow_down_black.svg");
const srcClose = require("../../assets/close_black.svg");

export class MenuSlim extends React.Component {
    constructor() {
        super();
        this.state = {
            show: false,
        };
        this.toggleMenu = this.toggleMenu.bind(this);
    }

    toggleMenu() {
        this.setState({
            show: !this.state.show,
        });
    }

    render() {
        const location = this.props.location;
        const show = this.state.show;
        return (
            <div>
                <div className="header-menu-slim" key="header-button">
                    <button
                        className="header-menu-btn"
                        onClick={this.toggleMenu}>
                        Menu
                    </button>
                </div>
                <OverlayMenu
                    key="overlay-menu"
                    show={show}
                    closeMenu={this.toggleMenu}
                    location={location}
                />
            </div>
        );
    }
}

export class OverlayMenu extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            links: NavLinks,
        };
    }

    render() {
        const closeMenu = this.props.closeMenu;
        const containerClasses = classnames("header-slim-overlay-container", {
            show: this.props.show,
        });
        const links = generateNavLinks(this.state.links, closeMenu);
        return (
            <div className={containerClasses}>
                <div className="header-slim-overlay" />
                <div className="header-slim-content">
                    <h1>Menu</h1>
                    <button onClick={closeMenu}>
                        <img src={srcClose} />
                    </button>
                    {links}
                </div>
            </div>
        );
    }
}

function generateNavLinks(links, closeMenu) {
    var items = links.map(function (link, index) {
        if (link.subLinks && link.subLinks.length > 0) {
            return (
                <SubMenu
                    key={index}
                    currentLink={link}
                    links={link.subLinks}
                    closeMenu={closeMenu}
                />
            );
        } else {
            return <LinkRow key={index} link={link} onClick={closeMenu} />;
        }
    });
    return items;
}

class LinkRow extends React.Component {
    render() {
        const link = this.props.link;
        const onClick = this.props.onClick;
        const linkClass = classnames({
            active: isPathAt(window.location.hash, link.url),
        });
        return (
            <div className="link-row">
                <Link className={linkClass} to={link.url} onClick={onClick}>
                    {link.name}
                </Link>
            </div>
        );
    }
}

class SubMenu extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            show: false,
        };
        this.toggleShow = this.toggleShow.bind(this);
    }

    toggleShow() {
        this.setState({
            show: !this.state.show,
        });
    }

    render() {
        const currentLink = this.props.currentLink;
        const links = this.props.links;
        const closeMenu = this.props.closeMenu;
        const subLinks = links.map(function (link, index) {
            return (
                <li key={link.id}>
                    <LinkRow link={link} onClick={closeMenu} />
                </li>
            );
        });
        const submenuClasses = classnames("submenu", {
            show: this.state.show,
        });

        return (
            <div className={submenuClasses}>
                <div className="submenu-head" onClick={this.toggleShow}>
                    {currentLink.name}
                    <img src={srcArrowDown} />
                </div>
                <ul>{subLinks}</ul>
            </div>
        );
    }
}
