"use strict";
require("./afh.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { Link } from "react-router-dom";

export class AskForHelpPage extends React.Component {
    state = {
        list: [],
    };

    componentDidMount() {
        API.get("api/askforhelp/all").then((res) => {
            const afh = res.data;
            this.setState({
                list: afh,
            });
        });
    }

    render() {
        const rows = this.state.list.map((row, index) => {
            return <AFHRow row={row} key={index} />;
        });

        return (
            <div id="view-afh">
                <h1>All Ask For Help Sessions</h1>

                <ul id="header">
                    <li className="li-small">Date</li>
                    <li className="li-med">Time</li>
                    <li className="li-med">Subject</li>
                    <li className="li-med">Title</li>
                    <li className="li-med">LocationId</li>
                    <li className="li-large">Notes</li>
                    <span className="edit-spacer"></span>
                </ul>
                {rows}
                <Link id="add-class" to={"/afh/add"}>
                    <button>Add Ask For Help</button>
                </Link>
            </div>
        );
    }
}

class AFHRow extends React.Component {
    render() {
        const row = this.props.row;
        const url = "/afh/" + row.id + "/edit";
        const date = moment(row.date);
        return (
            <ul id="afh-row">
                <li className="li-small">{date.format("M/D/YYYY")}</li>
                <li className="li-med">{row.timeString}</li>
                <li className="li-med">{row.subject}</li>
                <li className="li-med">{row.title}</li>
                <li className="li-med">{row.locationId}</li>
                <li className="li-large">{row.notes}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
