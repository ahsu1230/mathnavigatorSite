"use strict";
require("./programs.sass");
import React from "react";
import API from "../utils/api.js";
import { trackAnalytics } from "../utils/analyticsUtils.js";
import { ProgramCard } from "./programCard.js";

export class ProgramsPage extends React.Component {
    state = {
        semesters: [],
        programClassesMap: {},
        fullStates: [],
    };

    componentDidMount = () => {
        trackAnalytics("programs");
        API.get("api/classesbysemesters?published=true").then((res) => {
            const classesbysemesters = res.data || [];
            let semesters = [];
            let programClassesMap = {};
            classesbysemesters.forEach((element) => {
                semesters.push(element.semester);
                programClassesMap[element.semester.semesterId] =
                    element.programClasses;
            });

            this.setState({
                semesters: semesters,
                programClassesMap: programClassesMap,
            });
        });

        API.get("api/classes/full-states").then((res) => {
            this.setState({
                fullStates: res.data,
            });
        });
    };

    render = () => {
        const semesterSections = this.state.semesters.map((semester, index) => (
            <ProgramSection
                key={index}
                semester={semester}
                programClasses={
                    this.state.programClassesMap[semester.semesterId]
                }
                fullStates={this.state.fullStates}
            />
        ));

        return (
            <div id="view-programs">
                <h1>Program Catalog</h1>
                <p>
                    Math Navigator offers new programs and classes every
                    semester. Some programs will have multiple classes available
                    to accomodate different schedules. Programs are recurring
                    and will usually be offered again in the following semester,
                    so if you miss the enrollment period this time, make sure to
                    enroll next period!
                </p>
                {semesterSections}
            </div>
        );
    };
}

export class ProgramSection extends React.Component {
    render = () => {
        const semester = this.props.semester || {};
        const programClasses = this.props.programClasses || [];

        let programs = [];
        let programClassesMap = {};
        programClasses.forEach((programClass) => {
            if (!!programClass.classes) {
                programs.push(programClass.program);
                programClassesMap[programClass.program.programId] =
                    programClass.classes;
            }
        });

        const cards = programs.map((program, index) => {
            return (
                <ProgramCard
                    key={index}
                    semester={semester}
                    program={program}
                    classes={programClassesMap[program.programId]}
                    fullStates={this.props.fullStates}
                />
            );
        });

        return (
            <div className="section">
                <h2 className="section-title">{semester.title}</h2>
                {cards.length > 0 ? cards : <p>Coming soon...</p>}
            </div>
        );
    };
}
