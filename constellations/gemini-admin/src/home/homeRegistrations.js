"use strict";
require("./homeSection.sass");
import React from "react";
import API from "../api.js";
import moment from "moment";
import { getFullName } from "../common/userUtils.js";
import RowCardColumns from "../common/rowCards/rowCardColumns.js";

export class HomeTabSectionRegistrations extends React.Component {
    render() {
        const classReg = this.props.newUserClasses || [];
        const afhReg = this.props.newUserAfh || [];
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
                    {registrations.length > 0 ? (
                        <div className="class-section">
                            <ul>{registrations}</ul>
                        </div>
                    ) : (
                        <p className="empty">No new registrations recently.</p>
                    )}
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
