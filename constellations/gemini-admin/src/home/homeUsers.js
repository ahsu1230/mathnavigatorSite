"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

// does not work

const sectionDisplayNames = {
    user: "New Users",
};

export class HomeTabSectionUsers extends React.Component {
    state = {
        newUsers: [],
    };

    // counter to keep track of the number of new users => newUsers.length

    //new users
    /*   componentDidMount() {
        API.get("api/users").then((res) => {
            const users = res.data;
            this.setState({
                newUsers: users,
            });
        });
    } */

    render() {
        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">New Users</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>View All Users</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="header-flex">
                        <div className={"list-header" + " name"}>Name</div>
                        <div className={"list-header" + " email"}>Email</div>
                        <div className={"list-header" + " other"}>
                            Account Number
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
