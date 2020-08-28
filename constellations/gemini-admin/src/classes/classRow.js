"use strict";
require("./classRow.sass");
import React from "react";
import { Link } from "react-router-dom";
import { formatCurrency } from "../utils/userUtils.js";

export class ClassRow extends React.Component {
    render = () => {
        const classObj = this.props.classObj;
        const isUnpublished = this.props.isUnpublished;
        const url = "/classes/" + classObj.classId + "/edit";

        const price = classObj.pricePerSession
            ? formatCurrency(classObj.pricePerSession) + " per session"
            : formatCurrency(classObj.priceLump) + " total";

        return (
            <div id={isUnpublished ? "unpublished" : ""} className="row">
                {renderCheckbox(
                    classObj.classId,
                    this.props.isCollapsed,
                    isUnpublished,
                    this.props.isSelected,
                    this.props.onSelectRow
                )}
                <span className="small-column">
                    {isUnpublished ? "Unpublished" : "Published"}
                </span>
                <span className="large-column">{classObj.classId}</span>
                <span className="medium-column">{classObj.locationId}</span>
                <span className="large-column">{classObj.times}</span>
                <span className="medium-column">{price}</span>
                <span className="edit">
                    <Link to={url}>{"Edit >"}</Link>
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
