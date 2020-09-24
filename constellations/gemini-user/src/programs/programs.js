"use strict";
require("./programs.sass");
import React from "react";
import API from "../utils/api.js";
import { sortedSemesterInsert } from "../utils/semesterUtils.js";
import { ProgramCard } from "./programCard.js";

export class ProgramsPage extends React.Component {
    state = {
        semesters: [],
        programClassesMap: {},
        fullStates: [],
    };

    componentDidMount = () => {
        API.get("api/classesbysemesters?published=true").then((res) => {
            const classesbysemesters = res.data || [];
            let semesters = [];
            let programClassesMap = {};
            classesbysemesters.forEach((element) => {
                semesters = sortedSemesterInsert(semesters, element.semester);
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

        return <div id="view-programs">{semesterSections}</div>;
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

        const cards = programs.map((program, index) => (
            <ProgramCard
                key={index}
                semester={semester}
                program={program}
                classes={programClassesMap[program.programId]}
                fullStates={this.props.fullStates}
            />
        ));

        return (
            <div className="section">
                <h1 className="section-title">{semester.title}</h1>
                {cards.length > 0 ? cards : <p>Coming soon...</p>}
            </div>
        );
    };
}
