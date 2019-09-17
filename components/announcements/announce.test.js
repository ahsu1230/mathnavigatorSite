import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AnnouncePage } from "./announce.js";

describe("test", () => {
  test("renders", () => {
    const component = shallow(<AnnouncePage/>);
    expect(component.exists()).toBe(true);
  });
});
