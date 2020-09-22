"use strict";
require("./editPageWrapper.sass");
import React from "react";
import { Modal } from "../modals/modal.js";
import OkayModal from "../modals/okayModal.js";
import YesNoModal from "../modals/yesnoModal.js";

/*
 * A component that automatically handles the Header and Footer portion of an entity's "Edit-View" page.
 * This includes a title, and the buttons at the bottom of the page.
 * This component will handle the saving and deleting modals, but you must pass in the
 * onSave and onDelete callback functions.
 *
 * Available props for this Component:
 *
 * - title - display name of the header of the page (e.g. Edit Program)
 * - content - the body content HTML / component in between the header and footer
 * - prevPageUrl - the url to return to if the user "cancels" the edit.
 * - onSave - the save callback function to call when the user wants to save their updates.
 * - onDelete - the delete callback function to call when the user confirms deleting the entity.
 * - entityId - the id of the entity. Will be used in the modal messages (e.g. "Program ap_calc successfully saved!")
 * - entityName - the name of the entity. Will be used in the modal message (e.g. "Program successfully saved!")
 */
export default class EditPageWrapper extends React.Component {
    state = {
        showDeleteModal: false,
        showSaveModal: false,
    };

    dismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onClickCancel = () => {
        this.goBackToPrevPage();
    };

    onClickSave = () => {
        this.props
            .onSave()
            .then((res) => {
                this.setState({ showSaveModal: true });
            })
            .catch((err) => {
                this.dismissModal();
                alert(
                    "Error occured when saving: " + err.response.data.message
                );
            });
    };

    persistDelete = () => {
        this.props
            .onDelete()
            .then((res) => {
                this.goBackToPrevPage();
            })
            .catch((err) => {
                alert(
                    "Error occured when deleting: " + err.response.data.message
                );
            })
            .finally(this.dismissModal());
    };

    onSaved = () => {
        this.dismissModal();
        this.goBackToPrevPage();
    };

    goBackToPrevPage = () => {
        window.location.hash = this.props.prevPageUrl;
    };

    createDeleteModalText = () => {
        const entityId = this.props.entityId;
        const entityName = this.props.entityName;
        if (entityId) {
            return `Are you sure you want to delete '${entityId}'?`;
        } else if (entityName) {
            return `Are you sure you want to delete this ${entityName}?`;
        } else {
            return "Are you sure you want to delete?";
        }
    };

    createSaveModalText = () => {
        const entityId = this.props.entityId;
        const entityName = this.props.entityName;
        if (entityId) {
            return `'${entityId}' was successfully saved!`;
        } else if (entityName) {
            return `This ${entityName} was successfully saved!`;
        } else {
            return "Successfully saved!";
        }
    };

    renderModal = () => {
        let modalDiv;
        let modalContent;
        let showModal;
        let modalText;
        if (this.state.showDeleteModal) {
            showModal = this.state.showDeleteModal;
            modalText = this.createDeleteModalText();
            modalContent = (
                <YesNoModal
                    text={modalText}
                    onAccept={this.persistDelete}
                    onReject={this.dismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalText = this.createSaveModalText();
            modalContent = <OkayModal text={modalText} onOkay={this.onSaved} />;
        }
        if (modalContent) {
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showModal}
                    onDismiss={this.dismissModal}
                />
            );
        }
        return modalDiv;
    };

    render() {
        const title = this.props.title || "";
        const isEdit = this.props.isEdit || false;
        const content = this.props.content || <div></div>;
        const modalDiv = this.renderModal();
        return (
            <div className="edit-page-wrapper">
                {modalDiv}
                <h1>{title}</h1>
                {content}
                <div className="btn-footer">
                    <div>
                        <button className="cancel" onClick={this.onClickCancel}>
                            Cancel
                        </button>
                        {isEdit ? (
                            <button
                                className="delete"
                                onClick={this.onClickDelete}>
                                Delete
                            </button>
                        ) : (
                            <div></div>
                        )}
                    </div>
                    <div>
                        <button className="save" onClick={this.onClickSave}>
                            Save
                        </button>
                    </div>
                </div>
            </div>
        );
    }
}
