"use strict";
require("./registerSticky.sass");
import React from "react";
import srcCheckmark from "../../assets/checkmark_white.svg";
import RegisterSelected from "./registerSelected.js";

export default class RegisterSticky extends React.Component {
    canConfirm = () => {
        const validClass = this.props.choice == "class" && this.props.validClass;
        const validAfh = this.props.choice == "afh" && this.props.validAfh;
        const validStudent = this.props.validStudent;
        const validGuardian = this.props.choice == "afh" || (this.props.choice == "class" && this.props.validGuardian);
        return (validClass || validAfh) && validStudent && validGuardian;
    }

    onConfirm = () => {
        // const confirmed = this.canConfirm();
        const confirmed = true;
        if (confirmed && this.props.choice == "class") {
            window.location.hash = "/register-success/class";
            return;
        } else if (confirmed && this.props.choice == "afh") {
            window.location.hash = "/register-success/afh";
            return;
        }
    }

    render() {
        const canConfirm = this.canConfirm();
        return (
            <div className={"sticky" + (canConfirm ? " active" : "")}>
                <RegisterSelected/>
                <h5>Complete the form to submit<br/>your registration request!</h5>
                { this.props.choice == "class" && 
                    <StepperListItem
                        valid={this.props.validClass} 
                        message={"Select a class"}/>
                }
                { this.props.choice == "afh" && 
                    <StepperListItem 
                        valid={this.props.validAfh} 
                        message={"Select an ask-for-help session"}/>
                }
                <StepperListItem 
                    valid={this.props.validStudent}
                    message={"Fill out student information"}/>
                { this.props.choice == "class" && 
                    <StepperListItem 
                    valid={this.props.validGuardian}
                    message={"Fill out guardian information"}/>
                }
                
                <button className={canConfirm ? "active" : ""} 
                        onClick={this.onConfirm}>
                    <div className="highlight"></div>
                    Confirm Registration
                </button>

                {
                    canConfirm &&
                    <p>
                        Thank you for correctly filling out your information! 
                        Press <b>Confirm</b> to submit your request.
                    </p>
                }
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
                    <img src={srcCheckmark}></img>
                </div>
                <span>{this.props.message}</span>
            </div>
        );
    }
    
}