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
        const sessions = [
            {
                startsAt: time,
                endsAt: time,
                canceled: true,
                notes: "",
            },
            {
                startsAt: time2,
                endsAt: time2,
                canceled: false,
                notes: "Moved to gym",
            },
        ];
        component.setProps({ sessions: sessions });
        expect(component.find(".row").length).toBe(2);

        let row0 = component.find(".row").at(0);
        expect(row0.find("span").at(0).text()).toBe("1");
        expect(row0.find("span").at(1).text()).toBe(time.format("l"));
        expect(row0.find("span").at(2).text()).toBe(
            time.format("dddd h:mma") + " - " + time.format("h:mma")
        );
        expect(row0.find("span").at(3).text()).toBe("Canceled");

        let row1 = component.find(".row").at(1);
        expect(row1.find("span").at(0).text()).toBe("2");
        expect(row1.find("span").at(1).text()).toBe(time2.format("l"));
        expect(row1.find("span").at(2).text()).toBe(
            time2.format("dddd h:mma") + " - " + time2.format("h:mma")
        );
        expect(row1.find("span").at(3).text()).toBe("");
    });
});
