import React from "react";
import { shallow } from "enzyme";
import { ClassPage } from "./class.js";

describe("Class Page", () => {
    const component = shallow(<ClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
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
        const sessions = [];
        component.setState({
            classObj: classObj,
            semester: semester,
            program: program,
            location: location,
            sessions: sessions,
        });

        const breadcrumbs = component.find("ClassBreadcrumbs");
        const programInfo = component.find("ClassProgramInfo");
        const register = component.find("ClassRegister");
        const classInfo = component.find("ClassInfo");
        const schedule = component.find("ClassSchedule");

        expect(breadcrumbs.exists()).toBe(true);
        expect(breadcrumbs.prop("program")).toEqual(program);
        expect(breadcrumbs.prop("classObj")).toEqual(classObj);
        expect(breadcrumbs.prop("semester")).toEqual(semester);

        expect(programInfo.exists()).toBe(true);
        expect(programInfo.prop("program")).toEqual(program);

        expect(register.exists()).toBe(true);
        expect(register.prop("classObj")).toEqual(classObj);

        expect(classInfo.exists()).toBe(true);
        expect(classInfo.prop("classObj")).toEqual(classObj);
        expect(classInfo.prop("location")).toEqual(location);
        expect(classInfo.prop("sessions")).toEqual(sessions);

        expect(schedule.exists()).toBe(true);
        expect(schedule.prop("sessions")).toEqual(sessions);
    });
});
