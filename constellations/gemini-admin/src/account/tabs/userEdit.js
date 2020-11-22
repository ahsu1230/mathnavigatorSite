"use strict";
import React from "react";
import API from "../../api.js";
import EditPageWrapper from "../../common/editPages/editPageWrapper.js";
import { UserInputs } from "../utils/userInputs.js";

export class UserEditPage extends React.Component {
    state = {
        isEdit: false,
        userId: 0,
        firstName: "",
        middleName: "",
        lastName: "",
        email: "",
        phone: "",
        isGuardian: false,
        school: "",
        graduationYear: 0,
        notes: "",
    };

    componentDidMount() {
        const userId = this.props.userId;
        if (userId) {
            API.get("api/users/user/" + userId).then((res) => {
                const user = res.data;
                this.setState({
                    isEdit: true,
                    userId: user.id,
                    accountId: user.accountId,
                    firstName: user.firstName,
                    middleName: user.middleName || "",
                    lastName: user.lastName,
                    email: user.email,
                    phone: user.phone || "",
                    isGuardian: user.isGuardian || false,
                    school: user.school || "",
                    graduationYear: parseInt(user.graduationYear) || 2020,
                    notes: user.notes || "",
                });
            });
        }
    }

    handleChange = (event, value) => {
        this.setState({ [value]: event.target.value });
    };

    handleChangeIsGuardian = (value) => {
        this.setState({ isGuardian: value == "true" });
    };

    onSave = () => {
        const user = {
            id: this.state.userId,
            accountId: parseInt(this.props.accountId),
            firstName: this.state.firstName,
            middleName: this.state.middleName,
            lastName: this.state.lastName,
            email: this.state.email,
            phone: this.state.phone,
            isGuardian: this.state.isGuardian,
            school: this.state.school,
            graduationYear: parseInt(this.state.graduationYear),
            notes: this.state.notes,
            isAdminCreated: true,
        };

        if (this.state.isEdit) {
            return API.post("api/users/user/" + this.state.userId, user);
        } else {
            return API.post("api/users/create", user);
        }
    };

    renderContent = () => {
        return (
            <div>
                <h3>User for Account: {this.props.accountId}</h3>
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
            </div>
        );
    };

    render() {
        const accountId = this.props.accountId;
        const isEdit = this.state.isEdit;
        const title = isEdit ? "Edit User" : "Add User";
        return (
            <main>
                <EditPageWrapper
                    isEdit={isEdit}
                    title={title}
                    content={this.renderContent()}
                    prevPageUrl={"/account/" + accountId + "?view=edit-users"}
                    onSave={this.onSave}
                    entityName={"user"}
                />
            </main>
        );
    }
}
