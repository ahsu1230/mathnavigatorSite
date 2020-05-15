import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomePage } from "./home.js";

describe("test", () => {
    const component = shallow(<HomePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Math Navigator");
    });
});