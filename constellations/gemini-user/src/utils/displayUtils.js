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