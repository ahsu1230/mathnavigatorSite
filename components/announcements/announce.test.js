import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AnnouncePage } from "./announce.js";

describe("test", () => {
  const component = shallow(<AnnouncePage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Announcements");
    expect(component.find("AnnounceCard").length).toBe(0);
  });

  test("renders 1 card", () => {
    var list = {
      '8/1/2019': [{ message: 'asdf'}, { message: 'qwer'}]
    };
    component.setState({groupByDate: list});
    expect(component.find("AnnounceCard").length).toBe(1);

    var targetCard = component.find("AnnounceCard").at(0);
    expect(targetCard.prop('date')).toBe("8/1/2019");
    expect(targetCard.prop('announcements')[0].message).toBe("asdf");
    expect(targetCard.prop('announcements')[1].message).toBe("qwer");
  });

  test("renders 2 cards", () => {
    var list = {
      '8/1/2019': [{ message: 'asdf'}],
      '8/2/2019': [{ message: 'qwer'}]
    };
    component.setState({groupByDate: list});

    var cards = component.find("AnnounceCard");
    expect(cards.length).toBe(2);

    var targetCard = cards.at(0);
    expect(targetCard.prop('date')).toBe("8/1/2019");
    expect(targetCard.prop('announcements')[0].message).toBe("asdf");

    targetCard = cards.at(1);
    expect(targetCard.prop('date')).toBe("8/2/2019");
    expect(targetCard.prop('announcements')[0].message).toBe("qwer");
  });
});
