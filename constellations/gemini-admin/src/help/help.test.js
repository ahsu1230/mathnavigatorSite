import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HelpPage } from "./help.js";

describe("Help Page", () => {
    const component = shallow(<HelpPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Administrative Support");
        expect(component.find("h1").at(0).text()).toContain(
            "How to Create a Class"
        );
        expect(component.find("h1").at(1).text()).toContain(
            "How to Email Users"
        );
        expect(component.find("h1").at(2).text()).toContain(
            "How to use the Dashboard"
        );
    });
});
