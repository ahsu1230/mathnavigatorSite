import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchieveEditPage } from "./achieveEdit.js";

describe("test", () => {
    const component = shallow(<AchieveEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Add Achievement");
        expect(component.find("InputText").at(0).prop("label")).toBe("Year");
        expect(component.find("InputText").at(1).prop("label")).toBe("Message");
        expect(component.find("InputText").length).toBe(2);

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

        expect(component.find("InputText").at(0).prop("value")).toBe(2016);
        expect(component.find("InputText").at(1).prop("value")).toBe("Obama");
        expect(component.find("InputText").length).toBe(2);
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
        expect(modal.prop("onDismiss")).toBeDefined();
        // TODO: Test if prop "content" is a YesNoModal
        expect(modal.prop("content")).toBeDefined();
        // expect(modal.prop("content")).toContain(<YesNoModal/>);
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
        expect(modal.prop("onDismiss")).toBeDefined();
        // TODO: Test if prop "content" is an OkayModal
        expect(modal.prop("content")).toBeDefined();
        // expect(modal.prop("content")).toContain(<OkayModal/>);
    });
});
