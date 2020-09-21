import React from "react";
import { shallow } from "enzyme";
import { InputRadio } from "./inputRadio.js";

import ImgCheckmark from "../../../assets/checkmark_green.svg";

describe("InputRadio", () => {
    const component = shallow(<InputRadio/>);

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
            options: [
                { value: "option1", displayName: "Option1" }
            ]
        });

        expect(component.find("h2").text()).toBe("Some Label");
        expect(component.find("h4").text()).toContain("Some Description");
        expect(component.find("h4").text()).toContain("(required)");
        expect(component.find(".input-radio-wrapper input").prop("value")).toBe("option1");
        expect(component.find(".input-radio-wrapper span").text()).toBe("Option1");
        expect(component.find(".input-radio-wrapper").length).toBe(1);
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

        expect(component.find("h4").text()).toContain("Some Description");
        expect(component.find("h4").text()).toContain("(optional)");
    });

    test("render with empty options", () => {
        component.setProps({
            options: [],
            errorMessageIfEmpty: "No options!"
        });

        expect(component.find(".input-radio-wrapper").length).toBe(0);
        expect(component.find("h4.radio-error").text()).toBe("No options!");
        expect(component.find("img").exists()).toBe(false); // not chosen yet!
    });

    test("render valid value", () => {
        component.setProps({
            value: "option1",
            options: [
                { value: "option1", displayName: "Option1" },
                { value: "option2", displayName: "Option2" }
            ],
        });
        component.setState({
            chosen: true
        });

        expect(component.find(".input-radio-wrapper input").at(0).prop("value"))
            .toBe("option1");
        expect(component.find(".input-radio-wrapper span").at(0).text())
            .toBe("Option1");
        expect(component.find(".input-radio-wrapper input").at(1).prop("value"))
            .toBe("option2");
        expect(component.find(".input-radio-wrapper span").at(1).text())
            .toBe("Option2");
        expect(component.find(".input-radio-wrapper").length).toBe(2);
        expect(component.find("h4.radio-error").exists()).toBe(false);

        expect(component.find("img").exists()).toBe(true);
        expect(component.find("img").prop("src")).toEqual(ImgCheckmark);
    });
});
