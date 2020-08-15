"use strict";
require("./account.sass");
import React from "react";

export const chargeDisplayNames = {
    charge: "Charge",
    refund: "Refund",
    pay_check: "Paid (Check)",
    pay_cash: "Paid (Cash)",
    pay_paypal: "Paid (Paypal)",
};
export const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Computer Programming",
};
export const seasonOrder = ["spring", "summer", "fall", "winter"];

export function renderMultiline(lines) {
    return lines.map((line, index) => {
        return (
            <span key={index}>
                {line}
                <br />
            </span>
        );
    });
}
export function formatCurrency(amount) {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(amount);
}
export function fetchError(err) {
    alert("Could not fetch data: " + err);
}
