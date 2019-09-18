import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomePage } from "./home.js";

describe("test", () => {
  const component = shallow(<HomePage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.exists("HomeAnnounce"));
    expect(component.exists("HomeBanner"));
    expect(component.exists("HomeSectionPrograms"));
    expect(component.exists("HomeSectionSuccess"));
    expect(component.exists("HomeSectionStories"));
  });
});
