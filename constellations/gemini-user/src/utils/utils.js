"use strict";
/* 
    THIS FILE IS DEPRECATED.
    DO NOT ADD additional functions / constants to this file.
*/

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

export const chargeDisplayNames = {
    charge: "Charge",
    refund: "Refund",
    pay_check: "Paid (Check)",
    pay_cash: "Paid (Cash)",
    pay_paypal: "Paid (Paypal)",
};
