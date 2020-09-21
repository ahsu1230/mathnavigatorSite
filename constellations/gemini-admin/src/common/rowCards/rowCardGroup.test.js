import React from "react";
import { shallow } from "enzyme";
import RowCardGroup from "./rowCardGroup.js";

describe("RowCardGroup", () => {
    const component = shallow(<RowCardGroup/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
