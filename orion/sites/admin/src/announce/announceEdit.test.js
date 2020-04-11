import React from "react";
import moment from "moment";
import Enzyme, { shallow } from "enzyme";
import { AnnounceEditPage } from "./announceEdit.js";

describe("test", () => {
    const component = shallow(<AnnounceEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").text()).toContain("Add Announcement");
        expect(component.find("h4").at(0).text()).toBe("Author");
        expect(component.find("h4").at(1).text()).toBe("Message");
        expect(component.find("h4").length).toBe(2);
        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    // TODO: test renders with isEdit
    // TODO: test renders with delete button
    // TODO: test renders with both modals

    test("renders DateTime", () => {
        let now = moment();
        component.setState({ inputPostedAt: now });
        let announceEditDateTime = component.find("AnnounceEditDateTime");
        expect(announceEditDateTime.prop("postedAt").isSame(now)).toBe(true);
        expect(announceEditDateTime.prop("onMomentChange")).toBeDefined();
    });
});
