import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { SessionEditPage } from "./sessionEdit.js";

describe("test", () => {
    const component = shallow(<SessionEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");
        expect(wrapper.prop("title")).toContain("Edit session");
        expect(wrapper.prop("prevPageUrl")).toBe("sessions");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
