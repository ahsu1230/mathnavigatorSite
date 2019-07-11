'use strict';
var _ = require('lodash/core');
import { sessionMap } from './initPrograms.js';

export function getSessions(programClassKey) {
  return sessionMap[programClassKey];
}
