import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomePage } from "./home.js";

describe("Home Page", () => {
    const component = shallow(<HomePage />);

    const TAB_CLASSES = "classes";
    const TAB_USERS = "users";
    const TAB_REGISTRATIONS = "registrations";
    const TAB_UNPAID = "unpaid";

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Administrator Dashboard"
        );

        expect(component.find("TabButton").at(0).prop("section")).toBe(
            TAB_CLASSES
        );
        expect(component.find("TabButton").at(1).prop("section")).toBe(
            TAB_USERS
        );
        expect(component.find("TabButton").at(2).prop("section")).toBe(
            TAB_REGISTRATIONS
        );
        expect(component.find("TabButton").at(3).prop("section")).toBe(
            TAB_UNPAID
        );
        expect(component.find("TabButton").length).toBe(4);
    });
});
