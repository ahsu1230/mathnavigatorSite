import React from "react";
import Enzyme, { shallow } from "enzyme";
import { SemesterPage } from "./semester.js";

describe("Semester Page", () => {
    const component = shallow(<SemesterPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Semesters");
        expect(component.find("Link").text()).toContain("Add Semester");
        expect(component.find("SemesterRow").length).toBe(0);
    });
});
