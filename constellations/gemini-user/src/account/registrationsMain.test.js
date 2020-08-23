import React from "react";
import Enzyme, { shallow } from "enzyme";
import { RegistrationsTabMain } from "./registrations.js";

describe("Main Registrations Tab", () => {
    const props = {
        userClasses: [],
        upcomingAFHs: [],
        locationsById: [],
    };
    const component = shallow(<RegistrationsTabMain {...props} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").at(0).text()).toContain(
            "Currently Enrolled Classes"
        );
        expect(component.find("h2").at(1).text()).toContain(
            "Upcoming Ask For Help Sessions"
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

        expect(component.find("ul")).toHaveLength(2);
        let row0 = component.find("ul").at(0);
        expect(row0.text()).toContain("John Doe");
        expect(row0.text()).toContain("AP Calculus (Fall 2020)");
        expect(row0.text()).toContain("Enrolled on: 12/31/2019");

        let row1 = component.find("ul").at(1);
        expect(row1.text()).toContain("Jane Doe");
        expect(row1.text()).toContain("(No classes registered)");
    });

    test("renders afh sessions", () => {
        component.setProps({
            upcomingAFHs: [
                {
                    date: "2020-01-01T00:00:00Z",
                    locationId: "wchs",
                    subject: "math",
                    timeString: "3:00pm - 4:00pm",
                    title: "AP Calculus Practice Exam",
                },
            ],
            locationsById: {
                wchs: {
                    city: "Potomac",
                    state: "MD",
                    zipcode: "12345",
                    street: "1234 Gains Rd",
                    room: "Room 10",
                },
            },
        });

        expect(component.find("ul")).toHaveLength(4);
        let row0 = component.find("ul").at(3);
        expect(row0.text()).toContain("AP Calculus Practice Exam");
        expect(row0.text()).toContain("December 31st, 2019");
        expect(row0.text()).toContain("3:00pm - 4:00pm");
        expect(row0.text()).toContain("1234 Gains Rd");
        expect(row0.text()).toContain("Potomac, MD");
        expect(row0.text()).toContain("Room 10");
    });
});
