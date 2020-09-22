import React from "react";
import { shallow } from "enzyme";
import { ClassPage } from "./class.js";

describe("Classes Page", () => {
    const component = shallow(<ClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Classes");
        expect(header.prop("addUrl")).toBe("/classes/add");
    });

    test("renders 2 rows", () => {
        const classes = [
            {
                id: 1,
                programId: "ap_bc_calculus",
                semesterId: "2020_fall",
                classKey: "class1",
                classId: "ap_bc_calculus_2020_fall_class1",
                locationId: "churchill",
            },
            {
                id: 2,
                programId: "sat_math",
                semesterId: "2020_fall",
                classKey: "class2",
                classId: "sat_math_2020_fall_class2",
                locationId: "churchill",
            },
        ];
        component.setState({ classes: classes });

        let rows = component.find("RowCardColumns");
        let row0 = rows.at(0);
        expect(row0.prop("title")).toBe("ClassId");
        expect(row0.prop("subtitle")).toBe("ap_bc_calculus_2020_fall_class1");
        expect(row0.prop("editUrl")).toBe(
            "/classes/ap_bc_calculus_2020_fall_class1/edit"
        );

        let row1 = rows.at(1);
        expect(row1.prop("title")).toBe("ClassId");
        expect(row1.prop("subtitle")).toBe("sat_math_2020_fall_class2");
        expect(row1.prop("editUrl")).toBe(
            "/classes/sat_math_2020_fall_class2/edit"
        );

        expect(rows.length).toBe(2);
    });
});
