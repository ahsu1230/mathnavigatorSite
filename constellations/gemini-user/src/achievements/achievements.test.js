import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievementPage } from "./achievements.js";

describe("test", () => {
    const component = shallow(<AchievementPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Math Navigator Achievements"
        );
    });
});
