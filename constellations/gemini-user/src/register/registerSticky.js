"use strict";
require("./registerSticky.sass");
import React from "react";
import { Transition } from "react-transition-group";

import srcCheckmarkWhite from "../../assets/checkmark_white.svg";
import srcCheckmarkGreen from "../../assets/checkmark_light_blue.svg";
export default class RegisterSticky extends React.Component {
    state = {
        confirmClicked: false, // true if "Confirm" button was clicked on
        confirmTriggered: false, // true if "Confirm" was triggered by parent (from props)
    };

    canConfirm = () => {
        const validClass =
            this.props.choice == "class" && this.props.validClass;
        const validAfh = this.props.choice == "afh" && this.props.validAfh;
        const validStudent = this.props.validStudent;
        const validGuardian =
            this.props.choice == "afh" ||
            (this.props.choice == "class" && this.props.validGuardian);
        return (validClass || validAfh) && validStudent && validGuardian;
    };

    onConfirm = () => {
        const confirmed = this.canConfirm();
        if (confirmed && this.props.choice) {
            // start animation
            this.setState({ confirmClicked: true });
            // go to page (after timeout)
            setTimeout(() => {
                window.location.hash = "/register-success/" + this.props.choice;
            }, 1800);
            return;
        }
    };

    static getDerivedStateFromProps(nextProps, prevState) {
        if (nextProps.triggerConfirmed) {
            // go to page (after timeout)
            setTimeout(() => {
                window.location.hash = "/register-success/" + nextProps.choice;
            }, 1800);
        }
        return {
            // start animation
            confirmTriggered: nextProps.triggerConfirmed,
        };
    }

    render() {
        const canConfirm = this.canConfirm();
        return (
            <div className={"sticky" + (canConfirm ? " active" : "")}>
                <h5>
                    Complete the form to submit
                    <br />
                    your registration request!
                </h5>
                {this.props.choice == "class" && (
                    <StepperListItem
                        valid={this.props.validClass}
                        message={"Select a class"}
                    />
                )}
                {this.props.choice == "afh" && (
                    <StepperListItem
                        valid={this.props.validAfh}
                        message={"Select an ask-for-help session"}
                    />
                )}
                <StepperListItem
                    valid={this.props.validStudent}
                    message={"Fill out student information"}
                />
                {this.props.choice == "class" && (
                    <StepperListItem
                        valid={this.props.validGuardian}
                        message={"Fill out guardian information"}
                    />
                )}

                <button
                    className={canConfirm ? "active" : ""}
                    onClick={this.onConfirm}>
                    Confirm Registration
                </button>

                {canConfirm && (
                    <p>
                        Thank you for correctly filling out your information!
                        Press <b>Confirm</b> to submit your request.
                    </p>
                )}
                <Loader
                    in={
                        this.state.confirmClicked || this.state.confirmTriggered
                    }
                />
            </div>
        );
    }
}

class StepperListItem extends React.Component {
    render() {
        const isValid = this.props.valid;
        return (
            <div className={"step-container" + (isValid ? " active" : "")}>
                <div className="icon-container">
                    <img src={srcCheckmarkWhite}></img>
                </div>
                <span>{this.props.message}</span>
            </div>
        );
    }
}

const defaultBarStyle = {
    position: "absolute",
    bottom: 0,
    left: 0,
    height: "4px",
    width: "0px",
    transition: `width 1200ms cubic-bezier(0.4, 0.6, 0.9, 0.1)`,
    backgroundColor: "#056571",
};

const transitionBarStyles = {
    entering: { width: "100%" },
    entered: { width: "100%" },
    exiting: { width: 0 },
    exited: { width: 0 },
};

const defaultCheckStyle = {
    position: "absolute",
    bottom: 0,
    right: "4px",
    width: "20px",
    opacity: 0.0,
    transition: `all 100ms ease-in 1250ms`,
};

const transitionCheckStyles = {
    entering: { opacity: 1.0, bottom: "8px" },
    entered: { opacity: 1.0, bottom: "8px" },
    exiting: { opacity: 0.0, bottom: "0px" },
    exited: { opacity: 0.0, bottom: "0px" },
};

class Loader extends React.Component {
    render() {
        return (
            <div className="loader">
                <Transition in={this.props.in} timeout={1200}>
                    {(state) => (
                        <div
                            className="highlight"
                            style={{
                                ...defaultBarStyle,
                                ...transitionBarStyles[state],
                            }}></div>
                    )}
                </Transition>
                <Transition in={this.props.in} timeout={300}>
                    {(state) => (
                        <img
                            src={srcCheckmarkGreen}
                            style={{
                                ...defaultCheckStyle,
                                ...transitionCheckStyles[state],
                            }}
                        />
                    )}
                </Transition>
            </div>
        );
    }
}
