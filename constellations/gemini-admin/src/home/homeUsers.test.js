import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomeTabSectionUsers } from "./homeUsers.js";

describe("Home Users", () => {
    const component = shallow(<HomeTabSectionUsers />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("New Users");
        expect(component.find("Link").text()).toContain("View All Users");
    });
});
