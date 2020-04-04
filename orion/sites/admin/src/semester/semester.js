"use strict";
require("./semester.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";

export class SemesterPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: []
        };
    }

    render() {
        var numSemesters = 3;
        let fakeSemesterA = {
            semesterId: "2020_fall",
            title: "2020 Fall"
        };
        let fakeSemesterB = {
            semesterId: "2021_spring",
            title: "2021 Spring"
        };
        /* TODO: will use this to test text wrapping later */
        let fakeSemesterC = {
            semesterId: "2021_summer",
            title: "2020 Summer"
        };

        return (
            <div id="view-semester">
                <h1>All Semesters ({numSemesters})</h1>
                <ul className="semester-lists">
                    <li className="li-large"> Semester ID </li>
                    <li className="li-large"> Title </li>
                    <li className="li-small"> </li>
                </ul>
                <ul>
                    <SemesterRow semesterObj={fakeSemesterA} />
                    <SemesterRow semesterObj={fakeSemesterB} />
                    <SemesterRow semesterObj={fakeSemesterC} />
                </ul>
                <Link to={"/semesters/add"}>
                    {" "}
                    <button className="semester-button">
                        {" "}
                        Add Semester
                    </button>{" "}
                </Link>
            </div>
        );
    }
}

class SemesterRow extends React.Component {
    render() {
        const semesterId = this.props.semesterObj.semesterId;
        const title = this.props.semesterObj.title;
        const url = "/semester/" + "/edit";
        return (
            <ul className="semester-lists">
                <li className="li-large"> {semesterId} </li>
                <li className="li-large"> {title} </li>
                <Link to={url}> Edit </Link>
            </ul>
        );
    }
}
