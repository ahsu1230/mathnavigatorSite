import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AFHPage } from "./afh.js";
import { AfhForm } from './afhForm.js';

describe("test", () => {
  const component = shallow(<AFHPage/>);
  const formComponent = shallow(<AfhForm/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Ask For Help");
    expect(component.exists("AfhForm")).toBe(true);
  });

  test("renders form", () => {
    expect(formComponent.exists()).toBe(true);
    expect(formComponent.find("AfhSession").length).toBe(0);
    expect(formComponent.exists(".submit-container button")).toBe(true);
  });

  test("renders sessions", () => {
    var afh = [ { date: "8/1/2019" }, { date: "8/2/2019" }];
    formComponent.setState({ afh: afh});

    var sessions = formComponent.find("AfhSession");
    expect(sessions.length).toBe(2);
    expect(sessions.at(0).prop('afh').date).toBe("8/1/2019");
    expect(sessions.at(1).prop('afh').date).toBe("8/2/2019");
  });
});
