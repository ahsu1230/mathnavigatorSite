import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HashRouter as Router } from 'react-router-dom';
import { Header} from "./header.js";
import renderer from 'react-test-renderer';

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

  test("snapshot", () => {
    const tree = renderer.create(
      <Router>
        <Header history={fakeHistory}/>
      </Router>
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
