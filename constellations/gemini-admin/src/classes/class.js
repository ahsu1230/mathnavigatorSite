"use strict";
require("./class.sass");
import React from "react";
import API from "../api.js";
import { keys, size } from "lodash";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardColumns from "../utils/rowCardColumns.js";

export class ClassPage extends React.Component {
    state = {
        classes: [],
        fullStates: [],
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

        API.get("api/classes/full-states").then((res) => {
            this.setState({ fullStates: res.data });
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
        // If everything is selected, the SelectAll button deselects everything
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
        console.log("Publishing Classes ...");

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
                    Select All Unpublished Classes
                </button>
            );
        } else {
            return <div></div>;
        }
    };

    renderPublishButtonSection = () => {
        const numUnpublished = this.state.numUnpublished;
        const numSelected = size(this.state.selectedIds);

        let publish = <div></div>;
        if (numSelected > 0) {
            publish = (
                <button onClick={this.onClickPublish}>
                    Publish Selected Classes
                </button>
            );
        }

        if (numUnpublished > 0) {
            const firstWord = numUnpublished == 1 ? "class" : "classes";
            const secondWord = numSelected == 1 ? "class" : "classes";
            return (
                <div id="publish">
                    <p>
                        You have {numUnpublished} unpublished {firstWord}.
                        <br />
                        You have selected {numSelected} {secondWord} to publish.
                    </p>
                    {publish}
                </div>
            );
        } else {
            return <div id="publish"></div>;
        }
    };

    render = () => {
        const rows = this.state.classes.map((classObj, index) => {
            const isSelected = !!this.state.selectedIds[classObj.classId];
            const checkbox = renderCheckbox(
                classObj,
                this.state.numUnpublished,
                isSelected,
                this.onSelectRow
            );
            const fields = generateFields(classObj, this.state.fullStates);
            const texts = generateTexts(classObj);
            return (
                <div className="card-wrapper" key={index}>
                    {checkbox}
                    <RowCardColumns
                        title={"ClassId"}
                        subtitle={classObj.classId}
                        editUrl={"/classes/" + classObj.classId + "/edit"}
                        fieldsList={fields}
                        texts={texts}
                    />
                </div>
            );
        });

        return (
            <div id="view-class">
                <AllPageHeader
                    title={"All Classes (" + this.state.classes.length + ")"}
                    addUrl={"/classes/add"}
                    addButtonTitle={"Add Class"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />
                {this.renderSelectAllButton()}
                {this.renderPublishButtonSection()}

                <div className="cards">{rows}</div>
            </div>
        );
    };
}

function generateFields(classObj, fullStates) {
    const column1 = [
        {
            label: "ProgramId",
            value: classObj.programId,
        },
        {
            label: "SemesterId",
            value: classObj.semesterId,
        },
        {
            label: "LocationId",
            value: classObj.locationId,
        },
        {
            label: "Time",
            value: classObj.timesStr,
        },
    ];
    const column2 = [
        {
            label: "Published",
            value: !!classObj.publishedAt ? "ok" : "unpublished",
            highlightFn: () => !classObj.publishedAt,
        },
        {
            label: "Full State",
            value: fullStates[classObj.fullState],
            highlightFn: () => classObj.fullState != 0,
        },
        {
            label: "Price Lump Sum",
            value: classObj.priceLumpSum,
        },
        {
            label: "Price Per Session",
            value: classObj.pricePerSession,
        },
        {
            label: "Google Classroom",
            value: classObj.googleClassCode,
        },
    ];
    return [column1, column2];
}

function generateTexts(classObj) {
    return [
        {
            label: "Payment Notes",
            value: classObj.paymentNotes,
        },
    ];
}

function renderCheckbox(classObj, numUnpublished, isSelected, onSelectRow) {
    const isUnpublished = !classObj.publishedAt;
    if (numUnpublished == 0) {
        // collapse, no spacing
        return <div></div>;
    } else if (isUnpublished) {
        return (
            <input
                className="select"
                type="checkbox"
                checked={isSelected}
                onChange={() => onSelectRow(classObj.classId, isSelected)}
            />
        );
    } else {
        // Keep the spacing for checkbox in other rows
        return <div className="select"></div>;
    }
}
