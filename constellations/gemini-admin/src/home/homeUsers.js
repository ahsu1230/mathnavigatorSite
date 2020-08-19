"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";
import { EmptyMessage } from "./home.js";
import moment from "moment";

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
                    <div className="id">{row.id} </div>
                    <div className="name">{getFullName(row)} </div>
                    <div className="email">{row.email} </div>
                    <div className="fromNow">
                        {moment(row.createdAt).fromNow()}{" "}
                    </div>
                    <div className="view">
                        <Link to={"/users/" + row.id + "/edit"}>
                            {"View >"}
                        </Link>
                    </div>
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
                        <div className={"list-header" + " id"}>User ID</div>
                        <div className={"list-header" + " name"}>Name</div>
                        <div className={"list-header" + " email"}>Email</div>
                        <div className={"list-header" + " fromNow"}>
                            Created
                        </div>
                        <div className={"list-header" + " view"}> </div>
                    </div>
                    <EmptyMessage
                        section={"user"}
                        length={this.state.newUsers.length}
                    />
                    <ul>{newUsers}</ul>
                </div>
            </div>
        );
    }
}
