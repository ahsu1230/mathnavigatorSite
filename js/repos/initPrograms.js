'use strict';
var _ = require('lodash/core');
import jsons from './json/*.json';

export var locationMap = {};
export var preReqMap = {};
export var programMap = {};
export var classMapByKey = {};
export var sessionMap = {};

function init() {
  console.log('Initializing Programs...');
  locationMap = initLocations(jsons.locations);
  programMap = initPrograms(jsons.programs);
  classMapByKey = initClassesByKey(jsons.classes);
  sessionMap = initSessions(jsons.sessions);
  preReqMap = initPreReqs(jsons.prereqs);
  console.log('Programs done initializing.');
}

function initLocations(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.locationId;
    map[id] = obj;
  });
  return map;
}

function initPrograms(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.programId;
    map[id] = obj;
  });
  return map;
}

function initClassesByKey(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.key;
    map[id] = filterClassObj(obj);
  });
  return map;
}

function filterClassObj(obj) {
  obj.isAvailable = convertStrToBool(obj.isAvailable);
  obj.times = convertStringArray(obj.times);
  return obj;
}

function initSessions(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.key;
    if (id == "_") { return; }

    obj.canceled = convertStrToBool(obj.canceled);
    if (map[id] && map[id].length > 0) {
      map[id].push(obj);
    } else {
      map[id] = [obj];
    }
  });
  return map;
}

function initPreReqs(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.programId;
    obj.requiredProgramIds = convertStringArray(obj.requiredProgramIds);
    map[id] = obj;
  });
  return map;
}

function convertStrToBool(str) {
  if (str.toLowerCase() === "true") {
    return true;
  } else {
    return false;
  }
}

function convertStringArray(str) {
  if (!str || typeof str != 'string') {
    return [];
  }
  // We assume the following format:
  // "[a, b, c]"
  var newStr = str.substring(1, str.length - 1);
  var arr = newStr.split(", ");
  return arr;
}

init();
