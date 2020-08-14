"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";


export class HomeTabSectionRegistrations extends React.Component {
    state = {
        pendingReg: [],
        afhReg: [],
    };

    //pending registration for classes
    componentDidMount() {
        API.get("api/user-classes/new").then((res) => {
            const userClass = res.data;
            this.setState({
                pendingReg: userClass,
            });
        });
    }

    //afh registration
    componentDidMount() {
        API.get("api/userafhs/new").then((res) => {
            const userAfh = res.data;
            this.setState({
                afhReg: userAfh,
            });
        });
    }

    render() {
        let userClasses = this.state.pendingReg.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">{getFullName(row)} </div>
                    <div className="email">{row.email} </div>
                    <div className="other">{row.classId} </div>
                </li>
            );
        });

        let userAfh = this.state.afhReg.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">{getFullName(row)} </div>
                    <div className="email">{row.email} </div>
                    <div className="other">{row.title} </div>
                </li>
            );
        });


        return (
            <div id="registrations">
                <div className="sectionDetails">
                    <div className="container-class">
                        <h3 className="section-header">
                            Pending Registrations For Classes
                        </h3>
                        <button className="view-details">
                            <Link to={"/classes"}>View By Class</Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="container-flex">
                            <div className={"list-header" + " name"}>Name</div>
                            <div className={"list-header" + " email"}>Email</div>
                            <div className={"list-header" + " other"}>ClassId</div>
                        </div>

                        <ul>{userClasses}</ul>

                    </div>
                </div>

                <div className="sectionDetails">
                    <div className="container-class">
                        <h3 className="section-header">
                            New Registrations For AFH
                        </h3>
                        <button className="view-details">
                            <Link to={"/afh"}>View By AFH Session</Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="container-flex">
                            <div className={"list-header" + " name"}>Name</div>
                            <div className={"list-header" + " email"}>Email</div>
                            <div className={"list-header" + " other"}>Registered For</div>
                        </div>

                        <ul>{userAfh}</ul>

                    </div>
                </div>
            </div>
        );
    }
}
