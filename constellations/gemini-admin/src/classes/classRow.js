"use strict";
require("./classRow.styl");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class ClassRow extends React.Component {
    render() {
        const row = this.props.row;
        const isUnpublished = this.props.isUnpublished;

        return (
            <div id={isUnpublished ? "unpublished" : ""}>
                {renderCheckbox(
                    row.classId,
                    this.props.isCollapsed,
                    isUnpublished,
                    this.props.isSelected,
                    this.props.onSelectRow
                )}
                <span className="small">
                    {isUnpublished ? "Unpublished" : "Published"}
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
    isCollapsed,
    isUnpublished,
    isSelected,
    onSelectRow
) {
    if (isCollapsed) {
        // No space
        return <span></span>;
    } else if (isUnpublished) {
        return (
            <input
                className="checkbox"
                type="checkbox"
                checked={isSelected}
                onChange={() => onSelectRow(classId, isSelected)}
            />
        );
    } else {
        // Keeps the spacing
        return <span className="checkbox"></span>;
    }
}
