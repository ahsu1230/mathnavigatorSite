import React from "react";
import { shallow } from "enzyme";
import { AskForHelpPage } from "./afh.js";

describe("Ask For Help Page", () => {
    const component = shallow(<AskForHelpPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);

        const header = component.find("AllPageHeader");
        expect(header.prop("title")).toContain("All AskForHelp sessions");
        expect(header.prop("addUrl")).toBe("/afh/add");
    });

    test("renders cards", () => {
        const list = [
            {
                id: 1,
                title: "Afh1",
                subject: "math",
            },
            {
                id: 2,
                title: "Afh2",
                subject: "programming",
            },
        ];
        component.setState({ list: list });
        const cards = component.find("RowCardBasic");
        expect(cards.at(0).prop("title")).toBe("Afh1");
        expect(cards.at(0).prop("subtitle")).toBe("math");
        expect(cards.at(0).prop("editUrl")).toBe("/afh/1/edit");

        expect(cards.at(1).prop("title")).toBe("Afh2");
        expect(cards.at(1).prop("subtitle")).toBe("programming");
        expect(cards.at(1).prop("editUrl")).toBe("/afh/2/edit");

        expect(cards.length).toBe(list.length);
    });
});
