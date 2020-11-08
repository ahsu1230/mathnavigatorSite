import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomePage } from "./home.js";

describe("Home Page", () => {
    const component = shallow(<HomePage />);

    const TAB_CLASSES = "classes";
    const TAB_USERS = "users";
    const TAB_REGISTRATIONS = "registrations";
    const TAB_UNPAID = "unpaid";

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Administrator Dashboard"
        );

        expect(component.find("Tab")).toHaveLength(4);
        expect(component.find("TabPanel")).toHaveLength(4);
    });
});
