import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AskForHelpPage } from "./afh.js";

describe("Ask For Help Page", () => {
    const component = shallow(<AskForHelpPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Ask For Help Sessions");
        expect(component.find("li").at(0).text()).toContain("Date");
        expect(component.find("li").at(1).text()).toContain("Time");
        expect(component.find("li").at(2).text()).toContain("Subject");
        expect(component.find("li").at(3).text()).toContain("Title");
        expect(component.find("li").at(4).text()).toContain("LocationId");
        expect(component.find("li").at(5).text()).toContain("Notes");
    });
});
