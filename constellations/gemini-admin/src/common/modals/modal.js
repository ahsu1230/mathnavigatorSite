"use strict";
require("./modal.sass");
import React from "react";
const classnames = require("classnames");

/*
 * The wrapper component around a modalContent.
 * Pass HTML / Component as "content" props to this component to render a modal,
 * including the dismissing functionality.
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
