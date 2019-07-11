'use strict';
var _ = require('lodash/core');
import { classMap, programMap } from './initPrograms.js';

export function getAvailablePrograms() {
  var mapAvail = {};
  var mapSoon = {};

  _.forEach(classMap, function(obj) {
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
