import React from "react";
import { shallow } from "enzyme";
import { SemesterEditPage } from "./semesterEdit.js";

describe("Semester Edit Page", () => {
    const component = shallow(<SemesterEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");
        expect(wrapper.prop("title")).toContain("Add Semester");
        expect(wrapper.prop("prevPageUrl")).toBe("semesters");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
