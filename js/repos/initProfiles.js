'use strict';
var _ = require('lodash/core');

import profiles from './json/profiles.json';
import imgHashes from './../../assets/profiles/*.jpg';

export var profileList = [];

function init() {
  console.log("Initializing Profiles...");
  var keys = _.keys(imgHashes);
  _.forEach(profiles, function(profile) {
    var profileId = profile.id;
    var imgHash = imgHashes[profileId];
    profile.imgSrc = imgHash;
    profileList.push(profile);
  });
  console.log("Profiles done initializing.");
}

init();
