"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import moment from "moment";
import { EmptyMessage } from "./home.js";
import RowCardBasic from "../common/rowCards/rowCardBasic.js";

const TAB_CLASSES = "classes";

export class HomeTabSectionClasses extends React.Component {
    state = {
        unpubClasses: [],
    };

    componentDidMount() {
        API.get("api/classes/unpublished").then((res) => {
            const unpublishedList = res.data || [];
            this.setState({
                unpubClasses: unpublishedList,
            });
        });
    }

    render() {
        let unpublishedClasses = this.state.unpubClasses.map((row, index) => {
            return (
                <li key={index}>
                    <RowCardBasic
                        title={row.classId}
                        editUrl={"/classes/" + row.classId + "/edit"}
                        fields={[
                            {
                                label: "Created",
                                value: moment(row.createdAt).fromNow(),
                            },
                            {
                                label: "Last Updated",
                                value: moment(row.updatedAt).fromNow(),
                            },
                        ]}
                    />
                </li>
            );
        });

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">Unpublished Classes</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>
                            View All Classes to Publish{" "}
                        </Link>
                    </button>
                </div>

                <div className="class-section">
                    <EmptyMessage
                        section={TAB_CLASSES}
                        length={this.state.unpubClasses.length}
                    />
                    <ul>{unpublishedClasses}</ul>
                </div>
            </div>
        );
    }
}
