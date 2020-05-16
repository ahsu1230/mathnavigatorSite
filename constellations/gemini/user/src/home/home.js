"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";

export class HomePage extends React.Component {
    state = {
        programsList: [],
    };

    componentDidMount() {
        console.log("api attempt");
        API.get("api/programs/all").then((res) => {
            const programsList = res.data;
            console.log("api success!");
            this.setState({ programsList });
        });
    }

    render() {
        return (
            <div id="view-home">
                <h1>Math Navigator</h1>
                <p>{JSON.stringify(this.state.programsList)}</p>
            </div>
        );
    }
}
