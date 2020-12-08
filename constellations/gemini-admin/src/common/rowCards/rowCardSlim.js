"use strict";
require("./rowCard.sass");
import React from "react";
import { Link } from "react-router-dom";

export default class RowCardSlim extends React.Component {
    render() {
        const title = this.props.title;
        const fields = (this.props.fields || []).map((field, index) => (
            <div
                key={index}
                className={
                    field.highlightFn && field.highlightFn()
                        ? "highlighted"
                        : ""
                }>
                {field.value}
            </div>
        ));
        const inlineTitle = this.props.inlineTitle;

        return (
            <div className="row-card slim">
                {title && <p className="title">{title}</p>}

                <div className="row-fields-container">
                    {inlineTitle && (
                        <div className="inline-title">{inlineTitle}</div>
                    )}
                    {fields}
                    {this.props.editUrl && (
                        <Link className="edit-url" to={this.props.editUrl}>
                            Edit >
                        </Link>
                    )}
                </div>

                {this.props.text && <p className="text">{this.props.text}</p>}
            </div>
        );
    }
}
