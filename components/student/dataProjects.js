'use strict';
export const SiteDescription = "As the final project of our Website Development Program, students create a personal website using the skills they've learned from our sessions. These websites can be for personal portfolios and hobbies or even for their school clubs.";

const srcThumbnailAries = require('../../assets/students/projects/summer2019_aries.png');
const srcThumbnailDanielZ = require('../../assets/students/projects/summer2019_danielz.png');
const srcThumbnailJessica = require('../../assets/students/projects/summer2019_jessica.png');
const srcThumbnailRichard = require('../../assets/students/projects/summer2019_richard.png');
const srcThumbnailNoah = require('../../assets/students/projects/summer2019_noah.png');
const srcThumbnailDanielL = require('../../assets/students/projects/summer2019_daniell.png');
const srcThumbnailBobby = require('../../assets/students/projects/summer2019_bobby.png');

const summer2019 = [
  {
    student1: "Noah",
    grade: "9th Grade",
    school: "Clarksburg H.S.",
    title: "ShotByNoah",
    description: "I made a website to help inform others about me and my experiences with photography.",
    imgSrc: srcThumbnailNoah,
    url: "/2019_summer/noah_a/index.html"
  },
  {
    student1: "Daniel",
    grade: "10th Grade",
    school: "Thomas Wootton H.S.",
    title: "Alpha Sirius",
    description: "This website is about exploring space beyond the moon and into the stars. NASA has achieved a lot of space missions in the past and is wanting to take humans one step forward.",
    imgSrc: srcThumbnailDanielZ,
    url: "/2019_summer/daniel_z/indexProX.html"
  },
  {
    student1: "Aries",
    grade: "11th Grade",
    school: "Richard Montgomery H.S.",
    title: "Aries's Website",
    description: "This website is a concept testing website about equipment used in the military such as planes and vehicles.",
    imgSrc: srcThumbnailAries,
    url: "/2019_summer/aries_w/index.html"
  },
  {
    student1: "Daniel",
    grade: "11th Grade",
    school: "Winston Churchill H.S.",
    title: "Starbucks",
    description: "This is a website in which you can custom order your own coffee, and receive a receipt for your order.",
    imgSrc: srcThumbnailDanielL,
    url: "/2019_summer/daniel_l/index.html"
  },
  {
    student1: "Bobby",
    grade: "10th Grade",
    school: "Winston Churchill H.S.",
    title: "Froshmeme Central",
    description: "This is a website that is supposed to be something fun about a group that I am involved with on the XC team.",
    imgSrc: srcThumbnailBobby,
    url: "/2019_summer/bobby_d/html/index.html"
  },
  {
    student1: "Jessica",
    grade: "12th Grade",
    school: "Montgomery Blair H.S.",
    title: "Personal Portfolio",
    description: "",
    imgSrc: srcThumbnailJessica,
    url: undefined
  },

  {
    student1: "Richard",
    grade: "12th Grade",
    school: "Walt Whitman H.S.",
    title: "Whitman Math Team",
    description: "",
    imgSrc: srcThumbnailRichard,
    url: undefined
  },
  {
    student1: "Sean",
    grade: "8th Grade",
    school: "Takoma Park M.S.",
    title: "German Jet Fighters",
    description: "",
    imgSrc: undefined,
    url: undefined
  }
];


export const Projects = [
  {
    sectionTitle: "Summer 2019",
    projects: summer2019
  }
];
