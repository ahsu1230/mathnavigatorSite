import React from "react";
import Enzyme, { shallow } from "enzyme";
import { LocationPage } from "./location.js";

describe("Location Page", () => {
    const component = shallow(<LocationPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Locations");
        expect(header.prop("addUrl")).toBe("/locations/add");
    });
});
