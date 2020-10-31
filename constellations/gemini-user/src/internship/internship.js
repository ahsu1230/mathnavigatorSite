"use strict";
require("./internship.sass");
import React from "react";

import GithubImg from "./../../assets/intern_images/github.png";
import GolangImg from "./../../assets/intern_images/golang.png";
import ReactImg from "./../../assets/intern_images/react.png";
import NpmImg from "./../../assets/intern_images/npm.png";
import MySQLImg from "./../../assets/intern_images/mysql.png";
import DockerImg from "./../../assets/intern_images/docker.jpg";
import CircleImg from "./../../assets/intern_images/circleci.png";

import DevPullRequestImg from "./../../assets/intern_images/github_pull_request.png";
import DevSiteImg from "./../../assets/intern_images/develop_site.png";

import { PastInterns } from "./pastInterns.js";

export class InternshipPage extends React.Component {
    render() {
        const interns = PastInterns.map((intern, index) => {
            return (
                <div className="intern-container" key={index}>
                    <div className="name">{intern.name}</div>
                    <div className="school">{intern.school}</div>
                    <div className="college">{intern.college}</div>
                </div>
            );
        });

        return (
            <div id="view-intern">
                <h1 className="title">Software Development Internship</h1>
                <p className="intro">
                    Math Navigator occasionally offers Software Development
                    Internship for ambitious students who would like to pursue a
                    college major or a career in software development. In this
                    internship opportunity, students will learn the fundamentals
                    of modern technologies commonly used across
                    engineering-focused companies. And using these skills,
                    students help create this very site you are using!
                </p>

                <h1>Technology Stack</h1>
                <section className="tech-stack">
                    <div className="img-container github">
                        <img src={GithubImg} />
                    </div>

                    <div className="img-container react">
                        <img src={ReactImg} />
                    </div>

                    <div className="img-container golang">
                        <img src={GolangImg} />
                    </div>

                    <div className="img-container mysql">
                        <img src={MySQLImg} />
                    </div>

                    <div className="img-container docker">
                        <img src={DockerImg} />
                    </div>

                    <div className="img-container circleci">
                        <img src={CircleImg} />
                    </div>
                </section>

                <section className="structure1">
                    <h1>Internship Structure</h1>
                    <p>
                        The internship is split into various roles. Interns can
                        choose to join the front-end team to learn website
                        development with ReactJs or join the back-end team to
                        learn data management with Golang and MySQL. There are a
                        few students who have done internships for both teams!
                    </p>
                    <div className="role-container">
                        <div className="role-card">
                            <h4>Front-end Developer Intern</h4>
                            <div className="container">
                                <img src={NpmImg} className="npm" />
                                <img src={ReactImg} className="react" />
                            </div>
                        </div>
                        <div className="role-card">
                            <h4>Back-end Developer Intern</h4>
                            <div className="container">
                                <img src={GolangImg} className="golang" />
                                <img src={MySQLImg} className="mysql" />
                            </div>
                        </div>
                    </div>
                </section>

                <h1>Our Past Interns:</h1>
                <section className="past-interns">{interns}</section>

                <section>
                    <h1>Math Navigator Products</h1>
                    <p id="last-paragraph">Coming soon...</p>
                </section>

                <section id="last">
                    <p>
                        Internship opportunities will be announced when they are
                        available. When they are, student candidates must first
                        pass a series of coding assessments and interviews in
                        order to receive a position. This interview process
                        reflects the same process that famous technology
                        companies like Google, Facebook, Amazon, etc. have. We
                        encourage all students to attempt these assessments to
                        familiarize themselves with this interview process
                        structure if they ever want to pursue a career at any of
                        these companies.
                    </p>
                </section>
            </div>
        );
    }
}
