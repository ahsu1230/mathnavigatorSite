"use strict";
require("./modal.sass");
import React from "react";
const classnames = require("classnames");

/*
 * The wrapper component for all modals.
 * This component handles the pop-up modal as well as the dismissing functionality.
 * Pass HTML / Component as "body content" props to this component to render a full modal.
 *
 * Available props for this Component:
 *
 * - show - Can be true or false. Gives control to invoker of when to open/close modal.
 * - modalContent - HTML/react component of what the modal contains.
 * - modalClassName - class name for the "modal" wrapper itself to allow CSS customization
 * - onDismiss - function callback for when the modal is closed.
 * - withClose - if true, provide an "X" button on the top-right corner of the modal to allow closing the modal.
 * - persistent - if true, prevent user from dismissing modal by clicking on overlay.
 */
export class Modal extends React.Component {
    handleDismiss = () => {
        if (this.props.onDismiss) {
            this.props.onDismiss();
        }
    };

    render = () => {
        const modalContent = this.props.content;
        const persistent = this.props.persistent || false;
        const withClose = this.props.withClose || false;
        const modalViewClasses = classnames("modal-view", {
            show: this.props.show,
        });
        const modalOverlayClasses = classnames("modal-overlay", {
            show: this.props.show,
        });

        const onClickOverlay = persistent ? undefined : this.handleDismiss;
        var closeButton = <div></div>;
        if (withClose) {
            closeButton = (
                <button className="close-x" onClick={this.handleDismiss}>
                    Close
                </button>
            );
        }

        var modalClasses = classnames("modal", this.props.modalClassName);
        return (
            <div className={modalViewClasses}>
                <div
                    className={modalOverlayClasses}
                    onClick={onClickOverlay}></div>
                <div className={modalClasses}>
                    {closeButton}
                    {modalContent}
                </div>
            </div>
        );
    };
}
