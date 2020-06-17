"use strict";
require("./internship.sass");
import React from "react";
import ApCalcImgScr from "./../../assets/ap_calc.jpg";

export class InternshipPage extends React.Component {
    render() {
        return (
            <div id="view-intern">
                <h1>Software Development Internship</h1>
                <p>
                    Math Navigator occassionally offers Software Development
                    Internship for ambitious students who would like to pursuse
                    a college major or a career in software development. In this
                    internship oppotunity, students will learn the fundamentals
                    of modern technologies commonly used across
                    engineering-focused companies. And with these skills,
                    students help create this site that you are currently using!
                </p>
                <h1>Technology Stack</h1>

                <div class="container-main">
                    <div class="child-image">
                        <img class="image" src={ApCalcImgScr} />
                    </div>
                    <div class="child-caption">
                        <p>
                            Math Navigator occassionally offers Software
                            Development Internship for ambitious students who
                            would like to pursuse a college major or a career in
                            software development.
                        </p>
                    </div>
                </div>

                <div class="container-main">
                    <div class="child-image">
                        <img class="image" src={ApCalcImgScr} />
                    </div>
                    <div class="child-caption">
                        <p>
                            Math Navigator occassionally offers Software
                            Development Internship for ambitious students who
                            would like to pursuse a college major or a career in
                            software development.
                        </p>
                    </div>
                </div>

                <div class="container-main">
                    <div class="child-image">
                        <img class="image" src={ApCalcImgScr} />
                    </div>
                    <div class="child-caption">
                        <p>
                            Internship opportunities will be announced when they
                            are available. When they are, student candidates
                            must first pass a series of coding assessments and
                            interviews in order to receive a position. This
                            interview process reflects the same process that
                            famous technology companies like Google, Facebook,
                            Amazon, etc. have. We encourage all students to
                            attempt these assessments to familiarize themselves
                            with this interview process structure if they ever
                            want to pursue a career at any of these companies.
                        </p>
                    </div>
                </div>

                <h1>Internship Structure</h1>
                <p>
                    The internship is split into various roles. Interns can
                    choose to join the front-end team to learn website
                    development with ReactJs or join the back-end team to learn
                    data management with Golang and MySQL. There are a few
                    students who have done internships for both teams!
                </p>

                <table id="past-intern">
                    <caption>Past Interns:</caption> <br></br>
                    <tr>
                        <th>Name</th>
                        <th>School</th>
                    </tr>
                    <tr>
                        <td>Serena Xu</td>
                        <td>Winston Churchill HS 2020</td>
                    </tr>
                    <tr>
                        <td>Daniel Liu</td>
                        <td>Winston Churchill HS 2021</td>
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
                <p>
                    More in-depth information about what it's like doing an
                    internship. <br></br>- Talk about Admin tool (with pictures){" "}
                    <br></br>- Talk about team-organization with Asana?
                </p>
            </div>
        );
    }
}
