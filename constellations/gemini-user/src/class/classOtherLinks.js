"use strict";
import React from "react";

// Unused for now. Looks awkward to be one of the first things you see on page.
export class ClassOtherLinks extends React.Component {
    render() {
        const thisClass = this.props.classObj;
        const otherClasses = this.props.otherClasses.filter((classObj) => {
            return (
                classObj.classId != thisClass.classId &&
                classObj.fullState != FULL_STATE_VALUE
            );
        });

        let semesterMap = {};
        this.props.allSemesters.forEach((semester) => {
            semesterMap[semester.semesterId] = semester;
        });
        const rows = otherClasses.map((classObj, index) => {
            return (
                <div key={index}>
                    <Link
                        to={"/class/" + classObj.classId}
                        onClick={() => this.fetchData(classObj.classId)}>
                        {semesterMap[classObj.semesterId].title}{" "}
                        {capitalizeWord(classObj.classKey)}
                    </Link>
                    {displayTimeStringOneLine(classObj)}
                </div>
            );
        });

        let message = "You may also be interested in these similar classes.";
        if (thisClass.fullState == FULL_STATE_VALUE) {
            message =
                "Unfortunately this class is full. " +
                "Please consider enrolling in one of these available similar classes.";
        }

        const hide = otherClasses.length == 0 ? "hide" : "";

        return (
            <section id="other-classes" className={hide}>
                <p>{message}</p>
                {rows}
            </section>
        );
    }
}
