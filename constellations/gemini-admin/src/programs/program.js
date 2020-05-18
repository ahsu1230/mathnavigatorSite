"use strict";
require("./program.styl");
import React from "react";
import { Link } from "react-router-dom";
import { keys, size } from "lodash";
import API from "../api.js";
import { ProgramRow } from "./programRow.js";

export class ProgramPage extends React.Component {
    state = {
        programs: [],
        selectedIds: {},
        numUnpublished: 0
    }

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        API.get("api/programs/all").then((res) => {
            const programs = res.data;
            const numUnpublished = programs.filter(p => !p.publishedAt).length;
            this.setState({ 
                programs: programs,
                selectedIds: {},
                numUnpublished: numUnpublished
            });
        });
    }

    onSelectRow = (programId, currentlySelected) => {
        if (currentlySelected) {
            delete this.state.selectedIds[programId];
            this.setState({
                selectedIds: this.state.selectedIds
            });
        } else {
            this.state.selectedIds[programId] = true;
            this.setState({
                selectedIds: this.state.selectedIds
            });
        }
    }

    onClickSelectAll = () => {
        this.state.programs.forEach((program) => {
            if (!program.publishedAt) {
                this.onSelectRow(program.programId, false);
            }
        });
    }

    onClickPublish = () => {
        const publishList = keys(this.state.selectedIds);
        console.log("publishing...");
        const successCallback = (res) => {
            console.log("Successfully published programs!");
            this.fetchData();
        };
        const failCallback = (err) => {
            console.log("Publish failed. " + err);
        };
        API.post("api/programs/publish", publishList)
            .then((res) => successCallback(res))
            .catch((err) => failCallback(err));
    }

    render() {
        const rows = this.state.programs.map((row, index) => {
            const isSelected = !!this.state.selectedIds[row.programId];
            return <ProgramRow key={index} row={row} onSelectRow={this.onSelectRow} selected={isSelected} numUnpublished={this.state.numUnpublished}/>;
        });
        const numRows = rows.length;
        let numUnpublished = this.state.numUnpublished;
        let numSelected = size(this.state.selectedIds);

        return (
            <div id="view-program">
                <h1>All Programs ({numRows}) </h1>
                <ul id="list-heading">
                    {renderSelectAllButton(numUnpublished, this.onClickSelectAll)}
                    <li className="li-med">State</li>
                    <li className="li-med">ProgramKey</li>
                    <li className="li-med">Name</li>
                    <li className="li-small">Grade1</li>
                    <li className="li-small">Grade2</li>
                </ul>
                <ul id="list-rows">{rows}</ul>
                <div id="list-buttons">
                    <button>
                        <Link className="add-program" to={"/programs/add"}>
                            Add Program
                        </Link>
                    </button>
                    {renderUnpublishedButtonSection(numUnpublished, numSelected, this.onClickPublish)}
                </div>
            </div>
        );
    }
}

function renderUnpublishedButtonSection(numUnpublished, numSelected, onClickPublish) {
    if (numUnpublished > 0) {
        return (
            <div className="publish">
                <p>
                    You have {numUnpublished} unpublished items. <br />
                    You have selected {numSelected} items to publish.
                </p>
                <button
                    onClick={onClickPublish}>
                    Publish Selected
                </button>
            </div>      
        );
    } else {
        return (<div></div>);
    }
}

function renderSelectAllButton(numUnpublished, onClickSelectAll) {
    if (numUnpublished > 0) {
        return (
            <button
                className="li-checkbox"
                onClick={onClickSelectAll}>
                Select<br/>All
            </button>
        );
    } else {
        return (<div></div>);
    }
}