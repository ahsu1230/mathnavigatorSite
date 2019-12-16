import React from "react";
import Enzyme, { shallow } from "enzyme";
import {
  EmailModal,
  STATE_LOADING,
  STATE_SUCCESS,
  STATE_FAIL
} from "./emailModal.js";

describe("test", () => {
  const component = shallow(<EmailModal/>);

  test("renders", () => {
    expect(component.exists()).toBe(true);
  });

  test("renders loading", () => {
    component.setProps({ loadingState: STATE_LOADING });
    expect(component.find("SubmitLoading").prop('show')).toBe(true);
    expect(component.find("SubmitSuccess").prop('show')).toBe(false);
    expect(component.find("SubmitFail").prop('show')).toBe(false);
  });

  test("renders success", () => {
    component.setProps({ loadingState: STATE_SUCCESS });
    expect(component.find("SubmitLoading").prop('show')).toBe(false);
    expect(component.find("SubmitSuccess").prop('show')).toBe(true);
    expect(component.find("SubmitFail").prop('show')).toBe(false);
  });

  test("renders fail", () => {
    component.setProps({ loadingState: STATE_FAIL });
    expect(component.find("SubmitLoading").prop('show')).toBe(false);
    expect(component.find("SubmitSuccess").prop('show')).toBe(false);
    expect(component.find("SubmitFail").prop('show')).toBe(true);
  });
});
