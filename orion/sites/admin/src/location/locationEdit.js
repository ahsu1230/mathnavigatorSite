"use strict";
require("./locationEdit.styl");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class LocationEditPage extends React.Component {
    state = {
        oldLocId: "",
        inputLocId: "",
        inputStreet: "",
        inputCity: "",
        inputState: "",
        inputZip: "",
        inputRoom: "",
        isEdit: false,
        showDeleteModal: false,
        showSaveModal: false,
    };

    componentDidMount() {
        const locId = this.props.locId;
        if (locId) {
            API.get("api/locations/v1/location/" + locId).then((res) => {
                const location = res.data;
                this.setState({
                    oldLocId: location.locId,
                    inputLocId: location.locId,
                    inputStreet: location.street,
                    inputCity: location.city,
                    inputState: location.state,
                    inputZip: location.zipcode,
                    inputRoom: location.room || "",
                    isEdit: true,
                });
            });
        }
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onClickCancel = () => {
        window.location.hash = "locations";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onClickSave = () => {
        let location = {
            locId: this.state.inputLocId,
            street: this.state.inputStreet,
            city: this.state.inputCity,
            state: this.state.inputState,
            zipcode: this.state.inputZip,
            room: this.state.inputRoom,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save location: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/locations/v1/location/" + this.state.oldLocId,
                location
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/locations/v1/create", location)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    };

    onDeleted = () => {
        const locId = this.props.locId;
        API.delete("api/locations/v1/location/" + locId).then((res) => {
            window.location.hash = "locations";
        });
    };

    onSaved = () => {
        this.onDismissModal();
        window.location.hash = "locations";
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    render() {
        const title = this.state.isEdit ? "Edit Location" : "Add Location";

        const deleteButton = renderDeleteButton(
            this.state.isEdit,
            this.onClickDelete
        );
        const modalDiv = renderModalDiv(
            this.state.showDeleteModal,
            this.state.showSaveModal,
            this.onDismissModal,
            this.onSaved,
            this.onDeleted
        );

        return (
            <div id="view-location-edit">
                {modalDiv}
                <h2>{title}</h2>
                <h4>Location ID</h4>
                <input
                    value={this.state.inputLocId}
                    onChange={(e) => this.handleChange(e, "inputLocId")}
                />
                <h4>Street</h4>
                <input
                    value={this.state.inputStreet}
                    onChange={(e) => this.handleChange(e, "inputStreet")}
                />
                <h4>City</h4>
                <input
                    value={this.state.inputCity}
                    onChange={(e) => this.handleChange(e, "inputCity")}
                />
                <h4>State</h4>
                <input
                    value={this.state.inputState}
                    onChange={(e) => this.handleChange(e, "inputState")}
                />
                <h4>Zipcode</h4>
                <input
                    value={this.state.inputZip}
                    onChange={(e) => this.handleChange(e, "inputZip")}
                />
                <h4>Room</h4>
                <input
                    value={this.state.inputRoom}
                    onChange={(e) => this.handleChange(e, "inputRoom")}
                />
                <div className="buttons">
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    {deleteButton}
                </div>
            </div>
        );
    }
}

function renderDeleteButton(isEdit, onClickDelete) {
    let deleteButton = <div />;
    if (isEdit) {
        deleteButton = (
            <button className="btn-delete" onClick={onClickDelete}>
                Delete
            </button>
        );
    }
    return deleteButton;
}

function renderModalDiv(
    showDeleteModal,
    showSaveModal,
    onDismissModal,
    onSaved,
    onDeleted
) {
    let modalDiv;
    let modalContent;
    let showModal;
    if (showDeleteModal) {
        showModal = showDeleteModal;
        modalContent = (
            <YesNoModal
                text={"Are you sure you want to delete?"}
                onAccept={onDeleted}
                onReject={onDismissModal}
            />
        );
    }
    if (showSaveModal) {
        showModal = showSaveModal;
        modalContent = (
            <OkayModal text={"Location information saved!"} onOkay={onSaved} />
        );
    }
    if (modalContent) {
        modalDiv = (
            <Modal
                content={modalContent}
                show={showModal}
                onDismiss={onDismissModal}
            />
        );
    }
    return modalDiv;
}
