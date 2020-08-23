import React from "react";
import { shallow } from "enzyme";
import { ProgramPage } from "./program.js";

describe("Program Page", () => {
    const component = shallow(<ProgramPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Programs");
        expect(component.find("Link").text()).toContain("Add Program");
        expect(component.find("ProgramRow").length).toBe(0);
    });

    test("renders 2 rows", () => {
        const programs = [
            {
                programId: "ap_bc_calculus",
                name: "AP BC Calculus",
                grade1: 10,
                grade2: 12,
            },
            {
                programId: "sat_math",
                name: "SAT Math",
                grade1: 9,
                grade2: 12,
            },
        ];
        component.setState({ programs: programs });
        expect(component.find("ProgramRow").length).toBe(2);

        let row0 = component.find("ProgramRow").at(0);
        expect(row0.prop("program")).toHaveProperty(
            "programId",
            "ap_bc_calculus"
        );
        expect(row0.prop("program")).toHaveProperty("name", "AP BC Calculus");
        expect(row0.prop("program")).toHaveProperty("grade1", 10);
        expect(row0.prop("program")).toHaveProperty("grade2", 12);

        let row1 = component.find("ProgramRow").at(1);
        expect(row1.prop("program")).toHaveProperty("programId", "sat_math");
        expect(row1.prop("program")).toHaveProperty("name", "SAT Math");
        expect(row1.prop("program")).toHaveProperty("grade1", 9);
        expect(row1.prop("program")).toHaveProperty("grade2", 12);
    });
});
