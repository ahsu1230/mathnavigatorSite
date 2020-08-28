import React from "react";
import { shallow } from "enzyme";
import { ClassErrorPage } from "./classError.js";

describe("Class Error Page", () => {
    const component = shallow(<ClassErrorPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toBe("Page Not Found");
        expect(component.find("Link").text()).toBe("View our Programs");
    });

    test("renders classId", () => {
        component.setProps({ classId: "invalid_id" });
        expect(component.find("span").text()).toBe(
            "Class ID 'invalid_id' does not exist."
        );
    });
});
