"use strict";
import React from "react";

const seasonOrder = ["spring", "summer", "fall", "winter"];
export function sortBySemester(semA, semB) {
    // Takes two semester IDs, e.g. 2020_fall
    semA = semA.split("_");
    semB = semB.split("_");
    if (semA[0] < semB[0]) {
        return 1;
    } else if (semA[0] > semB[0]) {
        return -1;
    } else {
        return seasonOrder.indexOf(semA[1]) < seasonOrder.indexOf(semB[1])
            ? 1
            : -1;
    }
}

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
export function fetchError(err) {
    alert("Could not fetch data: " + err);
}

export const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Computer Programming",
};

export const chargeDisplayNames = {
    charge: "Charge",
    refund: "Refund",
    pay_check: "Paid (Check)",
    pay_cash: "Paid (Cash)",
    pay_paypal: "Paid (Paypal)",
};
export const formatCurrency = (amount) => {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(amount);
};

export const getFullStateName = (fullState, append = false) => {
    const state = ["", "ALMOST FULL", "FULL"][fullState];

    if (append && state) return " (" + state + ")";
    return state;
};
