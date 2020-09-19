"use strict";
require("./locationEdit.sass");
import React from "react";
import API from "../api.js";
import { InputText } from "../common/inputs/inputText.js";
import { InputRadio } from "../common/inputs/inputRadio.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

const ADDRESS_ONLINE = "online";
const ADDRESS_PHYSICAL = "physical";

export class LocationEditPage extends React.Component {
    state = {
        oldLocationId: "",
        inputLocationId: "",
        inputTitle: "",
        inputStreet: "",
        inputCity: "",
        inputState: "",
        inputZip: "",
        inputRoom: "",
        inputOnline: ADDRESS_PHYSICAL, // options are "physical" or "online"
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
                    oldLocationId: location.locationId,
                    inputLocationId: location.locationId,
                    inputTitle: location.title,
                    inputStreet: location.street || "",
                    inputCity: location.city || "",
                    inputState: location.state || "",
                    inputZip: location.zipcode || "",
                    inputRoom: location.room || "",
                    inputOnline: location.isOnline
                        ? ADDRESS_ONLINE
                        : ADDRESS_PHYSICAL,
                    isEdit: true,
                });
            });
        }
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onSave = () => {
        let location = {
            locationId: this.state.inputLocationId,
            title: this.state.inputTitle,
            street: this.state.inputStreet,
            city: this.state.inputCity,
            state: this.state.inputState,
            zipcode: this.state.inputZip,
            room: this.state.inputRoom,
            isOnline: this.state.inputOnline == ADDRESS_ONLINE,
        };

        if (this.state.isEdit) {
            return API.post(
                "api/locations/location/" + this.state.oldLocationId,
                location
            );
        } else {
            return API.post("api/locations/create", location);
        }
    };

    onDelete = () => {
        const locationId = this.props.locationId;
        return API.delete("api/locations/location/" + locationId);
    };

    renderPhysicalForm = () => {
        return (
            <div>
                <InputText
                    label="Street"
                    description="Enter a street number and street (e.g. 1234 Gains Rd, 5432 Victory Dr). This is only required for physical locations."
                    required={false}
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
                    label="City"
                    description="Enter a city (e.g. Potomac, Rockville). This is only required for physical locations."
                    required={false}
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
                    label="State"
                    description="Enter the 2 letter abbreviation of a state (e.g. MD). This is only required for physical locations."
                    required={false}
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
                    label="Zipcode"
                    description="Enter a zipcode. This is only required for physical locations."
                    required={false}
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
                    label="Room"
                    description="Enter a room description (e.g. 'Room 143' or 'Gymnasium')."
                    required={false}
                    value={this.state.inputRoom}
                    onChangeCallback={(e) => this.handleChange(e, "inputRoom")}
                />
            </div>
        );
    };

    renderContent = () => {
        const isOnline = this.state.inputOnline == ADDRESS_ONLINE;
        const form = isOnline ? <div></div> : this.renderPhysicalForm();
        return (
            <div>
                <InputText
                    label="Location ID"
                    description="Enter a location Id (e.g. wchs, home)"
                    required={true}
                    value={this.state.inputLocationId}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputLocationId")
                    }
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a location ID",
                        },
                    ]}
                />
                <InputText
                    label="Title"
                    description="Enter the name of this location"
                    required={true}
                    value={this.state.inputTitle}
                    onChangeCallback={(e) => this.handleChange(e, "inputTitle")}
                    validators={[
                        {
                            validate: (text) => text != "",
                            message: "You must input a location name",
                        },
                    ]}
                />
                <InputRadio
                    label="Is Online?"
                    description="Is this location online or a physical address?"
                    required={true}
                    value={this.state.inputOnline}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputOnline")
                    }
                    options={[
                        { value: ADDRESS_ONLINE, displayName: "Online" },
                        {
                            value: ADDRESS_PHYSICAL,
                            displayName: "Physical Address",
                        },
                    ]}
                />
                {form}
            </div>
        );
    };

    render() {
        const title = this.state.isEdit ? "Edit Location" : "Add Location";
        const content = this.renderContent();
        return (
            <div id="view-location-edit">
                <EditPageWrapper
                    isEdit={this.state.isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"locations"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityId={this.state.inputLocationId}
                    entityName={"location"}
                />
            </div>
        );
    }
}
