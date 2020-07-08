"use strict";
require("./classRow.sass");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class ClassRow extends React.Component {
    render = () => {
        const row = this.props.row;
        const isUnpublished = this.props.isUnpublished;

        return (
            <div id={isUnpublished ? "unpublished" : ""} className="row">
                {renderCheckbox(
                    row.classId,
                    this.props.isCollapsed,
                    isUnpublished,
                    this.props.isSelected,
                    this.props.onSelectRow
                )}
                <span className="column">
                    {isUnpublished ? "Unpublished" : "Published"}
                </span>
                <span className="large-column">{row.classId}</span>
                <span className="column">{row.locationId}</span>
                <span className="medium-column">
                    {moment(row.startDate).format("M/D/YYYY")}
                    {" - "}
                    {moment(row.endDate).format("M/D/YYYY")}
                </span>
                <span className="large-column">{row.times}</span>
                <span className="edit">
                    <Link to={"/classes/" + row.classId + "/edit"}>
                        {"Edit >"}
                    </Link>
                </span>
            </div>
        );
    };
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
        return <div></div>;
    } else if (isUnpublished) {
        return (
            <input
                className="select"
                type="checkbox"
                checked={isSelected}
                onChange={() => onSelectRow(classId, isSelected)}
            />
        );
    } else {
        // Keeps the spacing
        return <div className="select"></div>;
    }
}
