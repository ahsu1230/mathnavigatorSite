'use strict';
import {
  find,
  filter,
  forEach,
  map
} from 'lodash';
import {
  announceList,
  classMapByKey,
  classMapByProgramId,
  classMapBySemesterId,
  locationMap,
  preReqMap,
  programMap,
  sessionMap,
  semesterMap,
  semesterIds
} from './initPrograms.js';
import { profileList } from './initProfiles.js';

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
  return filter(classMapBySemesterId[semesterId], function (c) {
    return c.programId === programId;
  });
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

export function getProgramsBySemester(semesterId) {
  var map = {};
  forEach(classMapByKey, function(classObj) {
    var programId = classObj.programId;
    var programObj = programMap[programId];
    var semesterId = classObj.semesterId;
    programObj.semesterId = semesterId;

    var hasProgram = find(map[semesterId], {programId: programId, semesterId: semesterId});
    if (map[semesterId]) {
      if (!hasProgram) {
        map[semesterId].push(programObj);
      }
    } else {
      map[semesterId] = [programObj];
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

function createFullClassObj(classObj, programObj) {
  var className = classObj.className;
  classObj.fullClassName = programObj.title + (className ? (" " + className) : "");
  classObj.programTitle = programObj.title;
  classObj.grade1 = programObj.grade1;
  classObj.grade2 = programObj.grade2;
  classObj.description = programObj.description;
  return classObj;
}
