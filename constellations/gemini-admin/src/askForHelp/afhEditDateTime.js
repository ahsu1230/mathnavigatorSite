"use strict";
require("./afhEditDateTime.sass");
import React from "react";
import moment from "moment";

// React DatePicker
import "react-dates/initialize";
import { SingleDatePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";

// React TimePicker
import TimePicker from "react-times";
import "react-times/css/classic/default.css";

export class AfhEditDateTimes extends React.Component {
    state = {
        dateFocused: false,
    };

    onDateChange = (date) => {
        let dateMoment = moment(date);

        // only change date, month, year of old times
        let newStartsAt = moment(this.props.startsAt)
            .year(dateMoment.year())
            .month(dateMoment.month())
            .date(dateMoment.date());
        let newEndsAt = moment(this.props.endsAt)
            .year(dateMoment.year())
            .month(dateMoment.month())
            .date(dateMoment.date());
        this.props.onStartsAtChange(newStartsAt);
        this.props.onEndsAtChange(newEndsAt);
    };

    onFocusChange = (focusState) => {
        // For TimePicker. Do nothing
    };

    onTimeChange = (newTime, isStartsAt) => {
        const second = parseInt(newTime.second) || 0;
        const minute = parseInt(newTime.minute) || 0;
        let hour = parseInt(newTime.hour) || 0;

        // Use this if in 12-hour mode
        // 12-hour mode in TimePicker has a bug in it
        // const meridiem = newTime.meridiem;  // is either "AM" or "PM"
        // hour += (hour != 12 && meridiem == "PM" ? 12 : 0);

        // Only change hours, minutes, and seconds of old times
        if (isStartsAt) {
            let newStartsAt = moment(this.props.startsAt)
                .hour(hour)
                .minute(minute)
                .second(second)
                .millisecond(0);
            this.props.onStartsAtChange(newStartsAt);
        } else {
            let newEndsAt = moment(this.props.endsAt)
                .hour(hour)
                .minute(minute)
                .second(second)
                .millisecond(0);
            this.props.onEndsAtChange(newEndsAt);
        }
    };

    render() {
        const startsAt = moment(this.props.startsAt);
        const endsAt = moment(this.props.endsAt);

        return (
            <div id="afh-datetime-container">
                <h4>Select Date</h4>
                <SingleDatePicker
                    date={startsAt}
                    onDateChange={(date) => this.onDateChange(date)}
                    focused={this.state.dateFocused}
                    onFocusChange={({ focused }) =>
                        this.setState({ dateFocused: focused })
                    }
                    id="announce-date-picker"
                    showDefaultInputIcon
                />

                <h4>Select StartAt</h4>
                <TimePicker
                    time={startsAt.format("HH:mm")}
                    theme="classic"
                    showTimezone={true}
                    timeMode="24"
                    onFocusChange={this.onFocusChange}
                    onTimeChange={(time) => this.onTimeChange(time, true)}
                    minuteStep={15}
                    timeConfig={{
                        from: "8:00 AM",
                        to: "11:45 PM",
                        step: 5,
                        unit: "minutes",
                    }}
                />

                <h4>Select EndsAt</h4>
                <TimePicker
                    time={endsAt.format("HH:mm")}
                    theme="classic"
                    showTimezone={true}
                    timeMode="24"
                    onFocusChange={this.onFocusChange}
                    onTimeChange={(time) => this.onTimeChange(time, false)}
                    minuteStep={15}
                    timeConfig={{
                        from: "8:00 AM",
                        to: "11:45 PM",
                        step: 5,
                        unit: "minutes",
                    }}
                />
            </div>
        );
    }
}
