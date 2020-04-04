"use strict";
import axios from "axios";

var BASE_HOST;
if (process.env.NODE_ENV == "production") {
    BASE_HOST =
        "http://lb-prod-webserver-678749426.us-west-2.elb.amazonaws.com";
} else if (process.env.NODE_ENV == "development") {
    BASE_HOST =
        "http://lb-dev-webserver-2018209767.us-west-2.elb.amazonaws.com";
} else {
    BASE_HOST = "http://localhost:8080";
}

export default axios.create({
    baseURL: BASE_HOST
});
