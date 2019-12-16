import React from "react";
import Enzyme, { shallow } from "enzyme";
import { ProgramsPage, ProgramSection, ProgramCard } from "./programs.js";

describe("test", () => {
  const page = shallow(<ProgramsPage/>);

  test("page renders", () => {
    expect(page.exists()).toBe(true);
    expect(page.find("#view-program").exists()).toBe(true);
  });

  test("page renders with 1 semester, 1 program", () => {
    page.setState({
      semesterIds: ['summer_2019'],
      semesterMap: {
        'summer_2019': { title: "Summer 2019" }
      },
      programsBySemester: {
        'summer_2019': [ { title: "program1" } ]
      }
    });
    expect(page.find("ProgramSection").length).toBe(1);
  });

  test("page renders with 1 semester, 2 programs", () => {
    page.setState({
      semesterIds: ['summer_2019'],
      semesterMap: {
        'summer_2019': { title: "Summer 2019" }
      },
      programsBySemester: {
        'summer_2019': [ { title: "program1" }, { title: "program2"} ]
      }
    });
    expect(page.find("ProgramSection").length).toBe(1);
  });

  test("page renders with 2 semeters, 2 programs", () => {
    page.setState({
      semesterIds: ['summer_2019', 'fall_2019'],
      semesterMap: {
        'fall_2019': { id: 'fall_2019', title: "Fall 2019" },
        'summer_2019': { id: 'summer_2019', title: "Summer 2019" }
      },
      programsBySemester: {
        'fall_2019': [ { title: "program fall2"}],
        'summer_2019': [ { title: "program summer1" } ]
      }
    });
    expect(page.find("ProgramSection").length).toBe(2);
  });

  test("section renders with 2 programs", () => {
    const semester = { id: 'summer_2019', title: "Summer 2019"};
    const programs = [ { title: "program1"}, { title: "program2"} ];
    const section = shallow(<ProgramSection semester={semester} programs={programs}/>);
    expect(section.hasClass("section")).toBe(true);

    var cards = section.find("ProgramCard");
    expect(cards.length).toBe(2);
    expect(cards.at(0).prop('program').title).toMatch("program1");
    expect(cards.at(1).prop('program').title).toMatch("program2");
  });

  test("card renders with 1 class", () => {
    const semester = { id: 'summer_2019', title: "Summer 2019"};
    const program = { title: "program1" };
    const card = shallow(<ProgramCard semester={semester} program={program}/>);
    expect(card.hasClass("program-card-container")).toBe(true);

    card.setState({classes: [{ classKey: 'class1' }]});
    expect(card.find("Modal").exists()).toBe(false);
  });

  test("card renders with many classes", () => {
    const semester = { id: 'summer_2019', title: "Summer 2019"};
    const program = { title: "program1" };
    const card = shallow(<ProgramCard semester={semester} program={program}/>);
    expect(card.hasClass("program-card-container")).toBe(true);

    card.setState({
      classes: [
        { classKey: 'class1' },
        { classKey: 'class2' },
        { classKey: 'class3' }
      ]
    });
    expect(card.find("Modal").exists()).toBe(true);
  });
});
