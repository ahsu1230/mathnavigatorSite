"use strict";
import axios from "axios";

const orionBaseUrl =
    process.env.REACT_APP_ORION_HOST || "http://localhost:8001";

export default axios.create({
    baseURL: orionBaseUrl,
});
