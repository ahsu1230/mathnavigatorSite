import React from "react";
import { shallow } from "enzyme";
import RowCardColumns from "./rowCardColumns.js";

describe("RowCardColumns", () => {
    const component = shallow(<RowCardColumns />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find(".row-card").exists()).toBe(true);
        expect(component.find(".column-wrapper").exists()).toBe(true);
        expect(component.find(".column").length).toBe(0);
    });

    test("renders with title", () => {
        component.setProps({
            title: "Asdf",
            subtitle: "qwer",
        });
        expect(component.find("h2").text()).toBe("Asdf (qwer)");
    });

    test("renders with columns", () => {
        expect(component.find(".column").length).toBe(0);
        expect(component.find(".row-field-wrapper").length).toBe(0);
        component.setProps({
            fieldsList: [
                [
                    {
                        label: "FieldA",
                        value: "asdf",
                    },
                    {
                        label: "FieldB",
                        value: "qwer",
                    },
                ],
                [
                    {
                        label: "FieldC",
                        value: 10,
                    },
                    {
                        label: "FieldD",
                        value: 11,
                    },
                ],
            ],
        });

        let wrappers = component.find(".row-field-wrapper");
        expect(wrappers.at(0).find("span.label").text()).toBe("FieldA:");
        expect(wrappers.at(0).find("span.value").text()).toBe("asdf");
        expect(wrappers.at(1).find("span.label").text()).toBe("FieldB:");
        expect(wrappers.at(1).find("span.value").text()).toBe("qwer");
        expect(wrappers.at(2).find("span.label").text()).toBe("FieldC:");
        expect(wrappers.at(2).find("span.value").text()).toBe("10");
        expect(wrappers.at(3).find("span.label").text()).toBe("FieldD:");
        expect(wrappers.at(3).find("span.value").text()).toBe("11");
        expect(wrappers.length).toBe(4);

        expect(component.find(".column").length).toBe(2);
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
