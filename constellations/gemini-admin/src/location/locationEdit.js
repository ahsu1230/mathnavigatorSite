"use strict";
require("./locationEdit.sass");
import React from "react";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";
import { InputText } from "../utils/inputText.js";

export class LocationEditPage extends React.Component {
    state = {
        oldlocationId: "",
        inputlocationId: "",
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
        const locationId = this.props.locationId;
        if (locationId) {
            API.get("api/locations/location/" + locationId).then((res) => {
                const location = res.data;
                this.setState({
                    oldlocationId: location.locationId,
                    inputlocationId: location.locationId,
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
            locationId: this.state.inputlocationId,
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
                "api/locations/location/" + this.state.oldlocationId,
                location
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/locations/create", location)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    };

    onDeleted = () => {
        const locationId = this.props.locationId;
        API.delete("api/locations/location/" + locationId).then((res) => {
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
                <InputText
                    description="Location ID"
                    required={true}
                    value={this.state.inputlocationId}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputlocationId")
                    }
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a location ID",
                        },
                    ]}
                />
                <InputText
                    description="Street  (E.g. 1234 Gains Rd, 5432 Abc Drive W)"
                    required={true}
                    value={this.state.inputStreet}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputStreet")
                    }
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a street",
                        },
                    ]}
                />
                <InputText
                    description="City  (E.g. Potomac)"
                    required={true}
                    value={this.state.inputCity}
                    onChangeCallback={(e) => this.handleChange(e, "inputCity")}
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a city",
                        },
                    ]}
                />
                <InputText
                    description="State  (E.g. MD)"
                    required={true}
                    value={this.state.inputState}
                    onChangeCallback={(e) => this.handleChange(e, "inputState")}
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a state",
                        },
                    ]}
                />
                <InputText
                    description="Zipcode"
                    required={true}
                    value={this.state.inputZip}
                    onChangeCallback={(e) => this.handleChange(e, "inputZip")}
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a zipcode",
                        },
                    ]}
                />
                <InputText
                    description="Room"
                    required={false}
                    value={this.state.inputRoom}
                    onChangeCallback={(e) => this.handleChange(e, "inputRoom")}
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
