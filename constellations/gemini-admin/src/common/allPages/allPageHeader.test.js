import React from "react";
import { shallow } from "enzyme";
import AllPageHeader from "./allPageHeader.js";

describe("All Page Header", () => {
    const component = shallow(<AllPageHeader/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
