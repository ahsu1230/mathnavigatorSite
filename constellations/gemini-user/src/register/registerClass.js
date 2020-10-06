"use strict";
require("./register.sass");
import React from "react";
import {
    getFullTitle,
    isFullClass,
    displayPrice,
    displayTimeString,
} from "../utils/classUtils.js";
import { capitalizeWord } from "../utils/displayUtils.js";
import { createLocation } from "../utils/locationUtils.js";
import srcCheckmark from "../../assets/checkmark_light_blue.svg";

export default class RegisterSelectClass extends React.Component {
    onChangeClass = (e) => {
        this.props.onChangeStateValue("selectedClassId", e.target.value);
        this.props.onChangeStateValue("selectedAfhId", null);
    };

    render() {
        const classes = this.props.classes || [];
        const optionsClasses = classes.map((classObj, index) => {
            const program = this.props.programMap[classObj.programId];
            const semester = this.props.semesterMap[classObj.semesterId];
            const fullTitle =
                program.title +
                " " +
                capitalizeWord(classObj.classKey) +
                " " +
                semester.title;
            return (
                <option key={index} value={classObj.classId}>
                    {fullTitle}
                </option>
            );
        });

        let selected = <div></div>;
        if (this.props.classId) {
            const currentClass = this.props.classMap[this.props.classId];
            const program = this.props.programMap[currentClass.programId];
            const semester = this.props.semesterMap[currentClass.semesterId];
            const location = this.props.locationMap[currentClass.locationId];

            const fullTitle = getFullTitle(program, currentClass);
            const fullSection = isFullClass(currentClass) ? (
                <p className="error">
                    This class is full. Please select another class to enroll.
                </p>
            ) : (
                <div></div>
            );
            selected = (
                <div className="selection">
                    {fullSection}
                    You have selected to enroll into:
                    <h3>{fullTitle}</h3>
                    <h4>{semester.title}</h4>
                    {displayTimeString(currentClass)}
                    <p className="price">{displayPrice(currentClass)}</p>
                    <p className="payment-notes">{currentClass.paymentNotes}</p>
                    {createLocation(location)}
                </div>
            );
        }

        return (
            <section className="select class">
                <div
                    className={
                        "header-wrapper" + (this.props.classId ? " active" : "")
                    }>
                    <div className="title">
                        <div className="step-wrapper">1</div>
                        <h2>Select a class to enroll into:</h2>
                    </div>
                    {this.props.classId && (
                        <div>
                            <img src={srcCheckmark} />
                        </div>
                    )}
                </div>
                <select
                    value={this.props.classId || "none"}
                    onChange={this.onChangeClass}>
                    <option disabled value={"none"}>
                        -- Select a class --
                    </option>
                    {optionsClasses}
                </select>
                {selected}
            </section>
        );
    }
}
