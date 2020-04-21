import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AnnouncePage } from "./announce.js";

describe("test", () => {
    const component = shallow(<AnnouncePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Announcements");
        expect(component.find("Link").text()).toContain("Add Announcement");
        expect(component.find("AnnounceRow").length).toBe(0);
    });

    test("renders 2 rows", () => {
        const announcements = [
            {
                Id: 1,
                author: "Joe",
                message: "Awesome",
            },
            {
                Id: 2,
                author: "Schmoe",
                message: "Possum",
            },
        ];
        component.setState({ list: announcements });
        let rows = component.find("AnnounceRow");

        expect(rows.length).toBe(2);
        let row0 = rows.at(0);
        expect(row0.prop("row")).toHaveProperty("Id", 1);
        expect(row0.prop("row")).toHaveProperty("author", "Joe");
        expect(row0.prop("row")).toHaveProperty("message", "Awesome");

        let row1 = rows.at(1);
        expect(row1.prop("row")).toHaveProperty("Id", 2);
        expect(row1.prop("row")).toHaveProperty("author", "Schmoe");
        expect(row1.prop("row")).toHaveProperty("message", "Possum");
    });
});
