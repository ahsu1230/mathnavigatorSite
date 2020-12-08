import React from "react";
import { shallow } from "enzyme";
import RowCardBasic from "./rowCardBasic.js";

describe("RowCardBasic", () => {
    const component = shallow(<RowCardBasic />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find(".row-card").exists()).toBe(true);
    });

    test("renders with editUrl", () => {
        expect(component.find("Link").exists()).toBe(false);
        component.setProps({ editUrl: "/edit/program/1" });
        expect(component.find("Link").prop("to")).toBe("/edit/program/1");
        expect(component.find("Link").text()).toBe("Edit >");
    });

    test("renders with title", () => {
        component.setProps({
            title: "Asdf",
            subtitle: "qwer",
        });
        expect(component.find("h2").text()).toBe("Asdf (qwer)");
    });

    test("renders with fields", () => {
        expect(component.find(".row-field-wrapper").length).toBe(0);
        component.setProps({
            fields: [
                {
                    label: "FieldA",
                    value: "asdf",
                },
                {
                    label: "FieldB",
                    value: 10,
                },
            ],
        });
        let wrappers = component.find(".row-field-wrapper");
        expect(wrappers.at(0).find("span.label").text()).toBe("FieldA:");
        expect(wrappers.at(0).find("span.value").text()).toBe("asdf");
        expect(wrappers.at(1).find("span.label").text()).toBe("FieldB:");
        expect(wrappers.at(1).find("span.value").text()).toBe("10");
        expect(wrappers.length).toBe(2);
    });

    test("renders with highlighted field", () => {
        component.setProps({
            fields: [
                {
                    label: "FieldA",
                    value: "asdf",
                },
            ],
        });
        expect(component.find(".row-field-wrapper.highlighted").exists()).toBe(
            false
        );

        component.setProps({
            fields: [
                {
                    label: "FieldA",
                    value: "asdf",
                    highlightFn: () => true,
                },
            ],
        });
        expect(component.find(".row-field-wrapper.highlighted").exists()).toBe(
            true
        );
    });

    test("renders with texts", () => {
        expect(component.find("row-text-wrapper").length).toBe(0);
        component.setProps({
            texts: [
                {
                    label: "Message1",
                    value: "I am inevitable.",
                },
                {
                    label: "Message2",
                    value: "I am IronMan.",
                },
            ],
        });

        let wrappers = component.find(".row-text-wrapper");
        expect(wrappers.at(0).find("span").text()).toBe("Message1:");
        expect(wrappers.at(0).find("p").text()).toBe("I am inevitable.");
        expect(wrappers.at(1).find("span").text()).toBe("Message2:");
        expect(wrappers.at(1).find("p").text()).toBe("I am IronMan.");
        expect(wrappers.length).toBe(2);
    });
});
