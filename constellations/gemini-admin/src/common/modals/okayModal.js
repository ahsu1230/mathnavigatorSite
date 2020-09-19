"use strict";
require("./okayModal.sass");
import React from "react";

/*
 * The content of a modal with a message and one button (Okay).
 * This button may have a callback function you may pass in as props,
 * which will be invoked when the user clicks on the button.
 */
export default class OkayModal extends React.Component {
    render = () => {
        const text = this.props.text;
        const onOkay = this.props.onOkay;
        return (
            <div id="modal-view-okay">
                <p>{text}</p>
                <button onClick={onOkay}>OK</button>
            </div>
        );
    };
}
