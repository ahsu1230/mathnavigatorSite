"use strict";
require("./sessionEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";
import { YesNoModal } from "../modals/yesnoModal.js";

// React DatePicker
import "react-dates/initialize";
import "react-dates/lib/css/_datepicker.css";
import { SingleDatePicker } from "react-dates";

// React TimePicker
import "react-times/css/classic/default.css";
import TimePicker from "react-times";

export class SessionEditPage extends React.Component {
    state = {
        classId: this.props.classId,
        startsAt: moment(),
        endsAt: moment(),
        canceled: false,
        notes: "",
        dateFocused: false,
        showDeleteModal: false,
        showSaveModal: false,
    };

    componentDidMount = () => {
        const id = this.props.id;
        if (id) {
            API.get("api/sessions/session/" + id).then((res) => {
                const session = res.data;
                this.setState({
                    startsAt: moment(session.startsAt),
                    endsAt: moment(session.endsAt),
                    canceled: session.canceled,
                    notes: session.notes,
                });
            });
        }
    };

    onDateChange = (date) => {
        let newDate = moment(this.state.startsAt)
            .date(date.date())
            .month(date.month())
            .year(date.year());

        this.setState({
            startsAt: newDate,
        });
    };

    onTimeChange = (inputField, options) => {
        let { hour, minute, meridiem } = options;
        hour = parseInt(hour);
        minute = parseInt(minute);
        let newHour =
            meridiem == "PM" && parseInt(hour) < 12 ? hour + 12 : hour;

        if (inputField == "start") {
            const newStartTime = moment(this.state.startsAt)
                .hour(newHour)
                .minute(minute)
                .second(0);
            const newEndTime = moment(newStartTime).add(2, "h");

            this.setState({
                startsAt: newStartTime,
                endsAt: newEndTime,
            });
        } else if (inputField == "end") {
            const newEndTime = moment(this.state.endsAt)
                .hour(newHour)
                .minute(minute)
                .second(0);

            this.setState({
                endsAt: newEndTime,
            });
        }
    };

    onCanceledChange = () => {
        this.setState({
            canceled: !this.state.canceled,
        });
    };

    onNotesChange = (e) => {
        this.setState({
            notes: e.target.value,
        });
    };

    onClickCancel = () => {
        window.location.hash = "sessions";
    };

    onClickDelete = () => {
        this.setState({ showDeleteModal: true });
    };

    onClickSave = () => {
        let session = {
            classId: this.state.classId,
            startsAt: this.state.startsAt.toJSON(),
            endsAt: this.state.endsAt.toJSON(),
            canceled: this.state.canceled,
            notes: this.state.notes,
        };

        let successCallback = () => this.setState({ showSaveModal: true });
        let failCallback = (err) =>
            alert("Could not save session: " + err.response.data);

        API.post("api/sessions/session/" + this.props.id, session)
            .then(() => successCallback())
            .catch((err) => failCallback(err));
    };

    onSaved = () => {
        this.onDismissModal();
        window.location.hash = "sessions";
    };

    onDeleted = () => {
        API.delete("api/sessions/delete", { data: [parseInt(this.props.id)] })
            .then(() => {
                window.location.hash = "sessions";
            })
            .catch((err) => {
                alert("Could not delete session: " + err.response.data);
            })
            .finally(() => this.onDismissModal());
    };

    onDismissModal = () => {
        this.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
    };

    render = () => {
        let modalDiv;
        let modalContent;
        let showModal;
        if (this.state.showDeleteModal) {
            showModal = this.state.showDeleteModal;
            modalContent = (
                <YesNoModal
                    text={"Are you sure you want to delete?"}
                    onAccept={this.onDeleted}
                    onReject={this.onDismissModal}
                />
            );
        }
        if (this.state.showSaveModal) {
            showModal = this.state.showSaveModal;
            modalContent = (
                <OkayModal
                    text={"Session information saved!"}
                    onOkay={this.onSaved}
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
            <div id="view-session-edit">
                {modalDiv}
                <h2>Edit Session</h2>

                <div className="item">
                    <h4>Choose a Day</h4>
                    <SingleDatePicker
                        id="date-picker"
                        date={this.state.startsAt}
                        onDateChange={(date) => this.onDateChange(date)}
                        focused={this.state.dateFocused}
                        onFocusChange={({ focused }) =>
                            this.setState({
                                dateFocused: focused,
                            })
                        }
                        showDefaultInputIcon
                    />
                </div>

                <div id="edit-times">
                    <div className="item">
                        <h4>Start Time</h4>
                        <TimePicker
                            timezone="America/New_York"
                            time={this.state.startsAt.format("HH:mm")}
                            theme="classic"
                            onTimeChange={(options) =>
                                this.onTimeChange("start", options)
                            }
                            timeMode="12"
                            timeConfig={{
                                from: "8:00 AM",
                                to: "9:00 PM",
                                step: 15,
                                unit: "minutes",
                            }}
                        />
                    </div>

                    <div className="item">
                        <h4>End Time</h4>
                        <TimePicker
                            timezone="America/New_York"
                            time={this.state.endsAt.format("HH:mm")}
                            theme="classic"
                            onTimeChange={(options) =>
                                this.onTimeChange("end", options)
                            }
                            timeMode="12"
                            timeConfig={{
                                from: "8:00 AM",
                                to: "11:00 PM",
                                step: 15,
                                unit: "minutes",
                            }}
                        />
                    </div>
                </div>

                <div id="cancel-toggle" className="item">
                    <h4>Canceled</h4>
                    <input
                        type="checkbox"
                        value={this.state.canceled}
                        onChange={this.onCanceledChange}
                    />
                </div>

                <h4>Notes</h4>
                <input
                    value={this.state.notes}
                    onChange={(e) => this.onNotesChange(e)}
                />

                <div className="buttons">
                    <button className="btn-save" onClick={this.onClickSave}>
                        Save
                    </button>
                    <button className="btn-cancel" onClick={this.onClickCancel}>
                        Cancel
                    </button>
                    <button className="btn-delete" onClick={this.onClickDelete}>
                        Delete
                    </button>
                </div>
            </div>
        );
    };
}
