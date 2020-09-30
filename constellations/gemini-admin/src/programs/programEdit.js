"use strict";
require("./programEdit.sass");
import React from "react";
import API from "../api.js";
import { InputText, emptyValidator } from "../common/inputs/inputText.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

export class ProgramEditPage extends React.Component {
    state = {
        isEdit: false,
        oldProgramId: "",
        programId: "",
        title: "",
        grade1: 0,
        grade2: 0,
        subject: "",
        description: "",
        featured: "",
        allFeatured: [],
        allSubjects: [],
    };

    componentDidMount = () => {
        const programId = this.props.programId;
        API.get("api/subjects").then((res) => {
            const subjects = res.data;
            this.setState({
                allSubjects: subjects,
            });
        });

        if (programId) {
            API.get("api/programs/program/" + programId).then((res) => {
                const program = res.data;
                this.setState({
                    oldProgramId: program.programId,
                    programId: program.programId,
                    title: program.title,
                    grade1: program.grade1,
                    grade2: program.grade2,
                    subject: program.subject,
                    description: program.description,
                    isEdit: true,
                    featured: program.featured,
                });
            });
        }

        API.get("api/programs/featured").then((res) => {
            this.setState({ allFeatured: res.data });
        });
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onDelete = () => {
        const programId = this.props.programId;
        return API.delete("api/programs/program/" + programId);
    };

    onSave = () => {
        let program = {
            programId: this.state.programId,
            title: this.state.title,
            grade1: parseInt(this.state.grade1),
            grade2: parseInt(this.state.grade2),
            subject: this.state.subject,
            description: this.state.description,
            featured: this.state.featured,
        };

        if (this.state.isEdit) {
            return API.post(
                "api/programs/program/" + this.state.oldProgramId,
                program
            );
        } else {
            return API.post("api/programs/create", program);
        }
    };

    renderContent = () => {
        const subjectOptions = this.state.allSubjects.map((subject) => {
            return {
                value: subject,
                displayName: subject,
            };
        });

        return (
            <div>
                <InputText
                    label="Program Id"
                    value={this.state.programId}
                    onChangeCallback={(e) => this.handleChange(e, "programId")}
                    required={true}
                    description="Enter the program Id. Examples: ap_calculus, sat1, ap_java"
                    validators={[emptyValidator("program ID")]}
                />

                <InputText
                    label="Program Title"
                    value={this.state.title}
                    onChangeCallback={(e) => this.handleChange(e, "title")}
                    required={true}
                    description="Enter the program title. This title will be present to users. Example: AP Calculus, SAT2 Subject Math"
                    validators={[emptyValidator("program title")]}
                />

                <InputText
                    label="Grade1"
                    value={this.state.grade1}
                    onChangeCallback={(e) => this.handleChange(e, "grade1")}
                    required={true}
                    description="Enter the lower grade"
                    validators={[
                        emptyValidator("grade"),
                        {
                            validate: (grade1) =>
                                parseInt(grade1) >= 1 && parseInt(grade1) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade1) =>
                                this.state.grade2 >= parseInt(grade1),
                            message:
                                "Grade1 must be less than or equal to Grade2",
                        },
                    ]}
                />

                <InputText
                    label="Grade2"
                    value={this.state.grade2}
                    onChangeCallback={(e) => this.handleChange(e, "grade2")}
                    required={true}
                    description="Enter the higher grade"
                    validators={[
                        emptyValidator("grade"),
                        {
                            validate: (grade2) =>
                                parseInt(grade2) >= 1 && parseInt(grade2) <= 12,
                            message: "Grade must be between 1 and 12",
                        },
                        {
                            validate: (grade2) =>
                                this.state.grade1 <= parseInt(grade2),
                            message:
                                "Grade2 must be greater than or equal to Grade1",
                        },
                    ]}
                />

                <InputText
                    label="Description"
                    isTextBox={true}
                    value={this.state.description}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "description")
                    }
                    required={true}
                    description="Enter the description"
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a description",
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
                    label="Featured"
                    description="Some programs can have an optional 'feature' flag to differentiate them from other programs."
                    value={this.state.featured}
                    onChangeCallback={(e) => this.handleChange(e, "featured")}
                    options={this.state.allFeatured.map((featured) => {
                        return {
                            value: featured,
                            displayName: featured,
                        };
                    })}
                />
            </div>
        );
    };

    render = () => {
        const isEdit = this.state.isEdit;
        const title = isEdit ? "Edit Program" : "Add Program";
        const content = this.renderContent();

        return (
            <div id="view-program-edit">
                <EditPageWrapper
                    isEdit={isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"programs"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityId={this.state.programId}
                    entityName={"program"}
                />
            </div>
        );
    };
}
