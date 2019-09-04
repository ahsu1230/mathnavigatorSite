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
import { fetcher as FetchSessions } from './fetchers/fetchSessions.js';

/* Achievements */
export function getAchievementYears() {
  return FetchAchieve.getAchievementYears();
}

export function getAchievementsByYear(year) {
  return FetchAchieve.getAchievementsByYear(year);
}

export function getAchievementsByYears(years) {
  return FetchAchieve.getAchievementsByYears(years);
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
export function getClassesByProgram(programId) {
  return FetchClasses.getClassesByProgram(programId);
}

export function getClassesBySemester(semesterId) {
  return FetchClasses.getClassesBySemester(semesterId);
}

export function getClassesByProgramAndSemester(programId, semesterId) {
  return FetchClasses.getClassesByProgramAndSemester(programId, semesterId);
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
export function getProgramById(programId) {
  return FetchPrograms.getProgramById(programId);
}

export function getProgramByIds(arr) {
  return FetchPrograms.getProgramsByIds(arr);
}

export function getProgramByClassKey(classKey) {
  return new Promise(function(resolve, reject) {
  });

  // return getClassByKey(classKey).then(function() {
    //
  // });
}

export function getProgramsBySemester(semesterId) {
  return new Promise(function(resolve, reject) {
  });

  // var map = {};
  // forEach(classMapByKey, function(classObj) {
  //   var programId = classObj.programId;
  //   var programObj = programMap[programId];
  //   var semesterId = classObj.semesterId;
  //   programObj.semesterId = semesterId;
  //
  //   map[semesterId] = map[semesterId] || [];
  //   var hasProgram = find(map[semesterId], {programId: programId, semesterId: semesterId});
  //   if (!hasProgram) {
  //     map[semesterId].push(programObj);
  //   }
  // });
  // return map;
}

/* Semesters */
export function getSemesterIds() {
  return FetchSemesters.getSemesterIds();
}

export function getSemester(semesterId) {
  return FetchSemesters.getSemester(semesterId);
}

/* Sessions */
export function getSessions(classKey) {
  return FetchSessions.getSessionsByClass(classKey);
}
