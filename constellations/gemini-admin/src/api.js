"use strict";
import axios from "axios";
import moment from "moment";

const BASE_HOST = process.env.MATHNAV_ORION_HOST;

export default axios.create({
    baseURL: BASE_HOST,
});

/*
 * Given a list of API calls (promises), execute them in order.
 * If all succeed, execute successCallback. If any fail, execute the failCallback.
 */
export const executeApiCalls = (apiCalls, successCallback, failCallback) => {
    console.log("Reducing " + apiCalls.length);

    // Reduce and execute list of API calls
    reduceApiCalls(apiCalls)
        .then((results) => {
            console.log("All success!");
            successCallback(results);
        })
        .catch((results) => {
            console.log("One error?");
            failCallback(results);
        });
};

/*
 * Create a single reduced promise from a list of promises (API calls)
 */
export const reduceApiCalls = (apiCalls) => {
    console.log("Reducing " + apiCalls.length);
    let fnResolveTask = function (nextApi) {
        return new Promise((resolve, reject) => {
            nextApi
                .then((resp) => {
                    console.log("Success: " + moment().format("hh:mm:ss"));
                    resolve(resp.data);
                })
                .catch((res) => {
                    console.log("Failure: " + moment().format("hh:mm:ss"));
                    reject(res.response.data);
                });
        });
    };

    let sequence = apiCalls.reduce((accumulatorPromise, nextApi) => {
        console.log(`Loop! ${moment().format("hh:mm:ss")}`);
        return accumulatorPromise.then(() => {
            return fnResolveTask(nextApi);
        });
    }, Promise.resolve());
    return sequence;
};
