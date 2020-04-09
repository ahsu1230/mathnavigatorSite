"use strict";
require("./announceEdit.styl");
import React from "react";

import moment from "moment";
import "react-dates/initialize";
import { SingleDatePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";

import TimePicker from "react-times";
import "react-times/css/classic/default.css";

import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

export class AnnounceEditPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            announceId: 0,
            datePickerFocused: false,
            inputPostedAt: moment(), // always a moment object
            inputAuthor: "",
            inputMessage: "",
            isEdit: false,
        };

        this.handleChange = this.handleChange.bind(this);

        this.onClickCancel = this.onClickCancel.bind(this);
        this.onClickDelete = this.onClickDelete.bind(this);
        this.onClickSave = this.onClickSave.bind(this);

        this.onConfirmDelete = this.onConfirmDelete.bind(this);
        this.onSavedOk = this.onSavedOk.bind(this);
        this.onDismissModal = this.onDismissModal.bind(this);

        this.onMomentChange = this.onMomentChange.bind(this);
    }

    componentDidMount() {
        const announceId = this.props.announceId;
        if (announceId) {
            API.get("api/announcements/v1/announcement/" + announceId).then(
                (res) => {
                    const announce = res.data;
                    this.setState({
                        announceId: announce.id,
                        inputPostedAt: moment(announce.postedAt),
                        inputAuthor: announce.author,
                        inputMessage: announce.message,
                        isEdit: true,
                        showDeleteModal: false,
                        showSaveModal: false,
                    });
                }
            );
        }
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickSave() {
        let announcement = {
            postedAt: this.state.inputPostedAt.toJSON(),
            author: this.state.inputAuthor,
            message: this.state.inputMessage,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save announcement: " + err.response.data);
        if (this.state.isEdit) {
            API.post(
                "api/announcements/v1/announcement/" + this.state.announceId,
                announcement
            )
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        } else {
            API.post("api/announcements/v1/create", announcement)
                .then((res) => successCallback())
                .catch((err) => failCallback(err));
        }
    }

    onClickCancel() {
        window.location.hash = "announcements";
    }

    onClickDelete() {
        this.setState({ showDeleteModal: true });
    }

    onConfirmDelete() {
        const announceId = this.props.announceId;
        API.delete("api/announcements/v1/announcement/" + announceId).then(
            (res) => {
                window.location.hash = "announcements";
            }
        );
    }

    onSavedOk() {
        this.onDismissModal();
        window.location.hash = "announcements";
    }

    onDismissModal() {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    }

    onMomentChange(newMoment) {
        this.setState({ inputPostedAt: newMoment });
    }

    render() {
        const title = this.state.isEdit
            ? "Edit Announcement"
            : "Add Announcement";
        let deleteButton = <div></div>;
        if (this.state.isEdit) {
            deleteButton = (
                <button className="btn-delete" onClick={this.onClickDelete}>
                    Delete
                </button>
            );
        }

        let modalDiv;
        let modalContent;
        let showModal;
        if (this.state.showDeleteModal) {
            showModal = this.state.showDeleteModal;
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={this.onConfirmDelete}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalContent = (
                <OkayModal
                    text={"Announcement information saved!"}
                    onOkay={this.onSavedOk}
                />
            );
        }
        if (modalContent) {
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={showModal}
                    onDismiss={this.onDismissModal}
                />
            );
        }

        return (
            <div id="view-announce-edit">
                {modalDiv}
                <h2>{title}</h2>

                <AnnounceDateTimePicker
                    postedAt={this.state.inputPostedAt}
                    onMomentChange={this.onMomentChange}
                />

                <h4>Author</h4>
                <input
                    value={this.state.inputAuthor}
                    onChange={(e) => this.handleChange(e, "inputAuthor")}
                />

                <h4>Message</h4>
                <textarea
                    value={this.state.inputMessage}
                    onChange={(e) => this.handleChange(e, "inputMessage")}
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

class AnnounceDateTimePicker extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            dateFocused: false
        };
        this.onClickNow = this.onClickNow.bind(this);
        this.onDateChange = this.onDateChange.bind(this);
    }

    onClickNow() {
        const now = moment();
        this.props.onMomentChange(now);
    }

    onDateChange(date) {
        let dateMoment = moment(date);
        let newMoment = this.props.postedAt
                            .year(dateMoment.year())
                            .month(dateMoment.month())
                            .date(dateMoment.date());
        this.props.onMomentChange(newMoment);
    }

    onTimeChange(options) {
        let newHour = parseInt(options.hour);
        let hour = options.meridiem == "AM" ? newHour : newHour + 12;
        let minute = parseInt(options.minute);
        let newMoment = this.props.postedAt
                            .hour(hour)
                            .minute(minute);
        this.props.onMomentChange(newMoment);
    }

    onFocusChange(focusState) {
        // For TimePicker. Do nothing
    }

    render() {
        const postedAt = this.props.postedAt;
        const postedAtString = postedAt.format("dddd, MMMM Do YYYY, h:mm az").toString();

        return (
            <div id="announce-datetime-container">
                <div id="announce-date-picker">
                    <h4>Select Date</h4>
                    <SingleDatePicker
                        date={postedAt}
                        onDateChange={(date) => this.onDateChange(date)}
                        focused={this.state.dateFocused}
                        onFocusChange={({ focused }) =>
                            this.setState({ dateFocused: focused })
                        }
                        id="announce-date-picker"
                        showDefaultInputIcon
                    />
                </div>

                <div id="announce-time-picker">
                    <h4>Select Time</h4>
                    <TimePicker
                        time={postedAt.format("HH:mm")}
                        theme="classic"
                        showTimezone={true}
                        timeMode="12"
                        onFocusChange={this.onFocusChange.bind(this)}
                        onTimeChange={this.onTimeChange.bind(this)}
                        minuteStep={15}
                        timeConfig={{
                            from: "7:00 AM",
                            to: "11:45 PM",
                            step: 15,
                            unit: "minutes",
                        }}
                    />
                </div>
                <p>Or</p>
                <button className="btn-now" onClick={this.onClickNow}>Schedule For Now</button>

                <h4>Announcement will be published on <b>{postedAtString}</b></h4>
            </div>
        );
    }
}


// console.log("**************************");
// console.log(postedAt.toString());
// console.log(postedAt.toISOString());
// console.log(postedAt.toJSON());
// console.log(postedAt.toObject());

// console.log(postedAt.toDate());
// console.log(postedAt.toDate().toLocaleString());
// console.log(postedAt.toDate().toJSON());
// console.log("**************************");
// **************************
// Thu Apr 09 2020 12:00:00 GMT-0400
// 2020-04-09T16:00:00.000Z
// 2020-04-09T16:00:00.000Z
// {years: 2020, months: 3, date: 9, hours: 12, minutes: 0, …}years: 2020months: 3date: 9hours: 12minutes: 0seconds: 0milliseconds: 0__proto__: Object
// Thu Apr 09 2020 12:00:00 GMT-0400 (Eastern Daylight Time)
// 4/9/2020, 12:00:00 PM
// 2020-04-09T16:00:00.000Z
// **************************