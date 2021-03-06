"use strict";
require("./homeOnline.sass");
import React from "react";
const srcZoomClassroom =
    "http://d2dwqi4dzedhxu.cloudfront.net/images/zoom_classroom.jpg";

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
                        video conferencing, teachers and students can discuss
                        and share materials with ease. Despite the inconvenience
                        of schools and community centers closing down, we are
                        still dedicated to providing students with superior
                        education.
                    </p>
                </div>
            </div>
        );
    }
}
