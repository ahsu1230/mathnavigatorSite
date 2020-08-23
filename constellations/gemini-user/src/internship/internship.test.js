import React from "react";
import Enzyme, { shallow } from "enzyme";
import { InternshipPage } from "./internship.js";

describe("Internship Page", () => {
    const component = shallow(<InternshipPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Software Development Internship"
        );
        expect(component.find("h1").text()).toContain("Technology Stack");
        expect(component.find("h1").text()).toContain("Internship Structure");
        expect(component.find("h4").text()).toContain("Past Interns:");
        expect(component.find("h1").text()).toContain(
            "Math Navigator Products"
        );
    });
});
