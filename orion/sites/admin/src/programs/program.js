"use strict";
require("./program.styl");
import React from "react";
import ReactDOM from "react-dom";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { Link } from "react-router-dom";

export class ProgramPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/programs/v1/all").then((res) => {
            const programs = res.data;
            this.setState({ list: programs });
        });
    }

<<<<<<< HEAD
    onClickSelectAll() {
        var items = document.getElementsByName("unpublished");
        for (var i = 0; i < items.length; i++) {
            if (items[i].type == "checkbox") {
                items[i].checked = true;
            }
        }
    }

    onClickPublish() {
        console.log("clicked publish");
    }

=======
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
    render() {
        const rows = this.state.list.map((row, index) => {
            return <ProgramRow key={index} row={row} />;
        });
        const numRows = rows.length;
<<<<<<< HEAD
        let numUnpublished = 0;
        let numSelected = 0;
        return (
            <div id="view-program">
                <ul>
                    <h1>All Programs ({numRows}) </h1>
                    <p>
                        {" "}
                        You have {numUnpublished} unpublished items. You have
                        selected {numSelected} items to publish.{" "}
                    </p>
                </ul>
                <ul id="list-heading">
                    <button
                        className="li-small"
                        onClick={this.onClickSelectAll}
                    >
                        Select All
                    </button>
=======
        return (
            <div id="view-program">
                <h1>All Programs ({numRows})</h1>
                <ul id="list-heading">
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
                    <li className="li-med">ProgramKey</li>
                    <li className="li-med">Name</li>
                    <li className="li-small">Grade1</li>
                    <li className="li-small">Grade2</li>
                </ul>
                <ul id="list-rows">{rows}</ul>
<<<<<<< HEAD
                <ul id="list-buttons">
                    <div className="li-med">
                        <button>
                            <Link className="add-program" to={"/programs/add"}>
                                Add Program
                            </Link>
                        </button>
                    </div>
                    <div className="li-med">
                        <button
                            className="publish"
                            onClick={this.onClickPublish}
                        >
                            Publish
                        </button>
                    </div>
                </ul>
=======
                <Link className="add-program" to={"/programs/add"}>
                    Add Program
                </Link>
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
            </div>
        );
    }
}

class ProgramRow extends React.Component {
<<<<<<< HEAD
    renderCheckbox(isUnpublished) {
        let checkbox = <div> </div>;
        if (isUnpublished) {
            return (checkbox = (
                <input
                    className="li-small"
                    type="checkbox"
                    name="unpublished"
                    onClick={this.onClickBox}
                />
            ));
        } else {
            return (checkbox = <div className="li-small"></div>);
        }
    }

    render() {
        const row = this.props.row;
        const url = "/program/" + row.programId + "/edit";
        let checkbox = this.renderCheckbox(true);
        return (
            <li className="program-row">
                {checkbox}
=======
    render() {
        const row = this.props.row;
        const url = "/program/" + row.programId + "/edit";
        return (
            <li className="program-row">
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
                <div className="li-med">{row.programId}</div>
                <div className="li-med">{row.name}</div>
                <div className="li-small">{row.grade1}</div>
                <div className="li-small">{row.grade2}</div>
                <Link to={url}>Edit</Link>
            </li>
        );
    }
}
