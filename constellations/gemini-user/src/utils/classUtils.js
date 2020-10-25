"use strict";
import React from "react";
import { capitalizeWord, formatCurrency } from "./displayUtils";

export const getFullTitle = (programObj, classObj, semesterObj) => {
    const baseTitle =
        programObj.title + " " + capitalizeWord(classObj.classKey);
    if (semesterObj) {
        return baseTitle + " " + semesterObj.title;
    }
    return baseTitle;
};

export const isFullClass = (classObj) => {
    return classObj.fullState == 2;
};

export const displayPrice = (classObj) => {
    const isLump = !!classObj.priceLumpSum;
    const priceLabel = isLump ? "Total Price: " : "Price per session: ";
    const price = formatCurrency(
        isLump ? classObj.priceLumpSum : classObj.pricePerSession
    );
    return priceLabel + price;
};

export const displayTimeString = (classObj) => {
    const timesStr = classObj.timesStr || "";
    const times = timesStr.split(",");
    const timeLines = times.map((time, index) => {
        return (
            <div key={index} className="line">
                {time.trim()}
            </div>
        );
    });
    return <div className="class-times">{timeLines}</div>;
};

export const displayTimeStringOneLine = (classObj) => {
    const timesStr = classObj.timesStr || "";
    const times = timesStr
        .split(",")
        .map((str) => str.trim())
        .map((str) => <div className="class-times line">{str}</div>);
    return times;
};

export const displayFeaturedString = (program) => {
    if (program.featured == "popular") {
        return "This program is one of our most popular programs.";
    } else if (program.featured == "new") {
        return "This is a newly added program. Enroll now before it gets filled!";
    } else {
        return "";
    }
};
