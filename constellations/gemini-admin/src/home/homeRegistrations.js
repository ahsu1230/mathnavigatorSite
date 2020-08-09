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

    // counter to keep track of the number of registrations => pendingReg.length + afhReg.length

    //pending registration for classes
    /*  componentDidMount() {
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
    } */

    render() {
        // flexbox for headers (Name, Email, ClassId)
        // flexbox for headers (Names, Email, RegisteredAt)
        return (
            <div id="registrations">
                <div className="sectionDetails">
                    <div className="container-class">
                        <h3 className="section-header">
                            Pending Registrations For Classes
                        </h3>{" "}
                        <button className="view-details">
                            <Link to={"/classes"}>View By Class</Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="list-header">Name</div>
                    </div>
                </div>

                <div className="sectionDetails">
                    <div className="container-class">
                        <h3 className="section-header">
                            New Registrations For AFH
                        </h3>{" "}
                        <button className="view-details">
                            <Link to={"/classes"}>View By AFH Session</Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="list-header">Name</div>
                    </div>
                </div>
            </div>
        );
    }
}
