import React from "react";
import { shallow } from "enzyme";
import { ClassEditPage } from "./classEdit.js";

describe("Class Edit Page", () => {
    const component = shallow(<ClassEditPage />);
    var inputSelect = component.find("InputSelect");
    var inputText = component.find("InputText");

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Add Class");
        expect(inputSelect.at(0).prop("label")).toBe("ProgramId");
        expect(inputSelect.at(1).prop("label")).toBe("SemesterId");
        expect(inputSelect.at(2).prop("label")).toBe("LocationId");
        expect(inputSelect.at(3).prop("label")).toBe("Class Availability");
        expect(inputSelect.length).toBe(4);

        expect(inputText.at(0).prop("label")).toBe("ClassKey");
        expect(inputText.at(1).prop("label")).toBe("Display Time");
        expect(inputText.at(2).prop("label")).toBe("Google Classroom Code");
        expect(inputText.at(3).prop("label")).toBe("Price Lump");
        expect(inputText.at(4).prop("label")).toBe("Price Per Session");
        expect(inputText.at(5).prop("label")).toBe("Payment Notes");
        expect(inputText.length).toBe(6);

        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    test("renders input data", () => {
        component.setState({
            isEdit: true,

            classKey: "class1",
            times: "Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm",
            programId: "ap_bc_calculus",
            semesterId: "2020_fall",
            locationId: "Churchill",
            googleClassCode: "",
            pricePerSession: 50,
        });
        inputSelect = component.find("InputSelect");
        inputText = component.find("InputText");

        expect(inputSelect.at(0).prop("value")).toBe("Churchill");
        expect(inputSelect.at(1).prop("value")).toBe(0);
        expect(inputSelect.length).toBe(2);

        expect(inputText.at(0).prop("value")).toBe(
            "Wed. 5:30pm - 7:30pm, Fri. 2:00pm - 4:00pm"
        );
        expect(inputText.at(1).prop("value")).toBe("");
        expect(inputText.at(2).prop("value")).toBe(0);
        expect(inputText.at(3).prop("value")).toBe(50);
        expect(inputText.at(4).prop("value")).toBe("");
        expect(inputText.length).toBe(5);
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
