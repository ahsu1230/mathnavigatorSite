import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AccountInfo } from "./accountInfo.js";

describe("Account Info", () => {
    const props = { id: 1, email: "", users: [], transactions: [] };
    const component = shallow(<AccountInfo {...props} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").at(0).text()).toContain("Users in Account");
        expect(component.find("h2").at(1).text()).toContain(
            "Transaction History"
        );
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
                notes: null,
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
        const email = "johndoe@gmail.com";
        component.setProps({ users: users, email: email });
        expect(component.find("#account-users div.row")).toHaveLength(2);

        let row0 = component.find("#account-users div.row").at(0);
        expect(row0.text()).toContain("John Richard Doe");
        expect(row0.text()).toContain("johndoe@gmail.com");
        expect(row0.text()).toContain("(guardian, primary contact)");

        let row1 = component.find("#account-users div.row").at(1);
        expect(row1.text()).toContain("Jane Doe");
        expect(row1.text()).toContain("janedoe@gmail.com");
        expect(row1.text()).toContain("(student)");
    });

    test("renders 2 transactions", () => {
        const transactions = [
            {
                accountId: 1,
                amount: -1000,
                createdAt: "2020-01-01T00:00:00Z",
                id: 1,
                notes: null,
                type: "charge",
                updatedAt: "2020-01-01T00:00:00Z",
            },
            {
                accountId: 1,
                amount: 1000,
                createdAt: "2020-01-01T00:00:00Z",
                id: 2,
                notes: "Payment note",
                type: "pay_paypal",
                updatedAt: "2020-01-01T00:00:00Z",
            },
        ];
        component.setProps({ transactions: transactions });
        expect(component.find("#account-transactions div.row")).toHaveLength(3); // rows + header

        let row0 = component.find("#account-transactions div.row").at(1);
        expect(row0.find("span")).toHaveLength(4);
        expect(row0.text()).toContain("charge");
        expect(row0.text()).toContain("-$1,000.00");

        let row1 = component.find("#account-transactions div.row").at(2);
        expect(row0.find("span")).toHaveLength(4);
        expect(row1.text()).toContain("pay_paypal");
        expect(row1.text()).toContain("$1,000.00");
        expect(row1.text()).toContain("Payment note");
    });
});
