"use strict";
import axios from "axios";
import moment from "moment";

var BASE_HOST =
    process.env.NODE_ENV == "production" ? "" : "http://localhost:8001";

export default axios.create({
    baseURL: BASE_HOST,
});

export const executeApiCalls = (apiCalls, successCallback, failCallback) => {
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
    sequence
        .then((results) => {
            console.log("All success!");
            successCallback(results);
        })
        .catch((results) => {
            console.log("One error?");
            failCallback(results);
        });
};
