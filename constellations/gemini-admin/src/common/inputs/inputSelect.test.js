import React from "react";
import { shallow } from "enzyme";
import { InputSelect } from "./inputSelect.js";

describe("InputSelect", () => {
    const component = shallow(<InputSelect/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
