"use strict";
require("./home.sass");
import React from "react";
import { Link } from "react-router-dom";
import { programsText } from "./homeText.js";

export class HomeSectionPrograms extends React.Component {
    render() {
        return (
            <div className="section programs">
                <h2>Programs</h2>
                <div className="card-container">
                    <ProgramCard program={programsText[0]} />
                    <ProgramCard program={programsText[1]} />
                    <ProgramCard program={programsText[2]} />
                </div>

                <div>
                    <Link to="/programs">
                        <button>View More Programs</button>
                    </Link>
                </div>
            </div>
        );
    }
}

class ProgramCard extends React.Component {
    render() {
        var program = this.props.program;
        return (
            <div className="home-tile-card program">
                <div className="img-container">
                    <img src={program.imgSrc} />
                </div>
                <h3>{program.title}</h3>
                <p>{program.description}</p>
                <Link to="/programs">
                    <button className="inverted">Sign Up</button>
                </Link>
            </div>
        );
    }
}
