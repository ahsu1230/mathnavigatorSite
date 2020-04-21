"use strict";
require("./classAll.styl");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { Link } from "react-router-dom";

export class ClassAllPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/classes/v1/all").then((res) => {
            const classList = res.data;
            this.setState({ list: classList });
        });
    }

    render() {
        const rows = this.state.list.map((classObj, index) => {
            return (
                <li key={index}>
                    <ClassRow key={index} classObj={classObj} />
                </li>
            );
        });
        const numRows = rows.length;
        return (
            <div id="view-classes">
                <h1>All Classes ({numRows}) </h1>

                <ul className="class-list-row subheader">
                    <li className="li-med">ClassId</li>
                    <li className="li-small">LocationId</li>
                    <li className="li-small">Dates</li>
                    <li className="li-large">Times</li>
                </ul>

                <ul id="class-list">{rows}</ul>
                <Link to={"/classes/add"}>
                    <button className="class-button">Add Class</button>
                </Link>
            </div>
        );
    }
}

class ClassRow extends React.Component {
    render() {
        const classObj = this.props.classObj;
        const classId = classObj.classId;

        const url = "/classes/" + classId + "/edit";
        return (
            <ul className="class-list-row">
                <li className="li-med">{classId}</li>
                <li className="li-small">{classObj.locId}</li>
                <li className="li-small">
                    <div>{moment(classObj.startDate).format("M/D/YYYY")}</div>
                    <div>{moment(classObj.endDate).format("M/D/YYYY")}</div>
                </li>
                <li className="li-large">{classObj.times}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
