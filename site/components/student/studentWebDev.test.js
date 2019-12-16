import React from "react";
import Enzyme, { shallow } from "enzyme";
import { StudentWebDevPage } from "./studentWebDev.js";
import { HashRouter as Router } from 'react-router-dom';
import renderer from 'react-test-renderer';

describe("test", () => {
  const component = shallow(<StudentWebDevPage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h2").text()).toBe("Website Design & Development");
  });

  test("snapshot", () => {
    const tree = renderer.create(
      <Router>
        <StudentWebDevPage/>
      </Router>
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
