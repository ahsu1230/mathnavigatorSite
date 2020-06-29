"use strict";
require("./internship.sass");
import React from "react";

import DockerImg from "./../../assets/docker.jpg";
import GoLangImg from "./../../assets/golang.png";
import CircleImg from "./../../assets/circleci.png";
import MySQLImg from "./../../assets/mysql.png";
import ReactImg from "./../../assets/reactjs.png";

export class InternshipPage extends React.Component {
    render() {
        return (
            <div id="view-intern">
                <h1>Software Development Internship</h1>
                <p>
                    Math Navigator occasionally offers Software Development
                    Internship for ambitious students who would like to pursue a
                    college major or a career in software development. In this
                    internship opportunity, students will learn the fundamentals
                    of modern technologies commonly used across
                    engineering-focused companies. And with these skills,
                    students help create this site that you are currently using!
                </p>
                <h1>Technology Stack</h1>

                <div class="container-main">
                    <img src={ReactImg} />
                    <p>
                        ReactJs, created by Facebook, is a popular web framework
                        that helps web developers build powerful single page web
                        applications
                    </p>
                </div>

                <div class="container-main">
                    <img src={GoLangImg} />
                    <p>
                        Golang, created by Google, is an open source programming
                        language that makes it easy to build simple, efficient
                        and reliable software.
                    </p>
                </div>

                <div class="container-main">
                    <img src={MySQLImg} />
                    <p>
                        MySQL, developed by Oracle, is the most popular
                        relational (SQL) database management system and widely
                        used to store application data.
                    </p>
                </div>

                <div class="container-main">
                    <img src={DockerImg} />
                    <p>
                        Docker is the leading containerization platform that
                        helps software developers build applications and
                        efficiently deliver software.
                    </p>
                </div>

                <div class="container-main">
                    <img src={CircleImg} />
                    <p>
                        CircleCI is a modern continuous integration platform
                        that helps software developers easily test and deploy
                        code.
                    </p>
                </div>

                <h1>Internship Structure</h1>
                <p>
                    The internship is split into various roles. Interns can
                    choose to join the front-end team to learn website
                    development with ReactJs or join the back-end team to learn
                    data management with Golang and MySQL. There are a few
                    students who have done internships for both teams!
                </p>

                <div class="intern-main">
                    <h4>Past Interns:</h4>
                    <p>
                        Cathy Y. <br></br> Montgomery Blair HS 2020 <br></br>{" "}
                        Massachusetts Institute of Technology (MIT)
                    </p>
                    <p>
                        Jessica Y. <br></br> Montgomery Blair HS 2020 <br></br>{" "}
                        University of Maryland
                    </p>
                    <p>
                        Max Z. <br></br> Montgomery Blair HS 2021
                    </p>
                    <p>
                        Chujia G. <br></br> Montgomery Blair HS 2022
                    </p>
                    <p>
                        Frederick Z. <br></br> Montgomery Blair HS 2023
                    </p>
                    <p>
                        Tony W. <br></br> Richard Montgomery HS 2021
                    </p>
                    <p>
                        Serena X. <br></br> Winston Churchill HS 2020 <br></br>{" "}
                        University of Pennsylvania
                    </p>
                    <p>
                        Daniel L. <br></br> Winston Churchill HS 2021
                    </p>
                </div>

                <table id="past-intern">
                    <caption>Past Interns:</caption> <br></br>
                    <tr>
                        <th>Name</th>
                        <th>School</th>
                        <th>College</th>
                    </tr>
                    <tr>
                        <td>Cathy Y.</td>
                        <td>Montgomery Blair HS 2020</td>
                        <td>Massachusetts Institute of Technology (MIT)</td>
                    </tr>
                    <tr>
                        <td>Jessica Y.</td>
                        <td>Montgomery Blair HS 2020</td>
                        <td>University of Maryland</td>
                    </tr>
                    <tr>
                        <td>Max Z.</td>
                        <td>Montgomery Blair HS 2021</td>
                        <td> </td>
                    </tr>
                    <tr>
                        <td>Chujia G.</td>
                        <td>Montgomery Blair HS 2022</td>
                        <td> </td>
                    </tr>
                    <tr>
                        <td>Frederick Z.</td>
                        <td>Montgomery Blair HS 2023</td>
                        <td> </td>
                    </tr>
                    <tr>
                        <td>Tony W.</td>
                        <td>Richard Montgomery HS 2021</td>
                        <td> </td>
                    </tr>
                    <tr>
                        <td>Austin H.</td>
                        <td>Winston Churchill HS 2017</td>
                        <td>Carnegie Mellon University</td>
                    </tr>
                    <tr>
                        <td>Serena X.</td>
                        <td>Winston Churchill HS 2020</td>
                        <td>University of Pennsylvania</td>
                    </tr>
                    <tr>
                        <td>Daniel L.</td>
                        <td>Winston Churchill HS 2021</td>
                        <td> </td>
                    </tr>
                </table>

                <p>
                    Internship opportunities will be announced when they are
                    available. When they are, student candidates must first pass
                    a series of coding assessments and interviews in order to
                    receive a position. This interview process reflects the same
                    process that famous technology companies like Google,
                    Facebook, Amazon, etc. have. We encourage all students to
                    attempt these assessments to familiarize themselves with
                    this interview process structure if they ever want to pursue
                    a career at any of these companies.
                </p>

                <h1>Math Navigator Products</h1>
                <p id="last-paragraph">Coming soon...</p>
            </div>
        );
    }
}
