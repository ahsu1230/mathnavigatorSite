"use strict";
require("./userClassRow.sass");
import React from "react";
import API from "../api.js";

export class UserClassRow extends React.Component {
    state = {
        currentState: 0,
        newState: 0,
    };

    componentDidMount = () => {
        const state = this.props.userClass.state;
        this.setState({
            currentState: state,
            newState: state == 1 ? 3 : 1,
        });
    };

    formatState = (state) => {
        return ["Pending", "Accepted", "Trial", "Unenrolled", "Rejected"][
            state
        ];
    };

    onSelectState = (e) => {
        this.setState({
            newState: e.target.value,
        });
    };

    onSubmitState = () => {
        const userClass = this.props.userClass;
        const state = this.state.newState;

        if (state == 4) {
            this.onUserClassDelete(userClass);
        } else {
            this.onUserClassUpdate(userClass, state);
        }
    };

    onUserClassUpdate = (userClass, state) => {
        const newUserClass = {
            userId: parseInt(userClass.userId),
            classId: userClass.classObject.classId,
            accountId: parseInt(userClass.accountId),
            state: parseInt(state),
        };

        API.post("api/user-classes/user-class/" + userClass.id, newUserClass)
            .then(() =>
                this.setState({
                    currentState: state,
                    newState: state == 1 ? 3 : 1,
                })
            )
            .catch((err) => alert("Could not accept enrollment: " + err));
    };

    onUserClassDelete = (userClass) => {
        API.delete("api/user-classes/user-class/" + userClass.id)
            .then(() => this.props.updateCallback())
            .catch((err) => alert("Could not reject enrollment: " + err));
    };

    render = () => {
        const userClass = this.props.userClass;
        const enrollState = this.state.currentState;
        const state = ["Normal", "Almost Full", "Full"][
            userClass.classObject.fullState
        ];

        // 0, 1, 2, 3, 4 represent pending, accepted, trial, unenrolled, rejected respectively
        const options = {
            0: [1, 2, 4],
            1: [3],
            2: [1, 3],
            3: [1],
            4: [],
        };

        const stateOptions = options[enrollState].map((option, index) => {
            return (
                <option value={option} key={index}>
                    {this.formatState(option)}
                </option>
            );
        });

        return (
            <div className="row">
                <span className="large-column">
                    {userClass.classObject.classId}
                </span>
                <span className="column">{this.formatState(enrollState)}</span>
                <span className="column">{state}</span>
                <span className="space">
                    <select
                        className="dropdown"
                        value={this.state.newState}
                        onChange={(e) => this.onSelectState(e)}>
                        {stateOptions}
                    </select>
                </span>
                <span className="space">
                    <button onClick={this.onSubmitState}>Update</button>
                </span>
            </div>
        );
    };
}
