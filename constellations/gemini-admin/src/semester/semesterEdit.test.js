import React from "react";
import Enzyme, { shallow } from "enzyme";
import { SemesterEditPage } from "./semesterEdit.js";

describe("Semester Edit Page", () => {
    const component = shallow(<SemesterEditPage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("InputText").at(0).prop("label")).toBe(
            "Semester ID"
        );
        expect(component.find("InputText").at(1).prop("label")).toBe("Title");
        expect(component.find("InputText").length).toBe(2);

        expect(component.find("button").at(0).text()).toBe("Save");
        expect(component.find("button").at(1).text()).toBe("Cancel");
        expect(component.find("button").length).toBe(2);
    });

    test("renders input data", () => {
        component.setState({
            isEdit: true,
            inputSemesterId: "Fall_2020",
            inputTitle: "Fall 2020",
        });

        expect(component.find("InputText").at(0).prop("value")).toBe(
            "Fall_2020"
        );
        expect(component.find("InputText").at(1).prop("value")).toBe(
            "Fall 2020"
        );
        expect(component.find("InputText").length).toBe(2);
    });
});
