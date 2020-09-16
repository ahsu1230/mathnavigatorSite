"use strict";
require("./allPageHeader.sass");
import React from "react";
import { Link } from "react-router-dom";

export default class AllPageHeader extends React.Component {
    render() {
        return (
            <div className="all-page-header">
                <div className="title-container">
                    <h1>{this.props.title}</h1>
                    <Link to={this.props.addUrl}>
                        <button>{this.props.addButtonTitle}</button>
                    </Link>
                </div>
                <p className="description">{this.props.description}</p>
            </div>
        );
    }
}
