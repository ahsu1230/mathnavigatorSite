"use strict";
import React from "react";
import { InputText, emptyValidator } from "../../common/inputs/inputText.js";
import { InputRadio } from "../../common/inputs/inputRadio.js";
import { validateEmail, validatePhoneNumber } from "../../common/userUtils.js";

export class UserInputs extends React.Component {
    render() {
        const handleChange = this.props.handleChange;
        return (
            <section>
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
                            validate: (email) => validateEmail(email),
                            message: "Invalid email address",
                        },
                    ]}
                />

                <InputRadio
                    label={"Is this user a guardian?"}
                    value={this.props.isGuardian + ""}
                    onChangeCallback={(e) =>
                        this.props.handleChangeIsGuardian(e.target.value)
                    }
                    required={true}
                    options={[
                        {
                            value: "true",
                            displayName: "Yes",
                        },
                        {
                            value: "false",
                            displayName: "No",
                        },
                    ]}
                />

                {this.props.isGuardian && (
                    <InputText
                        label="Phone"
                        description="Enter your phone number"
                        required={true}
                        value={this.props.phone}
                        onChangeCallback={(e) => handleChange(e, "phone")}
                        validators={[
                            emptyValidator("phone number"),
                            {
                                validate: (phone) => validatePhoneNumber(phone),
                                message: "Invalid phone number",
                            },
                        ]}
                    />
                )}

                {!this.props.isGuardian && (
                    <InputText
                        label="School"
                        description="Enter your school"
                        value={this.props.school}
                        onChangeCallback={(e) => handleChange(e, "school")}
                    />
                )}

                {!this.props.isGuardian && (
                    <InputText
                        label="Graduation Year"
                        description="Enter your graduation year"
                        value={this.props.graduationYear}
                        onChangeCallback={(e) =>
                            handleChange(e, "graduationYear")
                        }
                    />
                )}

                <InputText
                    label="Notes"
                    isTextBox={true}
                    value={this.props.notes}
                    onChangeCallback={(e) => handleChange(e, "notes")}
                    description="Add any notes"
                />
            </section>
        );
    }
}
