import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievementPage } from "./achievements.js";

describe("test", () => {
    const component = shallow(<AchievementPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain(
            "Math Navigator Achievements"
        );
        expect(component.find("h3").at(0).text()).toContain(
            "Congratulations to our students!"
        );
        expect(component.find("YearList")).toHaveLength(0);
    });

    test("renders 1 card, 2 achievements", () => {
        const fakeAchievements = [
            {
                id: 2,
                year: 2019,
                message: "Possum",
            },
            {
                id: 3,
                year: 2019,
                message: "Dossum",
            },
        ];
        component.setState({ achieveList: fakeAchievements });
        expect(component.find("YearList")).toHaveLength(1);

        let row = component.find("YearList").at(0);
        expect(row.prop("year")).toBe("2019");
        let rowAchievements = row.prop("achievements");
        expect(rowAchievements[0]).toHaveProperty("message", "Possum");
        expect(rowAchievements[1]).toHaveProperty("message", "Dossum");
    });

    test("renders 2 cards, 3 achievements", () => {
        const fakeAchievements = [
            {
                id: 1,
                year: 2020,
                message: "Awesome",
            },
            {
                id: 2,
                year: 2019,
                message: "Possum",
            },
            {
                id: 3,
                year: 2019,
                message: "Dossum",
            },
        ];
        component.setState({ achieveList: fakeAchievements });
        expect(component.find("YearList")).toHaveLength(2);

        let row0 = component.find("YearList").at(0);
        expect(row0.prop("year")).toBe("2020");
        let rowAchievements0 = row0.prop("achievements");
        expect(rowAchievements0[0]).toHaveProperty("message", "Awesome");

        let row1 = component.find("YearList").at(1);
        expect(row1.prop("year")).toBe("2019");
        let rowAchievements1 = row1.prop("achievements");
        expect(rowAchievements1[0]).toHaveProperty("message", "Possum");
        expect(rowAchievements1[1]).toHaveProperty("message", "Dossum");
    });
});
