"use strict";
import axios from "axios";

const orionBasePath = process.env.ORION_HOST || "http://localhost:6001"

export default axios.create({
    baseURL: orionBasePath,
});
