"use strict";
require("./home.styl");
import React from "react";
import ReactDOM from "react-dom";
import API from "../api.js";

export class HomePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            programs: "",
            classes: "",
            locations: "",
            achieves: "",
            semesters:"",
            sessions: "",
            text: [],
        };
        this.onClickPrograms = this.onClickPrograms.bind(this);
    }

    componentDidMount() {
        API.get("api/v1/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                programs: unpublishedList.programs,
                classes: unpublishedList.classes,
                locations: unpublishedList.locations,
                achieves: unpublishedList.achieves,
                semesters: unpublishedList.semesters,
                sessions: unpublishedList.sessions,
            });
        });
    }

    onClickPrograms() {
        this.setState({text: this.state.programs});
    }

    render() {
        return (
            <div id="view-home">
                <div id="view-content">
                    <h2 id="unpublished-heading"> Unpublished Content </h2>
                    <ul>
                        <button onClick={this.onClickPrograms}>Programs</button>
                        <button>Classes</button>
                        <button>Locations</button>
                        <button>Achievements</button>
                        <button>Semesters</button>
                        <button>Sessions</button>
                    </ul>
                    <h2> Registrations </h2>
                    <ul>
                        <button>New Users</button>
                        <button>Questions</button>
                        <button>Complaints</button>
                    </ul>
                </div>
                <div id="box-and-button">
                    <div className="boxed">
                        {this.state.text}
                    </div>
                    <button id="go-to-page">Go to Page</button>
                </div>
            </div>
        );
    }
}
