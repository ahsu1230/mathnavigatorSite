"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";

export default class RowCardBasic extends React.Component {
    renderFields = (fields) => {
        return fields
            .filter((obj) => {
                return !_.isNil(obj.value);
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
                        <span>{obj.label}:</span>
                        <span>{obj.value}</span>
                    </div>
                );
            });
    };

    renderTexts = (texts) => {
        return texts
            .filter((textObj) => {
                return !_.isNil(textObj.value);
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
                <Link className="edit-url" to={this.props.editUrl}>
                    Edit
                </Link>
                {title}
                {rowFields}
                {rowTexts}
            </div>
        );
    }
}
