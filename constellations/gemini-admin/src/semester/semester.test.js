import React from "react";
import { shallow } from "enzyme";
import { SemesterPage } from "./semester.js";

describe("Semester Page", () => {
    const component = shallow(<SemesterPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All Semesters");
        expect(header.prop("addUrl")).toBe("/semesters/add");
    });

    test("renders cards", () => {
        const list = [
            {
                id: 1,
                year: 2020,
                season: "fall",
                semesterId: "2020_fall",
                title: "Fall 2020",
            },
            {
                id: 2,
                year: 2021,
                season: "winter",
                semesterId: "2021_winter",
                title: "Winter 2021",
            },
        ];
        component.setState({ list: list });
        const cards = component.find("RowCardBasic");
        expect(cards.at(0).prop("title")).toBe("Fall 2020");
        expect(cards.at(0).prop("subtitle")).toBe("2020_fall");
        expect(cards.at(0).prop("editUrl")).toBe("/semesters/2020_fall/edit");

        expect(cards.at(1).prop("title")).toBe("Winter 2021");
        expect(cards.at(1).prop("subtitle")).toBe("2021_winter");
        expect(cards.at(1).prop("editUrl")).toBe("/semesters/2021_winter/edit");

        expect(cards.length).toBe(list.length);
    });
});
