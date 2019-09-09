'use strict';
import { find } from 'lodash';

export const websiteName = "Math Navigator";

export function createPageTitle(title) {
  return websiteName + " - " + title;
}

const SubLinksPrograms = [
  { id: "program-catalog", name: "Catalog", url: "/programs" },
  { id: "announcements", name: "Announcements", url: "/announcements" },
  { id: "ask-for-help", name: "Ask For Help", url: "/ask-for-help" }
];

const SubLinksAchieve = [
  { id: "student-achieve", name: "Student Achievements", url: "/student-achievements" },
  { id: "student-webdev", name: "Student Web Development", url: "/student-webdev" },
  { id: "student-portfolios", name: "Student Websites", url: "/student-projects" }
];

export const NavLinks = [
  { id: "home", name: "Home", url: "/" },
  { id: "programs", name: "Programs", url: "/programs", subLinks: SubLinksPrograms },
  { id: "success", name: "Accomplishments", url: "/student-achievements", subLinks: SubLinksAchieve },
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
