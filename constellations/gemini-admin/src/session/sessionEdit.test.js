import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { SessionEditPage } from "./sessionEdit.js";

describe("test", () => {
    const component = shallow(<SessionEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Edit Session");

        expect(component.find("h4").at(0).text()).toBe("Choose a Day");
        expect(component.find("h4").at(1).text()).toBe("Start Time");
        expect(component.find("h4").at(2).text()).toBe("End Time");
        expect(component.find("h4").at(3).text()).toBe("Canceled");

        expect(component.find("InputText").prop("label")).toBe("Notes");

        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").at(2).text()).toBe("Delete");
        expect(component.find("button").length).toBe(3);
    });

    test("renders input data", () => {
        const time = moment();
        component.setState({
            classId: "ap_bc_calculus_2020_fall_class1",
            startsAt: time,
            endsAt: time,
            canceled: true,
            notes: "",
        });

        // TODO: fix
        // expect(component.find("SingleDatePicker").prop("date")).toBe(time);
        expect(component.find("TimePicker").at(0).prop("time")).toBe(
            time.format("HH:mm")
        );
        expect(component.find("TimePicker").at(1).prop("time")).toBe(
            time.format("HH:mm")
        );
        expect(component.find("input").at(0).prop("checked")).toBe(true);
        expect(component.find("InputText").at(0).prop("value")).toBe("");
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
