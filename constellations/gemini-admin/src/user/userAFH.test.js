import React from "react";
import Enzyme, { shallow } from "enzyme";
import { UserAFHPage } from "./userAFH.js";

describe("User AFH Page", () => {
    const component = shallow(<UserAFHPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2")).toHaveLength(4);
        expect(component.find("Link").text()).toContain("< Back to Users");
        expect(component.find("h2").at(1).text()).toContain("User Information");
        expect(component.find("h2").at(2).text()).toContain(
            "User AskForHelp Sessions"
        );
        expect(component.find("h2").at(3).text()).toContain(
            "Schedule AskForHelp for User"
        );
        expect(component.find("InputSelect")).toHaveLength(1);
    });

    test("renders test data", () => {
        const user = {
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
        };
        const userAFHs = [
            {
                date: "2020-01-01T00:00:00Z",
                id: 1,
                locationId: "wchs",
                notes: null,
                subject: "math",
                timeString: "3:00pm - 4:00pm",
                title: "AP Calculus Practice Exam",
            },
        ];
        component.setState({
            user: user,
            userAFHs: userAFHs,
            otherAFHs: userAFHs,
        });

        expect(component.find("p")).toHaveLength(3);
        expect(component.find("p").at(0).text()).toContain("John Richard Doe");
        expect(component.find("p").at(1).text()).toContain("johndoe@gmail.com");
        expect(component.find("p").at(2).text()).toContain("1234567890");

        expect(component.find("#user-afh div.row")).toHaveLength(2);
        let row0 = component.find("#user-afh div.row").at(1);
        expect(row0.text()).toContain("1/1/2020");
        expect(row0.text()).toContain("AP Calculus Practice Exam");
        expect(row0.text()).toContain("math");

        expect(component.find("InputSelect").prop("options")).toStrictEqual([
            {
                displayName: "1/1/2020 math 3:00pm - 4:00pm",
                value: 1,
            },
        ]);
        expect(component.find("button")).toHaveLength(1);
    });
});
