"use strict";
require("./classError.sass");
import React from "react";
import { Link } from "react-router-dom";
import srcBroken from "../../assets/compass_broken.png";

export class ClassErrorPage extends React.Component {
    render = () => {
        const classId = this.props.classId;
        const errorMsg = classId
            ? "Class ID '" + classId + "' does not exist."
            : "";

        return (
            <div id="view-error">
                <h1>Page Not Found</h1>
                <img src={srcBroken} />
                <div>
                    <Link to="/programs">View our Program Catalog</Link>
                    to find what you're looking for
                </div>
                <span>{errorMsg}</span>
            </div>
        );
    };
}
