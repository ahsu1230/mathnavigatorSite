import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AccountPage } from "./account.js";

describe("Account Page", () => {
    const component = shallow(<AccountPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Search Accounts");
        expect(component.find("AccountSearch")).toHaveLength(2);
        expect(component.find("Link").text()).toContain("Create New Account");
        expect(component.find("AccountInfo")).toHaveLength(1);
    });
});
