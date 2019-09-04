'use strict';

export var init = function() {
  if (process.env.NODE_ENV === 'production') {
    mixpanel.track("init");
  }
}

export var convertStringToBool = function(str) {
  var isBool = (typeof str == 'boolean');
  var isString = (typeof str == 'string');

  if (isBool) {
    return str;
  } else if (isString && str.toLowerCase() === "true") {
    return true;
  } else {
    return false;
  }
}

export var convertStringArray = function(str) {
  if (!str || (typeof str != 'string')) {
    return [];
  }

  // We assume the following format:
  // "[a, b, c]"
  var newStr = str.slice(0);
  var newStr = newStr.substring(1, newStr.length - 1);
  var arr = newStr.split(", ");
  return arr;
}

export function createFullClassObj(programObj, classObj) {
  var programId = classObj.programId;
  var className = classObj.className;
  classObj = assign({}, classObj, programObj);
  classObj.fullClassName = programObj.title + (className ? (" " + className) : "");
  classObj.programTitle = programObj.title;
  return classObj;
}
