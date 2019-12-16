import React from "react";
import Enzyme, { shallow } from "enzyme";
import { StudentProjectsPage } from "./studentProjects.js";
import { HashRouter as Router } from 'react-router-dom';
import renderer from 'react-test-renderer';

describe("test", () => {
  const component = shallow(<StudentProjectsPage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Student Website Projects");
  });

  test("snapshot", () => {
    const tree = renderer.create(
      <Router>
        <StudentProjectsPage/>
      </Router>
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
