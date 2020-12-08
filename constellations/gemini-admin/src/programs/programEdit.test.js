import React from "react";
import { shallow } from "enzyme";
import { ProgramEditPage } from "./programEdit.js";

describe("Program Edit Page", () => {
    const component = shallow(<ProgramEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");
        expect(wrapper.prop("title")).toContain("Add Program");
        expect(wrapper.prop("prevPageUrl")).toBe("programs");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
