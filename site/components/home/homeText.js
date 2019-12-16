'use strict';

const srcSat = require('../../assets/cb_sat.jpg');
const srcApCalc = require('../../assets/ap_calc.jpg');
const srcAmc = require('../../assets/amc.jpg');
const srcUserCheck = require('../../assets/user_check_white.svg');
const srcLightBulb = require('../../assets/lightbulb_white.svg');
const srcLibrary = require('../../assets/library_white.svg');

export const programsText = [
  {
    title: "SAT",
    description:
      "Our most popular programs. We provide dedicated programs for both the SAT1 Math Section and SAT Subject Test (Math Level 2).",
    imgSrc: srcSat,
    programId: ""
  },
  {
    title: "AP Calculus",
    description: "Get college credits early by acing the AP Calculus Exams. This program is dedicated to both Calculus AB and BC students.",
    imgSrc: srcApCalc,
    programId: ""
  },
  {
    title: "AMC",
    description: "The American Mathematics Competition administered by the Mathematics Association of America. Represent your school and compete with other mathematics athletes.",
    imgSrc: srcAmc,
    programId: ""
  }
];

export const successText = [
  {
    title: "Dedicated Professionalism",
    description: "We dedicate time to each individual student and provide plans to improve topics they may need more help in.",
    imgSrc: srcUserCheck
  },
  {
    title: "Strategical Thinking",
    description: "We want our students to work smart, not just hard. With Andy's problem solving strategies, students will learn how to tackle any problem.",
    imgSrc: srcLightBulb
  },
  {
    title: "Abundant Resources",
    description: "Andy makes all of the teaching material. With our huge collection of resources, students will be over-prepared for their next exams.",
    imgSrc: srcLibrary
  }
];
