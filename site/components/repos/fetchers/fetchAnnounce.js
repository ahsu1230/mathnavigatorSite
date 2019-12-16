'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';
import { convertStringArray, convertStringToBool } from './fetcherUtil.js';

var fetched = false;
var data;
var announceList;

export var getAnnouncements = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(announceList);
  });
}

export const fetcher = {
  getAnnouncements: getAnnouncements
}

/* Helper functions */
function fetch() {
  if (!fetched) {
    const arr = require('./json/announcements.json');
    data = arr;
    announceList = initAnnounce(data);
    fetched = true;
  }
}

function initAnnounce(arr) {
  var list = [];
  forEach(arr, function(obj) {
    obj.dateStr = obj.date;
    obj.date = new Date(obj.date);
    obj.important = convertStringToBool(obj.important);
    obj.onHomePage = convertStringToBool(obj.onHomePage);
    obj.classKeys = convertStringArray(obj.classKeys);
    list.push(obj);
  });
  return list.reverse(); // Newest post to oldest
}
