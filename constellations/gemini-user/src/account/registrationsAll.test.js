import React from "react";
import Enzyme, { shallow } from "enzyme";
import { RegistrationsTabAllClasses } from "./registrations.js";

describe("Main Registrations Tab", () => {
    const props = { userClasses: [] };
    const component = shallow(<RegistrationsTabAllClasses {...props} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").at(0).text()).toContain(
            "All Enrolled Classes"
        );
        expect(component.find("a")).toHaveLength(1);
    });

    test("renders classes", () => {
        component.setProps({
            userClasses: [
                {
                    name: "John Doe",
                    classes: [
                        {
                            enrollDate: "2020-01-01T00:00:00Z",
                            program: {
                                name: "AP Calculus",
                            },
                            semester: {
                                semesterId: "2020_fall",
                                title: "Fall 2020",
                            },
                        },
                    ],
                },
                {
                    name: "Jane Doe",
                    classes: [],
                },
            ],
        });

        expect(component.find("ul")).toHaveLength(1);
        let row0 = component.find("ul").at(0);
        expect(row0.text()).toContain("John Doe");
        expect(row0.text()).toContain("AP Calculus (Fall 2020)");
        expect(row0.text()).toContain("Enrolled on: 12/31/2019");
    });
});
