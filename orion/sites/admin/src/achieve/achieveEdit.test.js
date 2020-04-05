import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchieveEditPage } from "./achieveEdit.js";

describe("test", () => {
    const component = shallow(<AchieveEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Add Achievement");
        expect(component.find("h4").at(0).text()).toBe("Year");
        expect(component.find("h4").at(1).text()).toBe("Message");
        expect(component.find("h4").length).toBe(2);
        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    test("renders input data", () => {
        component.setState({
            isEdit: true,
            inputYear: 2016,
            inputMessage: "Obama",
        });

        let input0 = component.find("input").at(0);
        expect(input0.props().value).toBe(2016);

        let input1 = component.find("input").at(1);
        expect(input1.props().value).toBe("Obama");

        expect(component.find("input").length).toBe(2);
    });

    test("renders deleteModal", () => {
        // Initial state
        component.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
        expect(component.find("Modal").exists()).toBe(false);

        // Turn Modal on
        component.setState({
            showDeleteModal: true,
        });

        // After state
        let modal = component.find("Modal");
        expect(modal.prop("show")).toBe(true);
        expect(modal.prop("content")).toBeDefined();
        expect(modal.prop("onDismiss")).toBeDefined();
    });

    test("renders saveModal", () => {
        // Initial state
        component.setState({
            showDeleteModal: false,
            showSaveModal: false,
        });
        expect(component.find("Modal").exists()).toBe(false);

        // Turn Modal on
        component.setState({
            showSaveModal: true,
        });

        // After state
        let modal = component.find("Modal");
        expect(modal.prop("show")).toBe(true);
        expect(modal.prop("content")).toBeDefined();
        expect(modal.prop("onDismiss")).toBeDefined();
    });
});
