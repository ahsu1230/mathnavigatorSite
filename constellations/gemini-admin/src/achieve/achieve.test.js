import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievePage } from "./achieve.js";

describe("Achievement Page", () => {
    const component = shallow(<AchievePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Achievements");
        expect(component.find("Link").text()).toContain("Add Achievement");
        expect(component.find("AchieveRow").length).toBe(0);
    });

    test("renders 2 rows", () => {
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
        expect(component.find("AchieveRow").length).toBe(3);

        let row0 = component.find("AchieveRow").at(0);
        expect(row0.prop("achieve")).toHaveProperty("id", 1);
        expect(row0.prop("achieve")).toHaveProperty("year", 2020);
        expect(row0.prop("achieve")).toHaveProperty("position", 1);
        expect(row0.prop("achieve")).toHaveProperty("message", "Amazing");

        let row1 = component.find("AchieveRow").at(1);
        expect(row1.prop("achieve")).toHaveProperty("id", 2);
        expect(row1.prop("achieve")).toHaveProperty("year", 2020);
        expect(row1.prop("achieve")).toHaveProperty("position", 2);
        expect(row1.prop("achieve")).toHaveProperty("message", "Awesome");

        let row2 = component.find("AchieveRow").at(2);
        expect(row2.prop("achieve")).toHaveProperty("id", 3);
        expect(row2.prop("achieve")).toHaveProperty("year", 2019);
        expect(row2.prop("achieve")).toHaveProperty("position", 1);
        expect(row2.prop("achieve")).toHaveProperty("message", "Possum");
    });
});
