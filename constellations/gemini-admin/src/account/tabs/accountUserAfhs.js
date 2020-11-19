"use strict";
import React from "react";
import moment from "moment";
import API from "../../api.js";
import { getFullName } from "../../common/userUtils.js";
import { InputSelect } from "../../common/inputs/inputSelect.js";
import UserSelector from "./userSelector.js";

export default class UserAfhs extends React.Component {
    state = {
        userAfhs: [],
    };

    componentDidMount() {
        this.fetchUpdate();
    }

    componentDidUpdate(prevProps) {
        if (prevProps.selectedUser !== this.props.selectedUser) {
            this.fetchUpdate();
        }
    }

    fetchUpdate = () => {
        const selectedUser = this.props.selectedUser;
        const userId = selectedUser.id;
        API.get("api/user-afhs/users/" + userId)
            .then((res) => {
                const userAfhs = res.data;
                this.setState({
                    userAfhs: userAfhs,
                });
            })
            .catch((err) => console.log(err));
    };

    render() {
        const selectedUser = this.props.selectedUser;
        const userAfhs = this.state.userAfhs.map((userAfh, index) => {
            return (
                <div className="user-afh" key={index}>
                    <div>{moment(userAfh.createdAt).format("l")}</div>
                    <div>AfhId: {userAfh.afhId}</div>
                </div>
            );
        });
        return (
            <section>
                <h2>User AFH Registrations</h2>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={selectedUser.id}
                    onChange={this.props.onSwitchUser}
                />

                <div>
                    <div>Id: {selectedUser.id}</div>
                    <div>Name: {getFullName(selectedUser)}</div>
                    <div>Email: {selectedUser.email}</div>
                </div>

                {userAfhs.length > 0 && userAfhs}
                {userAfhs.length == 0 && (
                    <p>This user has no ask-for-help registrations.</p>
                )}
                <RegisterUserAfh user={this.props.selectedUser} />
            </section>
        );
    }
}

class RegisterUserAfh extends React.Component {
    state = {
        show: false,
        afhs: [],
        options: [],
        selectedAfhId: "",
    };

    componentDidMount() {
        API.get("api/askforhelp/all")
            .then((res) => {
                const afhs = res.data;
                const options = afhs.map((afh) => {
                    return {
                        value: afh.id,
                        displayName: afh.id,
                    };
                });
                this.setState({
                    afhs: afhs,
                    options: options,
                });
            })
            .catch((err) => console.log("Cannot get all afhs"));
    }

    toggleShow = () => {
        this.setState({
            show: !this.state.show,
        });
    };

    onSelectChange = (e) => {
        this.setState({
            selectedAfhId: e.target.value,
        });
    };

    onConfirm = () => {
        const userAfh = {
            userId: this.props.user.id,
            accountId: this.props.user.accountId,
            afhId: this.state.selectedAfhId,
        };
        API.post("api/user-afhs/create", userAfh)
            .then(() => window.alert("User successfully registered!"))
            .catch((err) => window.alert("Error occured " + err));
    };

    render() {
        return (
            <div>
                <button className="toggler" onClick={this.toggleShow}>
                    Register an AskForHelp for user
                </button>
                {this.state.show && (
                    <div>
                        <InputSelect
                            label="Select an AskForHelp session"
                            value={this.state.selectedAfhId}
                            onChangeCallback={this.onSelectChange}
                            hasNoDefault={true}
                            options={this.state.options}
                        />
                        <button className="confirm" onClick={this.onConfirm}>
                            Confirm User Registration for AskForHelp
                        </button>
                    </div>
                )}
            </div>
        );
    }
}
