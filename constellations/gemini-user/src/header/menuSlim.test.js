import React from "react";
import Enzyme, { shallow } from "enzyme";
import { MenuSlim, OverlayMenu } from "./menuSlim.js";

describe("test", () => {
    const component = shallow(<MenuSlim />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.childAt(0).hasClass("header-menu-slim")).toBe(true);
    });

    test("toggles menu", () => {
        var button = component.find(".header-menu-btn");
        expect(button.exists()).toBe(true);

        expect(component.state("show")).toBe(false);
        button.simulate("click");
        expect(component.state("show")).toBe(true);
        button.simulate("click");
        expect(component.state("show")).toBe(false);
    });

    test("shows links", () => {
        var overlayMenu = shallow(<OverlayMenu />);
        expect(overlayMenu.find("LinkRow").length).toBe(3);
        var submenus = overlayMenu.find("SubMenu");
        expect(submenus.length).toBe(2);
    });
});
