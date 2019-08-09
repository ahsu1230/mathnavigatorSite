'use strict';
import { find } from 'lodash';

export const websiteName = "Math Navigator";

export function createPageTitle(title) {
  return websiteName + " - " + title;
}

export const NavLinks = [
  { id: "home", name: "Home", url: "/" },
  { id: "programs", name: "Programs", url: "/programs" },
  { id: "announce", name: "Announcements", url: "/announcements" },
  { id: "achieve", name: "Achievements", url: "/student-achievements"},
  { id: "contact", name: "Contact", url: "/contact" }
];

export function getNavById(id) {
  return find(NavLinks, {id: id});
}

export function getNavByUrl(url) {
  return find(NavLinks, {url: url});
}

export function isPathAt(currentPath, url) {
  if (url == getNavById("home").url) {
    return currentPath == '/';
  } else {
    return currentPath.indexOf(url) >= 0;
  }
}
