"use strict";
require("./registerForm.sass");
import React from "react";
import {
    RegisterSectionBase,
    REGISTER_SECTION_FORM_STUDENT,
    REGISTER_SECTION_CONFIRM,
} from "./registerBase.js";
import RegisterInput from "./registerInput.js";

export default class RegisterSectionFormGuardian extends React.Component {
    onChangeInput = (e, fieldName) => {
        this.props.onChangeStateValue(fieldName, e.target.value);
    }

    validateAllFields = () => {
        return this.props.guardian.firstName 
            && this.props.guardian.lastName
            && this.props.guardian.email
            && this.props.guardian.phone;
    }

    renderContent = () => {
        return (
            <div className="content">
                <p> 
                    Please fill out your guardian information below.
                    We use this information to contact you about important class updates,
                    so please use a valid email you frequently use.
                    This information is for our purposes only and will NOT be shared with anyone.
                </p>
                <div className="names">
                    <RegisterInput
                        title="First Name"
                        value={this.props.guardian.firstName}
                        placeholder="i.e. Alice"
                        onChangeCallback={(e) => this.onChangeInput(e, "guardianFirstName")}
                        />
                    <RegisterInput
                        title="Last Name"
                        value={this.props.guardian.lastName}
                        placeholder="i.e. Kim"
                        onChangeCallback={(e) => this.onChangeInput(e, "guardianLastName")}
                        />
                </div>
                <RegisterInput
                    className="email"
                    title="Email"
                    value={this.props.guardian.email}
                    placeholder="i.e. alicekim@gmail.com"
                    onChangeCallback={(e) => this.onChangeInput(e, "guardianEmail")}
                    />

                <RegisterInput
                    className="phone"
                    title="Phone Number"
                    value={this.props.guardian.phone}
                    placeholder="i.e. (XXX) XXX - XXXX"
                    onChangeCallback={(e) => this.onChangeInput(e, "guardianPhone")}
                />
            </div>
        );
    }

    render() {
        return (
            <RegisterSectionBase
                sectionName="form-guardian"
                title={"Guardian Information"}
                nextAllowed={this.validateAllFields()}
                next={REGISTER_SECTION_CONFIRM}
                prev={REGISTER_SECTION_FORM_STUDENT}
                content={this.renderContent()}
                onChangeSection={this.props.onChangeSection}
            />
        );
    }
}