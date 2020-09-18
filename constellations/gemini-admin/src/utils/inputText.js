"use strict";
require("./inputText.sass");
import React from "react";
import Checkbox from "../../assets/checkmark_green.svg";

/*
 * Props for the InputText Component:
 *
 * label: A large gray label for the input, e.g. "Program ID".
 *
 * description: A more detailed description for what the input should be.
 * The format of the input could also go here, e.g. "Enter the program ID. Examples: ap_calculus, sat1, ap_java".
 *
 * required: Whether or not the input is required (omit if false).
 * If this input is omitted, the validators prop should not be used either.
 *
 * isTextbox: If true, the input is a large textbox, otherwise it is a single line (omit if false).
 *
 * value: This is the value of the input, e.g. this.state.programId.
 *
 * onChangeCallback: This is the function that is called when the input changes.
 *
 * validators: This is a list of functions that validate the input. The validators are checked in order.
 * The function in validate must return true or else the message will appear, e.g.
 *      validators={[
 *          emptyValidator("grade"),
 *          {
 *              validate: (grade) => parseInt(grade) >= 1 && parseInt(grade) <= 12,
 *              message: "Grade must be between 1 and 12",
 *          },
 *      ]}
 * The emptyValidator(label) function can be imported and checks to make sure the input exists.
 * It returns an error in the form "You must input a <label>" if the input is empty.
 */
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
            <div className="input-wrapper">
                <h2 className="input-label">{this.props.label}</h2>
                {formatDescription}
                <div className="inputs">
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
