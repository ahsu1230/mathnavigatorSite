import React from "react";
import { shallow } from "enzyme";
import { AskForHelpEditPage } from "./afhEdit.js";

describe("AskForHelp Edit Page", () => {
    const component = shallow(<AskForHelpEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");

        expect(wrapper.prop("title")).toContain("Add AskForHelp Session");
        expect(wrapper.prop("prevPageUrl")).toBe("afh");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
