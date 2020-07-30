"use strict";
require("./userInput.sass");
import React from "react";
import { InputText, emptyValidator } from "../utils/inputText.js";

export class UserInput extends React.Component {
    render = () => {
        const handleChange = this.props.handleChange;

        return (
            <div id="view-user-inputs">
                <InputText
                    label="First Name"
                    description="Enter your first name"
                    required={true}
                    value={this.props.firstName}
                    onChangeCallback={(e) => handleChange(e, "firstName")}
                    validators={[emptyValidator("name")]}
                />

                <InputText
                    label="Middle Name"
                    description="Enter your middle name if applicable"
                    value={this.props.middleName}
                    onChangeCallback={(e) => handleChange(e, "middleName")}
                />

                <InputText
                    label="Last Name"
                    description="Enter your last name"
                    required={true}
                    value={this.props.lastName}
                    onChangeCallback={(e) => handleChange(e, "lastName")}
                    validators={[emptyValidator("name")]}
                />

                <InputText
                    label="Email"
                    description="Enter your email address"
                    required={true}
                    value={this.props.email}
                    onChangeCallback={(e) => handleChange(e, "email")}
                    validators={[
                        emptyValidator("email"),
                        {
                            validate: (email) =>
                                /^[^( @)]+@[^( @)]+\.[^( @)]+$/.test(email),
                            message: "Invalid email address",
                        },
                    ]}
                />

                <InputText
                    label="Phone"
                    description="Enter your phone number"
                    required={true}
                    value={this.props.phone}
                    onChangeCallback={(e) => handleChange(e, "phone")}
                    validators={[
                        emptyValidator("phone number"),
                        {
                            validate: (num) => /^[\d\s+.()/-]{3,}$/.test(num),
                            message: "Invalid phone number",
                        },
                    ]}
                />

                <h2 className="guardian">Is this user a guardian?</h2>
                <input
                    type="checkbox"
                    checked={this.props.isGuardian}
                    onChange={this.props.onChangeGuardian}
                />
                <span>Yes</span>

                <InputText
                    label="School"
                    description="Enter your school"
                    value={this.props.school}
                    onChangeCallback={(e) => handleChange(e, "school")}
                />

                <InputText
                    label="Graduation Year"
                    description="Enter your graduation year"
                    value={this.props.graduationYear}
                    onChangeCallback={(e) => handleChange(e, "graduationYear")}
                />
            </div>
        );
    };
}
