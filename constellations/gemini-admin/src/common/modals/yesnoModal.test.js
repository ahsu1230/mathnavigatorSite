import React from "react";
import { shallow } from "enzyme";
import YesNoModal from "./yesnoModal.js";

describe("YesNoModal", () => {
    const component = shallow(<YesNoModal />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("#modal-view-yesno").exists()).toBe(true);
        expect(component.find("#modal-view-yesno button.reject").exists()).toBe(
            true
        );
        expect(component.find("#modal-view-yesno button.accept").exists()).toBe(
            true
        );
        expect(component.find("#modal-view-yesno button").length).toBe(2);
    });

    test("renders with text", () => {
        expect(component.find("#modal-view-yesno button.reject").text()).toBe(
            "No"
        );
        expect(component.find("#modal-view-yesno button.accept").text()).toBe(
            "Yes"
        );

        component.setProps({
            text: "Is this modal ok?",
            rejectText: "Definitely not!",
            acceptText: "Totally yes!",
        });

        expect(component.find("#modal-view-yesno button.reject").text()).toBe(
            "Definitely not!"
        );
        expect(component.find("#modal-view-yesno button.accept").text()).toBe(
            "Totally yes!"
        );
    });

    test("renders with onClick", () => {
        expect(
            component.find("#modal-view-yesno button.reject").prop("onClick")
        ).toBeUndefined();
        expect(
            component.find("#modal-view-yesno button.accept").prop("onClick")
        ).toBeUndefined();
        component.setProps({
            onAccept: jest.fn(),
            onReject: jest.fn(),
        });
        expect(
            component.find("#modal-view-yesno button.reject").prop("onClick")
        ).toBeDefined();
        expect(
            component.find("#modal-view-yesno button.accept").prop("onClick")
        ).toBeDefined();
    });
});
