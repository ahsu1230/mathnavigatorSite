"use strict";
require("./home.sass");
require("./homeBanner.sass");
import React from "react";
import { Link } from "react-router-dom";

export default class HomeBanner extends React.Component {
    render() {
        return (
            <div id="home-banner-container">
                <div id="banner-bg-img"></div>
                <div id="banner-bg-overlay"></div>
                <div id="banner-content">
                    <div className="texts">
                        <h1>Navigate to success with Math Navigator</h1>
                        <p>
                            Providing affordable, high quality education
                            <br />
                            to thousands of students for 15 years.
                        </p>
                        <p>Montgomery County, MD</p>
                    </div>
                    <div className="buttons">
                        <Link to="/register">
                            <button className="enroll">
                                Enroll into a Class
                            </button>
                        </Link>
                        <Link to="/programs">
                            <button className="browse">
                                Browse Program Catalog
                            </button>
                        </Link>
                    </div>
                </div>
            </div>
        );
    }
}
