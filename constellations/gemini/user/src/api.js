"use strict";
import axios from "axios";

const orionBaseUrl = process.env.ORION_HOST || "http://localhost:6001";

export default axios.create({
    baseURL: orionBaseUrl,
});
