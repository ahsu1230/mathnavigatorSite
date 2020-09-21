import React from "react";
import { shallow } from "enzyme";
import YesNoModal from "./YesNoModal.js";

describe("YesNoModal", () => {
    const component = shallow(<YesNoModal/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });
});
