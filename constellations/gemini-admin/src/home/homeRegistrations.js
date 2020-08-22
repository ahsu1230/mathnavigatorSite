"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { getFullName } from "../utils/userUtils.js";
import { EmptyMessage } from "./home.js";
import moment from "moment";

const TAB_REGISTRATIONS = "registrations";

export class HomeTabSectionRegistrations extends React.Component {
    state = {
        classReg: [],
        afhReg: [],
    };

    componentDidMount() {
        //pending registration for classes
        API.get("api/user-classes/new").then((res) => {
            const userClass = res.data;
            this.setState({
                classReg: userClass,
            });
        });

        //afh registration
        API.get("api/userafhs/new").then((res) => {
            const userAfh = res.data;
            this.setState({
                afhReg: userAfh,
            });
        });
    }

    render() {
        let userClasses = this.state.classReg.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">
                        <UserInfo userId={row.userId} className={"name"} />
                    </div>
                    <div className="email">
                        <UserInfo userId={row.userId} className={"email"} />
                    </div>
                    <div className="class-long">{row.classId} </div>
                    <div className="from-now">
                        {moment(row.updatedAt).fromNow()}{" "}
                    </div>
                </li>
            );
        });

        let userAfh = this.state.afhReg.map((row, index) => {
            return (
                <li className="container-flex" key={index}>
                    <div className="name">
                        <UserInfo userId={row.userId} className={"name"} />
                    </div>
                    <div className="email">
                        <UserInfo userId={row.userId} className={"email"} />
                    </div>
                    <div className="class-long">
                        <AfhInfo afhId={row.afhId} />
                    </div>
                    <div className="from-now">
                        {moment(row.updatedAt).fromNow()}{" "}
                    </div>
                </li>
            );
        });

        return (
            <div id="registrations">
                <div className="section-details">
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
                            <div className={"list-header name"}>Name</div>
                            <div className={"list-header email"}>Email</div>
                            <div className={"list-header class-long"}>
                                Class Id
                            </div>
                            <div className={"list-header from-now"}>
                                Registered
                            </div>
                        </div>
                        <EmptyMessage
                            section={TAB_REGISTRATIONS}
                            length={this.state.classReg.length}
                        />
                        <ul>{userClasses}</ul>
                    </div>
                </div>

                <div className="section-details">
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
                            <div className={"list-header name"}>Name</div>
                            <div className={"list-header email"}>Email</div>
                            <div className={"list-header class-long"}>
                                AFH Session
                            </div>
                            <div className={"list-header from-now"}>
                                Registered
                            </div>
                        </div>
                        <EmptyMessage
                            section={TAB_REGISTRATIONS}
                            length={this.state.afhReg.length}
                        />
                        <ul>{userAfh}</ul>
                    </div>
                </div>
            </div>
        );
    }
}

class UserInfo extends React.Component {
    state = {
        user: {},
    };
    componentDidMount() {
        API.get("api/users/user/" + this.props.userId).then((res) => {
            const userData = res.data;
            this.setState({
                user: userData,
            });
        });
    }

    render() {
        const userItem = this.props.className;
        let returnData =
            userItem == "name"
                ? getFullName(this.state.user)
                : this.state.user.email;

        return <div>{returnData}</div>;
    }
}

class AfhInfo extends React.Component {
    state = {
        afh: {},
    };
    componentDidMount() {
        API.get("api/askforhelp/afh/" + this.props.afhId).then((res) => {
            const afhData = res.data;
            this.setState({
                afh: afhData,
            });
        });
    }

    render() {
        let returnData = this.state.afh.title;

        return <div>{returnData}</div>;
    }
}
