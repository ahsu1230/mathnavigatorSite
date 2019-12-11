'use strict';
import axios from 'axios';

var BASE_HOST;
if (process.env.NODE_ENV == 'production') {
  BASE_HOST = "http://34.222.112.53"; // replace later
} else if (process.env.NODE_ENV == 'development') {
  BASE_HOST = "http://52.38.235.34"; // replace later
} else {
  BASE_HOST = "http://localhost:8080";
}

export default axios.create({
  baseURL: BASE_HOST
});
