import React from "react";
import Enzyme, { shallow } from "enzyme";
import Footer from "./footer.js";

describe("test", () => {
  const component = shallow(<Footer/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
  });

  test("shows links", () => {
    var ul = component.find("ul");
    expect(ul.exists()).toBe(true);
    expect(ul.children().length).toBe(4);
  });
});
