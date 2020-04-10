"use strict";
require("./semester.styl");
import React from "react";
import { Link } from "react-router-dom";
import API from "../api.js";

export class SemesterPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/semesters/v1/all").then((res) => {
            const semesters = res.data;
            this.setState({ list: semesters });
        });
    }

    render() {
        const numSemesters = this.state.list.length;
        const rows = this.state.list.map((s, i) => (
            <SemesterRow key={i} semesterObj={s} />
        ));
        return (
            <div id="view-semester">
                <h1>All Semesters ({numSemesters})</h1>
                <ul className="semester-lists">
                    <li className="li-med">Semester ID</li>
                    <li className="li-med">Title</li>
                </ul>
                <ul>{rows}</ul>
                <Link to={"/semesters/add"}>
                    <button className="semester-button">Add Semester</button>
                </Link>
            </div>
        );
    }
}

class SemesterRow extends React.Component {
    render() {
        const semesterId = this.props.semesterObj.semesterId;
        const title = this.props.semesterObj.title;
        const url = "/semesters/" + semesterId + "/edit";
        return (
            <ul className="semester-lists">
                <li className="li-med"> {semesterId} </li>
                <li className="li-med"> {title} </li>
                <Link to={url}> Edit </Link>
            </ul>
        );
    }
}
