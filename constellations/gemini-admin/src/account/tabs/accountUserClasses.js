"use strict";
import React from "react";
import moment from "moment";
import { assign, keyBy } from "lodash";
import API from "../../api.js";
import { InputSelect } from "../../common/inputs/inputSelect.js";
import { Modal } from "../../common/modals/modal.js";
import YesNoModal from "../../common/modals/yesnoModal.js";
import UserSelector from "./userSelector.js";

export default class UserClasses extends React.Component {
    state = {
        userClasses: [],
        userClassStates: [],
        userClassStateMap: {},
        selectedUserClass: {},
        showDeleteModal: false,
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
        API.get("api/user-classes/user/" + userId)
            .then((res) => {
                const userClasses = res.data;
                this.setState({
                    userClasses: userClasses,
                });
            })
            .catch((err) => console.log("Could not fetch user-class" + err));
        API.get("api/user-classes/states")
            .then((res) => {
                const userClassStates = res.data;
                this.setState({
                    userClassStates: userClassStates,
                    userClassStateMap: keyBy(
                        userClassStates,
                        (state) => state + "A"
                    ),
                });
            })
            .catch((err) => console.log("Could not fetch states " + err));
    };

    onClickDelete = (userClass) => {
        this.setState({
            showDeleteModal: true,
            selectedUserClass: userClass,
        });
    };

    onDismissModal = () => {
        this.setState({ showDeleteModal: false });
    };

    onChangeUserClassState = (userClass, newState) => {
        const newUserClass = assign(userClass, { state: newState });
        API.post("api/user-classes/user-class/" + userClass.id, newUserClass)
            .then((res) => this.fetchUpdate())
            .catch((err) =>
                window.alert("Could not change user class state. " + err)
            );
    };

    render() {
        const selectedUser = this.props.selectedUser;
        const userClasses = this.state.userClasses.map((userClass, index) => {
            return (
                <div className="user-class" key={index}>
                    <div>{moment(userClass.createdAt).format("l")}</div>
                    <div>{userClass.classId}</div>
                    <InputSelect
                        required={false}
                        value={userClass.state}
                        onChangeCallback={(e) =>
                            this.onChangeUserClassState(
                                userClass,
                                e.target.value
                            )
                        }
                        options={this.state.userClassStates.map((state) => {
                            return {
                                value: state,
                                displayName: state,
                            };
                        })}
                    />
                    <button onClick={() => this.onClickDelete(userClass)}>
                        Delete
                    </button>
                </div>
            );
        });

        return (
            <section>
                <h2>User Class Registrations</h2>
                <UserSelector
                    users={this.props.users}
                    selectedUserId={selectedUser.id}
                    onChange={this.props.onSwitchUser}
                />

                {userClasses.length > 0 && userClasses}
                {userClasses.length == 0 && (
                    <p>This user has no class registrations.</p>
                )}
                <RegisterUserClass user={this.props.selectedUser} />
                {
                    <DeleteUserClass
                        show={this.state.showDeleteModal}
                        userClass={this.state.selectedUserClass}
                        onDismissModal={this.onDismissModal}
                    />
                }
            </section>
        );
    }
}

class DeleteUserClass extends React.Component {
    persistDelete = () => {
        const userClassId = this.props.userClass.id;
        API.delete("api/user-classes/user-class/" + userClassId)
            .then((res) => window.alert("User-class successfully deleted!"))
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
                    "Are you sure you want to disassociate user " +
                    userId +
                    " and class " +
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

class RegisterUserClass extends React.Component {
    state = {
        show: false,
        classes: [],
        options: [],
        selectedClassId: "",
    };

    componentDidMount() {
        API.get("api/classes/all")
            .then((res) => {
                const classes = res.data;
                const options = classes.map((classObj) => {
                    return {
                        value: classObj.classId,
                        displayName: classObj.classId,
                    };
                });
                this.setState({
                    classes: classes,
                    options: options,
                });
            })
            .catch((err) => console.log("Cannot get all classes"));
    }

    toggleShow = () => {
        this.setState({
            show: !this.state.show,
        });
    };

    onSelectChange = (e) => {
        this.setState({
            selectedClassId: e.target.value,
        });
    };

    onConfirm = () => {
        const userClass = {
            userId: this.props.user.id,
            accountId: this.props.user.accountId,
            classId: this.state.selectedClassId,
            state: 0,
        };
        API.post("api/user-classes/create", userClass)
            .then(() => window.alert("User successfully registered!"))
            .catch((err) => window.alert("Error occured " + err));
    };

    render() {
        return (
            <div>
                <button className="toggler" onClick={this.toggleShow}>
                    Register a class for user
                </button>
                {this.state.show && (
                    <div>
                        <InputSelect
                            label="Select a class"
                            value={this.state.selectedClassId}
                            onChangeCallback={this.onSelectChange}
                            hasNoDefault={true}
                            options={this.state.options}
                        />
                        <button className="confirm" onClick={this.onConfirm}>
                            Confirm User Registration for Class
                        </button>
                    </div>
                )}
            </div>
        );
    }
}
