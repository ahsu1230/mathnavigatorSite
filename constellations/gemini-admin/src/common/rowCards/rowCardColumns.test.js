import React from "react";
import { shallow } from "enzyme";
import RowCardColumns from "./rowCardColumns.js";

describe("RowCardColumns", () => {
    const component = shallow(<RowCardColumns/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
