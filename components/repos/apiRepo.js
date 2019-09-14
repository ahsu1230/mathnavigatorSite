'use strict';
import { Promise } from 'bluebird';
import { forEach, find } from 'lodash';
import { fetcher as FetchAchieve } from './fetchers/fetchAchieve.js';
import { fetcher as FetchAfh } from './fetchers/fetchAfh.js';
import { fetcher as FetchAnnounce } from './fetchers/fetchAnnounce.js';
import { fetcher as FetchClasses } from './fetchers/fetchClasses.js';
import { fetcher as FetchOther } from './fetchers/fetchOther.js';
import { fetcher as FetchPrereqs } from './fetchers/fetchPrereqs.js';
import { fetcher as FetchProfiles } from './fetchers/fetchProfiles.js';
import { fetcher as FetchPrograms } from './fetchers/fetchPrograms.js';
import { fetcher as FetchSemesters } from './fetchers/fetchSemesters.js';
import { fetcher as FetchSessions } from './fetchers/fetchSessions.js';

/* Achievements */
export function getAchievementsByYears() {
  return FetchAchieve.getAchievementsByYears();
}

/* Announcements */
export function getAnnouncements() {
  return FetchAnnounce.getAnnouncements();
}

/* AskForHelp */
export function getAFH() {
  return FetchAfh.getAFH();
}

/* Classes */
export function getAllClasses() {
  return FetchClasses.getAllClasses();
}

export function getClass(classKey) {
  return FetchClasses.getClassByKey(classKey);
}

export function getClasses(classKeys) {
  return FetchClasses.getClassesByKeys(classKeys);
}

export function getClassesByProgram(programId) {
  return FetchClasses.getClassesByProgram(programId);
}

export function getClassesBySemester(semesterId) {
  return FetchClasses.getClassesBySemester(semesterId);
}

export function getClassesByProgramAndSemester(programId, semesterId) {
  return FetchClasses.getClassesByProgramAndSemester(programId, semesterId);
}

export function getAllClassesBySemesters() {
  return Promise.join(getAllSemesters(), getAllClasses(), getAllPrograms(),
    function(semesterMap, classMap, programMap) {
      var map = {};
      forEach(classMap, function(classObj) {
        var semesterId = classObj.semesterId;
        var programId = classObj.programId;
        var programObj = programMap[programId];
        classObj.programObj = programObj;

        map[semesterId] = map[semesterId] || [];
        map[semesterId].push(classObj);
      });
      return {
        semesterMap: semesterMap,
        classSemesterMap: map
      };
    }
  );
}

/* Key Values */
export function getKeyValue(key) {
  return FetchOther.getKeyValue(key);
}

/* Locations */
export function getLocation(locationId) {
  return FetchOther.getLocation(locationId);
}

/* Prereqs */
export function getPrereqs(programId) {
  return FetchPrereqs.getPrereqs(programId);
}

/* Profiles */
export function getProfileById(profileId) {
  return FetchProfiles.getProfile(profileId);
}

export function getAllProfiles() {
  return FetchProfiles.getAll();
}


/* Programs */
export function getAllPrograms() {
  return FetchPrograms.getAllPrograms();
}

export function getProgramById(programId) {
  return FetchPrograms.getProgramById(programId);
}

export function getProgramByIds(arr) {
  return FetchPrograms.getProgramsByIds(arr);
}

export function getProgramAndClass(classKey) {
  var targetClass;
  return getClass(classKey).then(classObj => {
    targetClass = classObj;
    return classObj.programId;
  })
  .then(programId => { return getProgramById(programId) })
  .then(programObj => {
    return {
      classObj: targetClass,
      programObj: programObj
    };
  });
}

export function getAllProgramsAndClasses() {
  return Promise.join(getAllClasses(), getAllPrograms(), (classes, programs) => {
    var map = {};
    forEach(classes, classObj => {
      var classKey = classObj.key;
      var programId = classObj.programId;
      var programObj = programs[programId];
      map[classKey] = {
        classObj: classObj,
        programObj: programObj
      };
    });
    return map;
  });
}

export function getProgramsBySemesters() {
  return Promise.join(getAllSemesters(), getAllClasses(), getAllPrograms(),
    function(semesterMap, classMap, programMap) {
      var map = {};
      forEach(classMap, function(classObj) {
        var semesterId = classObj.semesterId;
        var programId = classObj.programId;
        var programObj = programMap[programId];
        programObj.semesterId = semesterId;
        map[semesterId] = map[semesterId] || [];
        var hasProgram = find(map[semesterId], {programId: programId, semesterId: semesterId});
        if (!hasProgram) {
          map[semesterId].push(programObj);
        }
      });
      return {
        semesterMap: semesterMap,
        programSemesterMap: map
      };
    }
  );
}

/* Semesters */
export function getAllSemesters() {
  return FetchSemesters.getAllSemesters();
}

export function getSemester(semesterId) {
  return FetchSemesters.getSemester(semesterId);
}

export function getSemesterIds() {
  return FetchSemesters.getSemesterIds();
}

/* Sessions */
export function getSessions(classKey) {
  return FetchSessions.getSessionsByClass(classKey);
}
