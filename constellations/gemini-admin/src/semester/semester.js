"use strict";
require("./semester.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

export class SemesterPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/semesters/all").then((res) => {
            const semesters = res.data;
            this.setState({ list: semesters });
        });
    }

    render() {
        const numSemesters = this.state.list.length;
        const rows = this.state.list.map((semester, index) => {
            const fields = generateFields(semester);
            return (
                <RowCardBasic
                    key={index}
                    title={semester.title}
                    subtitle={semester.semesterId}
                    editUrl={"/semesters/" + semester.semesterId + "/edit"}
                    fields={fields}
                />
            );
        });

        return (
            <div id="view-semester">
                <AllPageHeader
                    title={"All Semesters (" + numSemesters + ")"}
                    addUrl={"/semesters/add"}
                    addButtonTitle={"Add Semester"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />

                <div className="cards-wrapper">{rows}</div>
            </div>
        );
    }
}

function generateFields(semester) {
    return [
        {
            label: "Year",
            value: semester.year,
        },
        {
            label: "Season",
            value: semester.season,
        },
    ];
}
