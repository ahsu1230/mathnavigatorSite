'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var data;
var prereqMap;

export const fetcher = {
  getPrereqs: function(programId) {
    return new Promise(function(resolve, reject) {
      fetch();
      resolve(prereqMap[programId]);
    });
  }
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/prereqs.json');
    data = arr;
    prereqMap = initPrereqs(data);
    fetched = true;
  }
}

function initPrereqs(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.programId;
    obj.requiredProgramIds = convertStringArray(obj.requiredProgramIds);
    map[id] = obj;
  });
  return map;
}
