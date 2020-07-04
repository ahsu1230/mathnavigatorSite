import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HashRouter as Router } from "react-router-dom";
import { HomePage } from "./home.js";
import renderer from "react-test-renderer";

describe("test", () => {
    const component = shallow(<HomePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.exists("HomeAnnounce"));
        expect(component.exists("HomeBanner"));
        expect(component.exists("HomeSectionPrograms"));
        expect(component.exists("HomeSectionSuccess"));
        expect(component.exists("HomeSectionStories"));
    });

    test("snapshot", () => {
        const tree = renderer
            .create(
                <Router>
                    <HomePage />
                </Router>
            )
            .toJSON();
        //expect(tree).toMatchSnapshot();
    });
});
