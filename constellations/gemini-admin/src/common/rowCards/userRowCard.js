"use strict";
import React from "react";
import moment from "moment";
import { getFullName } from "../userUtils.js";
import RowCardColumns from "./rowCardColumns.js";

/**
 * A utility component to depict a User's information.
 * Parameters:
 * - user: the user object
 * - editTitle: (optional) the display name of the Link
 * - editUrl: (optional) the url of the Link
 * - account: (optional) if present, display relevant account information
 */
export class UserRowCard extends React.Component {
    render() {
        const user = this.props.user;
        const account = this.props.account || {};
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
            { label: "AccountId", value: account.id },
            { label: "Account Primary Contact", value: account.primaryEmail },
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
