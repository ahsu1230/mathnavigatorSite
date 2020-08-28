import React from "react";
import { shallow } from "enzyme";
import { ClassPage } from "./class.js";

describe("Classes Page", () => {
    const component = shallow(<ClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Classes");
        expect(component.find("Link").text()).toContain("Add Class");
        expect(component.find("ClassRow").length).toBe(0);
    });

    test("renders 2 rows", () => {
        const classes = [
            {
                programId: "ap_bc_calculus",
                semesterId: "2020_fall",
                classKey: "class1",
                classId: "ap_bc_calculus_2020_fall_class1",
                locationId: "Churchill",
                times: "Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm",
                fullState: 0,
                pricePerSession: 50,
            },
            {
                programId: "ap_bc_calculus",
                semesterId: "2020_fall",
                classKey: "class2",
                classId: "ap_bc_calculus_2020_fall_class2",
                locationId: "Churchill",
                times: "Thur. 3:30pm - 5:30pm, Sat. 1:00pm - 3:00pm",
                fullState: 2,
                priceLump: 1000,
            },
        ];
        component.setState({ classes: classes });
        expect(component.find("ClassRow").length).toBe(2);

        let row0 = component.find("ClassRow").at(0);
        expect(row0.prop("classObj")).toHaveProperty(
            "classId",
            "ap_bc_calculus_2020_fall_class1"
        );
        expect(row0.prop("classObj")).toHaveProperty("locationId", "Churchill");
        expect(row0.prop("classObj")).toHaveProperty(
            "times",
            "Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm"
        );
        expect(row0.prop("classObj")).toHaveProperty("fullState", 0);
        expect(row0.prop("classObj")).toHaveProperty("pricePerSession", 50);

        let row1 = component.find("ClassRow").at(1);
        expect(row1.prop("classObj")).toHaveProperty(
            "classId",
            "ap_bc_calculus_2020_fall_class2"
        );
        expect(row1.prop("classObj")).toHaveProperty("locationId", "Churchill");
        expect(row1.prop("classObj")).toHaveProperty(
            "times",
            "Thur. 3:30pm - 5:30pm, Sat. 1:00pm - 3:00pm"
        );
        expect(row1.prop("classObj")).toHaveProperty("fullState", 2);
        expect(row1.prop("classObj")).toHaveProperty("priceLump", 1000);
    });
});
