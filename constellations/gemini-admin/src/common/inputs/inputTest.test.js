import React from "react";
import { shallow } from "enzyme";
import { InputText } from "./inputText.js";

describe("InputText", () => {
    const component = shallow(<InputText/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
