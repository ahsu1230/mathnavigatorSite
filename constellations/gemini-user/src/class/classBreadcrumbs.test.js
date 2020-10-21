import React from "react";
import { shallow } from "enzyme";
import { ClassBreadcrumbs } from "./classBreadcrumbs.js";

describe("Class Breadcrumbs", () => {
    const program = {
        programId: "program1",
        title: "Program1",
        grade1: 1,
        grade2: 2,
        featured: "popular",
    };

    const classObj = {
        programId: "program1",
        semesterId: "2020_fall",
        classKey: "classA",
        locationId: "zoom",
        classId: "program1_2020_fall_classA",
    };

    const semesterObj = {
        semesterId: "2020_fall",
        title: "Fall 2020",
    };

    const component = shallow(
        <ClassBreadcrumbs
            program={program}
            classObj={classObj}
            semester={semesterObj}
        />
    );

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("#breadcrumbs Link").text()).toBe(
            "Program Catalog"
        );
        expect(component.find("Link").text()).toBe("Program Catalog");
        expect(component.find("h1").text()).toBe("Program1 ClassA");
    });

    test("renders featured", () => {
        const featured = component.find(".featured");
        expect(featured.exists()).toBe(true);
        expect(featured.find("p").text()).toContain("most popular programs");
    });
});
