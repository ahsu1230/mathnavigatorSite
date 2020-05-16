"use strict";
import axios from "axios";

var BASE_HOST = process.env.NODE_ENV == "production" ? "" : "http://localhost:6001";

export default axios.create({
    baseURL: BASE_HOST,
});
