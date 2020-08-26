import React from "react";
import Enzyme, { shallow } from "enzyme";
import { SettingsTab } from "./settings.js";

describe("Settings Tab", () => {
    const props = { accountId: 1, users: [] };
    const component = shallow(<SettingsTab {...props} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").at(0).text()).toContain(
            "Your Account Information"
        );
        expect(component.find("h2").at(1).text()).toContain("User Information");
        expect(component.find("a")).toHaveLength(3);
    });

    test("renders 2 users", () => {
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
        component.setProps({
            users: users,
        });
        component.setState({ primaryEmail: "johndoe@gmail.com" });

        let row0 = component.find("ul.users-table").at(0);
        expect(row0.text()).toContain("John Doe");
        expect(row0.text()).toContain("johndoe@gmail.com");
        expect(row0.text()).toContain("1234567890");
        expect(row0.text()).toContain("Guardian (Primary Contact)");

        let row1 = component.find("ul.users-table").at(1);
        expect(row1.text()).toContain("Jane Doe");
        expect(row1.text()).toContain("janedoe@gmail.com");
        expect(row1.text()).toContain("1111111111");
        expect(row1.text()).toContain("Student");
        expect(row1.text()).toContain("Winston Churchill High School");
        expect(row1.text()).toContain("12th Grade, Graduation Year: 2020");
    });
});
