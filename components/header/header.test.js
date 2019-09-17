import React from "react";
import Enzyme, { shallow } from "enzyme";
import { Header} from "./header.js";

describe("test", () => {
  var fakeHistory = {
    location: '/home',
    listen: jest.fn()
  };
  const component = shallow(<Header history={fakeHistory}/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.exists("#view-header")).toBe(true);

    expect(component.state().location).toBe("/home");
  });
});
