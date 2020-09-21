import React from "react";
import { shallow } from "enzyme";
import { InputRadio } from "./inputRadio.js";

describe("InputRadio", () => {
    const component = shallow(<InputRadio/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
