'use strict';
import { preReqMap } from './initPrograms.js';

export function getPrereqs(programId) {
  return preReqMap[programId];
}
