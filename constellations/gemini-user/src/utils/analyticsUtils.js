"use strict";

export const trackAnalytics = (trackId, dataObj) => {
    const data = dataObj || {};
    if (process.env.NODE_ENV === "production") {
        mixpanel.track(trackId, data);
    } else {
        console.log("Track '" + trackId + "' " + JSON.stringify(data));
    }
};
