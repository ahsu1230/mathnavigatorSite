import React from "react";
import { shallow } from "enzyme";
import { ClassInfo } from "./classInfo.js";

describe("Class ClassInfo", () => {
    const classObj = {
        programId: "program1",
        semesterId: "2020_fall",
        classKey: "classA",
        locationId: "zoom",
        classId: "program1_2020_fall_classA",
        timesStr: "Friday 1:00pm - 3:00pm, Saturday 3:00pm - 5:00pm",
        pricePerSession: 42,
    };

    const locationObj = {
        locationId: "wchs",
        title: "Winston Churchill H.S.",
        isOnline: false,
    };

    const sessions = [];

    const component = shallow(
        <ClassInfo
            classObj={classObj}
            location={locationObj}
            sessions={sessions}
        />
    );

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find(".block.location h3").text()).toEqual("Location");
        expect(component.find(".block.class-times h3").text()).toEqual("Times");
        expect(component.find(".block.class-price h3").text()).toEqual(
            "Pricing"
        );
    });

    test("renders location", () => {
        expect(component.find(".block.location .line.title").text()).toEqual(
            "Winston Churchill H.S."
        );
    });

    test("renders times", () => {
        const times = component.find(".block.class-times .line");
        expect(times.at(0).text()).toEqual("Friday 1:00pm - 3:00pm");
        expect(times.at(1).text()).toEqual("Saturday 3:00pm - 5:00pm");
    });

    test("renders price", () => {
        expect(component.find(".block.class-price p").at(0).text()).toContain(
            "Price per session: "
        );
    });
});
