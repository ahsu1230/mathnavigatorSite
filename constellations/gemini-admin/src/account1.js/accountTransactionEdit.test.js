import React from "react";
import Enzyme, { shallow } from "enzyme";
import { TransactionEditPage } from "./accountTransactionEdit.js";

describe("Account Transaction Edit Page", () => {
    const component = shallow(<TransactionEditPage />);

    test("renders", () => {
        component.setState({ accountId: 1 });
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Add Transaction for Account No. 1"
        );
        expect(component.find("InputText")).toHaveLength(2);
        expect(component.find("select")).toHaveLength(1);

        expect(component.find("button").at(1).text()).toBe("Save");
        expect(component.find("button").at(0).text()).toBe("Cancel");
        expect(component.find("button")).toHaveLength(2);
    });

    test("renders input data", () => {
        component.setState({
            type: "charge",
            amount: "-1000",
            notes: "Payment notes",
        });
        expect(component.find("select").at(0).prop("value")).toBe("charge");
        expect(component.find("InputText").at(0).prop("value")).toBe("-1000");
        expect(component.find("InputText").at(1).prop("value")).toBe(
            "Payment notes"
        );
        expect(component.find("select")).toHaveLength(1);
        expect(component.find("InputText")).toHaveLength(2);
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
