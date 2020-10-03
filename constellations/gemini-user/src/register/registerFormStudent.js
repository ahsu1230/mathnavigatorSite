"use strict";
require("./registerForm.sass");
import React from "react";
import { 
    RegisterSectionBase,
    REGISTER_SECTION_SELECT,
    REGISTER_SECTION_FORM_GUARDIAN,
    REGISTER_SECTION_CONFIRM,
} from "./registerBase.js";
import RegisterInput from "./registerInput.js";
import { validateEmail } from "../utils/validators.js";

export default class RegisterSectionFormStudent extends React.Component {
    onChangeInput = (e, fieldName) => {
        this.props.onChangeStateValue(fieldName, e.target.value);
    }

    validateAllFields = () => {
        return this.props.student.firstName 
            && this.props.student.lastName
            && this.props.student.email
            && this.props.student.school
            && this.props.student.graduationYear > 2019
            && this.props.student.graduationYear < 2030
            && this.props.student.grade > 5
            && this.props.student.grade < 13;
    }

    renderContent = () => {
        return (
            <div className="content">
                <p> 
                    Please fill out your student information below.
                    We use this information to contact you about important class updates,
                    so please use a valid email you frequently use.
                    School information is used to help us estimate your current level and help us cater to your needs.
                    This information is for our purposes only and will NOT be shared with anyone.<br/>
                    All fields are required.
                </p>
                <div className="names">
                    <RegisterInput
                        title="First Name"
                        value={this.props.student.firstName}
                        placeholder="i.e. Alice"
                        onChangeCallback={(e) => this.onChangeInput(e, "studentFirstName")}
                        />
                    <RegisterInput
                        title="Last Name"
                        value={this.props.student.lastName}
                        placeholder="i.e. Kim"
                        onChangeCallback={(e) => this.onChangeInput(e, "studentLastName")}
                        />
                </div>

                <RegisterInput
                    className="email"
                    title="Email"
                    value={this.props.student.email}
                    placeholder="i.e. alicekim@gmail.com"
                    onChangeCallback={(e) => this.onChangeInput(e, "studentEmail")}
                    validators={[
                        { 
                            validate: () => { return this.props.student.email != "" && validateEmail(this.props.student.email)},
                            message: "You must input a valid email." 
                        }
                    ]}
                    />

                <RegisterInput
                    className="school"
                    title="School"
                    value={this.props.student.school}
                    placeholder="i.e. Winston Churchill High School"
                    onChangeCallback={(e) => this.onChangeInput(e, "studentSchool")}
                />
                
                <div className="grades">
                    <RegisterInput
                        className="grade"
                        title="Grade"
                        value={this.props.student.grade}
                        placeholder="i.e. Alice Kim"
                        onChangeCallback={(e) => this.onChangeInput(e, "studentGrade")}
                        validators={[
                            { 
                                validate: () => { 
                                    const grade = this.props.student.grade;
                                    return parseInt(grade) >= 6 && parseInt(grade) <= 12;
                                },
                                message: "You must input a valid grade." 
                            }
                        ]}
                        />
                    <RegisterInput
                        className="grad-year"
                        title="Graduation Year"
                        value={this.props.student.graduationYear}
                        onChangeCallback={(e) => this.onChangeInput(e, "studentGraduationYear")}
                        validators={[
                            { 
                                validate: () => { 
                                    const gradYear = this.props.student.graduationYear;
                                    return parseInt(gradYear) >= 2000 && parseInt(gradYear) <= 2030;
                                },
                                message: "You must input a valid graduation year." 
                            }
                        ]}
                        />
                </div>
                
            </div>
        );
    }

    render() {
        const nextSection = this.props.isAfh ? REGISTER_SECTION_CONFIRM : REGISTER_SECTION_FORM_GUARDIAN;
        return (
            <RegisterSectionBase
                sectionName="form-student"
                title={"Student Information"}
                nextAllowed={this.validateAllFields()}
                next={nextSection}
                prev={REGISTER_SECTION_SELECT}
                content={this.renderContent()}
                onChangeSection={this.props.onChangeSection}
            />
        );
    }
}