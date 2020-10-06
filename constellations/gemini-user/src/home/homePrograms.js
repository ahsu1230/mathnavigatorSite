"use strict";
require("./homeProgram.sass");
import React from "react";
import { Link } from "react-router-dom";
import { programsText } from "./homeText.js";

export default class HomeSectionPrograms extends React.Component {
    render() {
        return (
            <div className="section programs">
                <h2>Our Featured Programs</h2>
                <div className="card-container">
                    <ProgramCard program={programsText[0]} className="sat" />
                    <ProgramCard program={programsText[1]} className="ap" />
                    <ProgramCard program={programsText[2]} className="amc" />
                </div>

                <Link to="/programs" className="link-bar">
                    Visit our <b>Program Catalog</b> to view more programs to
                    enroll into.
                </Link>
            </div>
        );
    }
}

class ProgramCard extends React.Component {
    render() {
        var program = this.props.program;
        return (
            <div className="card">
                <div className="img-container">
                    <img src={program.imgSrc} />
                </div>
                <h3>{program.title}</h3>
                <p>{program.description}</p>

                <Link to="/programs">
                    <button>Enroll Now</button>
                </Link>
            </div>
        );
    }
}
