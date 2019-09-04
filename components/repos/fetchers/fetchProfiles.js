'use strict';
import { forEach, keys } from 'lodash';
import { Promise } from 'bluebird';

var fetched = false;
var dataProfiles;
var dataImgHashes;
var profileList = [];

export var getProfileList = function() {
  return new Promise(function(resolve, reject) {
    fetch();
    resolve(profileList);
  });
}

export const fetcher = {
  getProfileList: getProfileList
}

function fetch() {
  if (!fetched) {
    dataProfiles = require('./json/profiles.json');
    dataImgHashes = require('./../../assets/profiles/*.jpg');
    profileList = initProfiles(profiles, imgHashes);
    fetched = true;
  }
}

function initProfiles(profiles, imgHashes) {
  var list = [];
  console.log("Initializing Profiles...");
  forEach(profiles, function(profile) {
    var profileId = profile.id;
    var imgHash = imgHashes[profileId];
    profile.imgSrc = imgHash;
    list.push(profile);
  });
  console.log("Profiles done initializing.");
  return list;
}
