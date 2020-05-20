"use strict";
require("./achievements.sass");
import React from "react";
import API from "../utils/api.js";

export class AchievementPage extends React.Component {
    state = {
        achieveList: [],
    };

    componentDidMount() {
        console.log("api attempt ");
        API.get("api/achievements/all").then((res) => {
            const achieveList = res.data;
            console.log("api success!");
            this.setState({ achieveList });
        });
    }

    render() {
        return (
            <div id="view-achieve">
                <h1>Math Navigator Achievements</h1>
                <p>{JSON.stringify(this.state.achieveList)}</p>
            </div>
        );
    }
}
