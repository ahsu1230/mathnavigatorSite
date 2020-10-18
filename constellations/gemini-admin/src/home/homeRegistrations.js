"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import moment from "moment";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";
import { EmptyMessage } from "./home.js";

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
        API.get("api/user-afhs/new").then((res) => {
            const userAfh = res.data;
            this.setState({
                afhReg: userAfh,
            });
        });
    }

    render() {
        const classReg = this.state.classReg || [];
        const afhReg = this.state.afhReg || [];
        const allRegistrations = classReg.concat(afhReg);
        let registrations = allRegistrations.map((row, index) => {
            return (
                <li key={index}>
                    <RegistrationInfo row={row} />
                </li>
            );
        });

        return (
            <div id="registrations">
                <div className="section-details">
                    <div className="container-class">
                        <h3 className="section-header">
                            Pending Registrations
                        </h3>
                    </div>

                    <div className="class-section">
                        <EmptyMessage
                            section={TAB_REGISTRATIONS}
                            length={registrations.length}
                        />
                        <ul>{registrations}</ul>
                    </div>
                </div>
            </div>
        );
    }
}

class RegistrationInfo extends React.Component {
    state = {
        user: {},
        afhObj: {},
        classObj: {},
    };

    componentDidMount() {
        API.get("api/users/user/" + this.props.row.userId).then((res) => {
            const userData = res.data;
            this.setState({
                user: userData,
            });
        });

        if (this.props.row.afhId) {
            API.get("api/askforhelp/afh/" + this.props.row.afhId).then(
                (res) => {
                    const afhData = res.data;
                    this.setState({
                        afhObj: afhData,
                    });
                }
            );
        }

        if (this.props.row.classId) {
            API.get("api/classes/class/" + this.props.row.classId).then(
                (res) => {
                    const classData = res.data;
                    this.setState({
                        classObj: classData,
                    });
                }
            );
        }
    }

    createSecondColumn = (row) => {
        const classId = row.classId;
        const classObj = this.state.classObj;
        const afhObj = this.state.afhObj;
        if (this.props.row.afhId) {
            return [
                {
                    label: "AFH",
                    value: afhObj.title,
                },
                {
                    label: "Time",
                    value: afhObj.startsAt
                        ? moment(afhObj.startsAt).format("llll")
                        : undefined,
                },
            ];
        } else {
            // is a class
            console.log(JSON.stringify(classObj));
            return [
                {
                    label: "Class",
                    value: classId,
                },
                {
                    label: "Times",
                    value: classObj.timesStr,
                },
            ];
        }
    };

    render() {
        const row = this.props.row;
        const userName = getFullName(this.state.user);
        const userEmail = this.state.user.email;
        const classId = row.classId;
        return (
            <RowCardColumns
                title={userName}
                subtitle={
                    classId ? "Registered for class" : "Registered for AFH"
                }
                fieldsList={[
                    [
                        {
                            label: "Email",
                            value: userEmail,
                        },
                        {
                            label: "Created",
                            value: moment(row.createdAt).fromNow(),
                        },
                        {
                            label: "Last Updated",
                            value: moment(row.updatedAt).fromNow(),
                        },
                    ],
                    this.createSecondColumn(row),
                ]}
            />
        );
    }
}
