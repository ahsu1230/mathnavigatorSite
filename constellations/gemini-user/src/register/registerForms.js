"use strict";
require("./register.sass");
import React from "react";
import RegisterInput from "./registerInput.js";
import {
    validateEmail,
    validatePhone,
    validateGrade,
    validateGradYear,
} from "../utils/validators.js";
import srcCheckmark from "../../assets/checkmark_light_blue.svg";

export class RegisterFormStudent extends React.Component {
    onChangeInput = (e, fieldName) => {
        this.props.onChangeStateValue(fieldName, e.target.value);
    };

    render() {
        return (
            <section className="register-form student">
                <div
                    className={
                        "header-wrapper" + (this.props.valid ? " active" : "")
                    }>
                    <div className="title">
                        <div className="step-wrapper">2</div>
                        <h2>Student Information</h2>
                    </div>
                    {this.props.valid && (
                        <div>
                            <img src={srcCheckmark} />
                        </div>
                    )}
                </div>

                {/* <p className="instructions"> 
                    Please fill out your student information below.
                    <br/>
                    All fields are required.
                </p> */}
                <div className="names">
                    <RegisterInput
                        title="First Name"
                        value={this.props.student.firstName}
                        placeholder="i.e. Alice"
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "studentFirstName",
                                e.target.value
                            )
                        }
                    />
                    <RegisterInput
                        title="Last Name"
                        value={this.props.student.lastName}
                        placeholder="i.e. Kim"
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "studentLastName",
                                e.target.value
                            )
                        }
                    />
                </div>

                <RegisterInput
                    className="email"
                    title="Email"
                    value={this.props.student.email}
                    placeholder="i.e. alicekim@gmail.com"
                    onChangeCallback={(e) =>
                        this.props.onChangeStateValue(
                            "studentEmail",
                            e.target.value
                        )
                    }
                    validators={[
                        {
                            validate: () => {
                                return validateEmail(this.props.student.email);
                            },
                            message: "You must input a valid email.",
                        },
                    ]}
                />

                <RegisterInput
                    className="school"
                    title="School"
                    value={this.props.student.school}
                    placeholder="i.e. Winston Churchill High School"
                    onChangeCallback={(e) =>
                        this.props.onChangeStateValue(
                            "studentSchool",
                            e.target.value
                        )
                    }
                />

                <div className="grades">
                    <RegisterInput
                        className="grade"
                        title="Grade"
                        value={this.props.student.grade}
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "studentGrade",
                                parseInt(e.target.value) || 0
                            )
                        }
                        validators={[
                            {
                                validate: () => {
                                    const grade =
                                        parseInt(this.props.student.grade) || 0;
                                    return validateGrade(grade);
                                },
                                message: "You must input a valid grade number.",
                            },
                        ]}
                    />
                    <RegisterInput
                        className="grad-year"
                        title="Graduation Year"
                        value={this.props.student.graduationYear}
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "studentGraduationYear",
                                parseInt(e.target.value) || 0
                            )
                        }
                        validators={[
                            {
                                validate: () => {
                                    const gradYear =
                                        parseInt(
                                            this.props.student.graduationYear
                                        ) || 0;
                                    return validateGradYear(gradYear);
                                },
                                message:
                                    "You must input a valid graduation year.",
                            },
                        ]}
                    />
                </div>
            </section>
        );
    }
}

export class RegisterFormGuardian extends React.Component {
    render() {
        return (
            <section className="register-form guardian">
                <div
                    className={
                        "header-wrapper" + (this.props.valid ? " active" : "")
                    }>
                    <div className="title">
                        <div className="step-wrapper">3</div>
                        <h2>Guardian Information</h2>
                    </div>
                    {this.props.valid && (
                        <div>
                            <img src={srcCheckmark} />
                        </div>
                    )}
                </div>
                {/* <p className="instructions"> 
                    Please fill out your guardian information below.
                    We use this information to contact you about important class updates,
                    so please use a valid email you frequently use.<br/>
                </p> */}
                <div className="names">
                    <RegisterInput
                        title="First Name"
                        value={this.props.guardian.firstName}
                        placeholder="i.e. Alice"
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "guardianFirstName",
                                e.target.value
                            )
                        }
                        validators={[
                            {
                                validate: () => {
                                    return (
                                        this.props.student.firstName !=
                                            this.props.guardian.firstName ||
                                        this.props.student.lastName !=
                                            this.props.guardian.lastName
                                    );
                                },
                                message:
                                    "Your name must be different from the student's name.",
                            },
                        ]}
                    />
                    <RegisterInput
                        title="Last Name"
                        value={this.props.guardian.lastName}
                        placeholder="i.e. Kim"
                        onChangeCallback={(e) =>
                            this.props.onChangeStateValue(
                                "guardianLastName",
                                e.target.value
                            )
                        }
                    />
                </div>
                <RegisterInput
                    className="email"
                    title="Email"
                    value={this.props.guardian.email}
                    placeholder="i.e. alicekim@gmail.com"
                    onChangeCallback={(e) =>
                        this.props.onChangeStateValue(
                            "guardianEmail",
                            e.target.value
                        )
                    }
                    validators={[
                        {
                            validate: () => {
                                const email = this.props.guardian.email;
                                return validateEmail(email);
                            },
                            message: "You must input a valid email.",
                        },
                        {
                            validate: () => {
                                const studentEmail = this.props.student.email;
                                const email = this.props.guardian.email;
                                return email != "" && email != studentEmail;
                            },
                            message:
                                "You must input an email different from the student email.",
                        },
                    ]}
                />

                <RegisterInput
                    className="phone"
                    title="Phone Number"
                    value={this.props.guardian.phone}
                    placeholder="i.e. (XXX) XXX - XXXX"
                    onChangeCallback={(e) =>
                        this.props.onChangeStateValue(
                            "guardianPhone",
                            e.target.value
                        )
                    }
                    validators={[
                        {
                            validate: () => {
                                const phone = this.props.guardian.phone;
                                return phone != "" && validatePhone(phone);
                            },
                            message: "You must input a valid phone number.",
                        },
                    ]}
                />

                <RegisterInput
                    className="additional-info"
                    title="Additional Info (Optional)"
                    description="If there is any other information you would like to provide, please let us know here."
                    value={this.props.guardian.additionalInfo}
                    placeholder="(Optional)"
                    isTextArea={10}
                    onChangeCallback={(e) =>
                        this.props.onChangeStateValue(
                            "guardianAdditionalInfo",
                            e.target.value
                        )
                    }
                />
            </section>
        );
    }
}
