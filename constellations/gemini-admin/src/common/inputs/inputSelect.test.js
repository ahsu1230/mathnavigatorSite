import React from "react";
import { shallow } from "enzyme";
import { InputSelect } from "./inputSelect.js";

describe("InputSelect", () => {
    const component = shallow(<InputSelect />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("render with basic props", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: 0,
            onChangeCallback: onChangeMocked,
            required: true,
            label: "Some Label",
            description: "Some Description",
            options: [{ value: "option1", displayName: "Option1" }],
        });

        expect(component.find("h2").text()).toBe("Some Label");
        expect(component.find("h4").text()).toContain("Some Description");
        expect(component.find("h4").text()).toContain("(required)");
        expect(component.find("select").exists()).toBe(true);
        expect(component.find("select").prop("onChange")).toBeDefined();
        expect(component.find("option").prop("value")).toBe("option1");
        expect(component.find("option").text()).toBe("Option1");
        expect(component.find("option").length).toBe(1);
    });

    test("render as optional", () => {
        let onChangeMocked = jest.fn();
        component.setProps({
            value: 0,
            onChangeCallback: onChangeMocked,
            required: false,
            label: "Some Label",
            description: "Some Description",
        });

        expect(component.find("h4").text()).toContain("Some Description");
        expect(component.find("h4").text()).toContain("(optional)");
    });

    test("render with empty options", () => {
        component.setProps({
            options: [],
            errorMessageIfEmpty: "No options!",
        });

        expect(component.find("select option").length).toBe(0);
        expect(component.find("h4.select-error").text()).toBe("No options!");
        expect(component.find("img").exists()).toBe(false); // not chosen yet!
    });

    test("render with many options (has default option)", () => {
        component.setProps({
            hasNoDefault: false,
            value: "option2",
            options: [
                { value: "option1", displayName: "Option1" },
                { value: "option2", displayName: "Option2" },
                { value: "option3", displayName: "Option3" },
            ],
        });

        let options = component.find("option");
        expect(options.at(0).prop("value")).toBe("option1");
        expect(options.at(0).text()).toBe("Option1");
        expect(options.at(1).prop("value")).toBe("option2");
        expect(options.at(1).text()).toBe("Option2");
        expect(options.at(2).prop("value")).toBe("option3");
        expect(options.at(2).text()).toBe("Option3");
        expect(component.find("select option").length).toBe(3);
        expect(component.find("select").prop("value")).toBe("option2");
    });

    test("render with many options (does NOT have default option)", () => {
        component.setProps({
            hasNoDefault: true,
            options: [
                { value: "option1", displayName: "Option1" },
                { value: "option2", displayName: "Option2" },
                { value: "option3", displayName: "Option3" },
            ],
        });

        let options = component.find("option");
        let defaultOption = options.at(0);
        expect(defaultOption.prop("value")).toBe("");
        expect(defaultOption.text()).toContain("Select an option");

        expect(options.at(1).prop("value")).toBe("option1");
        expect(options.at(1).text()).toBe("Option1");
        expect(options.at(2).prop("value")).toBe("option2");
        expect(options.at(2).text()).toBe("Option2");
        expect(options.at(3).prop("value")).toBe("option3");
        expect(options.at(3).text()).toBe("Option3");
        expect(component.find("select option").length).toBe(4);
    });
});
