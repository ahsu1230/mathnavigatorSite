import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomeTabSectionRegistrations } from "./homeRegistrations.js";

describe("Home Registrations", () => {
    const component = shallow(<HomeTabSectionRegistrations />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").at(0).text()).toContain(
            "Pending Registrations For Classes"
        );
        expect(component.find("Link").at(0).text()).toContain("View By Class");
        expect(component.find("h3").at(1).text()).toContain(
            "New Registrations For AFH"
        );
        expect(component.find("Link").at(1).text()).toContain("View By AFH Session");
    });
});
