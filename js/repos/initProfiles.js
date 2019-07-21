'use strict';
import { forEach, keys } from 'lodash';

import profiles from './json/profiles.json';
import imgHashes from './../../assets/profiles/*.jpg';

export var profileList = [];

function init() {
  console.log("Initializing Profiles...");
  forEach(profiles, function(profile) {
    var profileId = profile.id;
    var imgHash = imgHashes[profileId];
    profile.imgSrc = imgHash;
    profileList.push(profile);
  });
  console.log("Profiles done initializing.");
}

init();
