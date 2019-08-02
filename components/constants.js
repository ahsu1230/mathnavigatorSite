'use strict';
import { find } from 'lodash';

export const NavLinks = [
  { id: "home", name: "Home", url: "/" },
  { id: "programs", name: "Programs", url: "/programs" },
  { id: "announce", name: "Announcements", url: "/announcements" },
  { id: "achieve", name: "Achievements", url: "/student-achievements"},
  { id: "contact", name: "Contact", url: "/contact" }
];

export function getNav(id) {
  return find(NavLinks, {id: id});
}

export function isPathAt(url) {
  var path = location.hash;
  if (url == getNav("home").url) {
    return path == '#/';
  } else {
    return path.indexOf(url) >= 0;
  }
}
