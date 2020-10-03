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

export const capitalizeWord = (word) => {
    word = word || "";
    return word.charAt(0).toUpperCase() + word.slice(1);
};

/**
 * We expect a string "?asdf=zxcv&qwer=asdf"
 * that comes from props.location.search
 * which contains the entire query param string.
 * 
 * Given this string, we return an object (key => value)
 * 
 * Limitations: no duplicate keys!
 */
export const parseQueryParams = (search) => {
    let pairs = search.substring(1).split("&"); // ["key=value", "key=value"]
    let keyValueMap = {};
    pairs.forEach(pair => {
        let kvs = pair.split("=");
        keyValueMap[kvs[0]] = kvs[1];
    });
    return keyValueMap;
}