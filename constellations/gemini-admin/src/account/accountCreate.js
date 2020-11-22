"use strict";
require("./accountCreate.sass");
import React from "react";
import API from "../api.js";
import EditPageWrapper from "../common/editPages/editPageWrapper.js";
import { UserInputs } from "./utils/userInputs.js";

export class AccountCreatePage extends React.Component {
    state = {
        firstName: "",
        middleName: "",
        lastName: "",
        email: "",
        phone: "",
        isGuardian: false,
        accountId: 0,
        school: "",
        graduationYear: 0,
        notes: "",
    };

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    handleChangeIsGuardian = (value) => {
        this.setState({ isGuardian: value == "true" });
    };

    generateRandomPassword = (length) => {
        let result = "";
        const characters =
            "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
        const charactersLength = characters.length;
        for (var i = 0; i < length; i++) {
            result += characters.charAt(
                Math.floor(Math.random() * charactersLength)
            );
        }
        return result;
    };

    onSave = () => {
        const account = {
            primaryEmail: this.state.email,
            password: this.generateRandomPassword(8),
        };
        const user = {
            firstName: this.state.firstName,
            middleName: this.state.middleName,
            lastName: this.state.lastName,
            email: this.state.email,
            phone: this.state.phone,
            isGuardian: this.state.isGuardian,
            school: this.state.school,
            graduationYear: parseInt(this.state.graduationYear),
            notes: this.state.notes,
        };
        return API.post("api/accounts/create", {
            account: account,
            user: user,
        });
    };

    renderContent = () => {
        return (
            <UserInputs
                handleChange={this.handleChange}
                handleChangeIsGuardian={this.handleChangeIsGuardian}
                firstName={this.state.firstName}
                middleName={this.state.middleName}
                lastName={this.state.lastName}
                email={this.state.email}
                phone={this.state.phone}
                isGuardian={this.state.isGuardian}
                school={this.state.school}
                graduationYear={this.state.graduationYear}
                notes={this.state.notes}
            />
        );
    };

    render() {
        return (
            <main id="view-create-account">
                <EditPageWrapper
                    isEdit={false}
                    title={"Create a new Account"}
                    content={this.renderContent()}
                    prevPageUrl={"/accounts"}
                    onSave={this.onSave}
                    entityName={"account"}
                />
            </main>
        );
    }
}
