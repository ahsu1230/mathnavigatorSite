import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { ClassSchedule } from "./classSchedule.js";

describe("Class Page", () => {
    const component = shallow(<ClassSchedule />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h3").text()).toContain("Schedule");
        expect(component.find(".row").length).toBe(0);
    });

    test("renders 2 rows", () => {
        const time = moment();
        const time2 = time.add(1, "weeks");
        const time3 = time.add(2, "weeks");
        const sessions = [
            {
                startsAt: time,
                endsAt: time,
                canceled: false,
                notes: "",
            },
            {
                startsAt: time2,
                endsAt: time2,
                canceled: false,
                notes: "Moved to gym",
            },
            {
                startsAt: time3,
                endsAt: time3,
                canceled: true,
                notes: "Holiday!",
            },
        ];
        component.setProps({ sessions: sessions });
        expect(component.find(".row").length).toBe(3);

        let row0 = component.find(".row").at(0);
        expect(row0.find("span").at(0).text()).toBe("1st");
        expect(row0.find("span").at(1).text()).toBe(time.format("l"));
        expect(row0.find("span").at(2).text()).toBe(
            time.format("dddd h:mma") + " - " + time.format("h:mma")
        );
        expect(row0.find("span").at(3).text()).toBe("");

        let row1 = component.find(".row").at(1);
        expect(row1.find("span").at(0).text()).toBe("2nd");
        expect(row1.find("span").at(1).text()).toBe(time2.format("l"));
        expect(row1.find("span").at(2).text()).toBe(
            time2.format("dddd h:mma") + " - " + time2.format("h:mma")
        );
        expect(row1.find("span").at(3).text()).toBe("Moved to gym");

        let row2 = component.find(".row").at(2);
        expect(row2.find("span").at(0).text()).toBe("");
        expect(row2.find("span").at(1).text()).toBe(time3.format("l"));
        expect(row2.find("span").at(2).text()).toBe("");
        expect(row2.find("span").at(3).text()).toBe("No Class. Holiday!");
    });
});
