"use strict";
/**
 * We expect a string "?asdf=zxcv&qwer=asdf"
 * that comes from props.location.search
 * which contains the entire query param string.
 *
 * Given this string, we return an object (key => value)
 *
 * Limitations: no duplicate keys!
 */
export const parseQueryParams = (search) => {
    let pairs = search.substring(1).split("&"); // ["key=value", "key=value"]
    let keyValueMap = {};
    pairs.forEach((pair) => {
        let kvs = pair.split("=");
        keyValueMap[kvs[0]] = kvs[1];
    });
    return keyValueMap;
};