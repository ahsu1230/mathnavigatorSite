import React from "react";
import Enzyme, { shallow } from "enzyme";
import { UserPage } from "./user.js";

describe("User Page", () => {
    const component = shallow(<UserPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("Search Users");
        expect(component.find("input#searchbar")).toHaveLength(1);
        expect(component.find("UserRow")).toHaveLength(0);
    });

    test("renders 2 rows", () => {
        const users = [
            {
                accountId: 1,
                createdAt: "2020-01-01T00:00:00Z",
                email: "johndoe@gmail.com",
                firstName: "John",
                graduationYear: null,
                id: 1,
                isGuardian: true,
                lastName: "Doe",
                middleName: "Richard",
                notes: "User notes",
                phone: "1234567890",
                school: null,
                updatedAt: "2020-01-01T00:00:00Z",
            },
            {
                accountId: 1,
                createdAt: "2020-01-01T00:00:00Z",
                email: "janedoe@gmail.com",
                firstName: "Jane",
                graduationYear: "2020",
                id: 2,
                isGuardian: false,
                lastName: "Doe",
                middleName: null,
                notes: null,
                phone: "1111111111",
                school: "Winston Churchill High School",
                updatedAt: "2020-01-01T00:00:00Z",
            },
        ];
        component.setState({ list: users });
        expect(component.find("UserRow")).toHaveLength(2);

        let row0 = component.find("UserRow").at(0);
        expect(row0.prop("row")).toHaveProperty("id", 1);
        expect(row0.prop("row")).toHaveProperty("accountId", 1);
        expect(row0.prop("row")).toHaveProperty("firstName", "John");
        expect(row0.prop("row")).toHaveProperty("middleName", "Richard");
        expect(row0.prop("row")).toHaveProperty("lastName", "Doe");
        expect(row0.prop("row")).toHaveProperty("email", "johndoe@gmail.com");
        expect(row0.prop("row")).toHaveProperty("phone", "1234567890");
        expect(row0.prop("row")).toHaveProperty("school", null);
        expect(row0.prop("row")).toHaveProperty("graduationYear", null);
        expect(row0.prop("row")).toHaveProperty("notes", "User notes");

        let row1 = component.find("UserRow").at(1);
        expect(row1.prop("row")).toHaveProperty("id", 2);
        expect(row1.prop("row")).toHaveProperty("accountId", 1);
        expect(row1.prop("row")).toHaveProperty("firstName", "Jane");
        expect(row1.prop("row")).toHaveProperty("middleName", null);
        expect(row1.prop("row")).toHaveProperty("lastName", "Doe");
        expect(row1.prop("row")).toHaveProperty("email", "janedoe@gmail.com");
        expect(row1.prop("row")).toHaveProperty("phone", "1111111111");
        expect(row1.prop("row")).toHaveProperty(
            "school",
            "Winston Churchill High School"
        );
        expect(row1.prop("row")).toHaveProperty("graduationYear", "2020");
        expect(row1.prop("row")).toHaveProperty("notes", null);
    });

    test("renders search query", () => {
        component.setState({ searchQuery: "search users" });
        expect(component.find("input#searchbar").prop("value")).toBe(
            "search users"
        );
    });
});
