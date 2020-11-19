"use strict";
require("./inputX.sass");
import React from "react";
import Checkmark from "../../../assets/checkmark_green.svg";

/*
 * Given a list of options, create a single component that manages a dropdown menu.
 *
 * Available props for the InputSelect Component:
 *
 * value: Current selected option
 *
 * onChangeCallback: Function that is called when the selection changes.
 *
 * required: Whether or not the selection is required (omit if false).
 *
 * hasNoDefault: If true, an extra option -- Select an option -- will be added as the default.
 *                  It cannot be reselected after another option is chosen.
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
 *  errorMessageIfEmpty: If options is empty, then this message will be displayed instead of the select.
 */
export class InputSelect extends React.Component {
    state = {
        chosen: false,
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

    renderSelect = (options, value) => {
        if (options.length) {
            var optionElements = options.map((option, index) => (
                <option key={index} value={option.value}>
                    {option.displayName}
                </option>
            ));

            var defaultOption = this.props.hasNoDefault ? (
                <option disabled selected value>
                    -- Select an option --
                </option>
            ) : null;

            value =
                this.props.hasNoDefault && !this.state.chosen
                    ? undefined
                    : value;

            return (
                <select value={value} onChange={(e) => this.onChangeSelect(e)}>
                    {defaultOption}
                    {optionElements}
                </select>
            );
        } else {
            return (
                <h4 className="select-error red">
                    {this.props.errorMessageIfEmpty}
                </h4>
            );
        }
    };

    render = () => {
        const required = this.props.required;
        const value = this.props.value;
        const options = this.props.options || [];
        const pass =
            options.length && (this.state.chosen || !this.props.hasNoDefault);

        var formatDescription = this.renderDescription(
            this.props.description,
            required,
            pass
        );
        var select = this.renderSelect(options, value);

        return (
            <div className="input-wrapper">
                <h2 className="input-label">{this.props.label}</h2>
                {formatDescription}
                <div className="inputs">
                    {select}
                    {pass && required ? <img src={Checkmark} /> : <span></span>}
                </div>
            </div>
        );
    };
}
