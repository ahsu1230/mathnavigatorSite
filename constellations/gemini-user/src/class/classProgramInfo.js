"use strict";
require("./classProgramInfo.sass");
import React from "react";

export class ClassProgramInfo extends React.Component {
    render() {
        const program = this.props.program;
        return (
            <section id="program-info">
                <h4>Program Description:</h4>
                <p className="grades">
                    For grades: {program.grade1} - {program.grade2}
                </p>
                <p className="description">{program.description}</p>
            </section>
        );
    }
}
