"use strict";
require("./editPageWrapper.sass");
import React from "react";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

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

    createDeleteModalText = (entityId, entityType) => {
        return "Are you sure you want to delete?";
    };

    createSaveModalText = (entityId, entityType) => {
        return "Successfully saved!";
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
                        <button className="delete" onClick={this.onClickDelete}>
                            Delete
                        </button>
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
