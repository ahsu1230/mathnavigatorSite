import React from "react";
import Enzyme, { shallow } from "enzyme";
import MenuWide from "./menuWide.js";

describe("test", () => {
    const component = shallow(<MenuWide />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.hasClass("header-menu-wide")).toBe(true);
    });

    test("shows links", () => {
        expect(component.children().length).toBe(5);
        expect(component.find("SubMenu").length).toBe(2);
        expect(component.find("MenuLink").length).toBe(3);
    });
});
