'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';

var fetched = false;
var data;

export var getAFH = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(data);
  });
}

export const fetcher = {
  getAFH: getAFH
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/askforhelp.json');
    data = arr;
    fetched = true;
  }
}
