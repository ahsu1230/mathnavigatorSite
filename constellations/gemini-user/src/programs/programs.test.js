import React from "react";
import { shallow } from "enzyme";
import { ProgramsPage, ProgramSection } from "./programs.js";

describe("Programs Page", () => {
    const component = shallow(<ProgramsPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("#star-legend").text()).toContain(
            "Featured Programs"
        );
        expect(component.find("ProgramSection").length).toBe(0);
    });

    test("renders 2 semesters", () => {
        const semesters = [
            {
                semesterId: "2020_fall",
                title: "Fall 2020",
            },
            {
                semesterId: "2020_winter",
                title: "Winter 2020",
            },
        ];
        component.setState({ semesters: semesters });
        expect(component.find("ProgramSection").length).toBe(2);
    });

    test("renders 2 programs with 1 and 2 classes", () => {
        const semesters = [
            {
                semesterId: "2020_fall",
                title: "Fall 2020",
            },
        ];
        const programClassesMap = {
            "2020_fall": [
                {
                    program: {
                        programId: "ap_bc_calculus",
                        name: "AP BC Calculus",
                        grade1: 10,
                        grade2: 12,
                        description: "Preparation for AP exam",
                    },
                    classes: [{ classId: "ap_bc_calculus_2020_fall_class1" }],
                },
                {
                    program: {
                        programId: "sat_math",
                        name: "SAT Math",
                        grade1: 9,
                        grade2: 12,
                        description: "SAT math preparation",
                    },
                    classes: [
                        { classId: "sat_math_2020_fall_class1" },
                        { classId: "sat_math_2020_fall_class2" },
                    ],
                },
            ],
        };
        component.setState({
            semesters: semesters,
            programClassesMap: programClassesMap,
        });
        expect(component.find("ProgramSection").length).toBe(1);

        const programSection = shallow(<ProgramSection />);
        programSection.setProps({
            semester: semesters[0],
            programClasses: programClassesMap["2020_fall"],
        });
        expect(programSection.find("ProgramCard").length).toBe(2);
    });
});
