"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";
import RowCardBasic from "./rowCardBasic.js";

// Given a title, subtitle, and a list of {fields, text}
// For example, title = Achievements, subtitle = (2020), list of [{position, message}, {position, message}, ...]

/*
 * The wrapper component for an all-page row card but with grouped fields.
 * This component handles formatting a card's title, "Edit" link,
 * and a list of fields/texts under a group.
 * This component extends from `RowCardBasic`.
 *
 * Available props for this Component:
 *
 * - title - The main title of the card (required)
 * - subtitle - The subtitle (optional) for the card
 * - editUrl - The link url when the user clicks on "Edit". Can be omitted.
 * - groupList - A list of objects. Every object basically represents the contents of a RowCardBasic
 *               (with many fields and many texts).
 * Example:
 * title = "Announcements on September 20th"
 * groupList = [
 *      {
 *          editUrl: /announcements/1
 *          fields: [...]       // objects with "label" and "value"
 *          texts: [...]        // objects with "label" and "value"
 *      },
 *      {
 *          editUrl: /announcements/2
 *          fields: [...]       // objects with "label" and "value"
 *          texts: [...]        // objects with "label" and "value"
 *      },
 * ]
 */
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
                    {obj.editUrl && (
                        <Link className="edit-url" to={obj.editUrl}>
                            Edit >
                        </Link>
                    )}
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
