import React from "react";
import { shallow } from "enzyme";
import { SessionPage } from "./session.js";

describe("Sessions Page", () => {
    const component = shallow(<SessionPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Select Class");
        expect(component.find("select").exists()).toBe(true);
    });

    test("renders select classId", () => {
        component.setState({ classId: "ap_bc_calculus_2020_fall_class1" });
        expect(component.find("select").prop("value")).toBe(
            "ap_bc_calculus_2020_fall_class1"
        );
    });
});
