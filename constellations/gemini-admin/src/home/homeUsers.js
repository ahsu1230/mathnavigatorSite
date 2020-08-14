"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";

export class HomeTabSectionUsers extends React.Component {
    state = {
        newUsers: [],
    };

    componentDidMount() {
        API.get("api/users/new").then((res) => {
            const users = res.data;
            this.setState({
                newUsers: users,
            });
        });
    }

    render() {
        let newUsers = this.state.newUsers.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">{getFullName(row)} </div>
                    <div className="email">{row.email} </div>
                    <div className="other">{row.id} </div>
                </li>
            );
        });

        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">New Users</h3>
                    <button className="view-details">
                        <Link to={"/users"}>View All Users</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="container-flex">
                        <div className={"list-header" + " name"}>Name</div>
                        <div className={"list-header" + " email"}>Email</div>
                        <div className={"list-header" + " other"}>User ID</div>
                    </div>

                    <ul>{newUsers}</ul>
                </div>
            </div>
        );
    }
}
