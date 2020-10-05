"use strict";
import React from "react";
import { capitalizeWord, formatCurrency } from "./displayUtils";

export const getFullTitle = (programObj, classObj) => {
    return programObj.title + capitalizeWord(classObj.classKey);
}

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
}

export const displayTimeString = (classObj) => {
    const timesStr = classObj.timesStr || "";
    const times = timesStr.split(",");
    const timeLines = times.map((time, index) => {
        return (<div key={index} className="line">{time.trim()}</div>);
    });
    return (
        <div className="class-times">
            {timeLines}
        </div>
    );
}