import React from "react";
import Enzyme, { shallow } from "enzyme";
import { ErrorPage } from "./error.js";
import { HashRouter as Router } from 'react-router-dom';
import renderer from 'react-test-renderer';

describe("test", () => {
  const component = shallow(<ErrorPage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Page Not Found");
  });

  test("class not found", () => {
    component.setProps({ classDNE: 'asdf'});
    expect(component.find("h6").text()).toBe("ClassKey 'asdf' does not exist.");
  });

  test("snapshot", () => {
    const tree = renderer.create(
      <Router>
        <ErrorPage/>
      </Router>
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
