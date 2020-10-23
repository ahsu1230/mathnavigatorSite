"use strict";
import { createBrowserHistory } from "history";

var customHistory = createBrowserHistory();

export const history = customHistory;

export const getCurrentPath = () => {
    // return window.location.hash;     // Use with HashRouter
    return window.location.pathname || ""; // Use with BrowserRouter
};

export const changePage = (url) => {
    // window.location.hash = url;      // Use with HashRouter

    // Use with BrowserRouter
    history.push({ pathname: url });
};
