import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { SessionAdd } from "./sessionAdd.js";

describe("Session Add Page", () => {
    const component = shallow(<SessionAdd />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("Add Sessions");
        expect(component.find("h4").at(0).text()).toBe("Choose a Day");
        expect(component.find("h4").at(1).text()).toBe("Repeat Every Week");
        expect(component.find("h4").at(2).text()).toBe("Notes");
        expect(component.find("h4").at(3).text()).toBe("Start Time");
        expect(component.find("h4").at(4).text()).toBe("End Time");
        expect(component.find("button").text()).toContain("Add Sessions");
        expect(component.find("SessionList").find("div").length).toBe(0);
    });
});
