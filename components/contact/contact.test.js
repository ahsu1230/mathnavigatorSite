import React from "react";
import Enzyme, { shallow } from "enzyme";
import { HashRouter as Router } from 'react-router-dom';
import { ContactPage } from "./contact.js";
import { ContactForm } from "./contactForm.js";
import { ContactInterestSection } from "./contactInterest.js";
import renderer from 'react-test-renderer';

describe("test", () => {

  test("renders", () => {
    const component = shallow(<ContactPage/>);
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toBe("Contact Us");
    expect(component.exists("ContactForm")).toBe(true);
  });

  test("parse location query", () => {
    var location = {
      search: "contact?interest=asdf"
    };
    var component = shallow(<ContactPage location={location}/>);
    var form = component.find("ContactForm");
    expect(form.prop('startingInterest')).toEqual(["asdf"]);
  });

  test("renders form", () => {
    const formComponent = shallow(<ContactForm/>);
    expect(formComponent.exists()).toBe(true);
    expect(formComponent.exists(".submit-container button")).toBe(true);

    const interestComponent = formComponent.find("ContactInterestSection");
    expect(interestComponent.exists()).toBe(true);
  });

  test("renders form with parameters", () => {
    var interested = ["asdf"];
    const formComponent = shallow(<ContactForm startingInterest={interested}/>);
    expect(formComponent.state('interestedPrograms')).toEqual(interested);

    const interestComponent = formComponent.find("ContactInterestSection");
    expect(interestComponent.prop('interested')).toEqual(interested);
  });

  test("renders interest", () => {
    const component = shallow(<ContactInterestSection/>);
    expect(component.exists()).toBe(true);
    expect(component.find("h2").text()).toBe("Interested Programs");
  });

  // test for updating interest section -> updates form.interestedPrograms

  test("snapshot", () => {
    const tree = renderer.create(
      <Router>
        <ContactPage/>
      </Router>
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
