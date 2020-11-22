"use strict";
require("./userClassesPage.sass");
import React from "react";
import moment from "moment";
import { assign, keyBy } from "lodash";
import { Link } from "react-router-dom";
import API from "../api.js";
import { InputSelect } from "../common/inputs/inputSelect.js";
import { getFullName } from "../common/userUtils.js";
import { Modal } from "../common/modals/modal.js";
import YesNoModal from "../common/modals/yesnoModal.js";
import SingleUserSearcher from "../common/accountUserSearcher/singleUserSearcher.js";

export class UserClassesPage extends React.Component {
    state = {
        allClasses: [],
        classMap: {},
        usersForClass: [],
        userClassStates: [],

        selectedClassId: "",
        selectedUserClass: "",
        showDeleteModal: false,
    };

    componentDidMount = () => {
        API.get("/api/classes/all")
            .then((res) => {
                const classes = res.data;
                this.setState({
                    allClasses: classes,
                    classMap: keyBy(classes, "classId"),
                });
            })
            .catch((err) => console.log("Could not fetch classes"));
        API.get("api/user-classes/states")
            .then((res) => {
                const userClassStates = res.data;
                this.setState({
                    userClassStates: userClassStates,
                });
            })
            .catch((err) => console.log("Could not fetch states " + err));
    };

    fetchUserClass = (userClass) => {
        this.onClassChange(userClass.classId);
    };

    onRefreshPage = () => {
        this.onClassChange(this.state.selectedClassId);
    };

    onClassChange = (nextClassId) => {
        this.setState({
            selectedClassId: nextClassId,
        });

        API.get("api/user-classes/class/" + nextClassId)
            .then((res) => {
                this.setState({ usersForClass: res.data });
            })
            .catch((err) => console.log("Could not fetch users"));
    };

    onClickRemoveUser = (userClass) => {
        this.setState({
            showDeleteModal: true,
            selectedUserClass: userClass,
        });
    };

    onDismissModal = () => {
        this.setState({ showDeleteModal: false });
    };

    render() {
        const selectedClassId = this.state.selectedClassId;
        const options = this.state.allClasses.map((classObj) => {
            return {
                value: classObj.classId,
                displayName: classObj.classId,
            };
        });
        const users = this.state.usersForClass.map((userClass, index) => (
            <UserRow
                key={index}
                userClass={userClass}
                userClassStates={this.state.userClassStates}
                onClickRemoveUser={this.onClickRemoveUser}
                fetchUpdateUserClass={this.fetchUserClass}
            />
        ));

        return (
            <div id="view-user-classes">
                <InputSelect
                    label="Select a Class"
                    value={selectedClassId}
                    onChangeCallback={(e) => this.onClassChange(e.target.value)}
                    options={options}
                    hasNoDefault={true}
                    errorMessageIfEmpty={
                        <span>
                            There are no Classes to choose from. Please add one{" "}
                            <Link to="/classes/add">here</Link>
                        </span>
                    }
                />

                {users.length > 0 && (
                    <div id="users">
                        <HelpInfo />
                        <h3>Users in Class</h3>
                        {users}
                    </div>
                )}
                {users.length == 0 && selectedClassId && (
                    <p>No Users currently registered for this class.</p>
                )}
                {selectedClassId && (
                    <AddUserClass
                        classId={selectedClassId}
                        onRefreshPage={this.onRefreshPage}
                    />
                )}
                {selectedClassId && (
                    <DeleteUserClass
                        show={this.state.showDeleteModal}
                        userClass={this.state.selectedUserClass}
                        onDismissModal={this.onDismissModal}
                        onRefreshPage={this.onRefreshPage}
                    />
                )}
            </div>
        );
    }
}

class HelpInfo extends React.Component {
    state = {
        show: false,
    };

    onToggle = () => {
        this.setState({ show: !this.state.show });
    };

    render() {
        const toggleTitle = this.state.show
            ? "Hide Help Information"
            : "Show Help Information";
        return (
            <div className="help">
                <button className="toggle" onClick={this.onToggle}>
                    {toggleTitle}
                </button>
                {this.state.show && (
                    <div className="description">
                        <p>
                            Below is a list of users (usually students) that are
                            in the selected class. Every user has an enrollment
                            state for this class.
                        </p>
                        <ul>
                            <li>
                                <i>Enrolled</i> means the user is enrolled into
                                the class with no complications. Usually, the
                                student pays the full class tuition if enrolled.
                            </li>
                            <li>
                                <i>Pending</i> means the user still awaiting
                                approval for enrollment. Student usually has not
                                paid the tuition yet.
                            </li>
                            <li>
                                <i>Trial</i> means the user is "trying out" the
                                class and will probably pay a discounted price
                                or an "entry fee".
                            </li>
                            <li>
                                <i>Dismissed</i> means the user was enrolled,
                                but had to drop the class for some reason.
                            </li>
                            <li>
                                <i>Removing a user</i> means all enrollment
                                information for between the user will be
                                deleted.
                            </li>
                        </ul>
                    </div>
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
        const userClass = this.props.userClass || {};
        const userId = userClass.userId;
        API.get("api/users/user/" + userId)
            .then((res) => {
                this.setState({ user: res.data });
            })
            .catch((err) => console.log("Could not find user " + userId));
    };

    onChangeUserClassState = (userClass, newState) => {
        const newUserClass = assign(userClass, { state: newState });
        API.post("api/user-classes/user-class/" + userClass.id, newUserClass)
            .then((res) => this.props.fetchUpdateUserClass(userClass))
            .catch((err) =>
                window.alert("Could not change user class state. " + err)
            );
    };

    onClickRemove = (userClass) => {
        this.props.onClickRemoveUser(userClass);
    };

    render() {
        const userClass = this.props.userClass || {};
        const userClassStates = this.props.userClassStates || [];
        const user = this.state.user;
        const viewUserUrl = "/account/" + user.accountId + "?view=edit-users";
        const viewAccountUrl =
            "/account/" + user.accountId + "?view=user-classes";
        return (
            <div className="user-row">
                <div className="user-info">
                    <div className="line name">
                        {getFullName(user)} (UserId {user.id})
                    </div>
                    <div className="line">{user.email}</div>
                    <div className="line">{user.school}</div>
                    <div className="line">
                        {"Graduation Year: " + user.graduationYear}
                    </div>
                </div>
                <div className="state">
                    <InputSelect
                        required={false}
                        value={userClass.state}
                        onChangeCallback={(e) =>
                            this.onChangeUserClassState(
                                userClass,
                                e.target.value
                            )
                        }
                        options={userClassStates.map((state) => {
                            return {
                                value: state,
                                displayName: state,
                            };
                        })}
                    />
                    <div className="line">
                        Last Updated: {moment(userClass.updatedAt).format("l")}
                    </div>
                </div>
                <div className="links">
                    <Link to={viewUserUrl}>View User Details</Link>
                    <Link to={viewAccountUrl}>View Account</Link>
                    <button
                        className="remove"
                        onClick={(e) => this.onClickRemove(userClass)}>
                        Remove User
                    </button>
                </div>
            </div>
        );
    }
}

class AddUserClass extends React.Component {
    state = {
        show: false,
        selectedUser: {},
    };

    onClickAdd = () => {
        this.setState({ show: true });
    };

    onClickConfirm = () => {
        const newUserClass = {
            classId: this.props.classId,
            userId: this.state.selectedUser.id,
            accountId: this.state.selectedUser.accountId,
            state: "enrolled",
        };
        API.post("api/user-classes/create", newUserClass)
            .then((res) => {
                window.alert("User-class successfully added");
                this.props.onRefreshPage();
            })
            .catch((err) =>
                window.alert("Could not enroll user into class. " + err)
            );
    };

    onFoundUser = (user) => {
        this.setState({ selectedUser: user });
    };

    render() {
        return (
            <div className="add-user-class">
                <button className="add" onClick={this.onClickAdd}>
                    Enroll a User into this Class
                </button>

                {this.state.show && (
                    <div>
                        <SingleUserSearcher onFoundUser={this.onFoundUser} />
                        {(this.state.selectedUser || {}).id && (
                            <button
                                className="confirm"
                                onClick={this.onClickConfirm}>
                                Confirm enrolling user into class
                            </button>
                        )}
                    </div>
                )}
            </div>
        );
    }
}

class DeleteUserClass extends React.Component {
    persistDelete = () => {
        const userClassId = this.props.userClass.id;
        API.delete("api/user-classes/user-class/" + userClassId)
            .then((res) => {
                window.alert("User-class successfully deleted!");
                this.props.onRefreshPage();
                this.props.onDismissModal();
            })
            .catch((err) => window.alert("Error deleting User-class " + err));
    };

    render() {
        const show = this.props.show;
        const onDismissModal = this.props.onDismissModal;
        const userId = this.props.userClass.userId;
        const classId = this.props.userClass.classId;
        const modalContent = (
            <YesNoModal
                text={
                    "Are you sure you want to remove user " +
                    userId +
                    " from class " +
                    classId +
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
