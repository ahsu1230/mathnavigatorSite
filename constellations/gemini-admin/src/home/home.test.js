import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomePage } from "./home.js";

describe("Home Page", () => {
    const component = shallow(<HomePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Administrator Dashboard"
        );

        expect(component.find("Tab")).toHaveLength(5);
        expect(component.find("TabPanel")).toHaveLength(5);
    });
});
