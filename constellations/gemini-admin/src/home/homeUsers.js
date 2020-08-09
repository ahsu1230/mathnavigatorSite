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
        // flexbox for headers (Name, Email, Account Number ? )

        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">New Users</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>View User Details ? </Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="list-header">Name</div>
                </div>
            </div>
        );
    }
}
