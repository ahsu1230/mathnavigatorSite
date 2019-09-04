'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var data;
var semesterIds;
var semesterMap;

export var getSemesterIds = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(semesterIds);
  });
}

export var getSemester = function(semesterId) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(semesterMap[semesterId]);
  });
}

export const fetcher = {
  getSemesterIds: getSemesterIds,
  getSemester: getSemester
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/semesters.json');
    data = arr;
    semesterIds = initSemesterIds(data);
    semesterMap = initSemesterMap(data);
    fetched = true;
  }
}

function initSemesterIds(arr) {
  var list = [];
  forEach(arr, function(obj) {
    list.push(obj.semesterId);
  });
  return list;
}

function initSemesterMap(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var id = obj.semesterId;
    map[id] = obj;
  });
  return map;
}
