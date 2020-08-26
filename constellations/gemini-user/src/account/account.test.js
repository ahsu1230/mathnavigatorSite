import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AccountPage } from "./account.js";

describe("Account Page", () => {
    const component = shallow(<AccountPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Your Account");
        expect(component.find("div#tab-container")).toHaveLength(1);
        expect(component.find("SettingsTab")).toHaveLength(1);
    });
});
