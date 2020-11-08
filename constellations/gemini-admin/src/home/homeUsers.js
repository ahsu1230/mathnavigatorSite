"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";

export class HomeTabSectionUsers extends React.Component {
    render() {
        const users = this.props.users || [];
        const list = users.map((row, index) => {
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
                    <div>
                        <Link to={"/users"}>View All Users</Link>
                    </div>
                </div>
                {list.length > 0 ? (
                    <div className="class-section">
                        <ul>{list}</ul>
                    </div>
                ) : (
                    <p className="empty">No new users recently.</p>
                )}
            </div>
        );
    }
}
