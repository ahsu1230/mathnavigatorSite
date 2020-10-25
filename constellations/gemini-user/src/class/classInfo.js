"use strict";
require("./classInfo.sass");
import React from "react";
import moment from "moment";
import { createLocation } from "../utils/locationUtils.js";
import { displayPrice, displayTimeStringOneLine } from "../utils/classUtils.js";

export class ClassInfo extends React.Component {
    render() {
        const classObj = this.props.classObj;
        const location = this.props.location;
        const sessions = this.props.sessions;
        let firstSession;
        let lastSession;
        if (sessions.length) {
            let firstSessionDate = moment(sessions[0].startsAt).format("l");
            let lastSessionDate = moment(
                sessions[sessions.length - 1].startsAt
            ).format("l");
            firstSession = <p>First session: {firstSessionDate}</p>;
            lastSession = <p>Last session: {lastSessionDate}</p>;
        } else {
            firstSession = <p>First session: To be determined</p>;
            lastSession = <p>Last session: To be determined</p>;
        }

        // Pricing information
        const priceString = displayPrice(classObj);

        return (
            <section id="class-info">
                <div className="block location">
                    <h3>Location</h3>
                    <div>{createLocation(location)}</div>
                </div>

                <div className="block class-times">
                    <h3>Times</h3>
                    {displayTimeStringOneLine(classObj)}
                    <div className="sessions">
                        {firstSession}
                        {lastSession}
                    </div>
                </div>

                <div className="block class-price">
                    <h3>Pricing</h3>
                    <div>
                        <p>{priceString}</p>
                        <p>{classObj.paymentNotes}</p>
                    </div>
                </div>
            </section>
        );
    }
}
