import React from "react";
import { shallow } from "enzyme";
import { AnnouncePage } from "./announce.js";

describe("test", () => {
    const component = shallow(<AnnouncePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Announcements");
        expect(header.prop("addUrl")).toBe("/announcements/add");
    });

    test("renders cards", () => {
        const announcements = [
            {
                id: 1,
                author: "Joe",
                message: "Awesome",
            },
            {
                id: 2,
                author: "Schmoe",
                message: "Possum",
            },
        ];
        component.setState({ list: announcements });
        let cards = component.find("RowCardBasic");
        let card0 = cards.at(0);
        // todo: check postedAt
        expect(card0.prop("editUrl")).toBe("/announcements/1/edit");

        let card1 = cards.at(1);
        // todo: check postedAt
        expect(card1.prop("editUrl")).toBe("/announcements/2/edit");
        expect(cards.length).toBe(2);
    });
});
