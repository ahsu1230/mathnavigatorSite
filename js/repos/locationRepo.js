'use strict';
var _ = require('lodash/core');
import { locationMap } from './initPrograms.js';

export function getLocation(locationId) {
  return locationMap[locationId];
}
