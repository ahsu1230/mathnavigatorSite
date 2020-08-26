import React from "react";
import Enzyme, { shallow } from "enzyme";
import { PasswordChange } from "./passwordChange.js";

describe("Password Change", () => {
    const component = shallow(<PasswordChange />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("a")).toHaveLength(1);

        component.setState({ tabOpen: true });
        expect(component.find("a")).toHaveLength(1);
        expect(component.find("ul")).toHaveLength(3);
        expect(component.find("button")).toHaveLength(2);
    });

    test("renders input data", () => {
        component.setState({
            oldPassword: "oldpassword",
            newPassword: "password",
            confirmPassword: "password",
        });
        expect(component.find("input").at(0).prop("value")).toBe("oldpassword");
        expect(component.find("input").at(1).prop("value")).toBe("password");
        expect(component.find("input").at(2).prop("value")).toBe("password");
    });
});
