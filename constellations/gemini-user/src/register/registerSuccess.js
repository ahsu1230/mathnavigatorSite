"use strict";
require("./registerSuccess.sass");
import React from "react";
import { Link } from "react-router-dom";
import { trackAnalytics } from "../utils/analyticsUtils.js";
import RegisterPayment from "./registerPayment.js";
import srcStar from "../../assets/star_white.svg";

export default class RegisterSuccessPage extends React.Component {
    componentDidMount() {
        trackAnalytics("register-success");
    }

    render() {
        const wasClass = this.props.registered == "class";

        let linkRestart = <div></div>;
        let linkBack = <div></div>;
        let instructions = <div></div>;
        if (wasClass) {
            instructions = (
                <p className="class">
                    Thank you for enrolling into our class. We look forward to
                    working with you! As a reminder, payment is due before the
                    first day of class. Failure to submit payment on-time may
                    result in losing this enrollment reservation. See you soon!
                </p>
            );
            linkBack = (
                <Link to="/programs" className="link-programs">
                    Back to Program Catalog
                </Link>
            );
            linkRestart = (
                <Link to="/register" className="link-register">
                    Enroll for another class
                </Link>
            );
        } else {
            instructions = (
                <p className="afh">
                    Thank you for attending our ask-for-help session. Before the
                    class begins, please be sure to have all relevant materials
                    prepared so you can use your time efficiently! If you
                    already have specific questions in mind, please email us at{" "}
                    <b>andymathnavigator@gmail.com</b> so your teacher can
                    prepare to answer your question. See you soon!
                </p>
            );
            linkBack = (
                <Link to="/ask-for-help" className="link-afh">
                    Back to Ask For Help
                </Link>
            );
            linkRestart = (
                <Link to="/register" className="link-register">
                    Register for another
                    <br />
                    Ask-for-Help session
                </Link>
            );
        }

        return (
            <div id="view-register-success">
                <div className="header-wrapper">
                    <div className="icon-wrapper">
                        <img src={srcStar} />
                    </div>
                    <h1>Success! We have received your request!</h1>
                </div>
                {instructions}
                {wasClass && <RegisterPayment />}
                <div>
                    {linkRestart}
                    <Link to="/" className="link-home">
                        Home
                    </Link>
                    {linkBack}
                </div>
            </div>
        );
    }
}
