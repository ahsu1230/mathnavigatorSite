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

        expect(component.find("TabButton").at(0).prop("section")).toBe(
            "clasess"
        );
        expect(component.find("TabButton").at(1).prop("section")).toBe("users");
        expect(component.find("TabButton").at(2).prop("section")).toBe(
            "registrations"
        );
        expect(component.find("TabButton").at(3).prop("section")).toBe(
            "unpaid"
        );
        expect(component.find("TabButton").length).toBe(4);
    });
});
