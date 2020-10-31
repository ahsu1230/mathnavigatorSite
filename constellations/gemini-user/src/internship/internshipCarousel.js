"use strict";
import React from "react";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader

import DevPullRequestImg from "./../../assets/intern_images/github_pull_request.png";
import DevSiteImg from "./../../assets/intern_images/develop_site.png";
import DevSiteAdminImg from "./../../assets/intern_images/develop_site_admin.png";
import DevMysqlImg from "./../../assets/intern_images/develop_mysql.png";
import DevDockerImg from "./../../assets/intern_images/develop_docker.png";
import DevCircleCiImg from "./../../assets/intern_images/develop_circleci.png";

const carouselArray = [
    {
        image: DevPullRequestImg,
        legend:
            "Create Pull Requests with Github and learn the fundamentals of Version Control.",
    },
    {
        image: DevSiteImg,
        legend:
            "Get practical hands-on experience by develop a product for real users!",
    },
    {
        image: DevMysqlImg,
        legend:
            "Learn MySQL basics and learn to navigate around our many database tables.",
    },
    {
        image: DevDockerImg,
        legend:
            "Work like a professional and run Docker containers on your computer with our easy-to-use scripts!",
    },
    {
        image: DevCircleCiImg,
        legend:
            "Watch codebases automatically build with CircleCi and run over 400 automated tests!",
    },
    {
        image: DevSiteAdminImg,
        legend:
            "There's more under the hood! Help build an admin website to help administrators manage data.",
    },
];

export default class InternshipCarousel extends React.Component {
    render() {
        const carouselImages = carouselArray.map((item, index) => {
            return (
                <div key={index}>
                    <img src={item.image} />
                    <p className="legend">{item.legend}</p>
                </div>
            );
        });
        return <Carousel>{carouselImages}</Carousel>;
    }
}
