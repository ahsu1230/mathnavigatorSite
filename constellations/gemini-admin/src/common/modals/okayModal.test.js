import React from "react";
import { shallow } from "enzyme";
import OkayModal from "./okayModal.js";

describe("OkayModal", () => {
    const component = shallow(<OkayModal />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("#modal-view-okay").exists()).toBe(true);
        expect(component.find("#modal-view-okay button").exists()).toBe(true);
    });

    test("renders with text", () => {
        component.setProps({ text: "Is this modal ok?" });
        expect(component.find("#modal-view-okay p").text()).toBe(
            "Is this modal ok?"
        );
        expect(component.find("#modal-view-okay button").text()).toBe("OK");
    });

    test("renders with onClick", () => {
        expect(
            component.find("#modal-view-okay button").prop("onClick")
        ).toBeUndefined();
        component.setProps({ onOkay: jest.fn() });
        expect(
            component.find("#modal-view-okay button").prop("onClick")
        ).toBeDefined();
    });
});
