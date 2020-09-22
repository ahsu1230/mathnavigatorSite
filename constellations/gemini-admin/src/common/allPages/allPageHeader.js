"use strict";
require("./allPageHeader.sass");
import React from "react";
import { Link } from "react-router-dom";

/*
 * A component that automatically handles the Header portion of an entity's "All-View" page.
 * This includes the header, description, and "Add" button.
 *
 * Available props for this Component:
 *
 * - title - display name of the header of the page (e.g. All Programs)
 * - addUrl - the url to the page when the user clicks on the "Add _____" button
 * - addButtonTitle - the display name of the button when user wants to add an entity
 * - description - a short description of what an entity represents.
 */
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
