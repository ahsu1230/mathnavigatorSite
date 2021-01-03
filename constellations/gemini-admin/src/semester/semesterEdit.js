"use strict";
require("./semesterEdit.sass");
import React from "react";
import API from "../api.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import { InputText, emptyValidator } from "../common/inputs/inputText.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

const ALL_SEASONS = ["winter", "spring", "summer", "fall"];

export class SemesterEditPage extends React.Component {
    state = {
        inputSeason: ALL_SEASONS[0],
        inputYear: 2020,
        oldSemesterId: "",
        semesterId: "",
        title: "",
        isEdit: false,
    };

    componentDidMount = () => {
        const semesterId = this.props.semesterId;
        if (semesterId) {
            API.get("api/semesters/semester/" + semesterId).then((res) => {
                const semester = res.data;
                this.setState({
                    inputSeason: semester.season,
                    inputYear: semester.year,
                    oldSemesterId: semester.semesterId,
                    semesterId: semester.semesterId,
                    title: semester.title,
                    isEdit: true,
                });
            });
        }
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onSave = () => {
        const semester = {
            season: this.state.inputSeason,
            year: parseInt(this.state.inputYear),
            // Only send season & year to API.
            // Backend will automatically generate semesterId + title
        };

        if (this.state.isEdit) {
            return API.post(
                "api/semesters/semester/" + this.props.semesterId,
                semester
            );
        } else {
            return API.post("api/semesters/create", semester);
        }
    };

    onDelete = () => {
        const semesterId = this.props.semesterId;
        // return API.delete("api/semesters/semester/" + semesterId);   // hard-delete
        return API.delete("api/semesters/archive/" + semesterId);       // soft-delete
    };

    renderContent = () => {
        const season = this.state.inputSeason;
        const year = this.state.inputYear;
        const semesterId = formSemesterId(season, year);
        const title = formSemesterTitle(season, year);

        return (
            <div>
                <InputSelect
                    label="Season"
                    description="Select which season this semester is in."
                    required={true}
                    value={season}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputSeason")
                    }
                    options={ALL_SEASONS.map((season) => {
                        return {
                            value: season,
                            displayName: season,
                        };
                    })}
                />

                <InputText
                    label="Year"
                    description="Input the year of this semester"
                    required={true}
                    value={year}
                    onChangeCallback={(e) => this.handleChange(e, "inputYear")}
                    validators={[
                        emptyValidator("year"),
                        {
                            validate: (year) =>
                                parseInt(year) >= 2000 && parseInt(year) < 2100,
                            message: "Must be a valid year!",
                        },
                    ]}
                />
                <div className="semester-line">
                    SemesterId: <span>{semesterId}</span>
                </div>
                <div className="semester-line">
                    Semester Title: <span>{title}</span>
                </div>
            </div>
        );
    };

    render = () => {
        const title = this.state.isEdit ? "Edit Semester" : "Add Semester";
        const content = this.renderContent();

        return (
            <div id="view-semester-edit">
                <EditPageWrapper
                    isEdit={this.state.isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"semesters"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityId={this.state.semesterId}
                    entityName={"semester"}
                />
            </div>
        );
    };
}

function formSemesterId(season, year) {
    return year + "_" + season;
}

function formSemesterTitle(season, year) {
    return (
        season.substring(0, 1).toUpperCase() + season.substring(1) + " " + year
    );
}
