import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievementPage } from "./achievements.js";

describe("test", () => {
  test("renders", () => {
    const component = shallow(<AchievementPage/>);
    expect(component.exists()).toBe(true);
  });
});
