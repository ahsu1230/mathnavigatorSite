import React from "react";
import { shallow } from "enzyme";
import { emptyValidator, InputText } from "./inputText.js";

describe("InputText", () => {
    const component = shallow(<InputText/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("render with basic props", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: "asdf",
            onChangeCallback: onChangeMocked,
            required: true,
            label: "Some Label",
            description: "Some Description",
            isTextBox: false,
        });

        expect(component.find("h2").text()).toBe("Some Label");
        expect(component.find("h4.description").text()).toContain("Some Description");
        expect(component.find("h4.description").text()).toContain("(required)");
        expect(component.find("input").exists()).toBe(true);
        expect(component.find("input").prop("value")).toBe("asdf");
        expect(component.find("textarea").exists()).toBe(false);
    });

    test("render as optional", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: 0,
            onChangeCallback: onChangeMocked,
            required: false,
            label: "Some Label",
            description: "Some Description"
        });

        expect(component.find("h4.description").text()).toContain("Some Description");
        expect(component.find("h4.description").text()).toContain("(optional)");
    });

    test("render as textarea", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: "asdf",
            onChangeCallback: onChangeMocked,
            required: true,
            label: "Some Label",
            description: "Some Description",
            isTextBox: true,
        });

        expect(component.find("input").exists()).toBe(false);
        expect(component.find("textarea").exists()).toBe(true);
        expect(component.find("textarea").prop("value")).toBe("asdf");
    });

    test("render with empty validator", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: "",
            onChangeCallback: onChangeMocked,
            required: true,
            validators: [ emptyValidator("some value") ]
        });

        expect(component.find("h4.red").text()).toContain("You must input a some value");
    });

    test("render with custom validator", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: "11",
            onChangeCallback: onChangeMocked,
            required: true,
            validators: [
                {
                    validate: (x) => parseInt(x) < 10,
                    message: "Number must be less than 10"
                }
            ]
        });
        expect(component.find("h4.red").text()).toContain("Number must be less than 10");

        // Change value to a valid one
        component.setProps({
            value: "5"
        });
        expect(component.find("h4.red").exists()).toBe(false);
        expect(component.find("img").exists()).toBe(true); // checkmark
    });
});
