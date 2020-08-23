import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AFHPage } from "./afh.js";

describe("AFH Page", () => {
    const component = shallow(<AFHPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Ask for Help");
        expect(component.find("h1").text()).toContain("Ask for Help Sessions by Subject");
        
        expect(component.find("TabButton").at(0).prop("subject")).toBe("math");
        expect(component.find("TabButton").at(1).prop("subject")).toBe("english");
        expect(component.find("TabButton").at(2).prop("subject")).toBe("programming");

    });
});