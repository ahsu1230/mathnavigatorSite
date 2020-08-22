"use strict";
require("./modal.sass");
import React from "react";
import srcClose from "../../assets/close_black.svg";

const classnames = require("classnames");

/*
 * Props for the Modal Component:
 *
 * content: This is displayed inside the modal.
 *
 * modalClassName: This prop allows the className of the div containing the content to be set.
 *
 * show: This is the boolean variable that determines if the modal is shown or not.
 *
 * persistent: If true, the modal can only be dismissed by clicking the close button.
 * If false (omit the prop), the modal can also be dismissed by clicking outside the modal.
 *
 * withClose: This is the boolean variable that determines if there is a close button.
 *
 * onDismiss: This is run when the modal is dismissed, either through clicking the close button
 * or clicking outside the modal if persistent is false.
 *
 * Example:
 *     <Modal
 *         content={modalContent}
 *         show={this.state.showModal}
 *         withClose={true}
 *         onDismiss={() => this.setState({ showModal: false })}
 *     />
 */
export class Modal extends React.Component {
    handleDismiss = () => {
        if (this.props.onDismiss) this.props.onDismiss();
    };

    render = () => {
        const modalContent = this.props.content;
        const persistent = this.props.persistent;
        const withClose = this.props.withClose;
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
                    <img src={srcClose} />
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
