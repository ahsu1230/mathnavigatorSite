import React from "react";
import { shallow } from "enzyme";
import { ClassEditPage } from "./classEdit.js";

describe("Class Edit Page", () => {
    const component = shallow(<ClassEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");

        expect(wrapper.prop("title")).toContain("Add Class");
        expect(wrapper.prop("prevPageUrl")).toBe("classes");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
