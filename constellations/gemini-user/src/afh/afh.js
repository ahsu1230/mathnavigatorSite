"use strict";
require("./afh.sass");
import React from "react";
import API from "../utils/api.js";

export class AFHPage extends React.Component {
    state = {
        currentTab: "",
    };

    openSubject = (subjectName) => {
        this.setState({
            currentTab: subjectName,
        });
    };

    render() {
        return (
            <div id="view-afh">
                <h1>Ask for Help</h1>
                <div className="description">
                    We provide free sessions for students to ask for additional
                    assistance on any of our program subjects. Please fill the
                    form to let us know you are coming. You must be registered
                    with one of our programs to attend.
                </div>

                <h1>Ask for Help Sessions by Subject</h1>
                <div class="tab">
                    <button
                        className="tabButton"
                        onClick={() => this.openSubject("Math")}>
                        Math
                    </button>
                    <button
                        className="tabButton"
                        onClick={() => this.openSubject("Eng")}>
                        English
                    </button>
                    <button
                        className="tabButton"
                        onClick={() => this.openSubject("Comp")}>
                        Computer Programming
                    </button>
                </div>

                <div
                    className={
                        ((id = "defaultOpen"),
                        this.state.currentTab == "Math" ? "showTab" : "hide")
                    }>
                    <p>testing 1</p>
                </div>

                <div
                    className={
                        this.state.currentTab == "Eng" ? "showTab" : "hide"
                    }>
                    <p>testing 2</p>
                </div>

                <div
                    className={
                        this.state.currentTab == "Comp" ? "showTab" : "hide"
                    }>
                    <p>testing 3</p>
                </div>
            </div>
        );
    }
}
