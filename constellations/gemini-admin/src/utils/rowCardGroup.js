"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";
import RowCardBasic from "../utils/rowCardBasic.js";

// Given a title, subtitle, and a list of {fields, text}
// For example, title = Achievements, subtitle = (2020), list of [{position, message}, {position, message}, ...]

export default class RowCardGroup extends RowCardBasic {
    render() {
        const title = this.renderTitle();
        const groupList = this.props.groupList || [];

        const rowGroup = groupList.map((obj, index) => {
            const fields = obj.fields || [];
            const texts = obj.texts || [];
            const rowFields = this.renderFields(fields);
            const rowTexts = this.renderTexts(texts);
            return (
                <div className="group" key={index}>
                    <Link className="edit-url" to={obj.editUrl}>
                        Edit >
                    </Link>
                    {rowFields}
                    {rowTexts}
                </div>
            );
        });

        return (
            <div className="row-card">
                {title}
                {rowGroup}
            </div>
        );
    }
}
