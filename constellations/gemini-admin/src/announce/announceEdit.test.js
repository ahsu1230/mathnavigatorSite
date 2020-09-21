import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { AnnounceEditPage } from "./announceEdit.js";

describe("test", () => {
    const component = shallow(<AnnounceEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const wrapper = component.find("EditPageWrapper");

        expect(wrapper.prop("title")).toContain("Add Announcement");
        expect(wrapper.prop("prevPageUrl")).toBe("announcements");
        expect(wrapper.prop("onSave")).toBeDefined();
        expect(wrapper.prop("onDelete")).toBeDefined();
    });
});
