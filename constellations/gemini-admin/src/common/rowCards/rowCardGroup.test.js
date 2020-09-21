import React from "react";
import { shallow } from "enzyme";
import RowCardGroup from "./rowCardGroup.js";

describe("RowCardGroup", () => {
    const component = shallow(<RowCardGroup />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find(".row-card").exists()).toBe(true);
        expect(component.find(".group").length).toBe(0);
    });

    test("renders with title", () => {
        component.setProps({
            title: "Asdf",
            subtitle: "qwer",
        });
        expect(component.find("h2").text()).toBe("Asdf (qwer)");
    });

    test("renders with groups", () => {
        component.setProps({
            groupList: [
                {
                    editUrl: "/group/1",
                    fields: [
                        {
                            label: "FieldA",
                            value: "asdf",
                        },
                    ],
                    texts: [
                        {
                            label: "MessageA",
                            value: "longer message qwer",
                        },
                    ],
                },
            ],
        });

        expect(component.find(".group").length).toBe(1);
        expect(component.find(".group Link").prop("to")).toBe("/group/1");
        expect(
            component.find(".group .row-field-wrapper span.label").text()
        ).toBe("FieldA:");
        expect(
            component.find(".group .row-field-wrapper span.value").text()
        ).toBe("asdf");
        expect(component.find(".group .row-text-wrapper span").text()).toBe(
            "MessageA:"
        );
        expect(component.find(".group .row-text-wrapper p").text()).toBe(
            "longer message qwer"
        );
    });
});
