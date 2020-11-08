"use strict";
require("./homeSection.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import RowCardBasic from "../common/rowCards/rowCardBasic.js";

export class HomeTabSectionClasses extends React.Component {
    render() {
        const unpublishedClasses = (this.props.unpublishedClasses || []).map(
            (row, index) => {
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
            }
        );

        return (
            <div className="section-details">
                <div className="container-class">
                    <h3 className="section-header">Unpublished Classes</h3>
                    <Link to={"/classes"}>View All Classes to Publish</Link>
                </div>

                {unpublishedClasses.length > 0 ? (
                    <div className="class-section">
                        <ul>{unpublishedClasses}</ul>
                    </div>
                ) : (
                    <p className="empty">
                        No unpublished classes at the moment. All classes are
                        publicly available!
                    </p>
                )}
            </div>
        );
    }
}
