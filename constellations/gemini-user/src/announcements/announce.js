"use strict";
require("./announce.sass");
import React from "react";
import API from "../utils/api.js";

export class AnnouncePage extends React.Component {
    state = {
        announcementList: [],
    };
    
    render() {
        return (
            <div id="view-achieve">
                <h1>Announcements</h1>
            </div>
        );
    }
}
