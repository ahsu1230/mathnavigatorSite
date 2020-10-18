"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";
import { EmptyMessage } from "./home.js";

const TAB_USERS = "users";

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
                <li key={index}>
                    <RowCardColumns
                        title={getFullName(row)}
                        editUrl={"/users/" + row.id + "/edit"}
                        fieldsList={[
                            [
                                {
                                    label: "Email",
                                    value: row.email,
                                },
                                {
                                    label: "Phone",
                                    value: row.phone,
                                },
                                {
                                    label: "Created",
                                    value: moment(row.createdAt).fromNow(),
                                },
                            ],
                            [
                                {
                                    label: "IsGuardian",
                                    value: row.isGuardian ? "true" : "false",
                                },
                                {
                                    label: "School",
                                    value: row.school,
                                },
                                {
                                    label: "GraduationYear",
                                    value: row.graduationYear,
                                },
                            ],
                        ]}
                    />
                </li>
            );
        });

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">New Users</h3>
                    <button className="view-details">
                        <Link to={"/users"}>View All Users</Link>
                    </button>
                </div>

                <div className="class-section">
                    <EmptyMessage
                        section={TAB_USERS}
                        length={this.state.newUsers.length}
                    />
                    <ul>{newUsers}</ul>
                </div>
            </div>
        );
    }
}
