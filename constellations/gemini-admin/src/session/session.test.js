import React from "react";
import moment from "moment";
import { shallow } from "enzyme";
import { SessionPage } from "./session.js";

describe("Sessions Page", () => {
    const component = shallow(<SessionPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Sessions per class");
        expect(component.find("h2").text()).toContain("Select Class");
        expect(component.find("select").exists()).toBe(true);
        expect(component.find("SessionAdd").exists()).toBe(true);

        expect(component.find("RowCardBasic").exists()).toBe(false);
        expect(component.find(".no-list").exists()).toBe(true);
    });

    test("renders select classId", () => {
        component.setState({ classId: "ap_bc_calculus_2020_fall_class1" });
        expect(component.find("select").prop("value")).toBe(
            "ap_bc_calculus_2020_fall_class1"
        );

        expect(component.find("SessionAdd").prop("classId")).toBe(
            "ap_bc_calculus_2020_fall_class1"
        );
    });

    test("renders several sessions", () => {
        const now = moment();
        component.setState({
            classId: "ap_bc_calculus_2020_fall_class1",
            sessions: [
                { id: 10, startsAt: now, endsAt: now.add(1, "hour") },
                {
                    id: 11,
                    startsAt: now.add(3, "hour"),
                    endsAt: now.add(4, "hour"),
                },
                {
                    id: 12,
                    startsAt: now.add(6, "hour"),
                    endsAt: now.add(7, "hour"),
                },
            ],
        });
        expect(component.find("RowCardBasic")).toHaveLength(3);
        expect(component.find(".no-list").exists()).toBe(false);
        expect(component.find("button.delete").exists()).toBe(true);
    });
});
