"use strict";
require("./registerSuccess.sass");
import React from "react";
import { Link } from "react-router-dom";
import { isEmpty } from "lodash";
import { RegisterSectionBase } from "./registerBase.js";

export default class RegisterSectionSuccess extends React.Component {
    renderContent = () => {
        const wasClass = !isEmpty(this.props.selectedClass);
        const wasAfh = !isEmpty(this.props.selectedAfh);

        let linkRestart = (<div></div>);
        let linkBack = (<div></div>);
        let instructions = (<div></div>);
        if (wasClass) {
            instructions = (
                <p className="class">
                    Thank you for enrolling into our class.<br/>
                    As a reminder, payment is due before the first day of class. 
                    Failure to submit payment on-time may result in losing this enrollment reservation.<br/>
                    We look forward to working with you!
                </p>
            );
            linkBack = (
                <Link to="/programs" className="link-programs">Back to Program Catalog</Link>
            );
            linkRestart = (
                <a className="link-register">Enroll for another class</a>
            );
        } else {
            instructions = (
                <p className="afh">
                    Thank you for attending our ask-for-help session.<br/>
                    Before the class begins, please be sure to have all relevant materials prepared 
                    so you can use your time efficiently! 
                    If you already have specific questions in mind, please email us at <b>andymathnavigator@gmail.com</b>{" "}
                    so your teacher can prepare to answer your question.<br/>
                    See you soon!
                </p>
            );
            linkBack = (
                <Link to="/ask-for-help" className="link-afh">Back to Ask For Help</Link>
            );
            linkRestart = (
                <a className="link-register">Register for another Ask for Help session</a>
            );
        }

        return (
            <div className="content">
                {instructions}
                <div>
                    {linkRestart}
                    <Link to="/" className="link-home">Home</Link>
                    {linkBack}
                </div>
            </div>
        );
    }

    render() {
        return (
            <RegisterSectionBase
                sectionName="success"
                title={"Success! We have received your request!"}
                content={this.renderContent()}
                />
        );
    }
}

class PaymentInstructions extends React.Component {
    render() {
        /* TODO: Payment instructions*/
        return (
            <div className="payment-container">
                <p>Payment instructions</p>
            </div>
        );
    }
}