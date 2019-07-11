'use strict';
var _ = require('lodash/core');

import profiles from './json/profiles.json';
import imgHashes from './../../assets/profiles/*.jpg';

function initAllProfiles() {
  console.log("Initializing Profiles...");
  var keys = _.keys(imgHashes);
  _.forEach(profiles, function(profile) {
    var profileId = profile.id;
    var imgHash = imgHashes[profileId];
    profile.imgSrc = imgHash;
  });
  console.log("Profiles done initializing.");
}

export function getProfileById(profileId) {
  return _.find(profiles, {id: profileId});
}

export function getAllProfiles() {
  return profiles;
}

initAllProfiles();
