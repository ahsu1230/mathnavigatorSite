// TODO: Add Select All and Deselect All Buttons

"use strict";
require("./emailPrograms.sass");
import axios from "axios";
import React from "react";
import API, { executeApiCalls } from "../api.js";
import { getFullName } from "../common/userUtils.js";

export class EmailProgramsPage extends React.Component {
    state = {
        selectProgramId: "",
        selectedProgramName: "",
        programs: [],
        classes: [],
        classesForProgram: [],
        userIds: [],
        usersInClass: [],
        selectedUsers: [],
    };

    componentDidMount = () => {
        const apiCalls = [
            API.get("api/programs/all"),
            API.get("api/classes/all"),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const programs = responses[0].data;
                    const classes = responses[1].data;
                    const hasClassId = responses.length > 3;
                    let classObj = hasClassId ? responses[3].data : {};
                    let selectedProgramId = hasClassId
                        ? classObj.programId
                        : programs[0].programId;

                    this.setState({
                        selectProgramId: selectedProgramId,
                        programs: programs,
                        classes: classes,
                    });
                })
            )
            .catch((err) => {
                console.log("Error: api call failed. " + err.message);
            });
    };

    handleProgramChange = (event) => {
        const programId = event.target.value;
        var selectedProgram = this.state.programs.find(
            (program) => program.programId === programId
        );
        selectedProgram = selectedProgram.name;
        let classesForProgram = this.state.classes.filter(
            (classes) => classes.programId == programId
        );
        const classes = classesForProgram.map((classes) => classes.classId);

        this.setState({
            selectProgramId: programId,
            classesForProgram: classes,
            selectedProgramName: selectedProgram,
            selectedUsers: [],
            usersInClass: [],
        });
    };

    onCheckClass = (e, classId) => {
        if (e.target.checked) {
            API.get("api/user-classes/class/" + classId)
                .then((response) => {
                    const users = response.data;
                    const userIds = users.map((user) => user.userId);
                    userIds.forEach((userId) => {
                        API.get("api/users/user/" + userId)
                            .then((response) => {
                                const user = response.data;
                                var users = this.state.usersInClass;
                                users.push(user);
                                this.setState({
                                    usersInClass: users,
                                });
                            })
                            .catch((err) => {
                                console.log(
                                    "Error: api call failed. " + err.message
                                );
                            });
                    });
                })
                .catch((err) => {
                    console.log("Error: api call failed. " + err.message);
                });
        } else {
            this.setState({
                usersInClass: [],
                selectedUsers: [],
            });
        }
    };

    onCheckUser = (e, userId) => {
        const users = this.state.usersInClass;
        if (e.target.checked) {
            var currentSelectedUsers = this.state.selectedUsers;
            const checkedUser = users.find((user) => user.id == userId);
            currentSelectedUsers.push(checkedUser);
            this.setState({
                selectedUsers: currentSelectedUsers,
            });
        } else {
            var currentSelectedUsers = this.state.selectedUsers;
            const uncheckedUser = users.find((user) => user.id == userId);
            currentSelectedUsers.splice(
                currentSelectedUsers.indexOf(uncheckedUser),
                1
            );
            this.setState({
                selectedUsers: currentSelectedUsers,
            });
        }
    };

    createUserRows = (users) => {
        const userRows = users.map((user, index) => {
            var status = user.isGuardian ? "Guardian" : "Student";
            return (
                <div className="row" key={index}>
                    <input
                        className="userSelect"
                        type="checkbox"
                        onChange={(e) => this.onCheckUser(e, user.id)}
                    />
                    <span className="column">{getFullName(user)}</span>
                    <span className="column">{user.email}</span>
                    <span className="column status">{status}</span>
                </div>
            );
        });
        return userRows;
    };

    render = () => {
        const programOptions = this.state.programs.map((program, index) => (
            <option key={index} value={program.programId}>
                {program.name}
            </option>
        ));

        const classOptions = this.state.classesForProgram.map(
            (classes, index) => (
                <div className="class" key={index}>
                    <input
                        type="checkbox"
                        onChange={(e) => this.onCheckClass(e, classes)}
                    />
                    <span className="classId">{classes}</span>
                </div>
            )
        );

        const userRows = this.createUserRows(this.state.usersInClass);

        return (
            <div id="view-program-emails">
                <section id="Select-program">
                    <h1>Generate Email to Program</h1>

                    <h2>Select a Program</h2>
                    <select
                        value={this.state.selectProgramId}
                        onChange={(e) => this.handleProgramChange(e)}>
                        {programOptions}
                    </select>

                    <h2>Select a class for {this.state.selectedProgramName}</h2>
                    {classOptions}

                    <section>
                        <h2>User Email List</h2>
                        <div id="userRows">{userRows}</div>
                    </section>

                    <section id="generated-email">
                        <div className="email-template">
                            <span id="template-title">
                                Generated Email Template
                            </span>
                            <h3>
                                To:{" "}
                                {this.state.selectedUsers
                                    .map((user) => user.email)
                                    .toString()}
                            </h3>
                            <h3>
                                Subject: Math Navigator{" "}
                                {this.state.selectedProgramName} Announcement
                            </h3>
                            <h3>Message: </h3>
                            <div className="generated-email">
                                <p>
                                    To all Math Navigator parents of{" "}
                                    {this.state.selectedProgramName},
                                </p>
                                <p>
                                    Lorem ipsum dolor sit amet, consectetur
                                    adipiscing elit, sed do eiusmod tempor
                                    incididunt ut labore et dolore magna aliqua.
                                    Ut enim ad minim veniam, quis nostrud
                                    exercitation ullamco laboris nisi ut aliquip
                                    ex ea commodo consequat. Duis aute irure
                                    dolor in reprehenderit in voluptate velit
                                    esse cillum dolore eu fugiat nulla pariatur.
                                    Excepteur sint occaecat cupidatat non
                                    proident, sunt in culpa qui officia deserunt
                                    mollit anim id est laborum.
                                </p>
                                <p>
                                    Best wishes from the Math Navigator Family
                                </p>
                            </div>
                        </div>
                    </section>
                </section>
            </div>
        );
    };
}
