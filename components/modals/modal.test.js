import React from "react";
import Enzyme, { shallow } from "enzyme";
import { Modal } from "./modal.js";

describe("test", () => {
  test("renders", () => {
    const component = shallow(<Modal/>);
    expect(component.exists()).toBe(true);
  });
});
