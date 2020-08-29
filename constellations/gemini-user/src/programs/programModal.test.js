import React from "react";
import { shallow } from "enzyme";
import { ProgramModal } from "./programModal.js";

describe("Program Modal", () => {
    const component = shallow(<ProgramModal />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Classes for ");
    });

    test("renders 2 classes", () => {
        const classes = [
            {
                programId: "sat_math",
                semesterId: "2020_fall",
                classKey: "class1",
                classId: "sat_math_2020_fall_class1",
                locationId: "Churchill",
                times: "Tue. 5:00pm - 7:00pm, Sat. 1:00pm - 3:00pm",
                fullState: 1,
                pricePerSession: 50,
            },
            {
                programId: "sat_math",
                semesterId: "2020_fall",
                classKey: "class2",
                classId: "sat_math_2020_fall_class2",
                locationId: "Churchill",
                times: "Fri. 4:30pm - 6:30pm, Sun. 3:00pm - 5:00pm",
                fullState: 2,
                priceLump: 1000,
            },
        ];
        const semester = {
            semesterId: "2020_fall",
            title: "Fall 2020",
        };
        const program = {
            programId: "sat_math",
            name: "SAT Math",
            grade1: 9,
            grade2: 12,
            description: "SAT Math Preparation",
        };
        component.setProps({
            semester: semester,
            program: program,
            classes: classes,
        });
        expect(component.find("ProgramClass").length).toBe(2);

        let row0 = component.find("ProgramClass").at(0);
        expect(row0.prop("classObj")).toHaveProperty("classKey", "class1");
        expect(row0.prop("classObj")).toHaveProperty(
            "classId",
            "sat_math_2020_fall_class1"
        );
        expect(row0.prop("classObj")).toHaveProperty("locationId", "Churchill");
        expect(row0.prop("classObj")).toHaveProperty(
            "times",
            "Tue. 5:00pm - 7:00pm, Sat. 1:00pm - 3:00pm"
        );
        expect(row0.prop("classObj")).toHaveProperty("fullState", 1);
        expect(row0.prop("classObj")).toHaveProperty("pricePerSession", 50);

        let row1 = component.find("ProgramClass").at(1);
        expect(row1.prop("classObj")).toHaveProperty("classKey", "class2");
        expect(row1.prop("classObj")).toHaveProperty(
            "classId",
            "sat_math_2020_fall_class2"
        );
        expect(row1.prop("classObj")).toHaveProperty("locationId", "Churchill");
        expect(row1.prop("classObj")).toHaveProperty(
            "times",
            "Fri. 4:30pm - 6:30pm, Sun. 3:00pm - 5:00pm"
        );
        expect(row1.prop("classObj")).toHaveProperty("fullState", 2);
        expect(row1.prop("classObj")).toHaveProperty("priceLump", 1000);
    });
});
