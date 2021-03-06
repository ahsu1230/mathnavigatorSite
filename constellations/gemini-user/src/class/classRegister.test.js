import React from "react";
import { shallow } from "enzyme";
import { ClassRegister } from "./classRegister.js";

describe("Class Register", () => {
    const classObj = {
        programId: "program1",
        semesterId: "2020_fall",
        classKey: "classA",
        locationId: "zoom",
        classId: "program1_2020_fall_classA",
        fullState: 0, // not full
    };
    const component = shallow(<ClassRegister classObj={classObj} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("button").text()).toEqual("Enroll");
    });

    test("renders fullState", () => {
        classObj.fullState = 0;
        component.setProps({ classObj: classObj });
        expect(component.find("h4").text()).toContain("please click on Enroll");

        classObj.fullState = 1;
        component.setProps({ classObj: classObj });
        expect(component.find("h4").text()).toContain("almost full");

        classObj.fullState = 2;
        component.setProps({ classObj: classObj });
        expect(component.find("h4").text()).toContain("class is full");
    });
});
