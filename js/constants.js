'use strict';
var _ = require('lodash/core');

export const NavLinks = [
  { id: "home", name: "Home", url: "/" },
  { id: "announce", name: "Announcements", url: "/announcements" },
  { id: "programs", name: "Programs", url: "/programs" },
  { id: "contact", name: "Contact", url: "/contact" }
];

export function getNav(id) {
  return _.find(NavLinks, {id: id});
}

export function isPathAt(url) {
  var path = location.hash;
  if (url == getNav("home").url) {
    return path == '#/';
  } else {
    return path.indexOf(url) >= 0;
  }
}
