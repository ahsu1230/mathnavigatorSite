import React from "react";
import { shallow } from "enzyme";
import { ProgramCard } from "./programCard.js";

describe("Program Card", () => {
    const component = shallow(<ProgramCard />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("renders 1 card", () => {
        const program = {
            programId: "ap_bc_calculus",
            title: "AP BC Calculus",
            grade1: 10,
            grade2: 12,
            description: "Preparation for AP exam",
        };
        component.setProps({ program: program });

        expect(component.find("h2").text()).toBe("AP BC Calculus");
        expect(component.find("h3").text()).toBe("Grades 10 - 12");
        expect(component.find("button").text()).toBe("View Details");
    });

    test("renders showModal", () => {
        // Initial state
        component.setState({ showModal: false });
        expect(component.find("Modal").exists()).toBe(false);

        // Turn Modal On
        component.setProps({ classes: [{}, {}] });
        component.setState({ showModal: true });

        // After state
        let modal = component.find("Modal");
        expect(modal.prop("show")).toBe(true);
        expect(modal.prop("onDismiss")).toBeDefined();
        // TODO: Test if prop "content" is a ProgramModal
        expect(modal.prop("content")).toBeDefined();
        // expect(modal.prop("content")).toContain(<ProgramModal />);
    });
});
