import React from "react";
import { shallow } from "enzyme";
import AllPageHeader from "./allPageHeader.js";

describe("All Page Header", () => {
    const component = shallow(<AllPageHeader/>);

    test("renders", () => {
        expect(component.exists()).toBe(true);
    });

    test("render with props", () => {
        component.setProps({
            title: "Some title",
            addUrl: "/link/me",
            addButtonTitle: "ClickMe",
            description: "Some description"
        });

        expect(component.find("h1").text()).toBe("Some title");
        expect(component.find("Link").prop("to")).toBe("/link/me");
        expect(component.find("button").text()).toBe("ClickMe");
        expect(component.find("p").text()).toBe("Some description");
    });
});
