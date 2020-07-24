"use strict";
require("./user.sass");
import React from "react";
import { Link } from "react-router-dom";
import { debounce } from "lodash";
import API from "../api.js";
import { getCurrentUserSearch, setCurrentUserSearch } from "../localStorage.js";
import DotsVertical from "../../assets/dots_vertical_gray.svg";

export class UserPage extends React.Component {
    state = {
        list: [],
        searchQuery: getCurrentUserSearch() || "",
        currentDropdown: -1,
    };

    componentDidMount = () => {
        this.searchUsers(this.state.searchQuery);
    };

    searchUsers = debounce((query) => {
        API.post("api/users/search", { query: query }).then((res) => {
            const users = res.data;
            this.setState({ list: users });
        });
    }, 200);

    onChangeSearch = (event) => {
        const query = event.target.value;

        setCurrentUserSearch(query);
        this.setState({ searchQuery: query }, () => {
            this.searchUsers(this.state.searchQuery);
        });
    };

    onClickDropdown = (event) => {
        if (event.target.id == this.state.currentDropdown) {
            this.setState({ currentDropdown: -1 });
        } else {
            this.setState({ currentDropdown: event.target.id });
        }
    };

    render = () => {
        const rows = this.state.list.map((row, index) => {
            return (
                <UserRow
                    row={row}
                    key={index}
                    dropdownClickCallback={this.onClickDropdown}
                    currentDropdown={this.state.currentDropdown}
                />
            );
        });

        return (
            <div id="view-user">
                <h1>Search Users</h1>

                <input
                    id="searchbar"
                    value={this.state.searchQuery}
                    onChange={this.onChangeSearch}
                    placeholder="Search for a User"
                />

                <ul id="header">
                    <li className="li-small">User ID</li>
                    <li className="li-small">Account ID</li>
                    <li className="li-med">Full Name</li>
                    <li className="li-med">Email</li>
                    <li className="li-med">Phone</li>
                    <li className="li-large">Notes</li>
                </ul>
                {rows}
            </div>
        );
    };
}

class UserRow extends React.Component {
    render = () => {
        const row = this.props.row;
        var fullName = row.firstName;
        if (row.middleName) {
            fullName += " " + row.middleName + " " + row.lastName;
        } else {
            fullName += " " + row.lastName;
        }

        return (
            <ul id="user-row">
                <li className="li-small">{row.id}</li>
                <li className="li-small">{row.accountId}</li>
                <li className="li-med">{fullName}</li>
                <li className="li-med">{row.email}</li>
                <li className="li-med">{row.phone}</li>
                <li className="li-large">{row.notes}</li>
                <Dropdown
                    id={row.id}
                    onClickCallback={this.props.dropdownClickCallback}
                    currentDropdown={this.props.currentDropdown}
                />
            </ul>
        );
    };
}

class Dropdown extends React.Component {
    render = () => {
        const editUrl = "/users/" + this.props.id + "/edit";
        const classUrl = "/users/" + this.props.id + "/class/edit";
        const afhUrl = "/users/" + this.props.id + "/afh/edit";
        return (
            <div
                className={
                    "dropdown " +
                    (this.props.id == this.props.currentDropdown
                        ? "dropdown-active"
                        : "")
                }>
                <img
                    src={DotsVertical}
                    onClick={this.props.onClickCallback}
                    id={this.props.id}
                />
                <div
                    className={
                        this.props.id == this.props.currentDropdown
                            ? "dropdown-content"
                            : "hide"
                    }>
                    <Link to={editUrl}>Edit</Link>
                    <Link to={classUrl}>Add Class</Link>
                    <Link to={afhUrl}>Add AskForHelp</Link>
                </div>
            </div>
        );
    };
}
