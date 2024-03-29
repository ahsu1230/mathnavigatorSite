"use strict";
require("./registerInput.sass");
import React from "react";

/**
 * A field for the register section.
 * Available properties include:
 * - title - label display string on top of input field
 * - description - instruction display string
 * - className - class name of component
 * - value - value of input field
 * - placeholder - placeholder text for input field
 * - isTextArea - will use a textarea element instead of input
 * - onChangeCallback - onChange function for input
 * - validators - an array of objects
 * i.e.
 * [
 *   {
 *     validate: () => {...},
 *     message: "...."
 *   }
 * ]
 */
export default class RegisterInput extends React.Component {
    validateField = () => {
        if (this.props.value === "") {
            return "";
        }

        const validators = this.props.validators || [];
        // find first validate function that is false
        let firstValidator = validators.find((validator) => {
            return !validator.validate();
        });
        return firstValidator ? firstValidator.message : "";
    };

    render() {
        const description = this.props.description ? (
            <p>{this.props.description}</p>
        ) : (
            <div></div>
        );
        const validateMessage = this.validateField();
        const displayMessage = validateMessage ? (
            <h4 className="error">{validateMessage}</h4>
        ) : (
            <div></div>
        );
        return (
            <div className={"register-input " + this.props.className}>
                <h3>{this.props.title}</h3>
                {description}
                {this.props.isTextArea ? (
                    <textarea
                        value={this.props.value}
                        placeholder={this.props.placeholder}
                        onChange={(e) => this.props.onChangeCallback(e)}
                    />
                ) : (
                    <input
                        type="text"
                        value={this.props.value}
                        placeholder={this.props.placeholder}
                        onChange={(e) => this.props.onChangeCallback(e)}
                    />
                )}
                {displayMessage}
            </div>
        );
    }
}
