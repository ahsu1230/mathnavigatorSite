"use strict";
import React from "react";
import moment from "moment";
import RowCardBasic from "./rowCardBasic.js";

/**
 * A utility component to depict an Account's information.
 * Parameters:
 * - account: the account object
 * - editTitle: (optional) the display name of the Link
 * - editUrl: (optional) the url of the Link
 */
export class AccountRowCard extends React.Component {
    render() {
        const account = this.props.account || {};
        const editTitle = this.props.editTitle || "View Account Details";
        const editUrl = this.props.editUrl || "/account/" + account.id;
        const fields = [
            {
                label: "Primary Contact",
                value: account.primaryEmail,
            },
            {
                label: "Account Created",
                value: moment(account.createdAt).format("l"),
            },
            {
                label: "Last Updated",
                value: moment(account.updatedAt).fromNow(),
            },
        ];

        return (
            <div className="user-row-card">
                <RowCardBasic
                    title={"AccountId: " + account.id}
                    editTitle={editTitle}
                    editUrl={editUrl}
                    fields={fields}
                />
            </div>
        );
    }
}
