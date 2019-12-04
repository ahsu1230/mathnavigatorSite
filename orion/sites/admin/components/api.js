'use strict';
import axios from 'axios';

const BASE_URL_DEV = "http://localhost:8080/";

export default axios.create({
  baseURL: BASE_URL_DEV
});
