"use strict";
require("./userAfhPage.sass");
import React from "react";
import moment from "moment";
import { keyBy } from "lodash";
import { Link } from "react-router-dom";
import API from "../api.js";
import { getFullName } from "../common/userUtils.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import { getAfhTitle } from "../common/displayUtils.js";
import { Modal } from "../common/modals/modal.js";
import YesNoModal from "../common/modals/yesnoModal.js";
import SingleUserSearcher from "../common/accountUserSearcher/singleUserSearcher.js";

export class UserAFHPage extends React.Component {
    state = {
        allAfhs: [],
        afhMap: {},
        usersForAfh: [],

        selectedAfhId: 0,
        selectedUserAfh: {},
        showDeleteModal: false,
    };

    componentDidMount = () => {
        API.get("/api/askforhelp/all")
            .then((res) => {
                const afhs = res.data;
                this.setState({
                    allAfhs: afhs,
                    afhMap: keyBy(afhs, "id"),
                });
            })
            .catch((err) => console.log("Could not fetch afh sessions"));
    };

    fetchUserAfh = (userAfh) => {
        this.onAfhChange(userAfh.afhId);
    };

    onRefreshPage = () => {
        this.onAfhChange(this.state.selectedAfhId);
    };

    onAfhChange = (nextAfhId) => {
        this.setState({
            selectedAfhId: nextAfhId,
        });

        API.get("api/user-afhs/afh/" + nextAfhId)
            .then((res) => {
                this.setState({ usersForAfh: res.data });
            })
            .catch((err) => console.log("Could not fetch users"));
    };

    onClickRemoveUser = (userAfh) => {
        this.setState({
            showDeleteModal: true,
            selectedUserAfh: userAfh,
        });
    };

    onDismissModal = () => {
        this.setState({ showDeleteModal: false });
    };

    render() {
        const selectedAfhId = this.state.selectedAfhId;
        const options = this.state.allAfhs.map((afh) => {
            const time =
                moment(afh.startsAt).format("MM/DD/yy hh:mm") +
                "-" +
                moment(afh.endsAt).format("hh:mm a");
            return {
                value: afh.id,
                displayName: getAfhTitle(afh),
            };
        });
        const users = this.state.usersForAfh.map((userAfh, index) => (
            <UserRow
                key={index}
                userAfh={userAfh}
                onClickRemoveUser={this.onClickRemoveUser}
                fetchUpdateUserAfh={this.fetchUserAfh}
            />
        ));

        return (
            <div id="view-user-afhs">
                <InputSelect
                    label="Select an AskForHelp session"
                    value={this.state.selectedAfhId}
                    onChangeCallback={(e) => this.onAfhChange(e.target.value)}
                    options={options}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no AskForHelp sessions to choose from.
                            Please add one <Link to="/afh/add">here</Link>
                        </span>
                    }
                />

                {users.length > 0 && (
                    <div id="users">
                        <h3>Users in AFH Session</h3>
                        {users}
                    </div>
                )}
                {users.length == 0 && !!this.state.selectedAfhId && (
                    <p>No Users currently registered for this AFH session.</p>
                )}
                {!!selectedAfhId && (
                    <AddUserAfh
                        afhId={selectedAfhId}
                        onRefreshPage={this.onRefreshPage}
                    />
                )}
                {!!selectedAfhId && (
                    <DeleteUserAfh
                        show={this.state.showDeleteModal}
                        userAfh={this.state.selectedUserAfh}
                        onDismissModal={this.onDismissModal}
                        onRefreshPage={this.onRefreshPage}
                    />
                )}
            </div>
        );
    }
}

class UserRow extends React.Component {
    state = {
        user: {},
    };

    componentDidMount = () => {
        const userAfh = this.props.userAfh || {};
        const userId = userAfh.userId;
        API.get("api/users/user/" + userId)
            .then((res) => {
                this.setState({ user: res.data });
            })
            .catch((err) => console.log("Could not find user " + userId));
    };

    onClickRemove = (userAfh) => {
        this.props.onClickRemoveUser(userAfh);
    };

    render() {
        const userAfh = this.props.userAfh || {};
        const user = this.state.user;
        const viewUserUrl = "/account/" + user.accountId + "?view=edit-users";
        const viewAccountUrl = "/account/" + user.accountId + "?view=user-afhs";
        return (
            <div className="user-row">
                <div className="user-info">
                    <div className="line name">
                        {getFullName(user)} (UserId {user.id})
                    </div>
                    <div className="line">{user.email}</div>
                    {user.phone && <div className="line">{user.phone}</div>}
                    {user.school && <div className="line">{user.school}</div>}
                    {user.graduationYear && (
                        <div className="line">
                            {"Graduation Year: " + user.graduationYear}
                        </div>
                    )}
                </div>
                <div className="state">
                    <div>
                        Registered on {moment(userAfh.updatedAt).format("l")}
                    </div>
                </div>
                <div className="links">
                    <Link to={viewUserUrl}>View User Details</Link>
                    <Link to={viewAccountUrl}>View Account</Link>
                    <button
                        className="remove"
                        onClick={(e) => this.onClickRemove(userAfh)}>
                        Remove User
                    </button>
                </div>
            </div>
        );
    }
}

class AddUserAfh extends React.Component {
    state = {
        show: false,
        selectedUser: {},
    };

    onClickAdd = () => {
        this.setState({ show: true });
    };

    onClickConfirm = () => {
        const newUserAfh = {
            afhId: parseInt(this.props.afhId),
            userId: this.state.selectedUser.id,
            accountId: this.state.selectedUser.accountId,
        };
        API.post("api/user-afhs/create", newUserAfh)
            .then((res) => {
                window.alert("User-afh successfully added");
                this.props.onRefreshPage();
            })
            .catch((err) =>
                window.alert("Could not register user into afh. " + err)
            );
    };

    onFoundUser = (user) => {
        this.setState({ selectedUser: user });
    };

    render() {
        return (
            <div className="add-user-afh">
                <button className="add" onClick={this.onClickAdd}>
                    Enroll a User into this AFH session
                </button>

                {this.state.show && (
                    <div>
                        <SingleUserSearcher onFoundUser={this.onFoundUser} />
                        {(this.state.selectedUser || {}).id && (
                            <button
                                className="confirm"
                                onClick={this.onClickConfirm}>
                                Confirm registering user into afh session
                            </button>
                        )}
                    </div>
                )}
            </div>
        );
    }
}

class DeleteUserAfh extends React.Component {
    persistDelete = () => {
        const userAfhId = parseInt(this.props.userAfh.id);
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
                    "Are you sure you want to remove user " +
                    userId +
                    " from afh session " +
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
