import React from "react";
import { shallow } from "enzyme";
import { AnnouncePage } from "./announce.js";
import moment from "moment";

describe("test", () => {
    const component = shallow(<AnnouncePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Announcements");
        expect(component.find("AnnouncementGroup")).toHaveLength(0);
    });

    test("renders 1 card, 2 announcements", () => {
        const now = moment();
        const earlier = now.subtract(30, "minutes");
        const fakeAnnouncements = [
            {
                id: 1,
                postedAt: now.toDate(),
                message: "Awesome",
            },
            {
                id: 2,
                postedAt: earlier.toDate(),
                message: "Possum",
            },
        ];
        component.setState({ announcementList: fakeAnnouncements });
        expect(component.find("AnnouncementGroup")).toHaveLength(1);

        let group = component.find("AnnouncementGroup").at(0);
        expect(group.prop("announcements")[0]).toHaveProperty(
            "message",
            "Awesome"
        );
        expect(group.prop("announcements")[1]).toHaveProperty(
            "message",
            "Possum"
        );
    });

    test("renders 2 cards, 3 announcements", () => {
        const now = moment();
        const earlier = moment(now).subtract(3, "days");
        const fakeAnnouncements = [
            {
                id: 1,
                postedAt: now.toDate(),
                message: "Awesome",
            },
            {
                id: 2,
                postedAt: now.toDate(),
                message: "Possum",
            },
            {
                id: 3,
                postedAt: earlier.toDate(),
                message: "Dossum",
            },
        ];
        component.setState({ announcementList: fakeAnnouncements });
        expect(component.find("AnnouncementGroup")).toHaveLength(2);

        let group0 = component.find("AnnouncementGroup").at(0);
        expect(group0.prop("announcements")[0]).toHaveProperty(
            "message",
            "Awesome"
        );
        expect(group0.prop("announcements")[1]).toHaveProperty(
            "message",
            "Possum"
        );

        let group1 = component.find("AnnouncementGroup").at(1);
        expect(group1.prop("announcements")[0]).toHaveProperty(
            "message",
            "Dossum"
        );
    });
});
