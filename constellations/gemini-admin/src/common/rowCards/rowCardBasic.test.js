import React from "react";
import { shallow } from "enzyme";
import RowCardBasic from "./rowCardBasic.js";

describe("RowCardBasic", () => {
    const component = shallow(<RowCardBasic/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
