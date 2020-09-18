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

export class AFHEditDateTime extends React.Component {
    render() {
        const postedAt = this.props.postedAt;
        const now = moment();
        let content = (
            <AFHDateTimePicker
                postedAt={postedAt || now}
                onMomentChange={this.props.onMomentChange}
            />
        );
        return <div id="afh-datetime-container">{content}</div>;
    }
}

class AFHDateTimePicker extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            dateFocused: false,
        };
        this.onDateChange = this.onDateChange.bind(this);
    }

    onDateChange(date) {
        let dateMoment = moment(date);
        let newMoment = this.props.postedAt
            .year(dateMoment.year())
            .month(dateMoment.month())
            .date(dateMoment.date());
        this.props.onMomentChange(newMoment);
    }

    onFocusChange(focusState) {
        // For TimePicker. Do nothing
    }

    render() {
        const postedAt = this.props.postedAt;
        return (
            <div id="afh-datetime-container">
                <div id="afh-date-picker">
                    <h3>Select Date</h3>
                    <SingleDatePicker
                        date={postedAt}
                        onDateChange={(date) => this.onDateChange(date)}
                        focused={this.state.dateFocused}
                        onFocusChange={({ focused }) =>
                            this.setState({ dateFocused: focused })
                        }
                        id="afh-date-picker"
                        showDefaultInputIcon
                    />
                </div>
            </div>
        );
    }
}
