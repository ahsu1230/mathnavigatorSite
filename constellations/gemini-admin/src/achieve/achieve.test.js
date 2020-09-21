import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievePage } from "./achieve.js";

describe("Achievement Page", () => {
    const component = shallow(<AchievePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Achievements");
        expect(header.prop("addUrl")).toBe("/achievements/add");
    });

    test("renders 3 rows", () => {
        const achievements = [
            {
                year: 2020,
                achievements: [
                    {
                        id: 1,
                        year: 2020,
                        position: 1,
                        message: "Amazing",
                    },
                    {
                        id: 2,
                        year: 2020,
                        position: 2,
                        message: "Awesome",
                    },
                ],
            },
            {
                year: 2019,
                achievements: [
                    {
                        id: 3,
                        year: 2019,
                        position: 1,
                        message: "Possum",
                    },
                ],
            },
        ];
        component.setState({ achievements: achievements });

        let cards = component.find("RowCardBasic");
        let row0 = cards.at(0);
        expect(row0.prop("title")).toBe("Achievement in 2020");
        expect(row0.prop("editUrl")).toBe("/achievements/1/edit");

        let row1 = cards.at(1);
        expect(row1.prop("title")).toBe("Achievement in 2020");
        expect(row1.prop("editUrl")).toBe("/achievements/2/edit");

        let row2 = cards.at(2);
        expect(row2.prop("title")).toBe("Achievement in 2019");
        expect(row2.prop("editUrl")).toBe("/achievements/3/edit");
        expect(cards.length).toBe(3);
    });
});
