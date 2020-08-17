"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import moment from "moment";
import { EmptyMessage } from "./home.js";

const sectionDisplayNames = {
    class: "Unpublished Classes",
};

export class HomeTabSectionClasses extends React.Component {
    state = {
        unpubClasses: [],
    };

    componentDidMount() {
        API.get("api/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                unpubClasses: unpublishedList.classes,
            });
        });
    }

    render() {
        let unpublishedClasses = this.state.unpubClasses.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="width50"> {row.classId} </div>
                    <div className="width50">
                        {" "}
                        {moment(row.updatedAt).fromNow()}{" "}
                    </div>
                </li>
            );
        });

        return (
            <div className="sectionDetails">
                <div className="container-class">
                    <h3 className="section-header">Unpublished Classes</h3>{" "}
                    <button className="view-details">
                        <Link to={"/classes"}>
                            View All Classes to Publish{" "}
                        </Link>
                    </button>
                </div>

                <div className="class-section">
                    <div className="container-flex">
                        <div className={"list-header" + " width50"}>
                            Class Id
                        </div>
                        <div className={"list-header" + " width50"}>
                            Last Updated
                        </div>
                    </div>
                    <EmptyMessage
                        section={"class"}
                        length={this.state.unpubClasses.length}
                    />
                    <ul>{unpublishedClasses}</ul>
                </div>
            </div>
        );
    }
}
