"use strict";
require("./announceEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { AnnounceEditDateTime } from "./announceEditDateTime.js";
import { InputText, emptyValidator } from "../utils/inputText.js";
import EditPageWrapper from "../utils/editPageWrapper.js";

export class AnnounceEditPage extends React.Component {
    state = {
        announceId: 0,
        datePickerFocused: false,
        inputPostedAt: null,
        inputAuthor: "",
        inputMessage: "",
        isEdit: false,
    };

    componentDidMount() {
        const announceId = this.props.announceId;
        if (announceId) {
            API.get("api/announcements/announcement/" + announceId).then(
                (res) => {
                    const announce = res.data;
                    this.setState({
                        announceId: announce.id,
                        inputPostedAt: moment(announce.postedAt),
                        inputAuthor: announce.author,
                        inputMessage: announce.message,
                        isEdit: true,
                    });
                }
            );
        }
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    onSave = () => {
        let announcement = {
            postedAt: this.state.inputPostedAt.toJSON(),
            author: this.state.inputAuthor,
            message: this.state.inputMessage,
        };

        if (this.state.isEdit) {
            return API.post(
                "api/announcements/announcement/" + this.state.announceId,
                announcement
            );
        } else {
            return API.post("api/announcements/create", announcement);
        }
    };

    onDelete = () => {
        const announceId = this.props.announceId;
        return API.delete("api/announcements/announcement/" + announceId);
    };

    onMomentChange = (newMoment) => {
        this.setState({ inputPostedAt: newMoment });
    };

    renderContent = () => {
        return (
            <div>
                <AnnounceEditDateTime
                    postedAt={this.state.inputPostedAt}
                    onMomentChange={this.onMomentChange}
                />

                <InputText
                    label="Author"
                    description="Input your name"
                    required={true}
                    value={this.state.inputAuthor}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputAuthor")
                    }
                    validators={[emptyValidator("author")]}
                />

                <InputText
                    label="Message"
                    description="Enter the announcement message"
                    isTextBox={true}
                    required={true}
                    value={this.state.inputMessage}
                    onChangeCallback={(e) =>
                        this.handleChange(e, "inputMessage")
                    }
                    validators={[emptyValidator("message")]}
                />
            </div>
        );
    };

    render() {
        const title = this.state.isEdit
            ? "Edit Announcement"
            : "Add Announcement";
        const content = this.renderContent();
        return (
            <div id="view-announce-edit">
                <EditPageWrapper
                    isEdit={this.state.isEdit}
                    title={title}
                    content={content}
                    prevPageUrl={"announcements"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityName={"announcement"}
                />
            </div>
        );
    }
}
