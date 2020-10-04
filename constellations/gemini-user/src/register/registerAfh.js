"use strict";
require("./register.sass");
import React from "react";
import moment from "moment";
import srcCheckmark from "../../assets/checkmark_light_blue.svg";

export default class RegisterSelectAfh extends React.Component {
    onChangeAfh = (e) => {
        this.props.onChangeStateValue("selectedClassId", null);
        this.props.onChangeStateValue("selectedAfhId", e.target.value);
    };

    render() {
        const afhs = this.props.afhs || [];
        const optionsAfh = afhs.map((afh, index) => {
            const afhTime = moment(afh.startsAt).format("MM/DD/YY h:mm a");
            const fullTitle = afh.title + " " + afhTime;
            return (
                <option key={index} value={afh.id}>
                    {fullTitle}
                </option>
            );
        });

        let selected = <div></div>;
        if (this.props.afhId) {
            const currentAfh = this.props.afhMap[this.props.afhId];
            const datetime =
                moment(currentAfh.startsAt).format("MM/DD/YY h:mm a") +
                " - " +
                moment(currentAfh.endsAt).format("h:mm a");
            const location = this.props.locationMap[currentAfh.locationId];
            selected = (
                <div className="selection">
                    You have selected to attend:
                    <div className="info">
                        <h3>{currentAfh.title}</h3>
                        <h4>{datetime}</h4>
                        <p>
                            Location: {location.title}
                            <br />
                            {location.street}
                            <br />
                            {location.city +
                                ", " +
                                location.state +
                                " " +
                                location.zipcode}
                            <br />
                            {location.room}
                        </p>
                    </div>
                </div>
            );
        }

        return (
            <section className="select afh">
                <div
                    className={
                        "header-wrapper" + (this.props.afhId ? " active" : "")
                    }>
                    <div className="title">
                        <div className="step-wrapper">1</div>
                        <h2>Select an Ask-for-Help session to attend:</h2>
                    </div>
                    {this.props.afhId && (
                        <div>
                            <img src={srcCheckmark} />
                        </div>
                    )}
                </div>
                <select
                    value={this.props.afhId || "none"}
                    onChange={this.onChangeAfh}>
                    <option disabled value={"none"}>
                        -- Select a session --
                    </option>
                    {optionsAfh}
                </select>
                {selected}
            </section>
        );
    }
}
