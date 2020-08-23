import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomeTabSectionClasses } from "./homeClasses.js";

describe("Home Classes", () => {
    const component = shallow(<HomeTabSectionClasses />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("Unpublished Classes");
        expect(component.find("Link").text()).toContain(
            "View All Classes to Publish"
        );
    });
});
