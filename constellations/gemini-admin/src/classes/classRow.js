"use strict";
require("./classRow.styl");
import React from "react";
import moment from "moment";
import { Link } from "react-router-dom";

export class ClassRow extends React.Component {
    formatTimes = (times) => {
        return times.replace("/,/g", "\nTest");
    };

    render() {
        const row = this.props.row;
        const unpublished = !row.publishedAt;
        const selected = this.props.selected;
        const publishedState = unpublished ? "Unpublished" : "Published";

        return (
            <div id={unpublished ? "unpublished" : ""}>
                {renderCheckbox(
                    row.classId,
                    unpublished,
                    selected,
                    this.props.onSelectRow
                )}
                <span className="small">{publishedState}</span>
                <span className="large">{row.classId}</span>
                <span className="small">{row.locationId}</span>
                <span className="medium">
                    {moment(row.startDate).format("M/D/YYYY")}
                    {" - "}
                    {moment(row.endDate).format("M/D/YYYY")}
                </span>
                <span className="large">{this.formatTimes(row.times)}</span>
                <Link className="edit" to={"/class/" + row.classId + "/edit"}>
                    Edit
                </Link>
            </div>
        );
    }
}

function renderCheckbox(classId, unpublished, selected, onSelectRow) {
    if (unpublished) {
        return (
            <input
                className="checkbox"
                type="checkbox"
                checked={selected}
                onChange={() => onSelectRow(classId, selected)}
            />
        );
    } else {
        return <span className="checkbox"></span>;
    }
}
