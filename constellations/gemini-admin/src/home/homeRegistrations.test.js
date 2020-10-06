import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomeTabSectionRegistrations } from "./homeRegistrations.js";

describe("Home Registrations", () => {
    const component = shallow(<HomeTabSectionRegistrations />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("Pending Registrations");
    });
});
