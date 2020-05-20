import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievePage } from "./achievements.js";

describe("test", () => {
    const component = shallow(<AchievePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Math Navigator Achievements"
        );
    });
});
