"use strict";
require("./home.sass");
import React from "react";
import { Link } from "react-router-dom";

export class HomeSectionStories extends React.Component {
    render() {
        return (
            <div className="section stories">
                <div className="story-banner-container">
                    <div className="story-banner-overlay"></div>
                    <div className="story-banner-img"></div>
                    <div className="text-container">
                        <h2>Succeeding Together</h2>
                        <Link to="/student-achievements">
                            <button>View our Student Achievements</button>
                        </Link>
                    </div>
                </div>
            </div>
        );
    }
}
