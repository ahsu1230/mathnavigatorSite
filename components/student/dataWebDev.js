'use strict';
const srcIg = require('./../../assets/students/webdev/evelyn_ig.png');
const srcTesla = require('../../assets/students/webdev/tesla.jpg');
const srcProfileWide = require('../../assets/students/webdev/gaoxing_profile_wide.png');
const srcProfileSlim = require('../../assets/students/webdev/jessica_profile_slim2.png');
const srcYelp = require('../../assets/students/webdev/daniel_yelp.png');

export const SiteDescription = "When enrolled in our Website Development Program, students will create multiple mini-projects as exercises to apply their knowledge in website development. These small projects will allow students to recreate popular website designs and to creatively design their own!";

export const sectionIg = {
  title: "Instagram Header",
  description: "Our first project in this class is recreating the header of Instagram's website. We teach students how to render images and position them to perfectly match the real website's header pixel by pixel!",
  imgSrc: srcIg,
  student1: "Evelyn",
  student2: "8th Grade"
};

export const sectionTesla = {
  title: "Tesla's Social Media",
  description: "Many large companies provide links to their social media pages in order to increase user engagement. But simply providing hyperlinks is boring - interactive icons are better. Hover over the social media icons under the Tesla logo to try it yourself!",
  imgSrc: srcTesla,
  student1: "Richard",
  student2: "12th Grade"
};

export const sectionProfile = {
  title: "Professional Profile Cards",
  info1: {
    description: "Web design is a crucial element for modern websites to succeed. We encourage students to make their websites not only functional, but also aesthetically pleasing. In this project, we study how to create design-forward cards to present people or creative content in a professional fashion.",
    imgSrc: srcProfileWide,
    student1: "Gaoxing",
    student2: "12th Grade"
  },
  info2: {
    description: "With the introduction of mobile phones and tablets, websites are no longer browsed only with desktops. So we have our students create responsive websites that alter their structure based on browser screen size. Make your browser window larger or smaller to see the difference!",
    imgSrc: srcProfileSlim,
    student1: "Jessica",
    student2: "12th Grade"
  }
};

export const sectionYelp = {
  title: "Yelp Search Results",
  description: "Using some advanced Javascript techniques, you can produce long lists of content like Yelp's search results. In our class, students simulate the same functionality to produce pixel-perfect Yelp content!",
  imgSrc: srcYelp,
  student1: "Daniel",
  student2: "11th Grade"
};
