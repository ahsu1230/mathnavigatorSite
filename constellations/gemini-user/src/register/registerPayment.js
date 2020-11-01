"use strict";
require("./registerPayment.sass");
import React from "react";

import srcZelle from "../../assets/payment_images/zelle.png";
import srcVenmo from "../../assets/payment_images/venmo.png";
import srcPaypal from "../../assets/payment_images/paypal.png";
import srcCash from "../../assets/payment_images/cash_white.svg";
import srcCheck from "../../assets/payment_images/bank_check.svg";

export default class RegisterPayment extends React.Component {
    render() {
        return (
            <div id="payment">
                <h3>Payment Instructions</h3>

                <div className="container paypal">
                    <div className="img-container">
                        <img src={srcPaypal} />
                    </div>
                    <div className="instructions">
                        <h4>Pay with Paypal</h4>
                        <p>
                            Please send your payment to Paypal user{" "}
                            <u>mathnavigator@yahoo.com</u> as a Friend.
                        </p>
                    </div>
                </div>

                <div className="container zelle">
                    <div className="img-container">
                        <img src={srcZelle} />
                    </div>
                    <div className="instructions">
                        <h4>Pay with Zelle</h4>
                        <p>
                            Please send your payment to Zelle user{" "}
                            <u>andymathnavigator@gmail.com</u>.
                        </p>
                    </div>
                </div>

                <div className="container venmo">
                    <div className="img-container">
                        <img src={srcVenmo} />
                    </div>
                    <div className="instructions">
                        <h4>Pay with Venmo</h4>
                        <p>
                            Please send your payment to Venmo user{" "}
                            <u>@MathNavigator</u>.
                        </p>
                    </div>
                </div>

                <div className="container check cash">
                    <div className="img-container">
                        <img src={srcCash} />
                    </div>
                    <div className="instructions">
                        <h4>Pay with Cash or Check</h4>
                        <p>
                            If submitting a check, please write the check for{" "}
                            <b>Math Navigator</b> and include which class you
                            are submitting payment for. For both check and cash
                            payments, email us at{" "}
                            <u>andymathnavigator@gmail.com</u> for further
                            instructions.
                        </p>
                    </div>
                </div>
            </div>
        );
    }
}
