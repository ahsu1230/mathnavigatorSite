"use strict";
require("./home.sass");
import React from "react";
import { Link } from "react-router-dom";
// import { HomeAnnounce } from "./homeAnnounce.js";
import { HomeSectionPrograms } from "./homePrograms.js";
import { HomeSectionStories } from "./homeStories.js";
import { HomeSectionSuccess } from "./homeSuccess.js";
import { HomeAnnounce } from "./homeAnnounce.js";

export class HomePage extends React.Component {
    render() {
        return (
            <div id="view-home">
                <div id="view-home-container">
                    {/* <HomeAnnounce /> */}
                    <HomeBanner />
                    <HomeSectionPrograms />
                    <HomeSectionSuccess />
                    <HomeSectionStories />
                    <HomeAnnounce />
                </div>
            </div>
        );
    }
}

class HomeBanner extends React.Component {
    render() {
        return (
            <div id="home-banner-container">
                <div id="banner-bg-img"></div>
                <div id="banner-bg-overlay"></div>
                <div id="banner-content">
                    <h2>Montgomery County, MD</h2>
                    <h1>
                        Providing affordable, high quality education
                        <br />
                        to thousands of students for 15 years.
                    </h1>
                </div>
                <Link to="/programs">
                    <button>Join a Program today</button>
                </Link>
            </div>
        );
    }
}
