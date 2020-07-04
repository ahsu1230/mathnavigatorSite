"use strict";
require("./textInput.sass");
import React from "react";
import Checkbox from "./checkmark_green.svg";

export class TextInput extends React.Component {
    render = () => {
        const description = this.props.description;
        const required = this.props.required;
        const value = this.props.value;

        var format = <h4 className="hidden"></h4>;
        if (!!description) {
            let ending = required ? (
                <span className="red">{" (required)"}</span>
            ) : (
                " (optional)"
            );
            format = (
                <h4>
                    {description} {ending}
                </h4>
            );
        }

        var pass = true;
        var error = <h4 className="hidden"></h4>;
        if (required) {
            this.props.validators.some((validator) => {
                if (!validator.validate(value)) {
                    error = <h4 className="red">{validator.message}</h4>;
                    pass = false;
                    return true;
                }
            });
        } else {
            pass = false;
        }

        return (
            <div id="text-input-wrapper">
                <h2>{this.props.label}</h2>
                {format}
                <div id="input-wrapper">
                    <input
                        className={pass ? "blue" : ""}
                        type="text"
                        value={value}
                        onChange={(e) => this.props.onChangeCallback(e)}
                    />
                    {pass ? <img src={Checkbox} /> : <img />}
                </div>
                {error}
            </div>
        );
    };
}
