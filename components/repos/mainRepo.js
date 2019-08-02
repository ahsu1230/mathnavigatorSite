'use strict';
import {
  assign,
  find,
  filter,
  forEach,
  keys,
  map
} from 'lodash';
import {
  achievementMap,
  announceList,
  classMapByKey,
  classMapByProgramId,
  classMapBySemesterId,
  keyValuesMap,
  locationMap,
  preReqMap,
  programMap,
  sessionMap,
  semesterMap,
  semesterIds
} from './initPrograms.js';
import { profileList } from './initProfiles.js';

/* Achievements */
export function getAchievementKeys() {
  return keys(achievementMap);
}

export function getAchievements(yearKey) {
  return achievementMap[yearKey];
}

/* Announcements */
export function getAnnounceList() {
  return announceList;
}

/* Classes */
export function getClasses(programId) {
  return classMapByProgramId[programId];
}

export function getClassesBySemester(semesterId) {
  return classMapBySemesterId[semesterId];
}

export function getClassesByProgramAndSemester(programId, semesterId) {
  var programObj = programMap[programId];
  return filter(classMapBySemesterId[semesterId], c => c.programId === programId);
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

/* Key Values */
export function getKeyValue(key) {
  return keyValuesMap[key];
}

/* Locations */
export function getLocation(locationId) {
  return locationMap[locationId];
}

/* Prereqs */
export function getPrereqs(programId) {
  return preReqMap[programId];
}


/* Profiles */
export function getProfileById(profileId) {
  return find(profileList, {id: profileId});
}

export function getAllProfiles() {
  return profileList;
}


/* Programs */
export function getProgramByIds(arr) {
  return map(arr, function(programId) {
    return programMap[programId];
  });
}

export function getProgramsBySemester() {
  var map = {};
  forEach(classMapByKey, function(classObj) {
    var programId = classObj.programId;
    var programObj = programMap[programId];
    var semesterId = classObj.semesterId;
    programObj.semesterId = semesterId;

    map[semesterId] = map[semesterId] || [];
    var hasProgram = find(map[semesterId], {programId: programId, semesterId: semesterId});
    if (!hasProgram) {
      map[semesterId].push(programObj);
    }
  });
  return map;
}


/* Semesters */
export function getSemesterIds() {
  return semesterIds;
}

export function getSemester(semesterId) {
  return semesterMap[semesterId];
}


/* Sessions */
export function getSessions(programClassKey) {
  return sessionMap[programClassKey];
}


/* Helpers */
export function createFullClassObj(classObj) {
  var programId = classObj.programId;
  var programObj = programMap[programId];
  var className = classObj.className;
  classObj = assign({}, classObj, programObj);
  classObj.fullClassName = programObj.title + (className ? (" " + className) : "");
  classObj.programTitle = programObj.title;
  return classObj;
}
