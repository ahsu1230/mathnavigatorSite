import React from "react";
import { shallow } from "enzyme";
import { Modal } from "./modal.js";

describe("Modal", () => {
    const component = shallow(<Modal/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
