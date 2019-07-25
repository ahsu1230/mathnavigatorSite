'use strict';

// https://www.w3resource.com/javascript/form/email-validation.php
const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
// https://www.w3resource.com/javascript/form/phone-no-validation.php
const phoneRegex = /^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/;

export const NameCheck = {
  errorMsg: "Must enter a name",
  validate: function(input) {
    return !!input || false;
  }
};

export const AgeCheck = {
  errorMsg: "Must be number",
  validate: function(input) {
    return input >= 0;
  }
};

export const SchoolCheck = {
  errorMsg: "Required",
  validate: function(input) {
    return !!input || false;
  }
}

export const GradeCheck = {
  errorMsg: "Must be 1-12",
  validate: function(input) {
    return input >= 1 && input <= 12;
  }
}

export const EmailCheck = {
  errorMsg: "Not a valid email",
  validate: function(input) {
    return emailRegex.test(input);
  }
};

export const PhoneCheck = {
  errorMsg: "Must be XXX-XXX-XXXX",
  validate: function(input) {
    return phoneRegex.test(input);
  }
};
