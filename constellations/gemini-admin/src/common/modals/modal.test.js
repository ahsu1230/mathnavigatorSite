import React from "react";
import { shallow } from "enzyme";
import { Modal } from "./modal.js";

describe("Modal", () => {
    const component = shallow(<Modal />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find(".modal-view").exists()).toBe(true);
        expect(component.find(".modal-overlay").exists()).toBe(true);
        expect(component.find(".modal").exists()).toBe(true);
    });

    test("renders show", () => {
        component.setProps({ show: false });
        expect(component.find(".modal-view.show").exists()).toBe(false);
        expect(component.find(".modal-overlay.show").exists()).toBe(false);

        component.setProps({ show: true });
        expect(component.find(".modal-view.show").exists()).toBe(true);
        expect(component.find(".modal-overlay.show").exists()).toBe(true);
    });

    test("renders withCloseButton", () => {
        component.setProps({ withClose: false });
        expect(component.find("button.close-x").exists()).toBe(false);

        component.setProps({ withClose: true });
        expect(component.find("button.close-x").prop("onClick")).toBeDefined();
    });

    test("renders with persistent overlay", () => {
        component.setProps({ persistent: false });
        expect(component.find(".modal-overlay").prop("onClick")).toBeDefined();

        component.setProps({ persistent: true });
        expect(
            component.find(".modal-overlay").prop("onClick")
        ).toBeUndefined();
    });

    // todo: test modalContent that gets passed in
});
