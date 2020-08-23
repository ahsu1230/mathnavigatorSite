import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { SessionList } from "./sessionList.js";

describe("Session List Page", () => {
    const component = shallow(<SessionList />);

    test("render 2 rows", () => {
        const time = moment();
        const time2 = time.add(3, "days");
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
        const rows = component.find(".row");
        expect(rows.length).toBe(3);

        let row0 = rows.at(1);
        expect(row0.find("span").at(0).text()).toBe(moment(time).format("l"));
        expect(row0.find("span").at(1).text()).toBe(
            moment(time).format("LT - LT")
        );
        expect(row0.find("span").at(2).text()).toBe("Canceled");
        expect(row0.find("span").at(3).text()).toBe("");

        let row1 = rows.at(2);
        expect(row1.find("span").at(0).text()).toBe(moment(time2).format("l"));
        expect(row1.find("span").at(1).text()).toBe(
            moment(time2).format("LT - LT")
        );
        expect(row1.find("span").at(2).text()).toBe("Scheduled");
        expect(row1.find("span").at(3).text()).toBe("Moved to gym");
    });
});
