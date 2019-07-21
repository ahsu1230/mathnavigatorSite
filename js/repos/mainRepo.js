'use strict';
import { find, forEach, map } from 'lodash';
import {
  announceList,
  classMapByKey,
  classMapByProgramId,
  locationMap,
  preReqMap,
  programMap,
  sessionMap
} from './initPrograms.js';
import { profileList } from './initProfiles.js';

export function getAnnounceList() {
  return announceList;
}

export function getLocation(locationId) {
  return locationMap[locationId];
}

export function getPrereqs(programId) {
  return preReqMap[programId];
}

export function getSessions(programClassKey) {
  return sessionMap[programClassKey];
}

export function getAvailablePrograms() {
  var mapAvail = {};
  var mapSoon = {};

  forEach(classMapByKey, function(obj) {
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
  return map(arr, function(programId) {
    return programMap[programId];
  });
}

export function getClasses(programId) {
  return classMapByProgramId[programId];
}

export function getAvailableClasses() {
  var pair = getAvailablePrograms();
  var progsAvail = pair.available;
  var progsSoon = pair.soon;

  var classesAvail = [];
  var classesSoon = [];

  forEach(progsAvail, function(progObj) {
    var programId = progObj.programId;
    var classList = classMapByProgramId[programId];
    var fullClassList = classList.map(classObj =>
      createFullClassObj(classObj, progObj)
    );
    classesAvail = classesAvail.concat(fullClassList);
  });

  forEach(progsSoon, function(progObj) {
    var programId = progObj.programId;
    var classList = classMapByProgramId[programId];
    var fullClassList = classList.map(classObj =>
      createFullClassObj(classObj, progObj)
    );
    classesSoon = classesSoon.concat(fullClassList);
  });

  return {
    "available": classesAvail,
    "soon": classesSoon
  }
}

export function getProgramClass(key) {
  var classObj = classMapByKey[key];
  if (!classObj) { return {}; }

  var programId = classObj.programId;
  var programObj = programMap[programId];
  return {
    programObj: programMap[programId],
    classObj: classObj
  };
}

export function getProfileById(profileId) {
  return find(profileList, {id: profileId});
}

export function getAllProfiles() {
  return profileList;
}

function createFullClassObj(classObj, programObj) {
  var className = classObj.className;
  classObj.fullClassName = programObj.title + (className ? (" " + className) : "");
  classObj.programTitle = programObj.title;
  classObj.grade1 = programObj.grade1;
  classObj.grade2 = programObj.grade2;
  classObj.description = programObj.description;
  return classObj;
}
