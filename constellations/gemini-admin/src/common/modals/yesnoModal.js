"use strict";
require("./yesnoModal.sass");
import React from "react";

/*
 * The content of a modal with a message and two buttons (Yes/No or Accept/Reject).
 * Both buttons have a callback function you may pass in as props,
 * which will be invoked when the user clicks on one of those buttons.
 */
export default class YesNoModal extends React.Component {
    render = () => {
        const text = this.props.text;
        const rejectText = this.props.rejectText || "No";
        const acceptText = this.props.acceptText || "Yes";
        const onReject = this.props.onReject;
        const onAccept = this.props.onAccept;
        return (
            <div id="modal-view-yesno">
                <p>{text}</p>
                <button className="reject" onClick={onReject}>
                    {rejectText}
                </button>
                <button className="accept" onClick={onAccept}>
                    {acceptText}
                </button>
            </div>
        );
    };
}
