import React from "react";
import Enzyme, { shallow } from "enzyme";
import { UserInput } from "./userInput.js";

describe("User Input", () => {
    const component = shallow(<UserInput />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("InputText").at(0).prop("label")).toBe(
            "First Name"
        );
        expect(component.find("InputText").at(1).prop("label")).toBe(
            "Middle Name"
        );
        expect(component.find("InputText").at(2).prop("label")).toBe(
            "Last Name"
        );
        expect(component.find("InputText").at(3).prop("label")).toBe("Email");
        expect(component.find("InputText").at(4).prop("label")).toBe("Phone");
        expect(component.find("InputText").at(5).prop("label")).toBe("School");
        expect(component.find("InputText").at(6).prop("label")).toBe(
            "Graduation Year"
        );
        expect(component.find("InputText")).toHaveLength(7);
    });

    test("renders input data", () => {
        component.setProps({
            firstName: "John",
            middleName: "Richard",
            lastName: "Doe",
            email: "johndoe@gmail.com",
            phone: "1234567890",
            isGuardian: true,
            accountId: 1,
            school: "Winston Churchill High School",
            graduationYear: "2020",
        });

        expect(component.find("InputText").at(0).prop("value")).toBe("John");
        expect(component.find("InputText").at(1).prop("value")).toBe("Richard");
        expect(component.find("InputText").at(2).prop("value")).toBe("Doe");
        expect(component.find("InputText").at(3).prop("value")).toBe(
            "johndoe@gmail.com"
        );
        expect(component.find("InputText").at(4).prop("value")).toBe(
            "1234567890"
        );
        expect(component.find("InputText").at(5).prop("value")).toBe(
            "Winston Churchill High School"
        );
        expect(component.find("InputText").at(6).prop("value")).toBe("2020");
        expect(component.find("InputText")).toHaveLength(7);
    });
});