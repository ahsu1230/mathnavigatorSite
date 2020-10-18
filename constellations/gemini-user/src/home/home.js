"use strict";
require("./home.sass");
import React from "react";
import HomeBanner from "./homeBanner.js";
import HomeSectionPrograms from "./homePrograms.js";
import HomeSectionStrategy from "./homeStrategy.js";
import HomeSectionFooter from "./homeFooter.js";
import HomeSectionOnline from "./homeOnline.js";
import { HomeAnnounce } from "./homeAnnounce.js";

export class HomePage extends React.Component {
    render() {
        return (
            <div id="view-home">
                <div id="view-home-container">
                    <HomeBanner />
                    <HomeSectionPrograms />
                    <HomeSectionStrategy />
                    <HomeSectionOnline />
                    <HomeSectionFooter />
                    <HomeAnnounce />
                </div>
            </div>
        );
    }
}
