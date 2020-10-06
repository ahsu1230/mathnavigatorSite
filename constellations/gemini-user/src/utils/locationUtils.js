"use strict";
import React from "react";

export const createLocation = (location) => {
    if (location.isOnline) {
        return createOnlineLocation(location);
    } else {
        return createPhysicalLocation(location);
    }
};

export const createPhysicalLocation = (location) => {
    return (
        <div className="loc physical">
            <div className="line title">{location.title}</div>
            <div className="line">{location.street}</div>
            <div className="line">
                {location.city + ", " + location.state + " " + location.zipcode}
            </div>
            {location.room && <div className="line">{location.room}</div>}
        </div>
    );
};

export const createOnlineLocation = (location) => {
    return (
        <div className="loc online">
            <div className="line title">{location.title}</div>
            {location.room && <div className="line">{location.room}</div>}
            <div className="line">Online only</div>
            <div className="line disclaimer">
                Invitation links will be
                <br />
                posted in Google Classroom
            </div>
        </div>
    );
};
