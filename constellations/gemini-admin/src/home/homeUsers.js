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
    componentDidMount() {
        API.get("api/users").then((res) => {
            const users = res.data;
            this.setState({
                newUsers: users,
            });
        });
    }

    render() {
        let unpublishedClasses = this.state.unpubClasses.map((row, index) => {
            return <li key={index}> {row.classId} </li>;
        });
        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">Unpublished Classes</h3>{" "}
                    <button id="publish">
                        <Link to={"/classes"}>View All Classes to Publish</Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="list-header">Class ID</div>
                    <ul>{unpublishedClasses}</ul>
                </div>
            </div>
        );
    }
}
