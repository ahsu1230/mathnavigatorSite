import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HashRouter as Router } from "react-router-dom";
import Footer from "./footer.js";
import renderer from "react-test-renderer";

describe("test", () => {
    const component = shallow(<Footer />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("shows links", () => {
        var ul = component.find("ul");
        expect(ul.exists()).toBe(true);
        expect(ul.children().length).toBe(5);
    });

    test("snapshot", () => {
        const tree = renderer
            .create(
                <Router>
                    <Footer />
                </Router>
            )
            .toJSON();
        expect(tree).toMatchSnapshot();
    });
});
