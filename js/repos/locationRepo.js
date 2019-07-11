'use strict';
import { locationMap } from './initPrograms.js';

export function getLocation(locationId) {
  return locationMap[locationId];
}
