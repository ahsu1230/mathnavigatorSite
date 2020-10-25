import React from "react";
import { shallow } from "enzyme";
import { ClassProgramInfo } from "./classProgramInfo.js";

describe("Class ProgramInfo", () => {
    const program = {
        programId: "program1",
        title: "Program1",
        grade1: 1,
        grade2: 2,
        description: "asdf",
    };
    const component = shallow(<ClassProgramInfo program={program} />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h4").text()).toBe("Program Description:");
        expect(component.find("p.grades").text()).toContain("1 - 2");
        expect(component.find("p.description").text()).toEqual(
            program.description
        );
    });
});
