"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

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
        let numClasses = this.state.unpubClasses.length;

        let publishedMessage = <div></div>;
        if (numClasses == 0) {
            publishedMessage = (
                <p>All classes have been successfully published!</p>
            );
        }

        let unpublishedClasses = this.state.unpubClasses.map((row, index) => {
            return <li key={index}> {row.classId} </li>;
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
                    <div className="list-header">Class ID</div>
                    {publishedMessage}
                    <ul>{unpublishedClasses}</ul>
                </div>
            </div>
        );
    }
}
