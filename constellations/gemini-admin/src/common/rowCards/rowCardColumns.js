"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";
import RowCardBasic from "./rowCardBasic.js";

/*
 * The wrapper component for an all-page row card but with multiple columns in one card.
 * This component handles formatting a card's title, "Edit" link, a list of columns of fields and a list of long texts.
 * This component extends from `RowCardBasic`
 *
 * Available props for this Component:
 *
 * - title - The main title of the card (required)
 * - subtitle - The subtitle (optional) for the card
 * - editTitle - The link display name. Default value is "Edit". Can be omitted.
 * - editUrl - The link url when the user clicks on "Edit". Can be omitted.
 * - fieldsList - A list of list of objects. Every inner-list represents a column.
 * And inside each inner-list is an object that contains a "label" and value (just like RowCardBasic).
 * Example:
 * fieldsList =
 * [ // First column
 *     { label: "FieldA", value: "asdf"},
 *     { label: "FieldB", value: "zxcv" }
 * ],
 * [ // Second column
 *     { label: "FieldC", value: "qwer"},
 * ]
 * note: in addition, highlightFn() is also valid for each of these objects just like in RowCardBasic.
 *
 * - texts - A list of objects. Each object contains a "label" and a "value". Example:
 * [
 *     { label: "MessageA", value: "asdf"},
 *     { label: "MessageB", value: "zxcv" }
 * ]
 */

export default class RowCardColumns extends RowCardBasic {
    render() {
        const title = this.renderTitle();
        const fields = this.props.fieldsList || [];
        const texts = this.props.texts || [];
        const rowTexts = this.renderTexts(texts);
        const editTitle = this.props.editTitle || "Edit >";

        const rowColumns = fields.map((column, index) => {
            const rowFields = this.renderFields(column);
            return (
                <div className="column" key={index}>
                    {rowFields}
                </div>
            );
        });

        return (
            <div className="row-card">
                {this.props.editUrl && (
                    <Link className="edit-url" to={this.props.editUrl}>
                        {editTitle}
                    </Link>
                )}
                {title}
                <div className="column-wrapper">{rowColumns}</div>
                {rowTexts}
            </div>
        );
    }
}
