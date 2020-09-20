"use strict";
require("./inputX.sass");
import React from "react";
import Checkbox from "../../../assets/checkmark_green.svg";

/*
 * Given a list of radio selections, render a single component
 * that manages this selection of radio input options.
 * Props for the InputRadio Component:
 *
 * value: Current selected option
 *
 * onChangeCallback: Function that is called when the selection changes.
 *
 * required: Whether or not the selection is required (omit if false).
 *
 * label: A large gray label for the select, e.g. "ProgramID".
 *
 * description (optional): A more detailed description for what the selection is.
 *
 * options: A list of option objects:
 *      [
 *          { value: "ap_calc", displayName: "AP Calculus" }
 *      ]
 *  becomes
 *      <option value="ap_calc">AP Calculus</option>
 *  value should not be an empty string.
 *
 * errorMessageIfEmpty: If options is empty, then this message will be displayed instead of the radios.
 */
export class InputRadio extends React.Component {
    state = {
        chosen: !!this.props.value,
    };

    onChangeSelect = (e) => {
        this.setState({ chosen: true });
        this.props.onChangeCallback(e);
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

    renderRadios = (options, value) => {
        if (options.length) {
            return options.map((option, index) => (
                <div key={index} className="input-radio-wrapper">
                    <input
                        className="input-radio"
                        type="radio"
                        value={option.value}
                        checked={value == option.value}
                        onChange={(e) => this.onChangeSelect(e)}
                    />
                    <span>{option.displayName}</span>
                </div>
            ));
        } else {
            return (
                <h4 className="radio-error red">
                    {this.props.errorMessageIfEmpty}
                </h4>
            );
        }
    };

    render = () => {
        const required = this.props.required;
        const value = this.props.value;
        const options = this.props.options;
        const pass = options.length && this.state.chosen;

        var formatDescription = this.renderDescription(
            this.props.description,
            required,
            pass
        );
        var radios = this.renderRadios(options, value);

        return (
            <div className="input-wrapper">
                <h2 className="input-label">{this.props.label}</h2>
                {formatDescription}
                <div className="inputs radios">
                    {radios}
                    {pass ? <img src={Checkbox} /> : <img />}
                </div>
            </div>
        );
    };
}
