import React from "react";
import { shallow } from "enzyme";
import { LocationEditPage } from "./locationEdit.js";

describe("Location Edit Page", () => {
    const component = shallow(<LocationEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");

        expect(wrapper.prop("title")).toContain("Add Location");
        expect(wrapper.prop("prevPageUrl")).toBe("locations");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
