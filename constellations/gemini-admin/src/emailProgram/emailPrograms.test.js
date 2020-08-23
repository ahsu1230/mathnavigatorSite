import React from "react";
import Enzyme, { shallow } from "enzyme";
import { EmailProgramsPage } from "./emailPrograms.js";

describe("Email Programs Page", () => {
    const component = shallow(<EmailProgramsPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Generate Email to Program"
        );
        expect(component.find("h2").text()).toContain("Select A Program");
    });
});
