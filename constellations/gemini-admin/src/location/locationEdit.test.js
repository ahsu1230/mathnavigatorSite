import React from "react";
import Enzyme, { shallow } from "enzyme";
import { LocationEditPage } from "./locationEdit.js";

describe("Location Edit Page", () => {
    const component = shallow(<LocationEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("InputText").at(0).prop("label")).toBe(
            "Location ID"
        );
        expect(component.find("InputText").at(1).prop("label")).toBe("Street");
        expect(component.find("InputText").at(2).prop("label")).toBe("City");
        expect(component.find("InputText").at(3).prop("label")).toBe("State");
        expect(component.find("InputText").at(4).prop("label")).toBe("Zipcode");
        expect(component.find("InputText").at(5).prop("label")).toBe("Room");
        expect(component.find("InputText").length).toBe(6);

        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    test("renders input data", () => {
        component.setState({
            isEdit: true,
            inputlocationId: "wchs",
            inputStreet: "1300 Gainsborough Road",
            inputCity: "Potomac",
            inputState: "MD",
            inputZip: "20854",
            inputRoom: "23",
        });

        expect(component.find("InputText").at(0).prop("value")).toBe("wchs");
        expect(component.find("InputText").at(1).prop("value")).toBe(
            "1300 Gainsborough Road"
        );
        expect(component.find("InputText").at(2).prop("value")).toBe("Potomac");
        expect(component.find("InputText").at(3).prop("value")).toBe("MD");
        expect(component.find("InputText").at(4).prop("value")).toBe("20854");
        expect(component.find("InputText").at(5).prop("value")).toBe("23");
        expect(component.find("InputText").length).toBe(6);
    });
});
