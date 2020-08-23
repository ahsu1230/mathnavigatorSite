import React from "react";
import Enzyme, { shallow } from "enzyme";
import { UserClassPage } from "./userClass.js";

describe("User Class Page", () => {
    const component = shallow(<UserClassPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2")).toHaveLength(3);
        expect(component.find("Link").text()).toContain("< Back to Users");
        expect(component.find("h2").at(1).text()).toContain("User Information");
        expect(component.find("h2").at(2).text()).toContain("User Classes");
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
        const userClasses = [
            {
                id: 1,
                userId: 1,
                classObject: {
                    id: 1,
                    publishedAt: null,
                    programId: "ap_calculus",
                    semesterId: "2020_fall",
                    classKey: null,
                    classId: "ap_calculus_2020_fall",
                    locationId: "wchs",
                    times: "Mon 1:00pm - 5:00pm",
                    googleClassCode: null,
                    fullState: 0,
                    pricePerSession: null,
                    priceLump: null,
                    paymentNotes: null,
                },
                accountId: 1,
                state: 0,
            },
        ];
        component.setState({
            user: user,
            userClasses: userClasses,
            otherClassIds: ["ap_calculus_2020_fall"],
        });

        expect(component.find("p")).toHaveLength(3);
        expect(component.find("p").at(0).text()).toContain("John Richard Doe");
        expect(component.find("p").at(1).text()).toContain("johndoe@gmail.com");
        expect(component.find("p").at(2).text()).toContain("1234567890");

        expect(component.find("UserClassRow")).toHaveLength(1);
        let row0 = component.find("UserClassRow").at(0);
        expect(row0.prop("userClass")).toStrictEqual(userClasses[0]);

        expect(component.find("InputSelect").prop("options")[0]).toHaveProperty(
            "displayName",
            "ap_calculus_2020_fall"
        );
        expect(component.find("InputSelect").prop("options")[0]).toHaveProperty(
            "value",
            "ap_calculus_2020_fall"
        );

        expect(component.find("button")).toHaveLength(1);
    });
});
