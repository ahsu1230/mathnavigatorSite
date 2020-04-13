"use strict";
require("./home.styl");
import React from "react";
import ReactDOM from "react-dom";

export class HomePage extends React.Component {
    render() {
        return (
            <div id="view-home">
                <div id="view-content">
                    <h2 id="unpublished-heading"> Unpublished Content </h2>
                    <ul>
                        <li>Programs</li>
                        <li>Locations</li>
                        <li>Achievements</li>
                    </ul>
                    <h2> Registrations </h2>
                    <ul>
                        <li>New Users</li>
                        <li>Questions</li>
                        <li>Complaints</li>
                    </ul>
                </div>
                <div id="box-and-button">
                    <div className="boxed">text</div>
                    <button id="go-to-page">Go to Page</button>
                </div>
            </div>
        );
    }
}
