"use strict";
require("./programs.sass");
import React from "react";
import API from "../utils/api.js";
import { ProgramCard } from "./programCard.js";

export class ProgramsPage extends React.Component {
    state = {
        semesters: [],
        programClassesMap: {},
    };

    componentDidMount = () => {
        // TODO: Change to below line when published_at is removed from all tables except classes
        // API.get("api/classesbysemesters?published=true").then((res) => {
        API.get("api/classesbysemesters").then((res) => {
            const classesbysemesters = res.data || [];
            let semesters = [];
            let programClassesMap = {};
            classesbysemesters.forEach((element) => {
                semesters = this.sortedInsert(semesters, element.semester);
                programClassesMap[element.semester.semesterId] =
                    element.programClasses;
            });

            this.setState({
                semesters: semesters,
                programClassesMap: programClassesMap,
            });
        });
    };

    sortedInsert = (array, value) => {
        let low = 0;
        let high = array.length;

        while (low < high) {
            let mid = (low + high) >>> 1;
            if (this.compareSemesters(array[mid].semesterId, value.semesterId))
                low = mid + 1;
            else high = mid;
        }
        array.splice(low, 0, value);
        return array;
    };

    compareSemesters = (semester1, semester2) => {
        const seasonMap = {
            spring: 0,
            summer: 1,
            fall: 2,
            winter: 3,
        };
        let [year1, season1] = semester1.split("_");
        let [year2, season2] = semester2.split("_");
        season1 = seasonMap[season1];
        season2 = seasonMap[season2];

        if (year1 == year2 && season1 < season2) return true;
        return year1 < year2;
    };

    render = () => {
        const semesterSections = this.state.semesters.map((semester, index) => (
            <ProgramSection
                key={index}
                semester={semester}
                programClasses={
                    this.state.programClassesMap[semester.semesterId]
                }
            />
        ));

        return (
            <div id="view-programs">
                <div id="star-legend">
                    <div className="star-container">
                        <div className="star-img"></div>
                    </div>
                    = Featured Programs
                </div>
                {semesterSections}
            </div>
        );
    };
}

class ProgramSection extends React.Component {
    render = () => {
        const semester = this.props.semester;
        const programClasses = this.props.programClasses || [];

        let programs = [];
        let programClassesMap = {};
        programClasses.forEach((programClass) => {
            // TODO: remove if statement once backend fixes the bug
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
            />
        ));

        return (
            <div className="section">
                <h1 className="section-title">{semester.title}</h1>
                {cards}
            </div>
        );
    };
}
