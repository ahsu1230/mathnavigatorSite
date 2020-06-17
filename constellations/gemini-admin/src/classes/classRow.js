"use strict";
require("./classRow.styl");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class ClassRow extends React.Component {
    render() {
        const row = this.props.row;
        const unpublished = this.props.unpublished;

        return (
            <div id={unpublished ? "unpublished" : ""}>
                {renderCheckbox(
                    row.classId,
                    this.props.collapsed,
                    unpublished,
                    this.props.selected,
                    this.props.onSelectRow
                )}
                <span className="small">
                    {unpublished ? "Unpublished" : "Published"}
                </span>
                <span className="large">{row.classId}</span>
                <span className="small">{row.locationId}</span>
                <span className="medium">
                    {moment(row.startDate).format("M/D/YYYY")}
                    {" - "}
                    {moment(row.endDate).format("M/D/YYYY")}
                </span>
                <span className="large">{row.times}</span>
                <Link to={"/class/" + row.classId + "/edit"}>Edit</Link>
            </div>
        );
    }
}

function renderCheckbox(
    classId,
    collapsed,
    unpublished,
    selected,
    onSelectRow
) {
    if (collapsed) {
        // No space
        return <span></span>;
    } else if (unpublished) {
        return (
            <input
                className="checkbox"
                type="checkbox"
                checked={selected}
                onChange={() => onSelectRow(classId, selected)}
            />
        );
    } else {
        // Keeps the spacing
        return <span className="checkbox"></span>;
    }
}
