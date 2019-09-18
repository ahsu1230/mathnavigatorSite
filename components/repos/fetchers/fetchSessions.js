'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var sessionMap;

export var getSessionsByClass = function(classKey) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(sessionMap[classKey]);
  });
}

export const fetcher = {
  getSessionsByClass: getSessionsByClass
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/sessions.json');
    sessionMap = initSessionMap(arr);
    fetched = true;
  }
}

function initSessionMap(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.classKey;
    if (!id || id === "_") { return; }

    obj.canceled = convertStringToBool(obj.canceled);
    map[id] = map[id] || [];
    map[id].push(obj);
  });
  return map;
}
