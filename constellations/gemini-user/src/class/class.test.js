import React from "react";
import { shallow } from "enzyme";
import { ClassPage } from "./class.js";

describe("Class Page", () => {
    const component = shallow(<ClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("ClassErrorPage").exists()).toBe(true);
    });

    test("renders class", () => {
        const classObj = {
            programId: "sat_math",
            semesterId: "2020_fall",
            classKey: "class1",
            classId: "sat_math_2020_fall_class1",
            locationId: "Churchill",
            times: "Tue. 5:00pm - 7:00pm, Sat. 1:00pm - 3:00pm",
            fullState: 1,
            pricePerSession: 50,
        };
        const semester = {
            semesterId: "2020_fall",
            title: "Fall 2020",
        };
        const program = {
            programId: "sat_math",
            name: "SAT Math",
            grade1: 9,
            grade2: 12,
            description: "SAT Math Preparation",
        };
        const location = {
            locationId: "Churchill",
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

        expect(component.find("h1").text()).toBe(
            "Programs > SAT Math Fall 2020 class1"
        );
        expect(component.find("Link").at(0).text()).toBe("Programs");
        expect(component.find("h4").at(0).text()).toBe("ALMOST FULL");
        expect(component.find("h4").at(1).text()).toBe("Grades: 9 - 12");
        expect(component.find("p").at(0).text()).toBe("SAT Math Preparation");

        expect(component.find("Link").at(1).text()).toBe("Register");
        expect(component.find("h3").at(0).text()).toBe("Location");
        expect(component.find("p").at(1).text()).toBe("11300 Gainsborough Rd");
        expect(component.find("p").at(2).text()).toBe("Potomac, MD 20854");
        expect(component.find("p").at(3).text()).toBe("");

        expect(component.find("h3").at(1).text()).toBe("Times");
        expect(component.find("p").at(4).text()).toBe("Tue. 5:00pm - 7:00pm");
        expect(component.find("p").at(5).text()).toBe("Sat. 1:00pm - 3:00pm");
        expect(component.find("p").at(6).text()).toBe("No sessions scheduled");

        expect(component.find("h3").at(2).text()).toBe("Pricing");
        expect(component.find("p").at(7).text()).toBe(
            "Price per session: $50.00"
        );

        expect(component.find("Link").at(2).text()).toBe("Contact Us");
        expect(component.find("Link").at(3).text()).toBe("< More Programs");
    });
});
