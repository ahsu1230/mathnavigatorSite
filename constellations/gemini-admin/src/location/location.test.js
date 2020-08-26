import React from "react";
import Enzyme, { shallow } from "enzyme";
import { LocationPage } from "./location.js";

describe("Location Page", () => {
    const component = shallow(<LocationPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Locations");
        expect(component.find("Link").text()).toContain("Add Location");
        expect(component.find("LocationRow").length).toBe(0);
    });
});
