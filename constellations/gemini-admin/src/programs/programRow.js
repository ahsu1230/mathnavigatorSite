"use strict";
require("./programRow.sass");
import React from "react";
import { Link } from "react-router-dom";

export class ProgramRow extends React.Component {
    render = () => {
        const program = this.props.program;
        const url = "/programs/" + program.programId + "/edit";

        return (
            <div className="row">
                <span className="medium-column">{program.programId}</span>
                <span className="medium-column">{program.name}</span>
                <span className="small-column">{program.grade1}</span>
                <span className="small-column">{program.grade2}</span>
                <span className="large-column">{program.description}</span>
                <Link to={url}>{"Edit >"}</Link>
            </div>
        );
    };
}
