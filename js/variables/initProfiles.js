'use strict';
var _ = require('lodash/core');

import imgHashes from './../../assets/profiles/*.jpg';

export const profiles = [
  {
    id: 'moosey',
    name: "Moosey",
    subtitle1: "University of Geniuses, College Park",
    subtitle2: "Chief Happiness Officer, Nextbit",
    quote: "Andy is probably the most amazing teacher I’ve ever met. He genuinely cares about his students and has a hilarious teaching style. This human feeds me, which is nice. I like food. I like the other humans better, but he's cool. "
  },
  {
    id: 'goosey',
    name: "Goosey",
    subtitle1: "Yale University",
    subtitle2: "Founder of GuudBoyz",
    quote: "Andy is probably the most amazing teacher I’ve ever met. He genuinely cares about his students and has a hilarious teaching style. This human feeds me, which is nice. I like food. I like the other humans better, but he's cool. "
  }
];

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
