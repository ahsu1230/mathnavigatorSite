import React from "react";
import Enzyme, { shallow } from "enzyme";
import { PaymentTab } from "./payment.js";

describe("Payment Tab", () => {
    const component = shallow(<PaymentTab />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h2").at(0).text()).toContain(
            "Account Balance: $0.00"
        );
        expect(component.find("a")).toHaveLength(1);
        expect(component.find("h2").at(1).text()).toContain(
            "Your Payment History"
        );
    });

    test("renders 2 transactions", () => {
        const transactions = [
            {
                accountId: 1,
                amount: -1000,
                createdAt: "2020-01-01T00:00:00Z",
                id: 1,
                paymentNotes: null,
                paymentType: "charge",
                updatedAt: "2020-01-01T00:00:00Z",
            },
            {
                accountId: 1,
                amount: 1000,
                createdAt: "2020-01-01T00:00:00Z",
                id: 2,
                paymentNotes: "Payment note",
                paymentType: "pay_paypal",
                updatedAt: "2020-01-01T00:00:00Z",
            },
        ];
        component.setState({ transactions: transactions });

        let row0 = component.find("ul").at(1);
        expect(row0.text()).toContain("12/31/2019");
        expect(row0.text()).toContain("Paid (Paypal)");
        expect(row0.text()).toContain("$1,000.00");
        expect(row0.text()).toContain("$0.00");

        let row1 = component.find("ul").at(2);
        expect(row1.text()).toContain("12/31/2019");
        expect(row1.text()).toContain("Charge");
        expect(row1.text()).toContain("-$1,000.00");
        expect(row1.text()).toContain("-$1,000.00");
    });
});
