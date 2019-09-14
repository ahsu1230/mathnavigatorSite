'use strict';
export const SiteDescription = "As the final project of our Website Development Program, students create a personal website using the skills they've learned from our sessions. These websites can be for personal portfolios and hobbies or even for their school clubs.";

const srcThumbnailAries = require('../../assets/students/projects/aries_summer2019.png');
const srcThumbnailDaniel = require('../../assets/students/projects/danielz_summer2019.png');
const srcThumbnailJessica = require('../../assets/students/projects/jessica_summer2019.png');
const srcThumbnailRichard = require('../../assets/students/projects/richard_summer2019.png');

const summer2019 = [
  {
    student1: "Jessica",
    grade: "12th Grade",
    school: "Montgomery Blair H.S.",
    title: "Personal Portfolio",
    description: "",
    imgSrc: srcThumbnailJessica
  },
  {
    student1: "Daniel",
    grade: "10th Grade",
    school: "Thomas Wootton H.S.",
    title: "Alpha Sirius",
    description: "",
    imgSrc: srcThumbnailDaniel
  },
  {
    student1: "Richard",
    grade: "12th Grade",
    school: "Walt Whitman H.S.",
    title: "Whitman Math Team",
    description: "",
    imgSrc: srcThumbnailRichard
  },
  {
    student1: "Aries",
    grade: "11th Grade",
    school: "Richard Montgomery H.S.",
    title: "Military Planes",
    description: "",
    imgSrc: srcThumbnailAries
  },
  {
    student1: "Noah",
    grade: "9th Grade",
    school: "Clarksburg H.S.",
    title: "Personal Portfolio",
    description: "",
    imgSrc: undefined
  },
  {
    student1: "Sean",
    grade: "8th Grade",
    school: "Takoma Park M.S.",
    title: "German Jet Fighters",
    description: "",
    imgSrc: undefined
  }
];


export const Projects = [
  {
    sectionTitle: "Summer 2019",
    projects: summer2019
  }
];
