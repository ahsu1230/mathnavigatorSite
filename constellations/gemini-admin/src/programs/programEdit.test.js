import React from "react";
import { shallow } from "enzyme";
import { ProgramEditPage } from "./programEdit.js";

describe("Program Edit Page", () => {
    const component = shallow(<ProgramEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Add Program");
        expect(component.find("InputText").at(0).prop("label")).toBe(
            "Program ID"
        );
        expect(component.find("InputText").at(1).prop("label")).toBe(
            "Program Name"
        );
        expect(component.find("InputText").at(2).prop("label")).toBe("Grade1");
        expect(component.find("InputText").at(3).prop("label")).toBe("Grade2");
        expect(component.find("InputText").at(4).prop("label")).toBe(
            "Description"
        );
        expect(component.find("InputText").length).toBe(5);

        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    test("renders input data", () => {
        component.setState({
            isEdit: true,
            programId: "ap_bc_calculus",
            name: "AP BC Calculus",
            grade1: 10,
            grade2: 12,
            description: "Preparation for the AP exam",
        });

        expect(component.find("InputText").at(0).prop("value")).toBe(
            "ap_bc_calculus"
        );
        expect(component.find("InputText").at(1).prop("value")).toBe(
            "AP BC Calculus"
        );
        expect(component.find("InputText").at(2).prop("value")).toBe(10);
        expect(component.find("InputText").at(3).prop("value")).toBe(12);
        expect(component.find("InputText").at(4).prop("value")).toBe(
            "Preparation for the AP exam"
        );
        expect(component.find("InputText").length).toBe(5);
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
