import React from "react";
import Enzyme, { shallow } from "enzyme";
import moment from "moment";
import { AnnounceEditDateTime } from "./announceEditDateTime.js";

describe("test", () => {
    const component = shallow(<AnnounceEditDateTime />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("renders past postedAt", () => {
        let postedAt = moment().subtract(7, "days");
        component.setProps({ postedAt });
        expect(component.find("h4").at(0).text()).toContain(
            "This post was published"
        );
        expect(component.find("AnnounceDateTimePicker").exists()).toBe(false);
    });

    test("renders future postedAt", () => {
        let postedAt = moment().add(7, "days");
        component.setProps({ postedAt });
        expect(component.find("h4").exists()).toBe(false);
        expect(component.find("AnnounceDateTimePicker").exists()).toBe(true);
    });
});
