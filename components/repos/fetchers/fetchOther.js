'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetchedKeyValue = false;
var fetchedLocation = false;
var dataLocation;
var dataKeyValue;
var keyValueMap;
var locationMap;

export var getLocation = function(locationId) {
  return new Promise(function(resolve, reject) {
    if (!fetchedLocation) {
      dataLocation = fetchLocations();
      locationMap = initLocations(dataLocation);
      fetchedLocation = true;
    }
    resolve(locationMap[locationId]);
  });
}

export var getAllLocations = function() {
  return new Promise(function(resolve, reject) {
    if (!fetchedLocation) {
      dataLocation = fetchLocations();
      locationMap = initLocations(dataLocation);
      fetchedLocation = true;
    }
    resolve(locationMap);
  });
}

export var getKeyValue = function(key) {
  return new Promise(function(resolve, reject) {
    if (!fetchedKeyValue) {
      dataKeyValue = fetchKeyValues();
      keyValueMap = initKeyValues(dataKeyValue);
      fetchedKeyValue = true;
    }
    resolve(keyValueMap[key]);
  });
}

export const fetcher = {
  getLocation: getLocation,
  getAllLocations: getAllLocations,
  getKeyValue: getKeyValue
}

/* Helper functions */
function fetchKeyValues() {
  const arr = require('./json/keyvalues.json');
  return arr;
}

function fetchLocations() {
  const arr = require('./json/locations.json');
  return arr;
}

/* Key Values */
function initKeyValues(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.key;
    map[id] = obj.value;
  });
  return map;
}


/* Locations */
function initLocations(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.locationId;
    map[id] = obj;
  });
  return map;
}
