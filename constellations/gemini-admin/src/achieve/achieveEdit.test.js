import React from "react";
import { shallow } from "enzyme";
import { AchieveEditPage } from "./achieveEdit.js";

describe("test", () => {
    const component = shallow(<AchieveEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");

        expect(wrapper.prop("title")).toContain("Add Achievement");
        expect(wrapper.prop("prevPageUrl")).toBe("achievements");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
