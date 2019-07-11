'use strict';
var _ = require('lodash/core');

import jsons from './json/*.json';

export var locationMap = {};
export var preReqMap = {};
export var programMap = {};
export var classMap = {};
export var sessionMap = {};

function init() {
  console.log('Initializing Programs...');
  locationMap = initLocations(jsons.locations);
  programMap = initPrograms(jsons.programs);
  classMap = initClasses(jsons.classes);
  sessionMap = initSessions(jsons.sessions);
  console.log(jsons.prereqs);
  preReqMap = initPreReqs(jsons.prereqs);
  console.log(JSON.stringify(preReqMap));
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

function initClasses(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.key;
    obj.times = convertStringArray(obj.times);
    map[id] = obj;
  });
  return map;
}

function initSessions(arr) {
  var map = {};
  _.forEach(arr, function(obj) {
    var id = obj.key;
    if (id == "_") { return; }

    obj.canceled = convertToBool(obj.canceled);
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
    console.log(obj.requiredProgramIds);
    obj.requiredProgramIds = convertStringArray(obj.requiredProgramIds);
    map[id] = obj;
  });
  return map;
}

function convertToBool(str) {
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
