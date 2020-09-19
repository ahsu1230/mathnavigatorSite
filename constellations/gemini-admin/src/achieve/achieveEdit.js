"use strict";
require("./achieveEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { InputText, emptyValidator } from "../common/inputs/inputText.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

export class AchieveEditPage extends React.Component {
    state = {
        isEdit: false,
        achieveId: 0,
        year: moment().year(),
        position: 0,
        message: "",
    };

    componentDidMount = () => {
        const id = this.props.id;
        if (id) {
            API.get("api/achievements/achievement/" + id).then((res) => {
                const achieve = res.data;
                this.setState({
                    isEdit: true,
                    year: achieve.year,
                    position: achieve.position,
                    message: achieve.message,
                });
            });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onSave = () => {
        const achieveId = this.props.id;
        let achieve = {
            year: parseInt(this.state.year),
            position: parseInt(this.state.position),
            message: this.state.message,
        };
        if (this.state.isEdit) {
            return API.post(
                "api/achievements/achievement/" + achieveId,
                achieve
            );
        } else {
            return API.post("api/achievements/create", achieve);
        }
    };

    onDelete = () => {
        return API.delete("api/achievements/achievement/" + this.props.id);
    };

    renderContent = () => {
        return (
            <div>
                <InputText
                    label="Year"
                    description="Enter the achievement year"
                    required={true}
                    value={this.state.year}
                    onChangeCallback={(e) => this.handleChange(e, "year")}
                    validators={[
                        {
                            validate: (number) => parseInt(number) > 2000,
                            message: "You must input a year greater than 2000",
                        },
                    ]}
                />

                <InputText
                    label="Position"
                    description="Enter the position (Lower position numbers are shown first in that year)"
                    required={true}
                    value={this.state.position}
                    onChangeCallback={(e) => this.handleChange(e, "position")}
                    validators={[
                        {
                            validate: (number) =>
                                Number.isInteger(parseInt(number)) &&
                                parseInt(number) > 0,
                            message: "You must input a positive integer",
                        },
                    ]}
                />

                <InputText
                    label="Message"
                    description="Enter the achievement message"
                    required={true}
                    isTextBox={true}
                    value={this.state.message}
                    onChangeCallback={(e) => this.handleChange(e, "message")}
                    validators={[emptyValidator("message")]}
                />
            </div>
        );
    };

    render = () => {
        const isEdit = this.state.isEdit;
        const title = isEdit ? "Edit Achievement" : "Add Achievement";

        return (
            <div id="view-achieve-edit">
                <EditPageWrapper
                    isEdit={isEdit}
                    title={title}
                    content={this.renderContent()}
                    prevPageUrl={"achievements"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityId={this.state.achieveId}
                    entityName={"achievement"}
                />
            </div>
        );
    };
}
