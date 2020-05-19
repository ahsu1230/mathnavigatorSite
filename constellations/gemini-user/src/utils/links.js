'use strict';
import { concat, find } from 'lodash';

export const Navigation = [
    { id: "home", name: "Home", url: "/" },
    {
        id: "programs",
        name: "Programs",
        url: "/programs",
        subLinks: SubLinksPrograms,
    },
    {
        id: "success",
        name: "Accomplishments",
        url: "/student-achievements",
        subLinks: SubLinksAchieve,
    },
    { id: "contact", name: "Contact", url: "/contact" },
];

const SubLinksPrograms = [
    { id: "program-catalog", name: "Catalog", url: "/programs" },
    { id: "announcements", name: "Announcements", url: "/announcements" },
    { id: "ask-for-help", name: "Ask For Help", url: "/ask-for-help" },
];

const SubLinksAchieve = [
    {
        id: "student-achieve",
        name: "Student Achievements",
        url: "/student-achievements",
    },
    {
        id: "student-webdev",
        name: "Student Web Development",
        url: "/student-webdev",
    },
    {
        id: "student-portfolios",
        name: "Student Websites",
        url: "/student-projects",
    },
];

const AllNavLinks = concat(Navigation, SubLinksPrograms, SubLinksAchieve);

export function getNavById(id) {
  return find(AllNavLinks, {id: id});
}

export function getNavByUrl(url) {
  return find(AllNavLinks, {url: url});
}

/* not really used */
export function isPathAt(currentPath, url) {
    if (url == getNavById("home").url) {
      // return currentPath == '/'; // Use with BrowserRouter
      return currentPath == '#/';
    } else {
      return currentPath.indexOf(url) >= 0;
    }
}