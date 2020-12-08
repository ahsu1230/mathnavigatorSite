import React from "react";
import { shallow } from "enzyme";
import { ProgramPage } from "./program.js";

describe("Program Page", () => {
    const component = shallow(<ProgramPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Programs");
        expect(header.prop("addUrl")).toBe("/programs/add");
    });

    test("renders 2 rows", () => {
        const programs = [
            {
                programId: "ap_bc_calculus",
                title: "AP BC Calculus",
                grade1: 10,
                grade2: 12,
            },
            {
                programId: "sat_math",
                title: "SAT Math",
                grade1: 9,
                grade2: 12,
            },
        ];
        component.setState({ programs: programs });

        let rows = component.find("RowCardBasic");

        let row0 = rows.at(0);
        expect(row0.prop("title")).toBe("AP BC Calculus");
        expect(row0.prop("subtitle")).toBe("ap_bc_calculus");

        let row1 = rows.at(1);
        expect(row1.prop("title")).toBe("SAT Math");
        expect(row1.prop("subtitle")).toBe("sat_math");
        expect(rows.length).toBe(2);
    });
});
