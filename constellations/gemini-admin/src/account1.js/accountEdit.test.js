import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AccountEditPage } from "./accountEdit.js";

describe("Account Edit Page", () => {
    const component = shallow(<AccountEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Create Primary Contact for New Account"
        );
        expect(component.find("InputText").at(0).prop("label")).toBe("Notes");
        expect(component.find("InputText")).toHaveLength(1);

        expect(component.find("button").at(1).text()).toBe("Save");
        expect(component.find("button").at(0).text()).toBe("Cancel");
        expect(component.find("button")).toHaveLength(2);
    });

    test("renders input data", () => {
        component.setState({
            notes: "User notes",
        });
        expect(component.find("InputText").at(0).prop("value")).toBe(
            "User notes"
        );
        expect(component.find("InputText")).toHaveLength(1);
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
