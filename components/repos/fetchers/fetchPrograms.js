'use strict';
import { forEach, keys, map } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var data;
var programMap;

export var getProgramById = function(programId) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(programMap[programId]);
  });
}

export var getProgramsByIds = function(programIds) {
  return new Promise(function(resolve, reject) {
    fetch();
    var programs = map(programIds, function(programId) {
      return programMap[programId];
    });
    resolve(programs);
  });
}

/* Helper Functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/classes.json');
    data = arr;
    programMap = initPrograms(data);
    fetched = true;
  }
}

function initPrograms(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.programId;
    map[id] = obj;
  });
  return map;
}
