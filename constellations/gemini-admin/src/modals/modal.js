"use strict";
require("./modal.sass");
import React from "react";
const classnames = require("classnames");

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
        var closeButton;
        if (withClose) {
            closeButton = (
                <button className="close-x" onClick={this.handleDismiss}>
                    Close
                </button>
            );
        } else {
            <div></div>;
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
