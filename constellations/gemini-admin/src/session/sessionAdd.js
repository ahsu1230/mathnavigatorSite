"use strict";
require("./sessionAdd.sass");
import React from "react";
import moment from "moment";

// React DatePicker
import "react-dates/initialize";
import "react-dates/lib/css/_datepicker.css";
import { SingleDatePicker } from "react-dates";

// React TimePicker
import "react-times/css/classic/default.css";
import TimePicker from "react-times";

export class SessionAdd extends React.Component {
    constructor(props) {
        super(props);
        // Default is 3:00 pm
        let today = moment().hour(15).minute(0);
        this.state = {
            dateFocused: false,
            numWeeks: 1,
            startsAt: today,
            endsAt: moment(today).add(2, "h"),
            notes: "",
        };
    }

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

    onRepeatChange = (e) => {
        this.setState({
            numWeeks: parseInt(e.target.value == "" ? 0 : e.target.value),
        });
    };

    onNotesChange = (e) => {
        this.setState({
            notes: e.target.value,
        });
    };

    onClickAddSessions = () => {
        var sessions = [];
        const numWeeks = this.state.numWeeks;
        const notes = this.state.notes;
        for (let i = 0; i < numWeeks; i++) {
            let startsAt = moment(this.state.startsAt).add(i, "w");
            let endsAt = moment(this.state.endsAt).add(i, "w");

            sessions.push({
                startsAt: startsAt,
                endsAt: endsAt,
                notes: notes,
            });
        }
        this.props.addSessions(sessions);
    };

    render = () => {
        return (
            <section id="add-sessions">
                <h3>Add Sessions to selected class</h3>

                <div className="container">
                    <div id="date-inputs">
                        <div className="row">
                            <div className="item">
                                <h4>Choose a Day</h4>
                                <SingleDatePicker
                                    id="date-picker"
                                    date={this.state.startsAt}
                                    onDateChange={(date) =>
                                        this.onDateChange(date)
                                    }
                                    focused={this.state.dateFocused}
                                    onFocusChange={({ focused }) =>
                                        this.setState({
                                            dateFocused: focused,
                                        })
                                    }
                                    showDefaultInputIcon
                                />
                            </div>

                            <div className="item num-weeks">
                                <h4>Repeat Every Week</h4>
                                <input
                                    id="repeat"
                                    value={this.state.numWeeks}
                                    onChange={(e) => this.onRepeatChange(e)}
                                />
                            </div>

                            <div className="item notes">
                                <h4>Notes</h4>
                                <textarea
                                    id="notes"
                                    value={this.state.notes}
                                    onChange={(e) => this.onNotesChange(e)}
                                />
                            </div>
                        </div>

                        <div className="row">
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
                    </div>

                    <div id="buttons">
                        <button
                            id="btn-add-sessions"
                            onClick={this.onClickAddSessions}>
                            Add Sessions
                        </button>
                    </div>
                </div>
            </section>
        );
    };
}
