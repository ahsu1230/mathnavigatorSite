"use strict";
require("./accountUserAfhs.sass");
import React from "react";
import moment from "moment";
import { keyBy } from "lodash";
import API from "../../api.js";
import { InputSelect } from "../../common/inputs/inputSelect.js";
import { getAfhTitle } from "../../common/displayUtils.js";
import { Modal } from "../../common/modals/modal.js";
import YesNoModal from "../../common/modals/yesnoModal.js";
import UserSelector from "./userSelector.js";

export default class UserAfhs extends React.Component {
    state = {
        userAfhs: [],
        allAfhs: [],
        afhMap: {},
        selectedUserAfh: {},
        showDeleteModal: false,
    };

    componentDidMount() {
        this.fetchUpdate();
        API.get("api/askforhelp/all")
            .then((res) => {
                const afhs = res.data;
                this.setState({
                    allAfhs: afhs,
                    afhMap: keyBy(afhs, "id"),
                });
            })
            .catch((err) => console.log(err));
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

    onClickDelete = (userAfh) => {
        this.setState({
            showDeleteModal: true,
            selectedUserAfh: userAfh,
        });
    };

    onDismissModal = () => {
        this.setState({ showDeleteModal: false });
    };

    render() {
        const selectedUser = this.props.selectedUser;
        const afhMap = this.state.afhMap;
        const userAfhs = this.state.userAfhs.map((userAfh, index) => {
            const afh = afhMap[userAfh.afhId];
            return (
                <UserAfhRow
                    key={index}
                    afh={afh}
                    userAfh={userAfh}
                    onDelete={this.onClickDelete}
                />
            );
        });
        return (
            <section className="account-tab user-afhs">
                <h3>User AFH Registrations</h3>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={selectedUser.id}
                    onChange={this.props.onSwitchUser}
                />

                {userAfhs.length > 0 && userAfhs}
                {userAfhs.length == 0 && (
                    <p>This user has no ask-for-help registrations.</p>
                )}
                <RegisterUserAfh
                    user={this.props.selectedUser}
                    onRefreshPage={this.fetchUpdate}
                />
                {
                    <DeleteUserAfh
                        show={this.state.showDeleteModal}
                        userAfh={this.state.selectedUserAfh}
                        onDismissModal={this.onDismissModal}
                        onRefreshPage={this.fetchUpdate}
                    />
                }
            </section>
        );
    }
}

class UserAfhRow extends React.Component {
    render() {
        const afh = this.props.afh;
        const userAfh = this.props.userAfh;
        const onDelete = this.props.onDelete;
        return (
            <div className="user-afh">
                <div>Id: {afh.id}</div>
                <div>{afh.title + " (" + afh.subject + ")"}</div>
                <div>Registered on {moment(userAfh.createdAt).format("l")}</div>
                <button className="delete" onClick={() => onDelete(userAfh)}>
                    Delete
                </button>
            </div>
        );
    }
}

class DeleteUserAfh extends React.Component {
    persistDelete = () => {
        const userAfhId = this.props.userAfh.id;
        API.delete("api/user-afhs/user-afh/" + userAfhId)
            .then((res) => {
                window.alert("User-afh successfully deleted!");
                this.props.onRefreshPage();
                this.props.onDismissModal();
            })
            .catch((err) => window.alert("Error deleting User-afh " + err));
    };

    render() {
        const show = this.props.show;
        const onDismissModal = this.props.onDismissModal;
        const userId = this.props.userAfh.userId;
        const afhId = this.props.userAfh.afhId;
        const modalContent = (
            <YesNoModal
                text={
                    "Are you sure you want to disassociate user " +
                    userId +
                    " and AskForHelp session " +
                    afhId +
                    "?"
                }
                onAccept={this.persistDelete}
                onReject={onDismissModal}
            />
        );

        return (
            <div>
                {show && (
                    <Modal
                        content={modalContent}
                        show={show}
                        onDismiss={onDismissModal}
                    />
                )}
            </div>
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
                        displayName: getAfhTitle(afh),
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
        const onRefreshPage = this.props.onRefreshPage;
        const userAfh = {
            userId: this.props.user.id,
            accountId: this.props.user.accountId,
            afhId: parseInt(this.state.selectedAfhId),
        };
        API.post("api/user-afhs/create", userAfh)
            .then(() => {
                window.alert("User successfully registered!");
                onRefreshPage();
            })
            .catch((err) => window.alert("Error occured " + err));
    };

    render() {
        return (
            <div className="register-user-afh">
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
                        {this.state.selectedAfhId && (
                            <div>
                                <p>
                                    By confirming, this user will be registered
                                    to the selected ask-for-help session.
                                </p>
                                <button
                                    className="confirm"
                                    onClick={this.onConfirm}>
                                    Confirm User Registration for AskForHelp
                                </button>
                            </div>
                        )}
                    </div>
                )}
            </div>
        );
    }
}
