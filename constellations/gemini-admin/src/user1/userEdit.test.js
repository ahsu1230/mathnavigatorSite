import React from "react";
import Enzyme, { shallow } from "enzyme";
import { UserEditPage } from "./userEdit.js";

describe("User Edit Page", () => {
    const component = shallow(<UserEditPage />);

    test("renders add page", () => {
        component.setState({ accountId: 1 });
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Add User");
        expect(component.find("UserInput")).toHaveLength(1);
        expect(component.find("InputText").at(0).prop("label")).toBe("Notes");

        expect(component.find("div#associated-account")).toHaveLength(1);
        expect(
            component.find("div#associated-account").find("h2").text()
        ).toContain("Associated Account");

        expect(component.find("button").at(1).text()).toBe("Save");
        expect(component.find("button").at(0).text()).toBe("Cancel");
        expect(component.find("button")).toHaveLength(2);
    });

    test("renders edit page", () => {
        component.setState({ isEdit: true });
        expect(component.find("h1").text()).toContain("Edit User");
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
