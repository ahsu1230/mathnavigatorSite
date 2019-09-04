'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var achieveYearMap;

var getAchievementYears = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(keys(achieveYearMap));
  });
}

var getAchievementsByYear = function(year) {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(achieveYearMap[year]);
  });
}

var getAchievementsByYears = function(years) {
  return new Promise(function(resolve, reject) {
    fetch();
    var map = {};
    forEach(years, function(year) {
      map[year] = achieveYearMap[year];
    });
    resolve(map);
  });
}

export const fetcher = {
  getAchievementYears: getAchievementYears,
  getAchievementsByYear: getAchievementsByYear,
  getAchievementsByYears: getAchievementsByYears
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/achievements.json');
    achieveYearMap = initMap(arr);
    fetched = true;
  }
}

function initMap(arr) {
  var map = {};
  forEach(arr, function(obj) {
    var key = obj.year;
    obj.highlight = convertStringToBool(obj.highlight);
    obj.classKeys = convertStringArray(obj.classKeys);
    map[key] = map[key] || [];
    map[key].push(obj);
  });
  return map;
}
