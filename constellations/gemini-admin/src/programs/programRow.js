"use strict";
require("./programRow.styl");
import React from "react";
import { Link } from "react-router-dom";

export class ProgramRow extends React.Component {
    render() {
        const row = this.props.row;
        const numUnpublished = this.props.numUnpublished;
        const isSelected = this.props.selected;

        const url = "/program/" + row.programId + "/edit";
        const isUnpublished = !row.publishedAt;
        const checkbox = renderCheckbox(
            row.programId,
            isUnpublished,
            isSelected,
            numUnpublished,
            this.props.onSelectRow
        );
        const publishedState = isUnpublished ? "Unpublished" : "Published";

        return (
            <li className={isUnpublished ? "unpublished" : ""}>
                {checkbox}
                <div className="li-med">{publishedState}</div>
                <div className="li-med">{row.programId}</div>
                <div className="li-med">{row.name}</div>
                <div className="li-small">{row.grade1}</div>
                <div className="li-small">{row.grade2}</div>
                <Link to={url}>Edit</Link>
            </li>
        );
    }
}

function renderCheckbox(
    programId,
    isUnpublished,
    currentlySelected,
    numUnpublished,
    onSelectRow
) {
    if (numUnpublished > 0) {
        if (isUnpublished) {
            return (
                <input
                    className="li-checkbox"
                    type="checkbox"
                    name="unpublished"
                    checked={currentlySelected}
                    onChange={() => onSelectRow(programId, currentlySelected)}
                />
            );
        } else {
            return <div className="li-checkbox"></div>;
        }
    } else {
        return <div></div>;
    }
}
