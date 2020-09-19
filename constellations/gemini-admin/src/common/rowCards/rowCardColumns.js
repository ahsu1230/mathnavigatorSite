"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";
import RowCardBasic from "./rowCardBasic.js";

// Given a title, subtitle, and a list of {fields, text}
// For example, title = Class, subtitle = (ap_calc_2020),
// list of list of fields objects. [ [{}, {}, {}], [fields2], [fields3] ] <- will be presented as columns
// list of texts

export default class RowCardColumns extends RowCardBasic {
    render() {
        const title = this.renderTitle();
        const fields = this.props.fieldsList || [];
        const texts = this.props.texts || [];
        const rowTexts = this.renderTexts(texts);

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
                <Link className="edit-url" to={this.props.editUrl}>
                    Edit >
                </Link>
                {title}
                <div className="column-wrapper">{rowColumns}</div>
                {rowTexts}
            </div>
        );
    }
}
