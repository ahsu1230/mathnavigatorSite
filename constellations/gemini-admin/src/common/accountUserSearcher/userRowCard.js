"use strict";
require("./accountUserSearcher.sass");
import React from "react";
import moment from "moment";
import { getFullName } from "../userUtils.js";
import RowCardColumns from "../rowCards/rowCardColumns.js";

export class UserRowCard extends React.Component {
    render() {
        const user = this.props.user;
        const editTitle = this.props.editTitle;
        const editUrl = this.props.editUrl;
        const firstColumn = [
            { label: "Id", value: user.id },
            {
                label: "Status",
                value: user.isGuardian ? "Guardian" : "Student",
            },
            { label: "Last Updated", value: moment(user.updatedAt).fromNow() },
        ];
        const secondColumn = [
            { label: "Phone", value: user.phone },
            { label: "School", value: user.school },
            { label: "GraduationYear", value: user.graduationYear },
        ];
        const texts = [
            {
                label: "Notes about user",
                value: user.notes,
            },
        ];
        return (
            <div className="user-row-card">
                <RowCardColumns
                    title={getFullName(user)}
                    subtitle={user.email}
                    editTitle={editTitle}
                    editUrl={editUrl}
                    fieldsList={[firstColumn, secondColumn]}
                    texts={texts}
                />
            </div>
        );
    }
}
