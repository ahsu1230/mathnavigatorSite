import React from "react";
import Enzyme, { shallow } from "enzyme";
import { InternshipPage } from "./internship.js";

describe("Internship Page", () => {
    const component = shallow(<InternshipPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        expect(component.find("h1").at(0).text()).toContain(
            "Software Development Internship"
        );
        expect(component.find("h1").at(1).text()).toContain("Technology Stack");
        expect(component.find("h1").at(2).text()).toContain(
            "Internship Structure"
        );
        expect(component.find("h1").at(3).text()).toContain("Past Interns:");
        expect(component.find("InternshipCarousel").exists()).toBe(true);
    });
});
