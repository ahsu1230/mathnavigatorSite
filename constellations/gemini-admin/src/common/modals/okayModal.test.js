import React from "react";
import { shallow } from "enzyme";
import OkayModal from "./okayModal.js";

describe("OkayModal", () => {
    const component = shallow(<OkayModal/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
