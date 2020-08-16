"use strict";
require("./achieve.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

export class AchievePage extends React.Component {
    state = {
        achievements: [],
    };

    componentDidMount = () => {
        API.get("api/achievements/years").then((res) =>
            this.setState({ achievements: res.data })
        );
    };

    render = () => {
        let length = 0;
        const rows = this.state.achievements.map((group, index) =>
            (group.achievements || []).map((achieve, index) => {
                length++;
                return <AchieveRow key={index} achieve={achieve} />;
            })
        );

        return (
            <div id="view-achieve">
                <h1>All Achievements ({length})</h1>
                <div className="row header">
                    <span className="medium-column">Year</span>
                    <span className="medium-column">Position</span>
                    <span className="large-column"> Message</span>
                    <span className="edit"></span>
                </div>
                {rows}
                <button>
                    <Link id="add-achievement" to={"/achievements/add"}>
                        Add Achievement
                    </Link>
                </button>
            </div>
        );
    };
}

class AchieveRow extends React.Component {
    render = () => {
        const achieve = this.props.achieve;
        const url = "/achievements/" + achieve.id + "/edit";
        return (
            <div className="row">
                <span className="medium-column">{achieve.year}</span>
                <span className="medium-column">{achieve.position}</span>
                <span className="large-column">{achieve.message}</span>
                <Link className="edit" to={url}>
                    {"Edit >"}
                </Link>
            </div>
        );
    };
}
