"use strict";
require("./inputText.sass");
import React from "react";
import Checkbox from "../../assets/checkmark_green.svg";

export class InputText extends React.Component {
    renderErrorMessage = (required, value) => {
        var pass = true;
        var errorMessage = <h4 className="hidden"></h4>;
        if (required) {
            this.props.validators.some((validator) => {
                if (!validator.validate(value)) {
                    errorMessage = <h4 className="red">{validator.message}</h4>;
                    pass = false;
                    return true;
                }
            });
        } else {
            pass = false;
        }
        return [errorMessage, pass];
    };

    renderDescription = (description, required, pass) => {
        var formatDescription = <h4 className="hidden"></h4>;
        if (!!description) {
            let ending = required ? (
                <span className={pass ? "" : "red"}>{" (required)"}</span>
            ) : (
                " (optional)"
            );
            formatDescription = (
                <h4>
                    {description} {ending}
                </h4>
            );
        }
        return formatDescription;
    };

    renderInput = (pass, value) => {
        var input = (
            <input
                className={pass ? "blue" : ""}
                type="text"
                value={value}
                onChange={(e) => this.props.onChangeCallback(e)}
            />
        );
        if (this.props.isTextBox) {
            input = (
                <textarea
                    className={pass ? "blue" : ""}
                    value={value}
                    onChange={(e) => this.props.onChangeCallback(e)}
                />
            );
        }
        return input;
    };

    render = () => {
        const required = this.props.required;
        const value = this.props.value;

        var [errorMessage, pass] = this.renderErrorMessage(required, value);
        var formatDescription = this.renderDescription(
            this.props.description,
            required,
            pass
        );
        var input = this.renderInput(pass, value);

        return (
            <div id="text-input-wrapper">
                <h2>{this.props.label}</h2>
                {formatDescription}
                <div id="input-wrapper">
                    {input}
                    {pass ? <img src={Checkbox} /> : <img />}
                </div>
                {errorMessage}
            </div>
        );
    };
}

export const emptyValidator = (label) => {
    return {
        validate: (x) => x != "",
        message: "You must input a " + label,
    };
};
