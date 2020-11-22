"use strict";
import React from "react";
import API from "../../api.js";
import EditPageWrapper from "../../common/editPages/editPageWrapper.js";

export class UserMovePage extends React.Component {
    state = {
        user: {},
        newAccountId: 0,
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
        };

        if (this.state.isEdit) {
            return API.post("api/users/user/" + this.state.userId, user);
        } else {
            return API.post("api/users/create", user);
        }
    };

    renderContent = () => {
        const currentAccountId = this.props.accountId;
        const newAccountId = this.state.newAccountId;
        const user = this.state.user;

        return (
            <div>
                <h3>Selected User & Current Account</h3>
                <p>{user.email}</p>

                <h3>Select another Account</h3>
                {/* Search and select Account */}

                <h3>Are you sure you want to move this user?</h3>
            </div>
        );
    };

    render() {
        const accountId = this.props.accountId;
        return (
            <main>
                <EditPageWrapper
                    isEdit={false}
                    title={"Move user to another account"}
                    content={this.renderContent()}
                    prevPageUrl={"/account/" + accountId + "?view=edit-users"}
                    onSave={this.onSave}
                    entityName={"user"}
                />
            </main>
        );
    }
}
