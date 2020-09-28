"use strict";
require("./homeFooter.sass");
import React from "react";
import { Link } from "react-router-dom";

export default class HomeSectionFooter extends React.Component {
    render() {
        return (
            <div className="section home-footer">
                <div className="content">
                    <h1>Try your first class with Math Navigator today!</h1>
                    <Link to="/enroll">
                        <button>Get Started</button>
                    </Link>
                </div>
            </div>
        );
    }
}
