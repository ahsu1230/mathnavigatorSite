"use strict";
require("./class.styl");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";
import { keys, size } from "lodash";
import { ClassRow } from "./classRow.js";

export class ClassPage extends React.Component {
    state = {
        classes: [],
        selectedIds: {},
        numUnpublished: 0,
    };

    componentDidMount = () => {
        this.fetchData();
    };

    fetchData = () => {
        API.get("api/classes/all").then((res) => {
            const classes = res.data;
            const numUnpublished = classes.filter((c) => !c.publishedAt).length;
            this.setState({
                classes: classes,
                selectedIds: {},
                numUnpublished: numUnpublished,
            });
        });
    };

    onSelectRow = (classId, selected) => {
        // Switches the checkbox state
        if (selected) {
            delete this.state.selectedIds[classId];
            this.setState({
                selectedIds: this.state.selectedIds,
            });
        } else {
            this.state.selectedIds[classId] = true;
            this.setState({
                selectedIds: this.state.selectedIds,
            });
        }
    };

    onClickSelectAll = () => {
        // If everything is selected, the Select All button deselects everything
        if (size(this.state.selectedIds) == this.state.numUnpublished) {
            this.state.classes.forEach((c) => {
                this.onSelectRow(c.classId, true);
            });
        } else {
            this.state.classes.forEach((c) => {
                if (!c.publishedAt) {
                    this.onSelectRow(c.classId, false);
                }
            });
        }
    };

    onClickPublish = () => {
        const publishList = keys(this.state.selectedIds);
        console.log("... Publishing Classes ...");

        API.post("api/classes/publish", publishList)
            .then(() => {
                console.log("Successfully published classes!");
                this.fetchData();
            })
            .catch((err) => {
                window.alert("Publish failed: " + err);
            });
    };

    renderSelectAllButton = () => {
        if (this.state.numUnpublished > 0) {
            return (
                <button id="select-all" onClick={this.onClickSelectAll}>
                    Select
                    <br />
                    All
                </button>
            );
        } else {
            return <div></div>;
        }
    };

    renderPublishButtonSection = () => {
        const numUnpublished = this.state.numUnpublished;
        const selected = this.state.selected;

        let publish = <div></div>;
        if (selected > 0) {
            publish = (
                <button onClick={this.onClickPublish}>Publish Selected</button>
            );
        }

        if (numUnpublished > 0) {
            // Use the correct word
            const firstWord = numUnpublished == 1 ? "class" : "classes";
            const secondWord = selected == 1 ? "class" : "classes";
            return (
                <div id="publish">
                    <p>
                        You have {numUnpublished} unpublished {firstWord}.{" "}
                        <br />
                        You have selected {selected} {secondWord} to publish.
                    </p>
                    {publish}
                </div>
            );
        } else {
            return <p></p>;
        }
    };

    render = () => {
        const rows = this.state.classes.map((row, index) => {
            const isSelected = !!this.state.selectedIds[row.classId];
            return (
                <li key={index}>
                    <ClassRow
                        row={row}
                        isCollapsed={this.state.numUnpublished == 0}
                        isUnpublished={!row.publishedAt}
                        isSelected={isSelected}
                        onSelectRow={this.onSelectRow}
                    />
                </li>
            );
        });

        let count = rows.length;
        let numUnpublished = this.state.numUnpublished;
        let selected = size(this.state.selectedIds);

        return (
            <div id="view-class">
                <h1>All Classes ({count}) </h1>

                <section id="class-rows">
                    <div id="header">
                        {this.renderSelectAllButton()}
                        <span className="small">State</span>
                        <span className="large">ClassId</span>
                        <span className="small">LocationId</span>
                        <span className="medium">StartDate - EndDate</span>
                        <span className="large">Times</span>
                    </div>
                    <ul id="rows">{rows}</ul>
                </section>

                <section id="footer">
                    <button>
                        <Link id="add-class" to={"/classes/add"}>
                            Add Class
                        </Link>
                    </button>
                    {this.renderPublishButtonSection()}
                </section>
            </div>
        );
    };
}
