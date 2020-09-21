import React from "react";
import { shallow } from "enzyme";
import EditPageWrapper from "./editPageWrapper.js";

describe("EditPageWrapper", () => {
    const component = shallow(<EditPageWrapper />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("render with props", () => {
        let onSaveMocked = jest.fn();
        let onDeleteMocked = jest.fn();

        component.setProps({
            title: "Some title",
            content: <div></div>,
            onSave: onSaveMocked,
            onDelete: onDeleteMocked,
            isEdit: true,
        });

        expect(component.find("h1").text()).toBe("Some title");

        let buttons = component.find("button");
        expect(buttons.at(0).text()).toBe("Cancel");
        expect(buttons.at(1).text()).toBe("Delete");
        expect(buttons.at(2).text()).toBe("Save");
    });
});
