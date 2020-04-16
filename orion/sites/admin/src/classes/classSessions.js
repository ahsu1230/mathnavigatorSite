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
            listSessionsLocal: [],
            listSessionsRemote: []
        }
        this.refreshSessionList = this.refreshSessionList.bind(this);

        this.onDateChange = this.onDateChange.bind(this);
        this.onAddSessions = this.onAddSessions.bind(this);
        this.onDeleteSession = this.onDeleteSession.bind(this);

        this.onFocusChange = this.onFocusChange.bind(this)
        this.onTimeChange = this.onTimeChange.bind(this);
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.classId != "_" && this.props.classId !== prevProps.classId) {
            this.refreshSessionList();
        }
    }

    refreshSessionList() {
        console.log("before render: " + this.props.classId);
        API.get("api/sessions/v1/class/" + this.props.classId).then((res) => {
            const sessions = res.data;
            console.log("returned: " + sessions.length);
            this.setState({ 
                listSessionsLocal: sessions,
                listSessionsRemote: sessions
            });
        });
    }

    handleChange(event, value) {
        this.setState({ [value]: event.target.value });
    }

    onAddSessions() {
        const sessions = this.state.listSessionsLocal;
        const newSession = {
            id: "new" + sessions.length, // must generate a fake id because not yet persisted to database
            classId: this.props.classId,
            startsAt: this.state.inputStartDateTime,
            endsAt: this.state.inputEndDateTime,
            canceled: false
        };

        // Save to local list
        const newList = _.concat(sessions, newSession);
        this.setState({ listSessionsLocal: newList });


        // API.post("api/sessions/v1/create", session).then((res) => {
        //     this.refreshSessionList();
        //     console.log("Sessions added! refreshing");
        // }).catch((err) => {
        //     debugger;
        //     alert("Did not add sessions: " + err);
        // });
    }

    onDeleteSession(sessionId) {
        let sessions = this.state.listSessionsLocal;
        _.remove(sessions, {id: sessionId});
        this.setState({
            listSessionsLocal: sessions
        });
    }

    onDateChange(date) {
        let newDate = this.state.inputStartDateTime
            .date(date.date())
            .month(date.month())
            .year(date.year());
        this.setState({
            inputStartDateTime: newDate
        });
    }

    onTimeChange(inputField, options) {
        let newHour = parseInt(options.hour);
        let hour = options.meridiem == "AM" ? newHour : newHour + 12;
        let minute = parseInt(options.minute);

        if (inputField == "inputStartDateTime") {
            const newStartDateTime = this.state.inputStartDateTime.hour(hour).minute(minute).second(0);
            const newEndDateTime = moment(newStartDateTime).add(2, "h")
            this.setState({
                inputStartDateTime: newStartDateTime,
                inputEndDateTime: newEndDateTime
            });
        } else {
            const newEndDateTime = this.state.inputEndDateTime.hour(hour).minute(minute).second(0);
            this.setState({
                inputEndDateTime: newEndDateTime
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
                        onFocusChange={({ focused }) => this.setState({ dateFocused: focused })}
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
                        timeMode="12"
                        onFocusChange={this.onFocusChange}
                        onTimeChange={(options) => this.onTimeChange("inputStartDateTime", options)}
                        minuteStep={15}
                        timeConfig={{
                            from: "7:00 AM",
                            to: "11:45 PM",
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
                        timeMode="12"
                        onFocusChange={this.onFocusChange}
                        onTimeChange={(options) => this.onTimeChange("inputEndDateTime", options)}
                        minuteStep={15}
                        timeConfig={{
                            from: "7:00 AM",
                            to: "11:45 PM",
                            step: 15,
                            unit: "minutes",
                        }}
                    />
                </div>
                
                <div className="repeat-block">
                    <span>Repeat Every Week: </span>
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
                    </span>
                </div>
                
                <button className="btn-add" onClick={this.onAddSessions}>
                    Add Sessions
                </button>
                <ClassSessionList 
                    classId={this.props.classId} 
                    sessions={this.state.listSessionsLocal}
                    onDeleteSession={this.onDeleteSession}
                />
            </div>
        );
    }
}