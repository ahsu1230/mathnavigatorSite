"use strict";
require("./home.sass");
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

    //unpublished classes
    componentDidMount() {
        API.get("api/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                unpubClasses: unpublishedList.classes,
            });
        });
    }

    render() {
        //same as sectionDisplayNames
        let currentSection = this.props.section;

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
