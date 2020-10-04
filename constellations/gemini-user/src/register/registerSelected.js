"use strict";
require("./register.sass");
import React from "react";
import moment from "moment";

export default class RegisterSelected extends React.Component {
    renderForClass = () => {
        const currentClass = this.props.classMap[this.props.classId];
        const program = this.props.programMap[currentClass.programId];
        const semester = this.props.semesterMap[currentClass.semesterId];
        const location = this.props.locationMap[currentClass.locationId];

        const fullTitle = program.title + " " + capitalizeWord(currentClass.classKey);
        const fullSection = isFullClass(currentClass) ? 
            <p className="error">This class is full. Please select another class to enroll.</p> : 
            <div></div>;
        return (
            <div className="selection">
                {fullSection}
                You have selected to enroll into:
                <h3>{fullTitle}</h3>
                <h4>{semester.title}</h4>
                <p className="times">Times: {currentClass.timesStr}</p>
                <p className="price">Prices: {currentClass.pricePerSession || currentClass.priceLumpSum}</p>
                <p className="payment-notes">{currentClass.paymentNotes}</p>
                <p className="location">
                    Location: {location.title}<br/>
                    {location.street}<br/>
                    {location.city + ", " + location.state + " " + location.zipcode}<br/>
                    {location.room}
                </p>
            </div>
        );
    }

    renderForAfh = () => {
        const currentAfh = this.props.afhMap[this.props.afhId]
        const datetime = moment(currentAfh.startsAt).format("MM/DD/YY h:mm a") + 
                            " - " + 
                            moment(currentAfh.endsAt).format("h:mm a");
        const location = this.props.locationMap[currentAfh.locationId];
        return (
            <div className="selection">
                You have selected to attend:
                <div className="info">
                    <h3>{currentAfh.title}</h3>
                    <h4>{datetime}</h4>
                    <p>
                        Location: {location.title}<br/>
                        {location.street}<br/>
                        {location.city + ", " + location.state + " " + location.zipcode}<br/>
                        {location.room}
                    </p>
                </div>
            </div>
        );
    }

    render() {
        return (
            <div>
                {this.props.classId && this.renderForClass()}
                {this.props.afhId && this.renderForAfh()}
            </div>
        );
    }
}