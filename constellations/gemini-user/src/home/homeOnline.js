"use strict";
require("./homeOnline.sass");
import React from "react";
import srcZoomClassroom from "../../assets/zoom_classroom.jpg";

export default class HomeSectionOnline extends React.Component {
    render() {
        return (
            <div className="section online">
                <h2>Our Online-only Solution</h2>
                <div className="content">
                    <img src={srcZoomClassroom} />
                    <p>
                        With the impact of Covid-19 and the community mandate to
                        practice safe social distancing, all Math Navigator
                        courses have been moved online!
                        <br />
                        <br />
                        Using Google Classroom to manage classwork and Zoom
                        video conferencing, teachers and students can share
                        materials and discuss hard problems without
                        interference. Despite the inconvenience of schools and
                        community centers closing down, we are still dedicated
                        to providing superior education.
                    </p>
                </div>
            </div>
        );
    }
}
