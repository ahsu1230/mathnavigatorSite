import React from "react";
import { shallow } from "enzyme";
import EditPageWrapper from "./editPageWrapper.js";

describe("EditPageWrapper", () => {
    const component = shallow(<EditPageWrapper/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
