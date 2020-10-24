"use strict";
require("./afhEdit.sass");
import React from "react";
import { Link } from "react-router-dom";
import moment from "moment";
import API from "../api.js";
import { AfhEditDateTimes } from "./afhEditDateTime.js";
import { InputText } from "../common/inputs/inputText.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

export class AskForHelpEditPage extends React.Component {
    state = {
        isEdit: false,
        afhId: 0,
        startsAt: moment(), // default is now
        endsAt: moment(), // default is now
        subject: "",
        title: "",
        locationId: "",
        notes: "",

        allSubjects: [],
        locations: [],
    };

    componentDidMount() {
        const afhId = this.props.afhId;
        API.get("api/subjects").then((res) => {
            const subjects = res.data;
            this.setState({
                allSubjects: subjects,
                subject: subjects[0],
            });
        });
        API.get("api/locations/all").then((res) => {
            const locations = res.data;
            this.setState({
                locations: locations,
                locationId: locations[0].locationId,
            });
        });
        if (afhId) {
            API.get("api/askforhelp/afh/" + afhId).then((res) => {
                const afh = res.data;
                this.setState({
                    isEdit: true,
                    afhId: afh.id,
                    startsAt: moment(afh.startsAt),
                    endsAt: moment(afh.endsAt),
                    subject: afh.subject,
                    title: afh.title,
                    locationId: afh.locationId,
                    notes: afh.notes || "",
                });
            });
        }
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onSave = () => {
        let afh = {
            id: this.state.afhId,
            startsAt: moment(this.state.startsAt),
            endsAt: moment(this.state.endsAt),
            subject: this.state.subject,
            title: this.state.title,
            locationId: this.state.locationId,
            notes: this.state.notes,
        };

        if (this.state.isEdit) {
            return API.post("api/askforhelp/afh/" + this.state.afhId, afh);
        } else {
            return API.post("api/askforhelp/create", afh);
        }
    };

    onDelete = () => {
        const afhId = this.state.afhId;
        return API.delete("api/askforhelp/afh/" + afhId);
    };

    onStartsAtChange = (moment) => {
        this.setState({ startsAt: moment });
    };

    onEndsAtChange = (moment) => {
        this.setState({ endsAt: moment });
    };

    renderContent = () => {
        const locationOptions = this.state.locations.map((loc) => {
            return {
                value: loc.locationId,
                displayName: loc.locationId,
            };
        });

        const subjectOptions = this.state.allSubjects.map((subject) => {
            return {
                value: subject,
                displayName: subject,
            };
        });

        return (
            <div>
                <AfhEditDateTimes
                    startsAt={this.state.startsAt}
                    endsAt={this.state.endsAt}
                    onStartsAtChange={this.onStartsAtChange}
                    onEndsAtChange={this.onEndsAtChange}
                />

                <InputText
                    label="Title"
                    value={this.state.title}
                    onChangeCallback={(e) => this.handleChange(e, "title")}
                    required={true}
                    description="Enter a title (i.e. AP Calculus Practice Exam, Computer Programming Office Hours, English Essay Review Session)"
                    validators={[
                        {
                            validate: (name) => name != "",
                            message: "You must input a title",
                        },
                    ]}
                />

                <InputSelect
                    label="Subject"
                    description="Select a subject"
                    value={this.state.subject}
                    onChangeCallback={(e) => this.handleChange(e, "subject")}
                    required={true}
                    options={subjectOptions}
                />

                <InputSelect
                    label="LocationId"
                    description="Select a location id"
                    value={this.state.selectLocationId}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "selectLocationId")
                    }
                    required={true}
                    options={locationOptions}
                    errorMessageIfEmpty={
                        <span>
                            There are no locations to choose from. Please add
                            one <Link to="/locations/add">here</Link>
                        </span>
                    }
                />

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.handleChange(e, "notes")}
                    description="Add any notes"
                />
            </div>
        );
    };

    render() {
        const title =
            (this.state.isEdit ? "Edit" : "Add") + " AskForHelp Session";
        const content = this.renderContent();

        return (
            <div id="view-afh-edit">
                <EditPageWrapper
                    isEdit={this.state.isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"afh"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityName={"ask-for-help session"}
                />
            </div>
        );
    }
}
