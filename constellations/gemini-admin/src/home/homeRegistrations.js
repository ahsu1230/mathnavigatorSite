"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

// does not work

const sectionDisplayNames = {
    registration: "New Registrations",
};

export class HomeTabSectionRegistrations extends React.Component {
    state = {
        pendingReg: [],
        afhReg: [],
    };

    // counter to keep track of the number of registrations => pendingReg.size + afhReg.size

    //pending registration for classes
    componentDidMount() {
        API.get("api/user-classes").then((res) => {
            const userClass = res.data;
            this.setState({
                pendingReg: userClass,
            });
        });
    }

    //afh registration
    componentDidMount() {
        API.get("api/userafhs").then((res) => {
            const userAfh = res.data;
            this.setState({
                afhReg: userAfh,
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
