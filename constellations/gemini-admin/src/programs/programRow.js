"use strict";
require("./programRow.styl");
import React from "react";
import { Link } from "react-router-dom";

export class ProgramRow extends React.Component {
    render() {
        const row = this.props.row;

        return (
            <div>
                <span className="medium">{row.programId}</span>
                <span className="medium">{row.name}</span>
                <span className="small">{row.grade1}</span>
                <span className="small">{row.grade2}</span>
                <Link to={"/program/" + row.programId + "/edit"}>Edit</Link>
            </div>
        );
    }
}
