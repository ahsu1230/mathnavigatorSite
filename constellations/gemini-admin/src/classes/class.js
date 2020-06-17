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
        unpublished: 0,
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        API.get("api/classes/all").then((res) => {
            const classes = res.data;
            const unpublished = classes.filter((c) => !c.publishedAt).length;
            this.setState({
                classes: classes,
                selectedIds: {},
                unpublished: unpublished,
            });
        });
    }

    onSelectRow = (classId, selected) => {
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
        this.state.classes.forEach((c) => {
            if (!c.publishedAt) {
                this.onSelectRow(c.classId, false);
            }
        });
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
                console.log("Publish failed: " + err);
            });
    };

    render() {
        const rows = this.state.classes.map((row, index) => {
            const isSelected = !!this.state.selectedIds[row.classId];
            return (
                <li key={index}>
                    <ClassRow
                        row={row}
                        onSelectRow={this.onSelectRow}
                        selected={isSelected}
                        unpublished={this.state.unpublished}
                    />
                </li>
            );
        });

        let count = rows.length;
        let unpublished = this.state.unpublished;
        let selected = size(this.state.selectedIds);

        return (
            <div id="view-class">
                <h1>All Classes ({count}) </h1>

                <section id="class-rows">
                    <div id="header">
                        {renderSelectAllButton(
                            unpublished,
                            this.onClickSelectAll
                        )}
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
                    {renderPublishButtonSection(
                        unpublished,
                        selected,
                        this.onClickPublish
                    )}
                </section>
            </div>
        );
    }
}

function renderSelectAllButton(unpublished, onClickSelectAll) {
    if (unpublished > 0) {
        return (
            <button id="select-all" onClick={onClickSelectAll}>
                Select
                <br />
                All
            </button>
        );
    } else {
        return <div></div>;
    }
}

function renderPublishButtonSection(unpublished, selected, onClickPublish) {
    if (unpublished > 0) {
        const firstWord = unpublished == 1 ? "class" : "classes";
        const secondWord = selected == 1 ? "class" : "classes";
        return (
            <div id="publish">
                <p>
                    You have {unpublished} unpublished {firstWord}. <br />
                    You have selected {selected} {secondWord} to publish.
                </p>
                <button onClick={onClickPublish}>Publish Selected</button>
            </div>
        );
    } else {
        return <p></p>;
    }
}
