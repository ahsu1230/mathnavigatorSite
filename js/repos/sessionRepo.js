'use strict';
import { sessionMap } from './initPrograms.js';

export function getSessions(programClassKey) {
  return sessionMap[programClassKey];
}
