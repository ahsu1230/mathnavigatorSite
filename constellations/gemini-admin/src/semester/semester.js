"use strict";
require("./semester.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    A Semester consists of a season and a year. The title and semesterId are automatically generated based on these values. 
    The only available values for season are winter, spring, summer, fall. 
    When displaying programs in the user website's "Program Catalog" page, all programs will be grouped / sorted by these semesters 
    in ascending order (most recent year, then winter -> spring -> summer -> fall).
`;

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
                    description={PAGE_DESCRIPTION}
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
