import React from "react";
import Enzyme, { shallow } from "enzyme";
import { ClassPage, ClassContent } from "./class.js";

describe("test", () => {

  var info = {
    classKey: 'asdf',
    classObj: { className: "some class", times: [] },
    location: { name: "school" },
    programObj: { title: "some program"}
  };

  test("renders", () => {
    const component = shallow(<ClassPage/>);
    expect(component.exists()).toBe(true);
    expect(component.exists("ErrorPage")).toBe(true);
    expect(component.exists("ClassContent")).toBe(false);

    var classKey = 'asdf';
    var classObj = {
      classKey: classKey
    };
    component.setState({ classKey: classKey, classObj: classObj});
    expect(component.exists("ErrorPage")).toBe(false);
    expect(component.exists("ClassContent")).toBe(true);
  });

  test("content renders", () => {
    const contentComponent = shallow(<ClassContent info={info}/>);
    expect(contentComponent.exists()).toBe(true);
    expect(contentComponent.find("h1").text())
      .toMatch("some program some class");
  });

  test("content renders with grades and description", () => {
    info.programObj.grade1 = "8";
    info.programObj.grade2 = "12";
    info.programObj.description = "some description";

    const contentComponent = shallow(<ClassContent info={info}/>);
    var text = contentComponent.find(".class-info-1").text();
    expect(text).toMatch("Grades: 8 - 12");
    expect(text).toMatch("some description");
  });

  test("content renders with location", () => {
    info.location = {
      name: "school",
      address1: "address1",
      address2: "address2",
      address3: "room3"
    };

    const contentComponent = shallow(<ClassContent info={info}/>);
    var targetComponent =
      contentComponent.find(".class-info-2")
        .find(".class-lines")
        .at(0);

    var text = targetComponent.text();
    expect(text).toMatch("school");
    expect(text).toMatch("address1");
    expect(text).toMatch("address2");
    expect(text).toMatch("room3");
  });

  test("content renders with times", () => {
    info.classObj = {times: ["day1 time1", "day2 time2"]};

    const contentComponent = shallow(<ClassContent info={info}/>);
    var targetComponent =
      contentComponent.find(".class-info-2")
        .find(".class-lines")
        .at(1);

    var text = targetComponent.text();
    expect(text).toMatch("day1 time1");
    expect(text).toMatch("day2 time2");
  });

  test("content renders with pricing (priceLump)", () => {
    info.classObj.priceLump = 100;
    info.classObj.pricePerSession = 0;

    const contentComponent = shallow(<ClassContent info={info}/>);
    var targetComponent =
      contentComponent.find(".class-info-2")
        .find(".class-lines")
        .at(2);

    var text = targetComponent.text();
    expect(text).toMatch("Price: $" + 100);
  });

  test("content renders with pricing (pricePerSession)", () => {
    info.classObj.priceLump = 0;
    info.classObj.pricePerSession = 100;

    const contentComponent = shallow(<ClassContent info={info}/>);
    var targetComponent =
      contentComponent.find(".class-info-2")
        .find(".class-lines")
        .at(2);

    var text = targetComponent.text();
    expect(text).toMatch("Price per session: $" + 100);
  });

  test("content renders with 0 sessions", () => {
    info.sessions = [];
    const contentComponent = shallow(<ClassContent info={info}/>);
    var targetComponent = contentComponent.find(".class-lines.not-avail");
    var text = targetComponent.text();
    expect(text).toMatch("No schedule available");
    expect(contentComponent.find("SessionLine").length).toBe(0);
  });

  test("content renders with many sessions", () => {
    info.sessions = [{date: "date0"}, {date: "date1"}];
    const contentComponent = shallow(<ClassContent info={info}/>);
    var lines = contentComponent.find("#view-schedule").find("SessionLine");
    expect(lines.length).toBe(2);
    expect(lines.at(0).prop('date')).toBe("date0");
    expect(lines.at(1).prop('date')).toBe("date1");
  });

  test("content renders with AYR sessions", () => {
    info.classObj.allYear = true;
    info.classObj.times = ["day1 time1"];

    const contentComponent = shallow(<ClassContent info={info}/>);
    expect(contentComponent.find("#view-schedule").text())
      .toMatch("Classes are held every week");
  });
});
