"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";
import { isNil } from "lodash";

/*
 * The wrapper component for a regular all-page row card.
 * This component handles formatting a card's title, "Edit" link, a list of fields and a list of long texts.
 *
 * Available props for this Component:
 *
 * - title - The main title of the card (required)
 * - subtitle - The subtitle (optional) for the card
 * - editUrl - The link url when the user clicks on "Edit". Can be omitted.
 * - fields - A list of objects. Each object contains a "label" and a "value". Example:
 * [
 *     { label: "FieldA", value: "asdf"},
 *     { label: "FieldB", value: "zxcv" }
 * ]
 * NOTE * In addition, every object can also have a `highlightFn` which is a boolean function.
 * If the function returns true, the whole field-value will be highlighted.
 * Example:
 * {
 *     label: "FieldA",
 *     value: "special",
 *     highlightFn: () => {...... return true; }
 * }
 *
 * - texts - A list of objects. Each object contains a "label" and a "value". Example:
 * [
 *     { label: "MessageA", value: "asdf"},
 *     { label: "MessageB", value: "zxcv" }
 * ]
 */
export default class RowCardBasic extends React.Component {
    renderFields = (fields) => {
        return fields
            .filter((obj) => {
                return !isNil(obj.value);
            })
            .map((obj, index) => {
                const highlight = obj.highlightFn ? !!obj.highlightFn() : false;
                return (
                    <div
                        className={
                            "row-field-wrapper" +
                            (highlight ? " highlighted" : "")
                        }
                        key={index}>
                        <span className="label">{obj.label}:</span>
                        <span className="value">{obj.value}</span>
                    </div>
                );
            });
    };

    renderTexts = (texts) => {
        return texts
            .filter((textObj) => {
                return !isNil(textObj.value);
            })
            .map((textObj, index) => {
                return (
                    <div className="row-text-wrapper" key={index}>
                        <span>{textObj.label}:</span>
                        <p>{textObj.value}</p>
                    </div>
                );
            });
    };

    renderTitle() {
        const subtitle = this.props.subtitle
            ? " (" + this.props.subtitle + ")"
            : "";
        return <h2>{this.props.title + subtitle}</h2>;
    }

    render() {
        const fields = this.props.fields || [];
        const texts = this.props.texts || [];
        const rowFields = this.renderFields(fields);
        const rowTexts = this.renderTexts(texts);
        const title = this.renderTitle();

        return (
            <div className="row-card">
                {this.props.editUrl && (
                    <Link className="edit-url" to={this.props.editUrl}>
                        Edit >
                    </Link>
                )}
                {title}
                {rowFields}
                {rowTexts}
            </div>
        );
    }
}
