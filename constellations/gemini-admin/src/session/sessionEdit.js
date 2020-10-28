"use strict";
require("./sessionEdit.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import { InputText } from "../common/inputs/inputText.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";

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
                    notes: session.notes || "",
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

            if (this.state.startsAt.isAfter(newEndTime)) {
                window.alert("End time cannot be before start time");
            } else {
                this.setState({
                    endsAt: newEndTime,
                });
            }
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

    onSave = () => {
        let session = {
            classId: this.state.classId,
            startsAt: this.state.startsAt.toJSON(),
            endsAt: this.state.endsAt.toJSON(),
            canceled: this.state.canceled,
            notes: this.state.notes,
        };

        return API.post("api/sessions/session/" + this.props.id, session);
    };

    onDelete = () => {
        return API.delete("api/sessions/delete", {
            data: [parseInt(this.props.id)],
        });
    };

    renderContent = () => {
        return (
            <div>
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
                    <h4>Check if no class</h4>
                    <input
                        type="checkbox"
                        checked={!!this.state.canceled}
                        onChange={this.onCanceledChange}
                    />
                </div>

                <InputText
                    label="Notes"
                    description="You may enter any additional information about this session here"
                    isTextBox={true}
                    value={this.state.notes}
                    onChangeCallback={(e) => this.onNotesChange(e)}
                />
            </div>
        );
    };

    render = () => {
        const sessionId = this.props.id;
        const classId = this.state.classId;
        const content = this.renderContent();
        return (
            <div id="view-session-edit">
                <EditPageWrapper
                    isEdit={true}
                    title={
                        "Edit session " +
                        this.state.sessionId +
                        " for " +
                        classId
                    }
                    content={content}
                    prevPageUrl={"sessions"}
                    onDelete={this.onDelete}
                    onSave={this.onSave}
                    entityName={"session"}
                />
            </div>
        );
    };
}
