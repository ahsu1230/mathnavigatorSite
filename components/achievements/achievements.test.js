import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievementPage } from "./achievements.js";

describe("test", () => {
  const component = shallow(<AchievementPage/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Our Student Achievements");
    expect(component.find("AchievementCard").length).toBe(0);
  });

  test("renders 1 card", () => {
    var years = ['2019'];
    var achievementsByYear = {
      2019: [{ message: 'asdf'}, {message: 'qwer'}]
    };
    component.setState({years: years, achievementsByYear: achievementsByYear});
    expect(component.find("AchievementCard").length).toBe(1);

    var targetCard = component.find("AchievementCard").at(0);
    expect(targetCard.prop('year')).toBe("2019");
    expect(targetCard.prop('achievements')[0].message).toBe("asdf");
    expect(targetCard.prop('achievements')[1].message).toBe("qwer");
  });

  test("renders 2 cards", () => {
    var years = ['2019', '2018'];
    var achievementsByYear = {
      2019: [{ message: 'asdf'}],
      2018: [{ message: 'qwer'}]
    };
    component.setState({years: years, achievementsByYear: achievementsByYear});
    expect(component.find("AchievementCard").length).toBe(2);

    var targetCard = component.find("AchievementCard").at(0);
    expect(targetCard.prop('year')).toBe("2019");
    expect(targetCard.prop('achievements')[0].message).toBe("asdf");

    targetCard = component.find("AchievementCard").at(1);
    expect(targetCard.prop('year')).toBe("2018");
    expect(targetCard.prop('achievements')[0].message).toBe("qwer");
  });
});
