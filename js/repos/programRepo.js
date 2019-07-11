'use strict';
var _ = require('lodash/core');
import { programMap, classMapByKey } from './initPrograms.js';

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

export function getClasses(programId) {
  return [];
}
