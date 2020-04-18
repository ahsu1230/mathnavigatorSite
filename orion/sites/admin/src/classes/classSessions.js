"use strict";
require("./classSessions.styl");
import React from "react";
import _ from "lodash";
import moment from "moment";
import API from "../api.js";
import { ClassSessionList } from "./classSessionList.js";
import { Modal } from "../modals/modal.js";
import { OkayModal } from "../modals/okayModal.js";

// React DatePicker
import "react-dates/initialize";
import { SingleDatePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";

// React TimePicker
import TimePicker from "react-times";
import "react-times/css/classic/default.css";

export class ClassSessions extends React.Component {
    constructor(props) {
        super(props);
        let now = moment();
        this.state = {
            dateFocused: false,
            inputNumRepeat: 1,
            inputStartDateTime: now,
            inputEndDateTime: now.add(2, "h"),
        };

        this.onClickAddSessions = this.onClickAddSessions.bind(this);
        this.onDateChange = this.onDateChange.bind(this);
        this.onFocusChange = this.onFocusChange.bind(this);
        this.onTimeChange = this.onTimeChange.bind(this);
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onClickAddSessions() {
        const startDateTime = moment(this.state.inputStartDateTime);
        const endDateTime = moment(this.state.inputEndDateTime);
        const newSession = {
            id: "new" + this.props.sessions.length, // must generate a fake id because not yet persisted to database
            classId: this.props.classId,
            startsAt: startDateTime,
            endsAt: endDateTime,
            canceled: false,
        };
        this.props.onAddSessions([newSession]);
    }

    onDateChange(date) {
        let newDate = this.state.inputStartDateTime
            .date(date.date())
            .month(date.month())
            .year(date.year());
        this.setState({
            inputStartDateTime: newDate,
        });
    }

    onTimeChange(inputField, options) {
        let { hour, minute, meridiem } = options;
        hour = parseInt(hour);
        minute = parseInt(minute);
        let newHour =
            meridiem == "PM" && parseInt(hour) < 12 ? hour + 12 : hour;

        if (inputField == "start") {
            const newStartDateTime = moment(this.state.inputStartDateTime)
                .hour(newHour)
                .minute(minute)
                .second(0);
            const newEndDateTime = moment(newStartDateTime).add(2, "h");
            this.setState({
                inputStartDateTime: newStartDateTime,
                inputEndDateTime: newEndDateTime,
            });
        } else if (inputField == "end") {
            const newEndDateTime = this.state.inputEndDateTime
                .hour(newHour)
                .minute(minute)
                .second(0);
            this.setState({
                inputEndDateTime: newEndDateTime,
            });
        }
    }

    onFocusChange(focusState) {
        // For TimePicker. Do nothing
    }

    render() {
        return (
            <div id="section-class-sessions">
                <h3>Add Sessions</h3>
                <div className="date-block">
                    <h4>Choose Single Day</h4>
                    <SingleDatePicker
                        date={this.state.inputStartDateTime}
                        onDateChange={(date) => this.onDateChange(date)}
                        focused={this.state.dateFocused}
                        onFocusChange={({ focused }) =>
                            this.setState({ dateFocused: focused })
                        }
                        id="sessions-date-picker"
                        showDefaultInputIcon
                    />
                </div>

                <div className="time-block start">
                    <h4>Start Time</h4>
                    <TimePicker
                        time={this.state.inputStartDateTime.format("HH:mm")}
                        theme="classic"
                        showTimezone={true}
                        onFocusChange={this.onFocusChange}
                        onTimeChange={(options) =>
                            this.onTimeChange("start", options)
                        }
                        timeMode="12"
                        minuteStep={15}
                        timeConfig={{
                            from: "8:00 AM",
                            to: "10:00 PM",
                            step: 15,
                            unit: "minutes",
                        }}
                    />
                </div>

                <div className="time-block end">
                    <h4>End Time</h4>
                    <TimePicker
                        time={this.state.inputEndDateTime.format("HH:mm")}
                        theme="classic"
                        showTimezone={true}
                        onFocusChange={this.onFocusChange}
                        onTimeChange={(options) =>
                            this.onTimeChange("end", options)
                        }
                        timeMode="12"
                        minuteStep={15}
                        timeConfig={{
                            from: "8:00 AM",
                            to: "10:00 PM",
                            step: 15,
                            unit: "minutes",
                        }}
                    />
                </div>

                <div className="repeat-block">
                    {/* <span>Repeat Every Week: </span>
                    <input
                        className="num-repeat"
                        value={this.state.inputNumRepeat}
                        onChange={(e) =>
                            this.handleChange(e, "inputNumRepeat")
                        }
                    />
                    <span>times starting from </span>
                    <span>
                        {
                            this.state.inputStartDateTime.format("dddd, MMMM Do YYYY").toString() + 
                            ", " + 
                            this.state.inputStartDateTime.format("h:mm a").toString() + 
                            " to " + 
                            this.state.inputEndDateTime.format("h:mm a").toString()
                        }
                    </span> */}
                </div>

                <button className="btn-add" onClick={this.onClickAddSessions}>
                    Add Sessions
                </button>
                <ClassSessionList
                    classId={this.props.classId}
                    sessions={this.props.sessions}
                    onDeleteSession={this.props.onDeleteSession}
                />
            </div>
        );
    }
}
