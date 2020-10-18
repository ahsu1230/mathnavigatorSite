import React from "react";
import { shallow } from "enzyme";
import { ClassPage } from "./class.js";

describe("Class Page", () => {
    const component = shallow(<ClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("#breadcrumbs Link").text()).toBe(
            "Program Catalog"
        );
        expect(component.find("#class-footer Link").text()).toBe(
            "< More Programs"
        );
    });

    test("renders class", () => {
        const classObj = {
            programId: "sat_math",
            semesterId: "2020_fall",
            classKey: "class1",
            classId: "sat_math_2020_fall_class1",
            locationId: "churchill",
            timesStr: "Tue. 5:00pm - 7:00pm, Sat. 1:00pm - 3:00pm",
            fullState: 1,
            pricePerSession: 50,
        };
        const semester = {
            semesterId: "2020_fall",
            title: "Fall 2020",
        };
        const program = {
            programId: "sat_math",
            title: "SAT Math",
            grade1: 9,
            grade2: 12,
            description: "SAT Math Preparation",
        };
        const location = {
            locationId: "churchill",
            title: "Winston Churchill High School",
            street: "11300 Gainsborough Rd",
            city: "Potomac",
            state: "MD",
            zipcode: "20854",
        };
        component.setState({
            classObj: classObj,
            semester: semester,
            program: program,
            location: location,
        });

        expect(component.find("h1").text()).toBe("SAT Math Class1");

        expect(component.find("Link").at(1).text()).toBe("Enroll");

        expect(component.find("h3").at(0).text()).toBe("Location");
        expect(component.find("h3").at(1).text()).toBe("Times");
        expect(component.find("h3").at(2).text()).toBe("Pricing");

        expect(component.find("#program-info h4").text()).toBe(
            "This class is almost full! Enroll now to reserve your spot."
        );
        expect(component.find("#program-info p.grades").text()).toBe(
            "Grades: 9 - 12"
        );
        expect(component.find("#program-info p.description").text()).toBe(
            "SAT Math Preparation"
        );

        expect(component.find("#class-location .line").at(0).text()).toBe(
            "Winston Churchill High School"
        );
        expect(component.find("#class-location .line").at(1).text()).toBe(
            "11300 Gainsborough Rd"
        );
        expect(component.find("#class-location .line").at(2).text()).toBe(
            "Potomac, MD 20854"
        );

        expect(component.find("#class-times .line").text()).toBe(
            "Tue. 5:00pm - 7:00pm & Sat. 1:00pm - 3:00pm"
        );
        expect(component.find("#class-pricing p").at(0).text()).toBe(
            "Price per session: $50.00"
        );
    });
});
