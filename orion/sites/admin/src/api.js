'use strict';
import axios from 'axios';

const BASE_URL_DEV = "http://34.222.51.70/"; // temporary IP

export default axios.create({
  baseURL: BASE_URL_DEV
});
