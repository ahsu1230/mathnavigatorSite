"use strict";
require("./error.sass");
import React from "react";
import { Link } from 'react-router-dom';
import srcBroken from "../../assets/compass_broken.png";

export class ErrorPage extends React.Component {
	render = () => {
        const classId = this.props.classId;
        const errorMsg = classId ? "Class ID '" + classId + "' does not exist." : "";

        return (
            <div id="view-error">
                <h1>Page Not Found</h1>
                <img src={srcBroken} />
                <div>
                    <Link to="/programs">View our Programs</Link>
                    To find what you're looking for
                </div>
                <span>{errorMsg}</span>
            </div>
        );
    }
}
