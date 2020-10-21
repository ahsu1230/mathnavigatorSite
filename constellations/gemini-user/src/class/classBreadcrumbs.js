"use strict";
require("./classBreadcrumbs.sass");
import React from "react";
import { Link } from "react-router-dom";
import { capitalizeWord } from "../utils/displayUtils.js";
import { getFullTitle, displayFeaturedString } from "../utils/classUtils.js";
import srcStar from "../../assets/star_green.svg";

export class ClassBreadcrumbs extends React.Component {
    render() {
        const classObj = this.props.classObj;
        const program = this.props.program;
        const semester = this.props.semester;
        const fullTitle = getFullTitle(program, classObj);
        return (
            <section id="breadcrumbs">
                <div>
                    <div className="links">
                        <Link to="/programs">Program Catalog</Link>
                        <span>&middot;</span>
                        <span>{semester.title}</span>
                    </div>
                    <h1>{fullTitle}</h1>
                </div>

                {program.featured != "none" && (
                    <div className="featured">
                        <div className="header">
                            <img src={srcStar} />
                            <span>{capitalizeWord(program.featured)}</span>
                        </div>
                        <p>{displayFeaturedString(program)}</p>
                    </div>
                )}
            </section>
        );
    }
}
