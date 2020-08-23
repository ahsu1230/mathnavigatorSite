import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HomeTabSectionAccounts } from "./homeAccounts.js";

describe("Home Accounts", () => {
    const component = shallow(<HomeTabSectionAccounts />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("Unpaid Accounts");
        expect(component.find("Link").text()).toContain("View All Accounts");
    });
});
