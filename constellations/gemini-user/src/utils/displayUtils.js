"use strict";

export const subjectDisplayNames = {
    math: "Math",
    english: "English",
    programming: "Computer Programming",
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

export const convertToOrdinal = (index) => {
    const rem = index % 10;
    if (rem == 1) {
        return index + "st";
    } else if (rem == 2) {
        return index + "nd";
    } else if (rem == 3) {
        return index + "rd";
    } else {
        return index + "th";
    }
};
