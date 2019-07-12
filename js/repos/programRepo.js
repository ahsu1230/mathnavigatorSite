'use strict';
var _ = require('lodash/core');
import {
  programMap,
  classMapByKey,
  classMapByProgramId
} from './initPrograms.js';

export function getAvailablePrograms() {
  var mapAvail = {};
  var mapSoon = {};

  _.forEach(classMapByKey, function(obj) {
    var programId = obj.programId;
    var program = programMap[programId];

    if (obj.isAvailable) {
      if (!mapAvail[programId]) {
        mapAvail[programId] = program;
      }
    } else {
      if (!mapSoon[programId]) {
        mapSoon[programId] = program;
      }
    }
  });

  return {
    "available" : mapAvail,
    "soon" : mapSoon
  };
}

export function getProgramByIds(arr) {
  return _.map(arr, function(programId) {
    return programMap[programId];
  });
}

export function getClasses(programId) {
  return classMapByProgramId[programId];
}

export function getProgramClass(key) {
  var classObj = classMapByKey[key];
  var programId = classObj.programId;
  var programObj = programMap[programId];
  return {
    programObj: programMap[programId],
    classObj: classObj
  };
}
