import React from "react";
import Enzyme, { shallow } from "enzyme";
import { EmailPayments } from "./emailPayments.js";

describe("Email Payment Page", () => {
    const component = shallow(<EmailPayments />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Generate Payment Reminder Email");
        expect(component.find("button").text()).toContain("Search");
    });
});
