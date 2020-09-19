"use strict";
require("./announceEditDateTime.sass");
import React from "react";
import moment from "moment";

// React DatePicker
import "react-dates/initialize";
import { SingleDatePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";

// React TimePicker
import TimePicker from "react-times";
import "react-times/css/classic/default.css";

export class AnnounceEditDateTime extends React.Component {
    render() {
        const postedAt = this.props.postedAt;
        const now = moment();
        let content;
        if (postedAt && now.isAfter(postedAt)) {
            let postedAtString = postedAt
                .format("dddd, MMMM Do YYYY, h:mm a")
                .toString();
            content = (
                <h4>
                    This post was published on <b>{postedAtString}</b>
                </h4>
            );
        } else {
            content = (
                <AnnounceDateTimePicker
                    postedAt={postedAt || now}
                    onMomentChange={this.props.onMomentChange}
                />
            );
        }
        return <div id="announce-datetime-container">{content}</div>;
    }
}

class AnnounceDateTimePicker extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            dateFocused: false,
        };
        this.onClickNow = this.onClickNow.bind(this);
        this.onDateChange = this.onDateChange.bind(this);
    }

    onClickNow() {
        const now = moment().add(2, "minutes");
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
        let newMoment = this.props.postedAt.hour(hour).minute(minute);
        this.props.onMomentChange(newMoment);
    }

    onFocusChange(focusState) {
        // For TimePicker. Do nothing
    }

    render() {
        const postedAt = this.props.postedAt;
        const postedAtString = postedAt
            .format("dddd, MMMM Do YYYY, h:mm az")
            .toString();
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
                <button className="btn-now" onClick={this.onClickNow}>
                    Schedule For Now
                </button>

                <h4>
                    Announcement will be published on <b>{postedAtString}</b>
                </h4>
            </div>
        );
    }
}
