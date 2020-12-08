'use strict';
import { filter, forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var data;
var classMap;
var classMapByProgram;
var classMapBySemester;

export var getAllClasses = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(classMap);
  });
}

export var getClassByKey = function(classKey) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(classMap[classKey]);
  });
}

export var getClassesByKeys = function(classKeys) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(classKeys.map(function(classKey) {
      return classMap[classKey];
    }));
  });
}

export var getClassesByProgram = function(programId) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(classMapByProgram[programId]);
  });
}

export var getClassesBySemester = function(semesterId) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(classMapBySemester[semesterId]);
  });
}

export var getClassesByProgramAndSemester = function(programId, semesterId) {
  return new Promise(function(resolve, reject) {
    fetch();
    var classes = classMapByProgram[programId];
    var targetClass = filter(classes, {semesterId: semesterId});
    resolve(targetClass);
  });
}

export const fetcher = {
  getAllClasses: getAllClasses,
  getClassByKey: getClassByKey,
  getClassesByKeys: getClassesByKeys,
  getClassesByProgram: getClassesByProgram,
  getClassesBySemester: getClassesBySemester,
  getClassesByProgramAndSemester: getClassesByProgramAndSemester
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/classes.json');
    data = arr;
    classMap = initClassesByKey(arr);
    classMapByProgram = initClassesByProgram(arr);
    classMapBySemester = initClassesBySemester(arr);
    fetched = true;
  }
}

function initClassesByKey(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.key;
    if (id) {
      map[id] = filterClassObj(obj);
    }
  });
  return map;
}

function initClassesByProgram(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.programId;
    if (id) {
      obj = filterClassObj(obj);
      map[id] = map[id] || [];
      map[id].push(obj);
    }
  });
  return map;
}

function initClassesBySemester(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.semesterId;
    if (id) {
      obj = filterClassObj(obj);
      map[id] = map[id] || [];
      map[id].push(obj);
    }
  });
  return map;
}

function filterClassObj(obj) {
  if (!obj.filtered) {
    obj.isAvailable = convertStringToBool(obj.isAvailable);
    obj.times = convertStringArray(obj.times);
    obj.allYear = convertStringToBool(obj.allYear);
    obj.fullState = convertStringToBool(obj.fullState);
    obj.filtered = true;
  }
  return obj;
}
